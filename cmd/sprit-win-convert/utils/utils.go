package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"win-sprit-converter/types"

	"github.com/gocarina/gocsv"
)

func convertDotsDate(DotsDate string) (time.Time, error) {
	date, err := time.Parse("02.01.2006", DotsDate)
	if err != nil {
		return date, err
	}
	return date, nil
}

func GetCurrentDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

func UnmarshalFuelings(path string) ([]types.SpritFuelings, error) {
	in, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %v", err)
	}
	defer in.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';' // This is our separator now
		return r
	})

	var spritFuelings []types.SpritFuelings
	if err := gocsv.UnmarshalFile(in, &spritFuelings); err != nil {
		return nil, fmt.Errorf("gocsv unmarshalling file: %v", err)
	}

	for i, fillUp := range spritFuelings {
		date, err := convertDotsDate(fillUp.DateDEFormat)
		if err != nil {
			return nil, fmt.Errorf("converting date: %v", err)
		}
		spritFuelings[i].Date = date
		spritFuelings[i].FuelPrice = fillUp.Costs / fillUp.Liters
	}

	return spritFuelings, nil
}
