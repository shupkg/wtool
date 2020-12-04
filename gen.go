package main

import (
	"log"
	"os"

	"github.com/shupkg/wproto/cmd"
	"github.com/shupkg/wproto/gen"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/compiler/protogen"
)

func init() {
	var c = &cobra.Command{Use: `gen [flags] files...`, Short: "生成代码"}
	var g = cmd.WithFlag(c.Flags())
	c.Run = func(_ *cobra.Command, args []string) {
		g.Files = append(g.Files, args...)
		if err := g.Run(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	addToRoot(c)
}

func runProtoc(){
	protogen.Options{}.Run(gen.Run)
}