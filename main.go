package main

import (
	"fmt"
	"github.com/be-edu/weather-1/weather/v2"
	"github.com/be-edu/weather-1/weather/v2/location"
)

func main() {
	makePrediction(51.509865, -0.118092, "London")
}

func makePrediction(lat float64, long float64, locationName string) error {
	gpsCoords := location.GPSCoords([2]float64{lat, long})
	pred, err := weather.Predict(gpsCoords)

	if err != nil {
		return err
	}

	fmt.Printf("The weather prediction for %s is: %v\n", locationName, pred.ToString())

	return nil
}
