package helper

import "math"

type Cube struct {
	Sisi float64
}

func NewCubePtr(sisi float64) *Cube {
	return &Cube{
		Sisi: sisi,
	}
}
func NewCube(sisi float64) Cube {
	return Cube{
		Sisi: sisi,
	}
}

func (cptr *Cube) VolumePtr() float64 {
	return math.Pow(cptr.Sisi, 3)
}

func (cptr *Cube) LuasPtr() float64 {
	return math.Pow(cptr.Sisi, 2) * 6
}

func (cptr *Cube) KelilingPtr() float64 {
	return cptr.Sisi * 12
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Sisi, 3)
}

func (c Cube) Luas() float64 {
	return math.Pow(c.Sisi, 2) * 6
}

func (c Cube) Keliling() float64 {
	return c.Sisi * 12
}
