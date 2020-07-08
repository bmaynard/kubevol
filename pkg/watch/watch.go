package watch

import (
	"sync"

	"github.com/bmaynard/kubevol/pkg/core"

	"k8s.io/client-go/kubernetes"
)

var (
	WatchNamespace            = "kubevol"
	WatchConfigMapTrackerName = "kubevol-configmap-tracker"
	WatchSecretTrackerName    = "kubevol-secret-tracker"
	mutex                     = &sync.Mutex{}
)

type Watch struct {
	kubeData  *core.KubeData
	f         *core.Factory
	clientset kubernetes.Interface
}

func NewWatch(f *core.Factory) *Watch {
	clientset, err := f.CoreClient()
	kubeData := core.NewKubeData(clientset)

	if err != nil {
		f.Logger.Fatal(err)
	}

	return &Watch{
		kubeData:  kubeData,
		f:         f,
		clientset: clientset,
	}
}
