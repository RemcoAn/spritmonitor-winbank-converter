package types

import (
	"fmt"
	"time"
)

type SpritFuelings struct {
	Date         time.Time
	Liters       float64 `csv:"Spritmenge"`
	FuelPrice    float64
	Tripmeter    float64 `csv:"Teil-Km"`
	DateDEFormat string  `csv:"Datum"`
	Costs        float64 `csv:"Kosten"`
	// Odometer  float64   `csv:"Km-Stand"`
	// Currency string `csv:"Währung"`
	// `csv:"Tankart"`
	// `csv:"Reifen"`
	// `csv:"Strecken"`
	// `csv:"Fahrweise"`
	// `csv:"Kraftstoff"`
	// `csv:"Bemerkung"`
	// Mileage float64 `csv:"Verbrauch"`
	// `csv:"BC-Verbrauch"`
	// `csv:"BC-Spritmenge"`
	// `csv:"BC-Geschwindigkeit"`
	// `csv:"Tankstelle"`
	// `csv:"Land"`
	// `csv:"Großraum"`
	// `csv:"Ort"`
}

func (f *SpritFuelings) CsvMarshal() string {
	date := f.Date.Format("02/01/2006")
	liters := f.Liters
	fuelPrice := f.FuelPrice
	tripMeter := f.Tripmeter
	str := fmt.Sprintf("%s,%.2f,%f,%.0f\n", date, liters, fuelPrice, tripMeter)
	return str
}
