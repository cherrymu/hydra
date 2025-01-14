// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	"github.com/ory/x/clidoc"

	"github.com/ory/hydra/cmd"
)

func main() {
	if err := clidoc.Generate(cmd.NewRootCmd(nil, nil, nil), os.Args[1:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
		os.Exit(1)
	}
	fmt.Println("All files have been generated and updated.")
}
