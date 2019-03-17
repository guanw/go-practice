package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "something",
	Short: "something",
	Long:  `something more`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("requires at least three arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("the command is run")
		fmt.Println(args)
		fmt.Println(args[0])
		fmt.Println(args[1])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
