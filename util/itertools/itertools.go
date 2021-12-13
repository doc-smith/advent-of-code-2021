package itertools

import "github.com/doc-smith/adventofcode2021/util/imath"

func ZipWhile(
	xs []int,
	ys []int,
	code func(x int, y int) bool,
) bool {
	for i := 0; i < imath.Min(len(xs), len(ys)); i++ {
		if !code(xs[i], ys[i]) {
			return false
		}
	}
	return true
}

func Zip(
	xs []int,
	ys []int,
	code func(x int, y int),
) {
	ZipWhile(xs, ys, func(x int, y int) bool {
		code(x, y)
		return true
	})
}
