package main

import (
	"log"
	"os"

	"github.com/shupkg/wproto/gen"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/compiler/protogen"
)

var root = &cobra.Command{Use: "wtool [command]"}

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	stat, err := os.Stdin.Stat()
	if len(os.Args) == 1 && err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		protogen.Options{}.Run(gen.Run)
	} else {
		if err := root.Execute(); err != nil {
			log.Fatal(err)
		}
	}
}

func addToRoot(c *cobra.Command) {
	c.Flags().BoolP("help", "h", false, "显示帮助信息")
	root.AddCommand(c)
}
