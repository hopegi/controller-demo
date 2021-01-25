package v1

import (
	ecsv1 "controller-demo/api/v1"
	"controller-demo/client/versiond/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

type EcsV1Interface interface {
	RESTClient() rest.Interface
	EcsBindingGetter
}

type EcsV1Client struct {
	restClient rest.Interface
}

func (c *EcsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *EcsV1Client) EcsBinding() EcsBindingInterface {
	return newEcsbindings(c)
}

func NewForConfig(c *rest.Config) (*EcsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &EcsV1Client{client}, nil

}

func setConfigDefaults(config *rest.Config) error {
	gv := ecsv1.GroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

func NewForConfigOrDie(c *rest.Config) *EcsV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *EcsV1Client {
	return &EcsV1Client{c}
}
