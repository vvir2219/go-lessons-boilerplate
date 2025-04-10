package main

import (
	"cmp"
	"curs1_boilerplate/cmd/generics/optional"
	"fmt"
)

func ForEach[T any](v []T, do func (T)) {
	for _, e := range v {
		do(e)
	}
}

func Map[T, Q any](v []T, do func(T) Q) []Q {
	res := make([]Q, len(v))
	for i, e := range(v) {
		res[i] = do(e)
	}
	return res
}

type Mapper[T, Q any] []T

// func (a Arr[T]) ForEach(do func (T)) {
// 	for _, e := range a {
// 		do(e)
// 	}
// }

func (a Mapper[T, Q]) Map(do func(T)Q) []Q {
	res := make([]Q, len(a))
	for i, e := range(a) {
		res[i] = do(e)
	}
	return res
}

type MyNumber int

type Number interface {
	~int | ~float32 | ~float64
}

func Sum[T Number](v []T) T {
	var sum T
	for _, e := range v {
		sum += e
	}
	return sum
}

func Max[T cmp.Ordered](v []T) T {
	var max T
	for _, e := range v {
		if e > max {
			max = e
		}
	}
	return max
}

type formatter interface {
	Format() string
}

func FormatAll[T formatter] (v []T) {
	for _, e := range v {
		fmt.Println(e.Format())
	}
}

type Point struct {
	x, y int
}
func (p Point) Format() string {
	return fmt.Sprintf("[%d, %d]", p.x, p.y)
}

type Valoare int
func (v Valoare) Format() string {
	return fmt.Sprintf("Am %d valoare", v)
}

func main() {
	l := make([]formatter, 0)
	l = append(l, Point{1, 2}, Point{2, 4}, Valoare(7))

	FormatAll(l)
	//
	// mapped := Map(v, func(e int) string {
	// 	return fmt.Sprintf("valoare %d", e)
	// })
	// fmt.Println(mapped)

	o := optional.Optional[int]{}
	fmt.Println(o.IsEmpty())

	o2 := optional.Some("sunathose")
	fmt.Println(o2.IsEmpty())
	fmt.Println(o2.Get())

	v := make([]MyNumber, 0)
	v = append(v, 1, 2, 3)
	fmt.Println(Max(v))

	mapped := Mapper[MyNumber, string](v).Map(func(i MyNumber) string {
		return fmt.Sprintf("valoare %d", i)
	})
	fmt.Println(mapped)

	fmt.Println(Sum(v))

	// ForEach(v, func(x int) {
	// 	fmt.Println(x)
	// })
	//
	// ss := make([]string, 0)
	// ss = append(ss, "ana", "are", "mere")
	//
	// ForEach(ss, func(x string) {
	// 	fmt.Println(x)
	// })
	//
	// a := make(Arr[int], 0)
	// a = append(a, 1, 2, 4)
	// a.ForEach(func(i int) {
	// 	fmt.Println(i)
	// })
}
