package cmd

import (
	"github.com/spf13/cobra"
)

var (
	baseCfgPath   string
	targetCfgPath string

	rootCmd = &cobra.Command{
		Use:   "environator",
		Short: "a tool to do stuff with terraform files",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&baseCfgPath, "base", "", "path to base yaml file")
	rootCmd.PersistentFlags().StringVar(&targetCfgPath, "target", "", "path to target yaml file")
}

func Execute() error {
	return rootCmd.Execute()
}
