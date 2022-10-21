package transform

import (
	"fmt"
)

func ExampleMatrix() {
	// https://ja.wikipedia.org/wiki/%E4%BD%99%E5%9B%A0%E5%AD%90%E8%A1%8C%E5%88%97
	mat := Matrix{
		-3, 2, -5,
		-1, 0, -2,
		3, -4, 1,
	}
	fmt.Println(mat.Adjugate())

	// https://manabitimes.jp/math/1153
	mat = Matrix{
		1, 1, -1,
		-2, 0, 1,
		0, 2, 1,
	}
	fmt.Println(mat.Inverse())

	// https://risalc.info/src/determinant-three-by-three.html
	fmt.Println(Matrix{0, 1, 2, 3, 4, 5, 6, 7, 8}.Determinant())
	fmt.Println(Matrix{1, 0, 4, 0, 2, 0, 4, 0, 3}.Determinant())

	// Output:
	// [-8 18 -4 -5 12 -1 4 -6 2]
	// [-0.5 -0.75 0.25 0.5 0.25 0.25 -1 -0.5 0.5]
	// 0
	// -26
}

func ExampleNewDeformation() {
	mat := NewDeformation(
		Point{1, 2},
		Point{133, 4},
		Point{5, 106},
		Point{127, 108},
	)

	out := func(x, y float64) {
		fmt.Printf("%0.3f %0.3f\n", x, y)
	}

	out(mat.Transform(0, 0))
	out(mat.Transform(1, 0))
	out(mat.Transform(0, 1))
	out(mat.Transform(1, 1))

	// Output:
	// 1.000 2.000
	// 133.000 4.000
	// 5.000 106.000
	// 127.000 108.000
}
