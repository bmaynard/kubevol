package cmd

import (
	"fmt"

	"github.com/bmaynard/kubevol/pkg/core"
	"github.com/bmaynard/kubevol/pkg/utils"
	"github.com/bmaynard/kubevol/pkg/watch"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

func NewSecretCommand(f *core.Factory, k *core.KubeData) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "secret",
		Short: "Find all pods that have a specific Secret attached",
		RunE: func(cmd *cobra.Command, args []string) error {
			pods := k.GetPods(namespace)
			fmt.Fprintf(cmd.OutOrStdout(), color.GreenString("There are %d pods in the cluster\n", len(pods.Items)))

			if objectName == "" {
				fmt.Fprintf(cmd.OutOrStdout(), color.GreenString("Searching for pods that have a Secret attached\n\n"))
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), color.GreenString("Searching for pods that have \"%s\" Secret attached\n\n", objectName))
			}

			ui := core.SetupTable(table.Row{"Namespace", "Pod Name", "Secret Name", "Volume Name", "Out of Date"}, cmd.OutOrStdout())
			secretTracker, secretTrackerErr := k.GetConfigMap(watch.WatchSecretTrackerName, watch.WatchNamespace)

			if secretTrackerErr != nil {
				f.Logger.Warn(secretTrackerErr)
			}

			for _, pod := range pods.Items {
				podName := pod.ObjectMeta.Name
				namespace := pod.ObjectMeta.Namespace

				for _, volume := range pod.Spec.Volumes {
					if volume.Secret != nil {
						if objectName == "" || (volume.Secret != nil && volume.Secret.SecretName == objectName) {
							secret, err := k.GetSecret(volume.Secret.SecretName, namespace)
							trackerName := watch.GetConfigMapKey(namespace, volume.Secret.SecretName)

							o := utils.OutOfDateObject{
								ObjectTime:  secret.ObjectMeta.CreationTimestamp.Time,
								PodTime:     pod.ObjectMeta.CreationTimestamp.Time,
								ObjectErr:   err,
								Tracker:     secretTracker,
								TrackerErr:  secretTrackerErr,
								TrackerName: trackerName,
							}

							outOfDate := utils.GetOutOfDateText(o)

							ui.AppendRow([]table.Row{
								{color.BlueString(namespace), podName, volume.Secret.SecretName, volume.Name, outOfDate},
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
