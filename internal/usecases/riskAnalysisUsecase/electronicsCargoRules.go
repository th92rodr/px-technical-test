package riskAnalysisUsecase

func (r *riskAnalysisUsecase) electronicsEvaluateDistance(input RiskInput) {
	if input.TotalDistance > 50000 {
		r.riskScore++
		r.reason = append(r.reason, "Carga de eletrônicos sensíveis com distância maior que 50000km")
	}
}
