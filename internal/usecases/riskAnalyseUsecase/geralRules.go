package riskAnalyseUsecase

import "fmt"

func (r riskAnalyseUsecase) evaluteWeatherCondition(input RiskInput) {
	if input.WeatherForecast == "Chuva Forte com Ventos" || input.WeatherForecast == "Neve/Gelo" {
		r.riskScore++
		r.reason = append(r.reason, fmt.Sprintf("Condição climática: %s", input.WeatherForecast))
		r.recommendations = append(r.recommendations, "Reavaliar rota devido a condições climáticas")
	}
}

func (r riskAnalyseUsecase) evaluateTransportCompanyHistory(input RiskInput) {
	if input.TrafficAccidentYearHistory > 10 {
		r.riskScore = r.riskScore + 2
		r.reason = append(r.reason, fmt.Sprintf("Número de sinistros reportados nos últimos 12 meses: %d", input.TrafficAccidentYearHistory))
	} else if input.TrafficAccidentYearHistory > 5 {
		r.riskScore++
		r.reason = append(r.reason, fmt.Sprintf("Número de sinistros reportados nos últimos 12 meses: %d", input.TrafficAccidentYearHistory))
	}
}

func (r riskAnalyseUsecase) evaluateCargoValueAndInsurance(input RiskInput) {
	if input.CargoValue > 200000 && input.HasInsurance == false {
		r.riskScore = 3
		r.reason = append(r.reason, "Cargo valiosa e sem seguro")
		r.recommendations = append(r.recommendations, "Contratar escolta armada")
	}
}
