package cmd

import (
	"github.com/spf13/cobra"
)

var (
	defaultCfgPath  string
	overrideCfgPath string

	rootCmd = &cobra.Command{
		Use:   "environator",
		Short: "a tool to do stuff with terraform files",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&defaultCfgPath, "default", "", "path to default yaml file")
	rootCmd.PersistentFlags().StringVar(&overrideCfgPath, "override", "", "path to override yaml file")
}

func Execute() error {
	return rootCmd.Execute()
}
