package cmd

import (
	"fmt"
	"os"

	"github.com/bmaynard/kubevol/pkg/core"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var name string

var rootCmd = &cobra.Command{
	Use:   "kubevol",
	Short: "Get information on your pods volumes",
	Long:  `Find all the pods that have volumes attached and which volumes are attached with the ability to filter by specific type and name.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubevol.yaml)")
	rootCmd.PersistentFlags().StringVar(&name, "name", "", "Name of the object you wish to filter by")
	initConfig()

	factory := core.NewDepsFactory()
	coreClient, err := factory.CoreClient(viper.GetString("kubeconfig"))

	if err != nil {
		panic(err.Error())
	}

	kubeData := core.NewKubeData("", coreClient)

	rootCmd.AddCommand(NewConfigMapCommand(*kubeData))
	rootCmd.AddCommand(NewSecretCommand(*kubeData))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kubevol" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kubevol")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
