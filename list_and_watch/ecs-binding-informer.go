package list_and_watch

import (
	"controller-demo/client/versiond"
	//"controller-demo/client/versiond/scheme"
	"fmt"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/runtime/serializer"

	//"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	//"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers/internalinterfaces"
	"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	ecsv1 "controller-demo/api/v1"
	"time"
)

type EcsBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() EcsBindingLister
}

type ecsBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func NewFilteredEcsBindingInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	cfg, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		fmt.Printf("error building kubernetes config:%s", err.Error())
	}
	//cfg.GroupVersion=&ecsv1.GroupVersion
	//cfg.APIPath="/apis"
	//cfg.NegotiatedSerializer=serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	//resource:= meta_v1.APIResource{
	//	Name:"EcsBinding",
	//	Namespaced:false,
	//
	//}
	//dClient,err:= dynamic.NewClient(cfg)
	clientset, err := versiond.NewForConfig(cfg)
	if err != nil {
		fmt.Printf("create newclient error %v\n", err)
	}
	return cache.NewSharedIndexInformer(&cache.ListWatch{
		ListFunc: func(options meta_v1.ListOptions) (object runtime.Object, e error) {
			if tweakListOptions != nil {
				tweakListOptions(&options)
			}
			return clientset.EcsV1().EcsBinding().List(options)
			//return dClient.ParameterCodec(scheme.ParameterCodec).Resource(&resource,"").List(options)
		},
		WatchFunc: func(options meta_v1.ListOptions) (i watch.Interface, e error) {
			if tweakListOptions != nil {
				tweakListOptions(&options)
			}
			//return dClient.ParameterCodec(scheme.ParameterCodec).Resource(&resource,"").Watch(options)
			return clientset.EcsV1().EcsBinding().Watch(options)
		},
	}, &ecsv1.EcsBinding{}, resyncPeriod, indexers)
}

func (e *ecsBindingInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEcsBindingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, e.tweakListOptions)
}

func (e *ecsBindingInformer) Informer() cache.SharedIndexInformer {
	return e.factory.InformerFor(&ecsv1.EcsBinding{}, e.defaultInformer)
}

func (e *ecsBindingInformer) Lister() EcsBindingLister {
	return NewEcsBindingLister(e.Informer().GetIndexer())
}
