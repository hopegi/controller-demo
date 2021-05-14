package main

import (
	"controller-demo/app"
	//"controller-demo/client/versiond/scheme"
	//"k8s.io/apimachinery/pkg/runtime/serializer"
	//"k8s.io/client-go/dynamic"
	//"k8s.io/client-go/tools/clientcmd"

	"fmt"

	//ecsv1 "controller-demo/api/v1"
	//meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"os/signal"
	"syscall"
)

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM, os.Kill}

func main() {

	//cfg, _ := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")

	//clientset,_:=versiond.NewForConfig(cfg)
	//list,err:= clientset.EcsV1().EcsBinding().List(v1.ListOptions{})
	//ecs,err:= clientset.EcsV1().EcsBinding().Get("ag-config-feaif",v1.GetOptions{})
	//if ecs.Annotations==nil{
	//	ecs.Annotations= map[string]string{}
	//}
	//ecs.Annotations["aaa"]="212"
	//_,err= clientset.EcsV1().EcsBinding().Update(ecs)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(len(list.Items))
	//

	//cfg.GroupVersion = &ecsv1.GroupVersion
	//cfg.APIPath = "/apis"
	////cfg.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	//dClient, err := dynamic.NewClient(cfg)
	//if err != nil {
	//	fmt.Println("new client error ", err)
	//}
	//if err != nil {
	//	return
	//}
	//list, err := dClient.
	//	//ParameterCodec(scheme.ParameterCodec).
	//	Resource(&meta_v1.APIResource{Name: "Ecsbinding", Namespaced: false}, "").
	//	//List(meta_v1.ListOptions{})
	//	Get("ag-config-feaif", meta_v1.GetOptions{})
	//if err != nil {
	//	fmt.Println("client call error ", err)
	//}
	//fmt.Println(*list)
	//list.SetName(list.GetName() + "-copy")
	//list.SetAPIVersion("hopegi.com/v1")
	//list.SetKind("EcsBinding")
	//list.SetResourceVersion("")
	//createRes, err := dClient.
	//	Resource(&meta_v1.APIResource{Name: "Ecsbinding", Namespaced: false}, "").
	//	Create(list)
	//if err != nil {
	//	fmt.Println("client create error ", err)
	//} else {
	//	fmt.Println("create res ", createRes)
	//}
	//return
	stopCh := SetupSignalHandler()
	if err := app.Run(stopCh); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func SetupSignalHandler() (stopCh <-chan struct{}) {
	//防止多次调用，第二次调用直接panics
	close(onlyOneSignalHandler)

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()
	return stop
}
