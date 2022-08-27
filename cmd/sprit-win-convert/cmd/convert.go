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

	convertCmd = &cobra.Command{
		Use:     "convert spritmonitor-fuelings.csv --latest-conversion=date",
		Example: "convert fuelings.csv --latest-conversion=25/06/2022",
		Long: `
if the --latest-conversion flag is not set, the complete fuelings file is converted.
`,
		Args:    cobra.MinimumNArgs(0),
		RunE:    spritWinConvert,
	}
)

func init() {
	convertCmd.PersistentFlags().StringVarP(&latestConversionArg, "latest-conversion", "", "", "Date of latest conversion")
	convertCmd.PersistentFlags().StringVarP(&outputPath, "output", "", "converted.csv", "Path of converted fuelings")

	convertCmd.MarkPersistentFlagRequired("latest-conversion")
	rootCmd.AddCommand(convertCmd)
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
