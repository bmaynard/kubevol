package cmd

import (
	"fmt"
	"os"

	"github.com/bmaynard/kubevol/pkg/core"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

func NewConfigMapCommand(k core.KubeData) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "configmap",
		Short: "Find all pods that have a specific ConfigMap attached",
		Run: func(cmd *cobra.Command, args []string) {
			pods := k.GetPods()
			fmt.Printf(color.GreenString("There are %d pods in the cluster\n", len(pods.Items)))

			if name == "" {
				fmt.Printf(color.GreenString("Searching for pods that have a ConfigMap attached\n\n"))
			} else {
				fmt.Printf(color.GreenString("Searching for pods that have \"%s\" ConfigMap attached\n\n", name))
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Namespace", "Pod Name", "ConfigMap Name", "Volume Name", "Out of Date"})

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
						if name == "" || (volume.ConfigMap != nil && volume.ConfigMap.LocalObjectReference.Name == name) {
							configMap := k.GetConfigMap(volume.ConfigMap.LocalObjectReference.Name, namespace)
							outOfDate := color.YellowString("Unknown")

							if configMap.ObjectMeta.CreationTimestamp.Time.After(podCreationTime) {
								outOfDate = color.RedString("Yes")
							}

							t.AppendRows([]table.Row{
								{color.BlueString(namespace), podName, volume.ConfigMap.LocalObjectReference.Name, volume.Name, outOfDate},
							})
						}
					}
				}
			}

			t.Render()
		},
	}

	return cmd
}
