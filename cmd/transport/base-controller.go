/*
Copyright 2023 The KubeStellar Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	transportoptions "github.com/kubestellar/kubestellar/cmd/transport/options"
	edgeclientset "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned"
	edgeinformers "github.com/kubestellar/kubestellar/pkg/client/informers/externalversions"
	"github.com/kubestellar/kubestellar/pkg/transport"
)

const (
	defaultResyncPeriod = time.Duration(0)
)

func Run(transportImplementation transport.Transport) {
	logger := klog.Background().WithName(transport.ControllerName)
	ctx := klog.NewContext(context.Background(), logger)

	options := transportoptions.NewTransportOptions()
	fs := pflag.NewFlagSet(transport.ControllerName, pflag.ExitOnError)
	klog.InitFlags(flag.CommandLine)
	fs.AddGoFlagSet(flag.CommandLine)
	options.AddFlags(fs)
	fs.Parse(os.Args[1:])

	fs.VisitAll(func(flg *pflag.Flag) {
		logger.Info(fmt.Sprintf("Command line flag '%s' - value '%s'", flg.Name, flg.Value)) // log all arguments
	})

	// get the config for WDS
	wdsRestConfig, err := clientcmd.BuildConfigFromFlags("", options.WdsKubeConfig)
	if err != nil {
		logger.Error(err, "unable to build WDS kubeconfig")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	wdsRestConfig.UserAgent = transport.ControllerName

	// get the config for Transport space
	transportRestConfig, err := clientcmd.BuildConfigFromFlags("", options.TransportKubeConfig)
	if err != nil {
		logger.Error(err, "unable to build transport kubeconfig")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	transportRestConfig.UserAgent = transport.ControllerName

	edgeClientset, err := edgeclientset.NewForConfig(wdsRestConfig)
	if err != nil {
		logger.Error(err, "Failed to create edge clientset for Workload Description Space (WDS)")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	wdsDynamicClient, err := dynamic.NewForConfig(wdsRestConfig)
	if err != nil {
		logger.Error(err, "Failed to create dynamic k8s clientset for Workload Description Space (WDS)")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	edgeSharedInformerFactory := edgeinformers.NewSharedInformerFactoryWithOptions(edgeClientset, defaultResyncPeriod)
	edgePlacementDecisionInformer := edgeSharedInformerFactory.Edge().V2alpha1().PlacementDecisions()

	transportController, err := transport.NewTransportController(ctx, edgePlacementDecisionInformer, transportImplementation, edgeClientset, wdsDynamicClient, transportRestConfig, options.WdsName)

	// notice that there is no need to run Start method in a separate goroutine.
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	edgeSharedInformerFactory.Start(ctx.Done())

	if err := transportController.Run(ctx, options.Concurrency); err != nil {
		logger.Error(err, "failed to run transport controller")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	logger.Info("transport controller stopped")
}