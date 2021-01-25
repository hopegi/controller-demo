package controller

import (
	ecsV1 "controller-demo/api/v1"
	"controller-demo/client/versiond"
	"controller-demo/list_and_watch"
	"fmt"
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	"sync"
	"time"
)

var WorkerCount = 1

type EcsBindingController struct {
	kubeClient              *versiond.Clientset
	clusterName             string
	ecsbingdingLister       list_and_watch.EcsBindingLister
	ecsbingdingListerSynced cache.InformerSynced
	broadcaster             record.EventBroadcaster
	recorder                record.EventRecorder

	ecsQueue workqueue.DelayingInterface //workqueue.NewNamedDelayingQueue("service")
	lock     sync.Mutex
	//cloudProvider cloudproviders.CloudProvider
}

func NewEcsBindingController(kubeclient *versiond.Clientset, informer list_and_watch.EcsBindingInformer, clusterName string) *EcsBindingController {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "ecsbinding_controller"})

	ec := &EcsBindingController{
		kubeClient:              kubeclient,
		clusterName:             clusterName,
		ecsbingdingLister:       informer.Lister(),
		ecsbingdingListerSynced: informer.Informer().HasSynced,
		broadcaster:             eventBroadcaster,
		recorder:                recorder,

		ecsQueue: workqueue.NewNamedDelayingQueue("EcsBinding"),
	}

	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ecsbindingg := obj.(*ecsV1.EcsBinding)
			fmt.Printf("controller: Add event, ecsbinding [%s]\n", ecsbindingg.Name)
			ec.syncEcsbinding(ecsbindingg)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			ecs1, ok1 := oldObj.(*ecsV1.EcsBinding)
			ecs2, ok2 := newObj.(*ecsV1.EcsBinding)
			if ok1 && ok2 && !reflect.DeepEqual(ecs1, ecs2) {
				fmt.Printf("controller: Update event, ecsbinding [%s]\n", ecs1.Name)
				ec.syncEcsbinding(ecs1)
			}
		},
		DeleteFunc: func(obj interface{}) {
			ecsbindingg := obj.(*ecsV1.EcsBinding)
			fmt.Printf("controller: Delete event, ecsbinding [%s]\n", ecsbindingg.Name)
			ec.syncEcsbinding(ecsbindingg)
		},
	})

	return ec
}

func (e *EcsBindingController) syncEcsbinding(bind *ecsV1.EcsBinding) {
	//key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(bind)
	//if err != nil {
	//	fmt.Printf("Couldn't get key for object %#v: %v \n", bind, err)
	//	return
	//}
	key := bind.Name
	fmt.Printf("working queue add ecsbinding %s\n", key)
	e.ecsQueue.Add(key)
}

func (e *EcsBindingController) Run(stopCh <-chan struct{}) {
	defer runtime.HandleCrash()
	defer e.ecsQueue.ShutDown()

	for i := 0; i < WorkerCount; i++ {
		go wait.Until(e.worker, time.Second*5, stopCh)
	}

	<-stopCh
}

func (e *EcsBindingController) worker() {
	for {
		func() {
			key, quit := e.ecsQueue.Get()
			if quit {
				return
			}
			defer e.ecsQueue.Done(key)
			err := e.handleQueue(key.(string))
			if err != nil {
				fmt.Println("controller: error handle ecsbinding  ", err)
			}
		}()
	}
}

func (e *EcsBindingController) handleQueue(key string) error {
	name := key
	e.lock.Lock()
	defer e.lock.Unlock()
	ecsbinding, err := e.ecsbingdingLister.Get(name)
	switch {
	case errors.IsNotFound(err):
		fmt.Println("finizer ecsbinding ", name)
		return err
	case err != nil:
		e.ecsQueue.Add(key)
		fmt.Println("Unable to retrieve ecsbingding  ", key, " error ", err)
		return err
	default:
		if ecsbinding.Annotations == nil {
			ecsbinding.Annotations = map[string]string{}
		}
		ecsbinding.Annotations["has-reconcile"] = "true"
		_, err := e.kubeClient.EcsV1().EcsBinding().Update(ecsbinding)
		return err
	}

	return nil
}
