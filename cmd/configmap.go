package cmd

import (
	"fmt"

	"github.com/bmaynard/kubevol/pkg/core"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

func NewConfigMapCommand(k core.KubeData) *cobra.Command {
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

			for _, pod := range pods.Items {
				podName := pod.ObjectMeta.Name
				namespace := pod.ObjectMeta.Namespace
				_, err := k.GetPod(podName, namespace)

				if err != nil {
					panic(err.Error())
				}

				podCreationTime := pod.ObjectMeta.CreationTimestamp.Time

				for _, volume := range pod.Spec.Volumes {
					if volume.ConfigMap != nil {
						if objectName == "" || (volume.ConfigMap != nil && volume.ConfigMap.LocalObjectReference.Name == objectName) {
							configMap := k.GetConfigMap(volume.ConfigMap.LocalObjectReference.Name, namespace)
							outOfDate := color.YellowString("Unknown")

							if configMap.ObjectMeta.CreationTimestamp.Time.After(podCreationTime) {
								outOfDate = color.RedString("Yes")
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
