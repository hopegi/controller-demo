package controller

import (
	"fmt"
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"

	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"

	"k8s.io/client-go/tools/record"
	"k8s.io/kubernetes/pkg/controller"
	"time"
)

type PodController struct {
	kubeClient      kubernetes.Interface
	clusterName     string
	podLister       corelisters.PodLister
	podListerSynced cache.InformerSynced
	broadcaster     record.EventBroadcaster
	recorder        record.EventRecorder
}

func mylog(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func NewPodController(kubeClient kubernetes.Interface, podInformer coreinformers.PodInformer, clusterName string) *PodController {

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	//eventBroadcaster.StartLogging(mylog)
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "pod_controller"})

	rc := &PodController{
		kubeClient:      kubeClient,
		clusterName:     clusterName,
		podLister:       podInformer.Lister(),
		podListerSynced: podInformer.Informer().HasSynced,
		broadcaster:     eventBroadcaster,
		recorder:        recorder,
	}
	return rc
}

func (p *PodController) Run(stopCh <-chan struct{}) {
	glog.Info("Starting pod controller\n")
	defer glog.Info("Shutting down pod controller\n")

	if !controller.WaitForCacheSync("pod", stopCh, p.podListerSynced) {
		return
	}

	if p.broadcaster != nil {
		p.broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: v1core.New(p.kubeClient.CoreV1().RESTClient()).Events("")})
	}

	go wait.NonSlidingUntil(func() {
		if err := p.reconcilePods(); err != nil {
			glog.Errorf("Couldn't reconcile pod: %v", err)
		}
	}, metav1.Duration{Duration: 10 * time.Second}.Duration, stopCh)

	<-stopCh
}

func (p *PodController) reconcilePods() error {
	//fmt.Println("reconcilePods ")
	glog.Infof("reconcilePods ")
	pods, err := p.podLister.List(labels.Everything())
	if err != nil {
		return fmt.Errorf("error listing pods: %v", err)
	}
	return p.reconcile(pods)
}

func (p *PodController) reconcile(pods []*v1.Pod) error {
	//fmt.Println("reconcile pods")
	glog.Infof("reconcile pods")
	for _, pod := range pods {
		fmt.Printf("pod name is %s.%s \n", (*pod).Namespace, (*pod).Name)
		if len(pod.Spec.ImagePullSecrets) > 0 {
			fmt.Println("ImagePullSecrets of ", pod.Name)
			for _, si := range pod.Spec.ImagePullSecrets {
				fmt.Printf(" %s \n", si.Name)
			}
		}

	}
	nodes, err := p.kubeClient.CoreV1().Nodes().List(metav1.ListOptions{LabelSelector: "node-role.kubernetes.io/master"})
	if err != nil {
		glog.Infof("get master error %v\n", err)
		return err
	}
	for _, n := range nodes.Items {

		n.Labels["hopegi/pod-count"] = fmt.Sprintf("%d", len(pods))
		_, err = p.kubeClient.CoreV1().Nodes().Update(&n)
		if err != nil {
			glog.Infof("label node error:%v ", err)
		} else {
			//if p.recorder!=nil {
			//	msg:=fmt.Sprintf("pod count is  %d",len(pods))
			//	p.recorder.Eventf(
			//		&v1.ObjectReference{
			//			Kind:"Node",
			//			Name: n.Name,
			//			UID: n.UID,
			//			Namespace:"",
			//		}, v1.EventTypeNormal, "SuccessCalculatePod",msg)
			//}
		}
	}
	if p.recorder != nil {
		msg := fmt.Sprintf("pod count is  %d", len(pods))
		for _, pod := range pods {
			p.recorder.Eventf(&v1.ObjectReference{
				Kind:      "Pod",
				Name:      pod.Name,
				UID:       pod.UID,
				Namespace: pod.Namespace,
			}, v1.EventTypeNormal, "SuccessCalculatePod", msg)
		}
	}
	return nil
}
