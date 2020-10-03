package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ultimate-go-korean/kospell/kospell"
)

var rootCmd = &cobra.Command{
	Use:   "kospell",
	Short: "kospell is a CLI to https://speller.cs.pusan.ac.kr",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("no argument is given. Please provide an input")
		}
		out, err := kospell.Check(args[0])
		if err != nil {
			return err
		}
		kospell.PrintDiff(args[0], out)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
