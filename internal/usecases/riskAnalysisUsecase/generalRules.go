package riskAnalysisUsecase

import "fmt"

func (r *riskAnalysisUsecase) evaluteWeatherCondition(input RiskInput) {
	if input.WeatherForecast == HeavyRain || input.WeatherForecast == Snow {
		r.riskScore++
		r.reason = append(r.reason, fmt.Sprintf("Condição climática: %s", input.WeatherForecast))
		r.recommendations = append(r.recommendations, "Reavaliar rota devido a condições climáticas")
	}
}

func (r *riskAnalysisUsecase) evaluateTransportCompanyHistory(input RiskInput) {
	if input.TrafficAccidentYearHistory > 10 {
		r.riskScore = r.riskScore + 2
		r.reason = append(r.reason, fmt.Sprintf("Número de sinistros reportados nos últimos 12 meses: %d", input.TrafficAccidentYearHistory))
	} else if input.TrafficAccidentYearHistory > 5 {
		r.riskScore++
		r.reason = append(r.reason, fmt.Sprintf("Número de sinistros reportados nos últimos 12 meses: %d", input.TrafficAccidentYearHistory))
	}
}

func (r *riskAnalysisUsecase) evaluateCargoValueAndInsurance(input RiskInput) {
	if input.CargoValue > 200000 && input.HasInsurance == false {
		r.riskScore = 3
		r.reason = append(r.reason, fmt.Sprintf("Carga com valor de %f e sem seguro", input.CargoValue))
		r.recommendations = append(r.recommendations, "Contratar escolta armada")
	}
}
