package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(put)
}

var put = &cobra.Command{
	Use:   "put",
	Short: "Upload a blob",
	Long:  `All software has versions. This is Hugo's`,
	Args: func(command *cobra.Command, args []string) error {
		if ConfigFile == nil {
			// configFile, err := os.Open(*configPath)
			// if err != nil {
			// 	log.Fatalln(err)
			// }
			fmt.Printf("Config file not found")
		}
		if len(args) < 3 {
			return errors.New("requires at least three arg")
		}
		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
