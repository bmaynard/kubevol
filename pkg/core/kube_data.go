package core

import (
	"fmt"

	"github.com/fatih/color"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KubeData struct {
	coreClient kubernetes.Interface
}

func NewKubeData(coreClient kubernetes.Interface) *KubeData {

	return &KubeData{coreClient}
}

func (o KubeData) GetPods(namespace string) *v1.PodList {
	pods, err := o.coreClient.CoreV1().Pods(namespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return pods
}

func (o KubeData) GetPod(podName string, namespace string) (*v1.Pod, error) {
	pod, err := o.coreClient.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		return nil, fmt.Errorf(color.RedString("Pod %s in namespace %s not found\n"), podName, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		return nil, fmt.Errorf(color.RedString("Error getting pod %s in namespace %s: %v\n"),
			podName, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		return nil, err
	}

	return pod, nil
}

func (o KubeData) GetConfigMap(configMapName string, namespace string) (*v1.ConfigMap, error) {
	return o.coreClient.CoreV1().ConfigMaps(namespace).Get(configMapName, metav1.GetOptions{})
}

func (o KubeData) GetSecret(secretName string, namespace string) (*v1.Secret, error) {
	return o.coreClient.CoreV1().Secrets(namespace).Get(secretName, metav1.GetOptions{})
}
