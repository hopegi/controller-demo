package app

import (
	"controller-demo/client/versiond"
	"controller-demo/controller"
	"controller-demo/list_and_watch"
	"github.com/golang/glog"
	//"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Run(stopCh <-chan struct{}) error {
	run := func(stopCh <-chan struct{}) {
		err := StartController(stopCh)
		if err != nil {
			glog.Fatalf("error running service controllers: %v", err)
		}
		select {}
	}
	///原本在此处有leader选举
	run(stopCh)

	panic("unreachable")
}

func StartController(stopCh <-chan struct{}) error {
	cfg, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		glog.Fatalf("error building kubernetes config:%s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	factory := informers.NewSharedInformerFactory(kubeClient, 0)

	podInformer := factory.Core().V1().Pods()
	pc := controller.NewPodController(kubeClient, podInformer, "k8s-cluster")
	go pc.Run(stopCh)

	dynamicKubeClient, _ := versiond.NewForConfig(cfg)
	ecsBindingInformer := list_and_watch.New(factory, nil).EcsBinding()
	ec := controller.NewEcsBindingController(dynamicKubeClient, ecsBindingInformer, "k8s-cluster")
	go ec.Run(stopCh)

	factory.Start(stopCh)

	return nil
}
