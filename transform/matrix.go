package transform

type Matrix [9]float64

func NewIdentityMatrix() Matrix {
	return Matrix{1, 0, 0, 0, 1, 0, 0, 0, 1}
}

// 行列式
func (m Matrix) Determinant() float64 {
	var f float64
	for i := 0; i < 3; i++ {
		f += m[i] * m[(i+1)%3+3] * m[(i+2)%3+6]
		f -= m[i] * m[(i+2)%3+3] * m[(i+1)%3+6]
	}
	return f
}

// 余因子行列
func (m Matrix) Adjugate() Matrix {
	p00 := m[0]
	p01 := m[1]
	p02 := m[2]
	p10 := m[3]
	p11 := m[4]
	p12 := m[5]
	p20 := m[6]
	p21 := m[7]
	p22 := m[8]
	return Matrix{
		p11*p22 - p12*p21,
		p02*p21 - p01*p22,
		p01*p12 - p02*p11,
		p12*p20 - p10*p22,
		p00*p22 - p02*p20,
		p02*p10 - p00*p12,
		p10*p21 - p11*p20,
		p01*p20 - p00*p21,
		p00*p11 - p01*p10,
	}
}

func (m Matrix) MulScalar(s float64) Matrix {
	n := Matrix{}
	for i, v := range m {
		n[i] = v * s
	}
	return n
}

// 逆行列
func (m Matrix) Inverse() Matrix {
	return m.Adjugate().MulScalar(1 / m.Determinant())
}

func (m Matrix) Transform(x, y float64) (float64, float64) {
	v := [3]float64{x, y, 1}
	w := [3]float64{}
	for i := range v {
		for j, t := range v {
			w[i] += m[i*3+j] * t
		}
	}
	return w[0] / w[2], w[1] / w[2]
}

func (m *Matrix) MulMatrix(n Matrix) Matrix {
	o := Matrix{}
	for i := range o {
		r := i / 3
		c := i % 3
		for j := 0; j < 3; j++ {
			o[i] += m[r*3+j] * n[j*3+c]
		}
	}
	return o
}

func NewDeformation(s, x, y, e Point) Matrix {
	sx := x.Sub(s)
	sy := y.Sub(s)
	ex := x.Sub(e)
	ey := y.Sub(e)
	es := s.Sub(e)

	dx := ex.Dot(es) / ex.Dot(ey)
	dy := es.Dot(ey) / ex.Dot(ey)

	m := NewIdentityMatrix()
	m = m.MulMatrix(Matrix{
		1, 0, s.X,
		0, 1, s.Y,
		0, 0, 1,
	})
	m = m.MulMatrix(Matrix{
		sx.X, sy.X, 0,
		sx.Y, sy.Y, 0,
		0, 0, 1,
	})
	m = m.MulMatrix(Matrix{
		dy, 0, 0,
		0, dx, 0,
		dy - 1, dx - 1, 1,
	})
	return m
}
