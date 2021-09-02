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
	fileDefaultsOnly  string = "defaults_only.yaml"
	fileOverridesOnly string = "overrides_only.yaml"
	fileOverrides     string = "overrides.yaml"
)

var (
	splitCmd = &cobra.Command{
		Use:           "split",
		Short:         "generate two diff and an overwrite YAML file based on a default and override input YAML file",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          doSplit,
	}
)

func doSplit(cmd *cobra.Command, args []string) error {
	if defaultCfgPath == "" {
		return errors.New("--default filepath flag not set")
	}

	if overrideCfgPath == "" {
		return errors.New("--override filepath flag not set")
	}

	defaultBytes, err := os.ReadFile(defaultCfgPath)
	if err != nil {
		return fmt.Errorf("failed to read default file: %v", err)
	}

	overrideBytes, err := os.ReadFile(overrideCfgPath)
	if err != nil {
		return fmt.Errorf("failed to read override file: %v", err)
	}

	var defaultCfg map[string]string
	err = yaml.Unmarshal(defaultBytes, &defaultCfg)
	if err != nil {
		return fmt.Errorf("error unmarshalling defaults; is it a flat YAML file?: %v", err)
	}

	var overrideCfg map[string]string
	err = yaml.Unmarshal(overrideBytes, &overrideCfg)
	if err != nil {
		return fmt.Errorf("error unmarshalling overrides; is it a flat YAML file?: %v", err)
	}

	overridesOnly, defaultsOnly, overrides, err := split.Do(defaultCfg, overrideCfg)
	if err != nil {
		return err
	}

	tOut, err := yaml.Marshal(overridesOnly)
	if err != nil {
		return err
	}

	bOut, err := yaml.Marshal(defaultsOnly)
	if err != nil {
		return err
	}

	ovOut, err := yaml.Marshal(overrides)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileOverridesOnly, tOut, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileDefaultsOnly, bOut, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileOverrides, ovOut, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
