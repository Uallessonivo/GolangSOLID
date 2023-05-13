package isp

type Trabalho interface {
	Entrar()
	ComecarATrabalhar()
	ContinuarATrabalhar()
	Sair()
}

type HumanoTasks interface {
	Trabalho
	PausaProCafe()
	Almocar()
}
type RoboTasks interface {
	Trabalho
	VerificarOleo()
	CarregarBateria()
}

type (
	Robo   struct{}
	Humano struct{}
)

// Robo
func (r *Robo) Entrar()              {}
func (r *Robo) ComecarATrabalhar()   {}
func (r *Robo) ContinuarATrabalhar() {}
func (r *Robo) Sair()                {}
func (r *Robo) CarregarBateria()     {}
func (r *Robo) VerificarOleo()       {}

// Humano
func (r *Humano) Entrar()              {}
func (r *Humano) ComecarATrabalhar()   {}
func (r *Humano) ContinuarATrabalhar() {}
func (r *Humano) Sair()                {}
func (r *Humano) PausaProCafe()        {}
func (r *Humano) Almocar()             {}
