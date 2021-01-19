package utils

import (
	"context"
	flowmessage "github.com/cloudflare/goflow/v3/pb"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net"
)

func enrich(fmsg *flowmessage.FlowMessage) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	srcPodIp := net.IP(fmsg.SrcAddr).String()
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{FieldSelector: "status.podIP=" + srcPodIp})
	if err != nil {
		log.Infof("K8S: Could not find source pod with IP %s: %v", srcPodIp, err)
	}
	if len(pods.Items) == 1 {
		pod := pods.Items[0]
		fmsg.K8SSrcPodName = pod.Status.PodIP
		fmsg.K8SSrcPodNamespace = pod.Namespace
		fmsg.K8SSrcPodNode = pod.Status.HostIP
	} else if len(pods.Items) > 1 {
		pod := pods.Items[0]
		fmsg.K8SSrcPodNode = pod.Status.HostIP
	}
	dstPodIp := net.IP(fmsg.DstAddr).String()
	pods, err = clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{FieldSelector: "status.podIP=" + dstPodIp})
	if err != nil {
		log.Infof("K8S: Could not find destination pod with IP %s: %v", dstPodIp, err)
	}
	if len(pods.Items) == 1 {
		pod := pods.Items[0]
		fmsg.K8SDstPodName = pod.Status.PodIP
		fmsg.K8SDstPodNamespace = pod.Namespace
		fmsg.K8SDstPodNode = pod.Status.HostIP
	} else if len(pods.Items) > 1 {
		pod := pods.Items[0]
		fmsg.K8SDstPodNode = pod.Status.HostIP
	}
}
