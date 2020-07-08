package watch

import (
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (w Watch) UpateConfigMapTracker(old, new interface{}) {
	cm := new.(*apiv1.ConfigMap)

	if cm.Name == WatchConfigMapTrackerName {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	cmTracker, err := w.kubeData.GetConfigMap(WatchConfigMapTrackerName, WatchNamespace)

	now := time.Now()
	currentTime := fmt.Sprintf("%d", now.Unix())

	if err != nil {
		record := &apiv1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: WatchConfigMapTrackerName,
			},
			Data: map[string]string{
				cm.Name: currentTime,
			},
		}

		_, err := w.clientset.CoreV1().ConfigMaps(WatchNamespace).Create(record)

		if err != nil {
			w.f.Logger.Errorf("Error creating tracker configmap: %s", err)
		} else {
			w.f.Logger.Infof("Created tracker configmap and added configmap: \"%s\"", cm.Name)
		}

	} else {
		if cmTracker.Data == nil {
			cmTracker.Data = make(map[string]string)
		}

		cmTracker.Data[cm.Name] = currentTime
		_, err := w.clientset.CoreV1().ConfigMaps(WatchNamespace).Update(cmTracker)

		if err != nil {
			w.f.Logger.Errorf("Unable to update tracker configmap: %w", err)
		} else {
			w.f.Logger.Infof("Updated tracker for configmap: \"%s\"", cm.Name)
		}
	}
}

func (w Watch) DeleteConfigMapTracker(obj interface{}) {
	cm := obj.(*apiv1.ConfigMap)

	if cm.Name == WatchConfigMapTrackerName {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	cmTracker, err := w.kubeData.GetConfigMap(WatchConfigMapTrackerName, WatchNamespace)

	if err != nil {
		w.f.Logger.Info("Unable find tracker configmap")
		return
	}

	delete(cmTracker.Data, cm.Name)
	_, dErr := w.clientset.CoreV1().ConfigMaps(WatchNamespace).Update(cmTracker)

	if dErr != nil {
		w.f.Logger.Errorf("Unable to delete configmap from tracker; Error: %w", err)
	} else {
		w.f.Logger.Infof("Deleted configmap: \"%s\" from tracker", cm.Name)
	}

}
