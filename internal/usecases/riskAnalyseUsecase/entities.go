package riskAnalyseUsecase

type CargoType string

const (
	Electronics = "Eletrônicos Sensíveis"
	Food        = "Alimentos Perecíveis"
	Chemicals   = "Produtos Químicos Perigosos"
	DryCargo    = "Carga Geral Seca"
	Auto        = "Automotiva"
)

type WeatherCondition string

const (
	Stable    = "Estável"
	LightRain = "Chuva Moderada"
	HeavyRain = "Chuva Forte com Ventos"
	Snow      = "Neve/Gelo"
)

type Score string

const (
	LowScore      = "Baixo"
	MediumScore   = "Médio"
	HighScore     = "Alto"
	CriticalScore = "Crítico"
)

type RiskInput struct {
	OperationId string

	// Route
	Origin          string
	Destiny         string
	TotalDistance   int
	CargoType       CargoType
	CargoValue      float64
	WeatherForecast WeatherCondition

	// Transport Company
	TrafficAccidentYearHistory int
	HasInsurance               bool
}

type RiskOutput struct {
	RiskScore       Score
	Reasons         []string
	Recommendations []string
}

type RuleFunc func(RiskInput)
type CargoTypeRuleFunc func(RiskInput)
