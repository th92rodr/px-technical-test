package main

import (
	"encoding/json"
	"fmt"

	r "px-technical-test/internal/usecases/riskAnalyseUsecase"
)

func main() {
	riskAnalyseUsecase := r.New()

	input := r.RiskInput{
		OperationId: "OP_20250910_001",

		Origin:          "São Paulo",
		Destiny:         "Rio de Janeiro",
		TotalDistance:   430,
		CargoType:       r.Chemicals,
		CargoValue:      75000.50,
		WeatherForecast: "Estável",

		TrafficAccidentYearHistory: 2,
		HasInsurance:               true,
	}

	result := riskAnalyseUsecase.RiskAnalyse(input)

	formattedResult, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(formattedResult))
}
