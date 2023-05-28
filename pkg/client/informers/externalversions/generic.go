//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeStellar Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package informers

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	edgev1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/edge/v1alpha1"
	metav1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/meta/v1alpha1"
)

type GenericClusterInformer interface {
	Cluster(logicalcluster.Name) GenericInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() kcpcache.GenericClusterLister
}

type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericClusterInformer struct {
	informer kcpcache.ScopeableSharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.informer
}

// Lister returns the GenericClusterLister.
func (f *genericClusterInformer) Lister() kcpcache.GenericClusterLister {
	return kcpcache.NewGenericClusterLister(f.Informer().GetIndexer(), f.resource)
}

// Cluster scopes to a GenericInformer.
func (f *genericClusterInformer) Cluster(clusterName logicalcluster.Name) GenericInformer {
	return &genericInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().ByCluster(clusterName),
	}
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	lister   cache.GenericLister
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return f.lister
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error) {
	switch resource {
	// Group=edge.kcp.io, Version=V1alpha1
	case edgev1alpha1.SchemeGroupVersion.WithResource("customizers"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Edge().V1alpha1().Customizers().Informer()}, nil
	case edgev1alpha1.SchemeGroupVersion.WithResource("edgeplacements"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Edge().V1alpha1().EdgePlacements().Informer()}, nil
	case edgev1alpha1.SchemeGroupVersion.WithResource("edgesyncconfigs"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Edge().V1alpha1().EdgeSyncConfigs().Informer()}, nil
	case edgev1alpha1.SchemeGroupVersion.WithResource("singleplacementslices"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Edge().V1alpha1().SinglePlacementSlices().Informer()}, nil
	case edgev1alpha1.SchemeGroupVersion.WithResource("syncerconfigs"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Edge().V1alpha1().SyncerConfigs().Informer()}, nil
	// Group=meta.kcp.io, Version=V1alpha1
	case metav1alpha1.SchemeGroupVersion.WithResource("apiresources"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Meta().V1alpha1().APIResources().Informer()}, nil
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedScopedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=edge.kcp.io, Version=V1alpha1
	case edgev1alpha1.SchemeGroupVersion.WithResource("customizers"):
		informer := f.Edge().V1alpha1().Customizers().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case edgev1alpha1.SchemeGroupVersion.WithResource("edgeplacements"):
		informer := f.Edge().V1alpha1().EdgePlacements().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case edgev1alpha1.SchemeGroupVersion.WithResource("edgesyncconfigs"):
		informer := f.Edge().V1alpha1().EdgeSyncConfigs().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case edgev1alpha1.SchemeGroupVersion.WithResource("singleplacementslices"):
		informer := f.Edge().V1alpha1().SinglePlacementSlices().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	case edgev1alpha1.SchemeGroupVersion.WithResource("syncerconfigs"):
		informer := f.Edge().V1alpha1().SyncerConfigs().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	// Group=meta.kcp.io, Version=V1alpha1
	case metav1alpha1.SchemeGroupVersion.WithResource("apiresources"):
		informer := f.Meta().V1alpha1().APIResources().Informer()
		return &genericInformer{lister: cache.NewGenericLister(informer.GetIndexer(), resource.GroupResource()), informer: informer}, nil
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
