package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/freshly/environator/internal/split"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

const (
	fileDefaultsOnly   string = "defaults_only.yaml"
	fileOverwritesOnly string = "overwrites_only.yaml"
	fileOverwrites     string = "merge_overwrite.yaml"
)

var (
	splitCmd = &cobra.Command{
		Use:           "split",
		Short:         "split .tf file based on inputs",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          doSplit,
	}
)

func doSplit(cmd *cobra.Command, args []string) error {
	if defaultCfgPath == "" || overrideCfgPath == "" {
		return errors.New("must provide paths to target and base config files")
	}

	bb, err := os.ReadFile(defaultCfgPath)
	if err != nil {
		return fmt.Errorf("failed to read base config: %v", err)
	}

	tb, err := os.ReadFile(overrideCfgPath)
	if err != nil {
		return fmt.Errorf("failed to read prod config: %v", err)
	}

	var baseCfg map[string]string
	err = yaml.Unmarshal(bb, &baseCfg)
	if err != nil {
		return fmt.Errorf("error unmarshalling --base config: %v", err)
	}

	var targetCfg map[string]string
	err = yaml.Unmarshal(tb, &targetCfg)
	if err != nil {
		return fmt.Errorf("error unmarshalling --target config: %v", err)
	}

	fromOnly, toOnly, overrides, err := split.Do(targetCfg, baseCfg)
	if err != nil {
		return err
	}

	tOut, err := yaml.Marshal(toOnly)
	if err != nil {
		return err
	}

	ovOut, err := yaml.Marshal(overrides)
	if err != nil {
		return err
	}

	bOut, err := yaml.Marshal(fromOnly)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileOverwritesOnly, tOut, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileOverwrites, ovOut, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileDefaultsOnly, bOut, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
