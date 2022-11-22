package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "",
	Long: `A command line utility to interact with Azure Storage`,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return errors.New("requires at least one arg")
	// 	}
	// 	return fmt.Errorf("invalid color specified: %s", args[0])
	// },
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var ConfigFile *string

func init() {
	rootCmd.PersistentFlags().StringVarP(ConfigFile, "config", "c", "", "config file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
