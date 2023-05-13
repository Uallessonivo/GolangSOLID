package lsp

type Quadrilatero interface {
	GetArea() float64
	GetWidth() float64
	GetHeight() float64
	SetWidth(width float64)
	SetHeight(height float64)
}

type Quadrado struct {
	width, height float64
}

type Retangulo struct {
	width, height float64
}

func (r *Retangulo) GetArea() float64 {
	return r.height * r.width
}

func (r *Retangulo) GetWidth() float64 {
	return r.width
}

func (r *Retangulo) GetHeight() float64 {
	return r.height
}

func (r *Retangulo) SetWidth(width float64) {
	r.width = width
}

func (r *Retangulo) SetHeight(heigth float64) {
	r.height = heigth
}

func (r *Quadrado) GetArea() float64 {
	// devemos usar o r.height ou r.width aqui ?
	return r.height * r.height
}

func (r *Quadrado) GetWidth() float64 {
	return r.width
}

func (r *Quadrado) GetHeight() float64 {
	return r.height
}

func (r *Quadrado) SetWidth(width float64) {
	// os dois lados tem de se manter o mesmo valor
	r.width = width
	r.height = width
}

func (r *Quadrado) SetHeight(heigth float64) {
	// ou devemos fazer r.SetWidth(heigth)?
	r.width = heigth
	r.height = heigth
}
