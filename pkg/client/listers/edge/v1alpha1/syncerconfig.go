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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	edgev1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/edge/v1alpha1"
)

// SyncerConfigClusterLister can list SyncerConfigs across all workspaces, or scope down to a SyncerConfigLister for one workspace.
// All objects returned here must be treated as read-only.
type SyncerConfigClusterLister interface {
	// List lists all SyncerConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.SyncerConfig, err error)
	// Cluster returns a lister that can list and get SyncerConfigs in one workspace.
	Cluster(clusterName logicalcluster.Name) SyncerConfigLister
	SyncerConfigClusterListerExpansion
}

type syncerConfigClusterLister struct {
	indexer cache.Indexer
}

// NewSyncerConfigClusterLister returns a new SyncerConfigClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewSyncerConfigClusterLister(indexer cache.Indexer) *syncerConfigClusterLister {
	return &syncerConfigClusterLister{indexer: indexer}
}

// List lists all SyncerConfigs in the indexer across all workspaces.
func (s *syncerConfigClusterLister) List(selector labels.Selector) (ret []*edgev1alpha1.SyncerConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*edgev1alpha1.SyncerConfig))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get SyncerConfigs.
func (s *syncerConfigClusterLister) Cluster(clusterName logicalcluster.Name) SyncerConfigLister {
	return &syncerConfigLister{indexer: s.indexer, clusterName: clusterName}
}

// SyncerConfigLister can list all SyncerConfigs, or get one in particular.
// All objects returned here must be treated as read-only.
type SyncerConfigLister interface {
	// List lists all SyncerConfigs in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.SyncerConfig, err error)
	// Get retrieves the SyncerConfig from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*edgev1alpha1.SyncerConfig, error)
	SyncerConfigListerExpansion
}

// syncerConfigLister can list all SyncerConfigs inside a workspace.
type syncerConfigLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all SyncerConfigs in the indexer for a workspace.
func (s *syncerConfigLister) List(selector labels.Selector) (ret []*edgev1alpha1.SyncerConfig, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.SyncerConfig))
	})
	return ret, err
}

// Get retrieves the SyncerConfig from the indexer for a given workspace and name.
func (s *syncerConfigLister) Get(name string) (*edgev1alpha1.SyncerConfig, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("SyncerConfig"), name)
	}
	return obj.(*edgev1alpha1.SyncerConfig), nil
}

// NewSyncerConfigLister returns a new SyncerConfigLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewSyncerConfigLister(indexer cache.Indexer) *syncerConfigScopedLister {
	return &syncerConfigScopedLister{indexer: indexer}
}

// syncerConfigScopedLister can list all SyncerConfigs inside a workspace.
type syncerConfigScopedLister struct {
	indexer cache.Indexer
}

// List lists all SyncerConfigs in the indexer for a workspace.
func (s *syncerConfigScopedLister) List(selector labels.Selector) (ret []*edgev1alpha1.SyncerConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.SyncerConfig))
	})
	return ret, err
}

// Get retrieves the SyncerConfig from the indexer for a given workspace and name.
func (s *syncerConfigScopedLister) Get(name string) (*edgev1alpha1.SyncerConfig, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("SyncerConfig"), name)
	}
	return obj.(*edgev1alpha1.SyncerConfig), nil
}
