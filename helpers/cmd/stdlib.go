package cmd

import (
	"context"
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/helpers/genny/stdlib"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var stdlibCmd = &cobra.Command{
	Use:   "stdlib",
	Short: "generates the standard libary helpers",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		run := genny.WetRunner(ctx)
		if dryRun {
			run = genny.DryRunner(ctx)
		}
		g, err := stdlib.New(&stdlib.Options{})
		if err != nil {
			return err
		}
		// g.Transformer(gogen.FmtTransformer())
		run.With(g)

		run.WithRun(func(r *genny.Runner) error {

			return r.Exec(exec.Command("goimports", "-w", "./stdlib"))
		})
		return run.Run()
	},
}

func init() {
	rootCmd.AddCommand(stdlibCmd)
}
