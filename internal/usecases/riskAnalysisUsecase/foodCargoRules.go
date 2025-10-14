package riskAnalysisUsecase

func (r *riskAnalysisUsecase) foodEvaluateDistance(input RiskInput) {
	if input.TotalDistance > 300 {
		r.riskScore++
		r.reason = append(r.reason, "Carga de alimentos perecíveis com distância maior que 300km")
	}
}
