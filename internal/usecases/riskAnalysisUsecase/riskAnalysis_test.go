package riskAnalysisUsecase

import (
	"fmt"
	"testing"
)

func TestRiskAnalysis(t *testing.T) {
	tests := []struct {
		name           string
		input          RiskInput
		expectedOutput RiskOutput
	}{
		{
			name: "Risco Alto com Fatores Múltiplos",
			input: RiskInput{
				OperationId:                "OP_20250910_002",
				Origin:                     "Porto Alegre",
				Destiny:                    "Curitiba",
				TotalDistance:              700,
				CargoType:                  Electronics,
				CargoValue:                 65000.00,
				WeatherForecast:            HeavyRain,
				TrafficAccidentYearHistory: 7,
				HasInsurance:               false,
			},
			expectedOutput: RiskOutput{
				RiskScore: HighScore,
				Reasons: []string{
					fmt.Sprintf("Condição climática: %s", HeavyRain),
					"Número de sinistros reportados nos últimos 12 meses: 7"},
				Recommendations: []string{"Reavaliar rota devido a condições climáticas"},
			},
		},
		{
			name: "Risco Crítico por Falta de Seguro",
			input: RiskInput{
				OperationId:                "OP_20250910_003",
				Origin:                     "Belo Horizonte",
				Destiny:                    "Salvador",
				TotalDistance:              1500,
				CargoType:                  DryCargo,
				CargoValue:                 250000.00,
				WeatherForecast:            Stable,
				TrafficAccidentYearHistory: 3,
				HasInsurance:               false,
			},
			expectedOutput: RiskOutput{
				RiskScore:       CriticalScore,
				Reasons:         []string{fmt.Sprintf("Carga com valor de %f e sem seguro", 250000.00)},
				Recommendations: []string{"Contratar escolta armada"},
			},
		},
		{
			name: "Risco Baixo",
			input: RiskInput{
				OperationId:                "OP_20250910_004",
				Origin:                     "Joinville",
				Destiny:                    "Blumenau",
				TotalDistance:              50,
				CargoType:                  Food,
				CargoValue:                 25000.00,
				WeatherForecast:            Stable,
				TrafficAccidentYearHistory: 0,
				HasInsurance:               true,
			},
			expectedOutput: RiskOutput{
				RiskScore:       LowScore,
				Reasons:         []string{},
				Recommendations: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()

			output := r.RiskAnalysis(tt.input)

			if output.RiskScore != tt.expectedOutput.RiskScore {
				t.Errorf("expected riskScore=%s, got %s", tt.expectedOutput.RiskScore, output.RiskScore)
			}
			if len(output.Reasons) != len(tt.expectedOutput.Reasons) {
				t.Errorf("expected reason=%v, got %v", tt.expectedOutput.Reasons, output.Reasons)
			}
			if len(output.Recommendations) != len(tt.expectedOutput.Recommendations) {
				t.Errorf("expected recommendations=%v, got %v", tt.expectedOutput.Recommendations, output.Recommendations)
			}
		})
	}
}
