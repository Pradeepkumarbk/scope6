/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	openebs_io_v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	versioned "github.com/openebs/maya/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openebs/maya/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/maya/pkg/client/listers/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CStorPoolInformer provides access to a shared informer and lister for
// CStorPools.
type CStorPoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CStorPoolLister
}

type cStorPoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCStorPoolInformer constructs a new informer for CStorPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCStorPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCStorPoolInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCStorPoolInformer constructs a new informer for CStorPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCStorPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().CStorPools().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().CStorPools().Watch(options)
			},
		},
		&openebs_io_v1alpha1.CStorPool{},
		resyncPeriod,
		indexers,
	)
}

func (f *cStorPoolInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCStorPoolInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cStorPoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openebs_io_v1alpha1.CStorPool{}, f.defaultInformer)
}

func (f *cStorPoolInformer) Lister() v1alpha1.CStorPoolLister {
	return v1alpha1.NewCStorPoolLister(f.Informer().GetIndexer())
}
