package cmd

import (
	"encoding/json"
	. "github.com/phodal/coca/core/ast/full"
	. "github.com/phodal/coca/core/ast/identifier"
	"github.com/phodal/coca/core/infrastructure/coca_file"
	"github.com/spf13/cobra"
)

type AnalysisCmdConfig struct {
	Path        string
	ForceUpdate bool
}

var (
	analysisCmdConfig AnalysisCmdConfig
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := analysisCmdConfig.Path

		identifierApp := NewJavaIdentifierApp()
		iNodes := identifierApp.AnalysisPath(importPath)

		identModel, _ := json.MarshalIndent(iNodes, "", "\t")
		coca_file.WriteToCocaFile("identify.json", string(identModel))

		var classes []string = nil

		for _, node := range iNodes {
			classes = append(classes, node.Package+"."+node.ClassName)
		}

		callApp := NewJavaFullApp()

		callNodes := callApp.AnalysisPath(importPath, classes, iNodes)
		cModel, _ := json.MarshalIndent(callNodes, "", "\t")
		coca_file.WriteToCocaFile("deps.json", string(cModel))
	},
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Path, "path", "p", ".", "example -p core/main")
	analysisCmd.PersistentFlags().BoolVarP(&analysisCmdConfig.ForceUpdate, "force", "f", false, "force update -f")
}
