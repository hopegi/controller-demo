package list_and_watch

import (
	ecsv1 "controller-demo/api/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type EcsBindingLister interface {
	List(selector labels.Selector) ([]*ecsv1.EcsBinding, error)
	Get(name string) (*ecsv1.EcsBinding, error)
}

type ecsBindingLister struct {
	indexer cache.Indexer
}

func NewEcsBindingLister(indexer cache.Indexer) EcsBindingLister {
	return &ecsBindingLister{
		indexer: indexer,
	}
}

func (e *ecsBindingLister) List(selector labels.Selector) (ret []*ecsv1.EcsBinding, err error) {
	err = cache.ListAll(e.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*ecsv1.EcsBinding))
	})
	return
}

func (e *ecsBindingLister) Get(name string) (*ecsv1.EcsBinding, error) {
	obj, exists, err := e.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("ecsbind"), name)
	}
	return obj.(*ecsv1.EcsBinding), nil
}
