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

package v2alpha1

import (
	time "time"

	brokerv2alpha1 "github.com/artemiscloud/activemq-artemis-operator/pkg/apis/broker/v2alpha1"
	versioned "github.com/artemiscloud/activemq-artemis-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/artemiscloud/activemq-artemis-operator/pkg/client/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/artemiscloud/activemq-artemis-operator/pkg/client/listers/broker/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ActiveMQArtemisScaledownInformer provides access to a shared informer and lister for
// ActiveMQArtemisScaledowns.
type ActiveMQArtemisScaledownInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.ActiveMQArtemisScaledownLister
}

type activeMQArtemisScaledownInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewActiveMQArtemisScaledownInformer constructs a new informer for ActiveMQArtemisScaledown type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewActiveMQArtemisScaledownInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredActiveMQArtemisScaledownInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredActiveMQArtemisScaledownInformer constructs a new informer for ActiveMQArtemisScaledown type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredActiveMQArtemisScaledownInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BrokerV2alpha1().ActiveMQArtemisScaledowns(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BrokerV2alpha1().ActiveMQArtemisScaledowns(namespace).Watch(options)
			},
		},
		&brokerv2alpha1.ActiveMQArtemisScaledown{},
		resyncPeriod,
		indexers,
	)
}

func (f *activeMQArtemisScaledownInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredActiveMQArtemisScaledownInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *activeMQArtemisScaledownInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&brokerv2alpha1.ActiveMQArtemisScaledown{}, f.defaultInformer)
}

func (f *activeMQArtemisScaledownInformer) Lister() v2alpha1.ActiveMQArtemisScaledownLister {
	return v2alpha1.NewActiveMQArtemisScaledownLister(f.Informer().GetIndexer())
}
