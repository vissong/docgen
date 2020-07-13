//go:generate go run generate-asset.go

package main

import (
	"github.com/vissong/docgen/cmd"
)

func main() {
	cmd.Execute()
}
