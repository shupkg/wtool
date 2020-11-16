package main

import (
	"log"
	"os"

	"github.com/shupkg/wproto/cmd"
	"github.com/spf13/cobra"
)

func init() {
	var g = cmd.New()
	var cGen = &cobra.Command{
		Use:   `gen [flags] files...`,
		Short: "生成代码",
		Run: func(cmd *cobra.Command, args []string) {
			g.Files = append(g.Files, args...)
			if err := g.Exec(); err != nil {
				log.Println(err)
				os.Exit(1)
			}
		},
	}

	flag := cGen.Flags()
	flag.SortFlags = false
	flag.StringVarP(&g.Out, "out", "o", g.Out, "输出目录")
	flag.StringVarP(&g.Root, "root", "r", g.Root, "根包名")
	flag.StringVar(&g.Protoc, "protoc", g.Protoc, "protoc 路径")
	flag.StringSliceVar(&g.Files, "files", g.Files, "proto文件(夹)")
	flag.StringSliceVarP(&g.Language, "language", "l", g.Language, "生成语言, 格式为语言=输出路径, 如: --language go=.")
	flag.BoolP("help", "h", false, "显示帮助信息")
	root.AddCommand(cGen)
}

