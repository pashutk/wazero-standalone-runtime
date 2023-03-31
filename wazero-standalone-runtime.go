package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	entrypointPtr := flag.String("entrypoint", "_start", "entrypoint function name")
	flag.Parse()

	filename := flag.Args()[0]

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("failed to open %v: %v", filename, err)
	}
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		log.Panicln(err)
	}

	mod, err := r.Instantiate(ctx, content)
	if err != nil {
		log.Panicln(err)
	}

	start := mod.ExportedFunction(*entrypointPtr)
	result, err := start.Call(ctx, 1, 1)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("%d\n", result[0])
}
