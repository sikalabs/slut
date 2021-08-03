package root

import (
	"github.com/spf13/cobra"
)

var RootCmdFlagJson bool

var RootCmd = &cobra.Command{
	Use:   "slut",
	Short: "SikaLabs Utils",
}

func init() {
	RootCmd.PersistentFlags().BoolVar(
		&RootCmdFlagJson,
		"json",
		false,
		"Formatu output to JSON",
	)
}
