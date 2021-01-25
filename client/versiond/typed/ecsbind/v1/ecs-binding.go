package v1

import (
	"controller-demo/api/v1"
	"controller-demo/client/versiond/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type EcsBindingGetter interface {
	EcsBinding() EcsBindingInterface
}

type EcsBindingInterface interface {
	Create(*v1.EcsBinding) (*v1.EcsBinding, error)
	Update(*v1.EcsBinding) (*v1.EcsBinding, error)
	Delete(string, *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	List(metav1.ListOptions) (*v1.EcsBindingList, error)
	Get(name string, options metav1.GetOptions) (*v1.EcsBinding, error)
	Watch(options metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EcsBinding, err error)
}

type ecsBinding struct {
	client rest.Interface
}

func newEcsbindings(c *EcsV1Client) *ecsBinding {
	return &ecsBinding{
		client: c.RESTClient(),
	}
}

func (c *ecsBinding) Create(ecsbinding *v1.EcsBinding) (*v1.EcsBinding, error) {
	result := &v1.EcsBinding{}
	err := c.client.Post().
		Resource("ecsbinding").
		Body(ecsbinding).
		Do().
		Into(result)
	return result, err
}

func (c *ecsBinding) Update(ecsbinding *v1.EcsBinding) (*v1.EcsBinding, error) {
	result := &v1.EcsBinding{}
	err := c.client.Put().
		Resource("ecsbinding").
		Name(ecsbinding.Name).
		Body(ecsbinding).
		Do().
		Into(result)
	return result, err
}

func (c *ecsBinding) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ecsbinding").
		Name(name).
		Body(options).
		Do().
		Error()
}

func (c *ecsBinding) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Resource("ecsbinding").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

func (c *ecsBinding) List(opts metav1.ListOptions) (*v1.EcsBindingList, error) {
	result := &v1.EcsBindingList{}
	err := c.client.Get().
		Resource("ecsbinding").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return result, err
}

func (c *ecsBinding) Get(name string, options metav1.GetOptions) (*v1.EcsBinding, error) {
	result := &v1.EcsBinding{}
	err := c.client.Get().
		Resource("ecsbinding").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return result, err
}

func (c *ecsBinding) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		//Namespace(c.ns).
		Resource("ecsbinding").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

func (c *ecsBinding) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EcsBinding, err error) {
	result = &v1.EcsBinding{}
	err = c.client.Patch(pt).
		Resource("ecsbinding").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
