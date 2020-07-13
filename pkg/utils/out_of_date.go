package utils

import (
	"strconv"
	"time"

	"github.com/fatih/color"
	v1 "k8s.io/api/core/v1"
)

type OutOfDateObject struct {
	ObjectTime  time.Time
	PodTime     time.Time
	ObjectErr   error
	Tracker     *v1.ConfigMap
	TrackerName string
	TrackerErr  error
}

func GetOutOfDateText(o OutOfDateObject) string {
	if o.ObjectErr != nil {
		return color.RedString("Object missing")
	}

	if o.ObjectTime.After(o.PodTime) {
		return color.RedString("Yes")
	}

	if o.TrackerErr != nil {
		return color.YellowString("Unknown")
	}

	if updatedTime, ok := o.Tracker.Data[o.TrackerName]; ok {

		parsedTime, err := strconv.ParseInt(updatedTime, 10, 64)
		if err == nil && o.ObjectTime.Before(time.Unix(parsedTime, 0)) {
			return color.RedString("Yes")
		} else {
			return color.GreenString("No")
		}
	}

	return color.YellowString("Unknown")
}
