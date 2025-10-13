package riskAnalyseUsecase

func (r riskAnalyseUsecase) foodEvaluateDistance(input RiskInput) {
	if input.TotalDistance > 300 {
		r.riskScore++
		r.reason = append(r.reason, "Carga de alimentos pereciveis com distância de mais de 300km")
	}
}
