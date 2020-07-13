package watch

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/bmaynard/kubevol/pkg/core"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	WatchNamespace            = "kubevol"
	WatchConfigMapTrackerName = "kubevol-configmap-tracker"
	WatchSecretTrackerName    = "kubevol-secret-tracker"
	mutex                     = &sync.Mutex{}
	startTime                 = time.Now()
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

func (w Watch) UpdateTracker(trackerConfigmapName string, ns string, nm string) {
	mutex.Lock()
	defer mutex.Unlock()

	cmTracker, err := w.kubeData.GetConfigMap(trackerConfigmapName, WatchNamespace)
	trackerName := GetConfigMapKey(ns, nm)

	now := time.Now()
	currentTime := fmt.Sprintf("%d", now.Unix())

	if err != nil {
		record := &apiv1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: trackerConfigmapName,
			},
			Data: map[string]string{
				trackerName: currentTime,
			},
		}

		_, err := w.clientset.CoreV1().ConfigMaps(WatchNamespace).Create(record)

		if err != nil {
			w.f.Logger.Errorf("Error creating tracker configmap: \"%s\"; Error: \"%v\"", err)
		} else {
			w.f.Logger.Infof("Created tracker configmap and added object: \"%s\"; type: \"%s\"", nm, trackerConfigmapName)
		}

	} else {
		if cmTracker.Data == nil {
			cmTracker.Data = make(map[string]string)
		}

		cmTracker.Data[trackerName] = currentTime
		_, err := w.clientset.CoreV1().ConfigMaps(WatchNamespace).Update(cmTracker)

		if err != nil {
			w.f.Logger.Errorf("Unable to update tracker for type: \"%s\"; error: \"%v\"", trackerConfigmapName, err)
		} else {
			w.f.Logger.Infof("Updated tracker for object: \"%s\"; type: \"%s\"", nm, trackerConfigmapName)
		}
	}
}

func (w Watch) DeleteTracker(trackerConfigmapName string, ns string, nm string) {
	mutex.Lock()
	defer mutex.Unlock()

	cmTracker, err := w.kubeData.GetConfigMap(trackerConfigmapName, WatchNamespace)
	trackerName := GetConfigMapKey(ns, nm)

	if err != nil {
		w.f.Logger.Info("Unable find tracker configmap: \"%s\"", trackerConfigmapName)
		return
	}

	delete(cmTracker.Data, trackerName)
	_, dErr := w.clientset.CoreV1().ConfigMaps(WatchNamespace).Update(cmTracker)

	if dErr != nil {
		w.f.Logger.Errorf("Unable to delete object from tracker; Type: \"%s\"; Error: \"%v\"", trackerConfigmapName, err)
	} else {
		w.f.Logger.Infof("Deleted: \"%s\" from tracker; type: \"%s\"", nm, trackerConfigmapName)
	}
}

func GetConfigMapKey(namespace string, name string) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s|%s", namespace, name)))
	return hex.EncodeToString(hash[:])
}
