package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"win-sprit-converter/cmd/sprit-win-convert/utils"
)

var (
	latestConversionArg string
	outputPath          string

	rootCmd = &cobra.Command{
		Use:     "sprit-win-convert spritmonitor-fuelings.csv --latest-conversion=date",
		Example: "sprit-win-convert fuelings.csv --latest-conversion=25/06/2022",
		Short:   "fuelings converter between spritmonitor and winbank",
		Long: `
CLI tool to convert spritmonitor fuelings to winbank format. This enables importing the fuelings from spritmonitor.de in winbank for specific vehicles.
The latest conversion date can be found in the latest converted csv file and is used to export the fuelings since that specified date. 		
`,
		Args: cobra.MinimumNArgs(1),
		RunE: spritWinConvert,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&latestConversionArg, "latest-conversion", "", "", "Date of latest conversion")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "", "converted.csv", "Path of converted fuelings")

	rootCmd.MarkPersistentFlagRequired("latest-conversion")
}

func Execute() error {
	return rootCmd.Execute()
}

func spritWinConvert(_ *cobra.Command, args []string) error {
	latestConversion, err := time.Parse("02/01/2006", latestConversionArg)
	if err != nil {
		return fmt.Errorf("parsing latest conversion date: %v", err)
	}

	allFuelingsPath, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("resolving absolute filepath of %s: %v", args[0], err)
	}

	allFuelings, err := utils.UnmarshalFuelings(allFuelingsPath)
	if err != nil {
		return fmt.Errorf("unmarshalling fuelings: %v", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file: %v", err)
	}

	for _, fillUp := range allFuelings {
		if fillUp.Date.After(latestConversion) {
			_, err := f.Write([]byte(fillUp.CsvMarshal()))
			log.Infof("Writing fueling from %v to %s", fillUp.Date.String(), outputPath)
			if err != nil {
				return fmt.Errorf("writing fillup to file: %v", err)
			}
		}
	}

	return nil
}
