package main

import (
	"log"
	"os"

	"github.com/shupkg/wbin/embed"
	"github.com/spf13/cobra"
)

func init() {
	var packer = embed.New()
	var cmd = &cobra.Command{
		Use:     "embed [flags] files...",
		Aliases: []string{"bin"},
		Short:   "将文件嵌入代码中",
		Long: `将文件嵌入代码中
如果传入的是目录， 路径为与该目录的相对路径
如果传入的是文件， 路径为其文件名
所有路径以"/"开头
`,
		Run: func(cmd *cobra.Command, args []string) {
			packer.Files = append(packer.Files, args...)
			if err := packer.Run(); err != nil {
				log.Println(err)
				os.Exit(1)
			}
		},
	}
	var embedFlags = cmd.Flags()
	embedFlags.SortFlags = false
	embedFlags.StringVarP(&packer.Import, "import", "i", packer.Import, "导入「.Fs, .File」的包名")
	embedFlags.StringVarP(&packer.Out, "out", "o", packer.Out, "输出文件")
	embedFlags.StringVar(&packer.Var, "var", packer.Var, "变量命令前缀")
	embedFlags.StringSliceVar(&packer.Filters, "exclude", packer.Filters, "过滤规则，正则")
	embedFlags.StringSliceVar(&packer.Files, "file", packer.Files, "要嵌入的文件")
	embedFlags.BoolVarP(&packer.Force, "force", "f", packer.Force, "覆盖输出文件（如果已存在）")
	embedFlags.BoolVarP(&packer.Verbose, "verbose", "v", packer.Verbose, "输出过程信息")
	embedFlags.BoolP("help", "h", false, "该命令的帮助信息")
	root.AddCommand(cmd)
}
