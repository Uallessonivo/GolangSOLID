package ocp

type (
	FuncPj         struct{}
	FuncCLT        struct{}
	FuncEstagiario struct{}
)

type Funcionario interface {
	CalcularSalario() float64
}

func CalcularSalarioFunc() {
	funcionario := &FuncCLT{}
	funcionario.CalcularSalario()

	estag := &FuncEstagiario{}
	estag.CalcularSalario()
}

func (f *FuncPj) CalcularSalario() float64 {
	// Calcular salario
	return 0.0
}

func (f *FuncCLT) CalcularSalario() float64 {
	// Calcular salario
	return 0.0
}

func (f *FuncEstagiario) CalcularSalario() float64 {
	// Calcular salario
	return 0.0
}

func FolhaDePagamento(funcionario Funcionario) float64 {
	return funcionario.CalcularSalario()
}
