package cmd

import (
	"github.com/spf13/cobra"
)

var structureScope string
var structureTag string

// structureCmd represents the tag command
var structureCmd = &cobra.Command{
	Use:   "structure",
	Short: "Generates a structure diagram",
	Long: `Generates a diagram that shows the hierarchical structure of items

[1] Main elements of interest

The 'eldest' element with the specified tag,

[2] Relevant associated elements

Those that are associated to one of the main elements of interest, where either:
- The parent is an ancestor of scope
- It is a root element.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Generate the diagram
		diagram, err = arch.StructureView(structureScope, structureTag)
	},
}

func init() {
	generateCmd.AddCommand(structureCmd)

	fs := structureCmd.Flags()
	fs.StringVarP(&structureScope, "scope", "s", "", "scope for the tag view")
	fs.StringVarP(&structureTag, "tag", "t", "", "tag to filter by")
}
