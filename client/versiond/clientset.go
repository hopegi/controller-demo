package versiond

import (
	ecsbindv1 "controller-demo/client/versiond/typed/ecsbind/v1"
	"github.com/golang/glog"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	EcsV1() ecsbindv1.EcsV1Interface
}

type Clientset struct {
	*discovery.DiscoveryClient
	ecsV1 *ecsbindv1.EcsV1Client
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

func (c *Clientset) EcsV1() ecsbindv1.EcsV1Interface {
	return c.ecsV1
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.ecsV1, err = ecsbindv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil

}

func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.ecsV1 = ecsbindv1.NewForConfigOrDie(c)
	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

func New(c rest.Interface) *Clientset {

	var cs Clientset
	cs.ecsV1 = ecsbindv1.New(c)
	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
