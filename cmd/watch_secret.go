package cmd

import (
	"github.com/bmaynard/kubevol/pkg/core"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	w "github.com/bmaynard/kubevol/pkg/watch"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewWatchSecretCommand(f core.Factory) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "watch-secret",
		Short: "Watch for updates to Secrets",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientset, err := f.CoreClient()

			if err != nil {
				f.Logger.Fatal(err)
			}

			informer := cache.NewSharedIndexInformer(
				&cache.ListWatch{
					ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
						return clientset.CoreV1().Secrets("").List(options)
					},
					WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
						return clientset.CoreV1().Secrets("").Watch(options)
					},
				},
				&apiv1.Secret{},
				0, //Skip resync
				cache.Indexers{},
			)

			watcher := w.NewWatch(&f)

			informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
				UpdateFunc: watcher.UpateSecretTracker,
				DeleteFunc: watcher.DeleteSecretTracker,
			})

			stopCh := make(chan struct{})
			defer close(stopCh)
			informer.Run(stopCh)
			return nil
		},
	}

	return cmd
}
