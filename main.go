package main

import (
	"encoding/json"
	"fmt"

	r "px-technical-test/internal/usecases/riskAnalysisUsecase"
)

func main() {
	riskAnalysisUsecase := r.New()

	input := r.RiskInput{
		OperationId: "OP_20250910_001",

		Origin:          "São Paulo",
		Destiny:         "Rio de Janeiro",
		TotalDistance:   430,
		CargoType:       r.Chemicals,
		CargoValue:      75000.50,
		WeatherForecast: r.Stable,

		TrafficAccidentYearHistory: 2,
		HasInsurance:               true,
	}

	result := riskAnalysisUsecase.RiskAnalysis(input)

	formattedResult, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(formattedResult))
}
