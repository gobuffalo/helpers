package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/helpers/genny/docs"
	"github.com/spf13/cobra"
)

var dryRun bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "helpers",
	Short: "generates the README",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		run := genny.WetRunner(ctx)
		if dryRun {
			run = genny.DryRunner(ctx)
		}
		g, err := docs.New(&docs.Options{})
		if err != nil {
			return err
		}
		run.With(g)
		return run.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "dry run")
}
