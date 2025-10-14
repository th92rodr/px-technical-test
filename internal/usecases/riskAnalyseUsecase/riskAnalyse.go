package riskAnalyseUsecase

type RiskAnalyseUsecase interface {
	RiskAnalyse(RiskInput) RiskOutput
}

type riskAnalyseUsecase struct {
	riskScore       int
	reason          []string
	recommendations []string

	rules          []RuleFunc
	cargoTypeRules map[CargoType][]CargoTypeRuleFunc
}

func New() RiskAnalyseUsecase {
	r := &riskAnalyseUsecase{
		riskScore:       0,
		reason:          []string{},
		recommendations: []string{},
	}
	r.rules = []RuleFunc{
		r.evaluteWeatherCondition,
		r.evaluateTransportCompanyHistory,
		r.evaluateCargoValueAndInsurance,
	}
	r.cargoTypeRules = map[CargoType][]CargoTypeRuleFunc{
		Electronics: []CargoTypeRuleFunc{r.electronicsEvaluateDistance},
		Food:        []CargoTypeRuleFunc{r.foodEvaluateDistance},
	}
	return r
}

func (u *riskAnalyseUsecase) RiskAnalyse(input RiskInput) RiskOutput {
	for _, fn := range u.rules {
		fn(input)
	}

	cargoTypeRules := []CargoTypeRuleFunc{}

	switch input.CargoType {
	case Electronics:
		cargoTypeRules = u.cargoTypeRules[Electronics]
	case Food:
		cargoTypeRules = u.cargoTypeRules[Food]
	case Chemicals:
		cargoTypeRules = u.cargoTypeRules[Chemicals]
	case DryCargo:
		cargoTypeRules = u.cargoTypeRules[DryCargo]
	case Auto:
		cargoTypeRules = u.cargoTypeRules[Auto]
	}

	for _, fn := range cargoTypeRules {
		fn(input)
	}

	var finalScore Score
	switch u.riskScore {
	case 0:
		finalScore = LowScore
	case 1:
		finalScore = MediumScore
	case 2:
		finalScore = HighScore
	case 3:
		finalScore = CriticalScore
	}

	return RiskOutput{
		RiskScore:       finalScore,
		Reasons:         u.reason,
		Recommendations: u.recommendations,
	}
}
