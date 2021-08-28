package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/zlyzol/tendermint-0.32.3/version"
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}
