package watch

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func GetConfigMapKey(namespace string, name string) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s|%s", namespace, name)))
	return hex.EncodeToString(hash[:])
}
