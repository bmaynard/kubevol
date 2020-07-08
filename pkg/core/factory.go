package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Factory struct {
	Logger     *Logger
	kubeclient kubernetes.Interface
}

func NewDepsFactory() *Factory {
	return &Factory{
		Logger: NewLogger(),
	}
}

func (f *Factory) CoreClient() (kubernetes.Interface, error) {
	if f.kubeclient == nil {
		clientset, err := f.getKubernetesConfigClient(viper.GetString("kubeconfig"))

		if err != nil {
			return nil, err
		}

		f.kubeclient = clientset
	}

	return f.kubeclient, nil
}

func (f *Factory) getKubeconfigRESTClient(kubeconfig string) (kubernetes.Interface, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("Building Core clientset: %s", err)

	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Building Core clientset: %s", err)
	}

	return clientset, nil
}

func (f *Factory) getKubernetesConfigClient(kubeconfig string) (kubernetes.Interface, error) {

	if kubeconfig == "" {
		kubeconfig = filepath.Join(homeDir(), ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		return f.getKubeconfigRESTClient("")
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return f.getKubeconfigRESTClient("")
	}

	return clientset, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
