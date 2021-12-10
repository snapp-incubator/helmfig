package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "helmfig",
	Short: "Quick tool to generate configmap and values.yaml for helm charts",
	Long:  `Helmfig enables you to create configmap and appropriate values.yaml file from your example config of your project to helmify it quickly`,
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
