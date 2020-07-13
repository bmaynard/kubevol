package watch

import (
	apiv1 "k8s.io/api/core/v1"
)

func (w Watch) AddSecretTracker(obj interface{}) {
	o := obj.(*apiv1.Secret)

	if o.Name == WatchSecretTrackerName {
		return
	}

	if o.CreationTimestamp.Time.Before(startTime) {
		return
	}

	w.UpdateTracker(WatchSecretTrackerName, o.Namespace, o.Name)
}

func (w Watch) UpateSecretTracker(old, new interface{}) {
	o := new.(*apiv1.Secret)

	if o.Name == WatchSecretTrackerName {
		return
	}

	w.UpdateTracker(WatchSecretTrackerName, o.Namespace, o.Name)
}

func (w Watch) DeleteSecretTracker(obj interface{}) {
	o := obj.(*apiv1.Secret)

	if o.Name == WatchSecretTrackerName {
		return
	}

	w.DeleteTracker(WatchSecretTrackerName, o.Namespace, o.Name)
}
