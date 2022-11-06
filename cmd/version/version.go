package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// StartCmd version
var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "get version info",
		Example: "sme-delay-service version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println("v2022.11.06")
	return nil
}
