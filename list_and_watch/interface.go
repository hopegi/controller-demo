package list_and_watch

import "k8s.io/client-go/informers/internalinterfaces"

type Interface interface {
	EcsBinding() EcsBindingInformer
}

type version struct {
	factory internalinterfaces.SharedInformerFactory
	//namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func New(factory internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{
		factory:          factory,
		tweakListOptions: tweakListOptions,
	}
}

func (v *version) EcsBinding() EcsBindingInformer {
	return &ecsBindingInformer{
		factory:          v.factory,
		tweakListOptions: v.tweakListOptions,
	}
}
