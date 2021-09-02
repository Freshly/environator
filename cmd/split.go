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
	fileLeftOnly   string = "left_only.yaml"
	fileRightOnly  string = "right_only.yaml"
	fileOverwrites string = "overwrites.yaml"
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
	if baseCfgPath == "" || targetCfgPath == "" {
		return errors.New("must provide paths to target and base config files")
	}

	bb, err := os.ReadFile(baseCfgPath)
	if err != nil {
		return fmt.Errorf("failed to read base config: %v", err)
	}

	tb, err := os.ReadFile(targetCfgPath)
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

	err = os.WriteFile(fileRightOnly, tOut, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileOverwrites, ovOut, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileLeftOnly, bOut, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
