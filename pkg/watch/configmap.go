package watch

import (
	apiv1 "k8s.io/api/core/v1"
)

func (w Watch) AddConfigMapTracker(new interface{}) {
	o := new.(*apiv1.ConfigMap)

	if o.Name == WatchConfigMapTrackerName || o.Name == WatchSecretTrackerName {
		return
	}

	if o.CreationTimestamp.Time.Before(startTime) {
		return
	}

	w.UpdateTracker(WatchConfigMapTrackerName, o.Namespace, o.Name)
}

func (w Watch) UpateConfigMapTracker(old, new interface{}) {
	o := new.(*apiv1.ConfigMap)

	if o.Name == WatchConfigMapTrackerName || o.Name == WatchSecretTrackerName {
		return
	}

	w.UpdateTracker(WatchConfigMapTrackerName, o.Namespace, o.Name)
}

func (w Watch) DeleteConfigMapTracker(obj interface{}) {
	o := obj.(*apiv1.ConfigMap)

	if o.Name == WatchConfigMapTrackerName || o.Name == WatchSecretTrackerName {
		return
	}

	w.DeleteTracker(WatchConfigMapTrackerName, o.Namespace, o.Name)
}
