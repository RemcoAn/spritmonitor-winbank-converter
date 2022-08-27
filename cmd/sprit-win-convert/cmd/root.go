package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "sprit-win-convert cmd",
		Example: `
sprit-win-convert convert fuelings.csv --latest-conversion=25/06/2022
sprit-win-convert version
`,
		Short: "fuelings data converter between spritmonitor and winbank",
		Long: `
CLI tool to convert spritmonitor fuelings data to winbank format. This enables importing the fuelings from spritmonitor.de in winbank for specific vehicles.
The latest conversion date can be found in the latest converted csv file and is used to export the fuelings since that specified date. 		
`,
		Args: cobra.MinimumNArgs(0),
	}
)

func Execute() error {
	return rootCmd.Execute()
}
