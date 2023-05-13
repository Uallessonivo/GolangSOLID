package ocp

import "reflect"

type (
	FuncionarioCLT struct{}
	Estagiario     struct{}
)

func FolhaDePag[TipoFuncionario FuncionarioCLT | Estagiario](funcionario TipoFuncionario) float64 {
	if reflect.TypeOf(FuncionarioCLT{}) == reflect.TypeOf(funcionario) {
		// Calcular salario + beneficios
		return 0.0
	} else {
		// Calcular bolsa auxilio do Estagiario
		return 0.0
	}
}
