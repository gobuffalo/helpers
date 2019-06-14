package main

import (
	"context"
	"flag"
	"log"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/helpers/genny/docs"
)

func main() {
	var dryRun bool

	flag.BoolVar(&dryRun, "dry-run", false, "dry run")
	flag.Parse()

	ctx := context.Background()
	run := genny.WetRunner(ctx)
	if dryRun {
		run = genny.DryRunner(ctx)
	}
	g, err := docs.New(&docs.Options{})
	if err != nil {
		log.Fatal(err)
	}
	run.With(g)
	if err := run.Run(); err != nil {
		log.Fatal(err)
	}

}
