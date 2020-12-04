package main

import (
	"log"
	"os"

	"github.com/shupkg/wbin/cmd"
	"github.com/spf13/cobra"
)

func init() {
	var c = &cobra.Command{
		Use:     "embed [flags] files...",
		Aliases: []string{"bin"},
		Short:   "将文件嵌入代码中",
		Long: `将文件嵌入代码中
如果传入的是目录， 路径为与该目录的相对路径
如果传入的是文件， 路径为其文件名
所有路径以"/"开头
`,
	}

	var p = cmd.WithFlag(c.Flags())
	c.Run = func(_ *cobra.Command, args []string) {
		p.Files = append(p.Files, args...)
		if err := p.Run(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	addToRoot(c)
}
