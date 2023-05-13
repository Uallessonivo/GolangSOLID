package lsp

type RetanguloType interface {
	Area

	GetWidth() float64
	GetHeight() float64
	SetWidth(width float64)
	SetHeight(height float64)
}

type QuadradoType interface {
	Area

	SetSize(size float64)
	GetSize() float64
}

type Quadrado_ struct {
	size float64
}

type Retangulo_ struct {
	width, height float64
}

func (r *Retangulo_) GetArea() float64 {
	return r.width * r.width
}

func (r *Retangulo_) GetWidth() float64 {
	return r.width
}

func (r *Retangulo_) GetHeight() float64 {
	return r.height
}

func (r *Retangulo_) SetWidth(width float64) {
	r.width = width
}

func (r *Retangulo_) SetHeight(heigth float64) {
	r.height = heigth
}

func (r *Quadrado_) GetArea() float64 {
	return r.size * r.size
}

func (r *Quadrado_) GetSize() float64 {
	return r.size
}

func (r *Quadrado_) SetSize(size float64) {
	r.size = size
}
