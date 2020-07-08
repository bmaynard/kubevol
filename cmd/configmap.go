package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bmaynard/kubevol/pkg/core"
	"github.com/bmaynard/kubevol/pkg/watch"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

func NewConfigMapCommand(f *core.Factory, k *core.KubeData) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "configmap",
		Short: "Find all pods that have a specific ConfigMap attached",
		RunE: func(cmd *cobra.Command, args []string) error {
			pods := k.GetPods(namespace)
			fmt.Fprintf(cmd.OutOrStdout(), color.GreenString("There are %d pods in the cluster\n", len(pods.Items)))

			if objectName == "" {
				fmt.Fprintf(cmd.OutOrStdout(), color.GreenString("Searching for pods that have a ConfigMap attached\n\n"))
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), color.GreenString("Searching for pods that have \"%s\" ConfigMap attached\n\n", objectName))
			}

			ui := core.SetupTable(table.Row{"Namespace", "Pod Name", "ConfigMap Name", "Volume Name", "Out of Date"}, cmd.OutOrStdout())
			configmapTracker, err := k.GetConfigMap(watch.WatchConfigMapTrackerName, watch.WatchNamespace)

			if err != nil {
				f.Logger.Error(err)
			}

			for _, pod := range pods.Items {
				podName := pod.ObjectMeta.Name
				namespace := pod.ObjectMeta.Namespace
				_, err := k.GetPod(podName, namespace)

				if err != nil {
					f.Logger.Error(err)
				}

				podCreationTime := pod.ObjectMeta.CreationTimestamp.Time

				for _, volume := range pod.Spec.Volumes {
					if volume.ConfigMap != nil {
						if objectName == "" || (volume.ConfigMap != nil && volume.ConfigMap.LocalObjectReference.Name == objectName) {
							configMap, err := k.GetConfigMap(volume.ConfigMap.LocalObjectReference.Name, namespace)
							trackerName := watch.GetConfigMapKey(namespace, volume.ConfigMap.LocalObjectReference.Name)
							var outOfDate string

							if configmapTracker.CreationTimestamp.Time.Before(configMap.ObjectMeta.CreationTimestamp.Time) {
								outOfDate = color.GreenString("No")
							} else {
								outOfDate = color.YellowString("Unknown")
							}

							if err != nil || configMap.ObjectMeta.CreationTimestamp.Time.After(podCreationTime) {
								outOfDate = color.RedString("Yes")
							}

							if updatedTime, ok := configmapTracker.Data[trackerName]; ok {
								parsedTime, err := strconv.ParseInt(updatedTime, 10, 64)
								if err == nil && configMap.ObjectMeta.CreationTimestamp.Time.Before(time.Unix(parsedTime, 0)) {
									outOfDate = color.RedString("Yes")
								} else {
									outOfDate = color.RedString("No")
								}
							}

							ui.AppendRow([]table.Row{
								{color.BlueString(namespace), podName, volume.ConfigMap.LocalObjectReference.Name, volume.Name, outOfDate},
							})
						}
					}
				}
			}

			ui.Render()
			return nil
		},
	}

	return cmd
}
