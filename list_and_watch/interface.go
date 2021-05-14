package list_and_watch

import (
	"controller-demo/client/versiond"
	"k8s.io/client-go/informers/internalinterfaces"
)

type Interface interface {
	EcsBinding() EcsBindingInformer
}

type version struct {
	clientSet *versiond.Clientset
	factory   internalinterfaces.SharedInformerFactory
	//namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func New(factory internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc, clientSet *versiond.Clientset) Interface {
	return &version{
		factory:          factory,
		tweakListOptions: tweakListOptions,
		clientSet:        clientSet,
	}
}

func (v *version) EcsBinding() EcsBindingInformer {
	return &ecsBindingInformer{
		factory:          v.factory,
		tweakListOptions: v.tweakListOptions,
		clientSet:        v.clientSet,
	}
}
