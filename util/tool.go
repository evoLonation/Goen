package util

type Addable interface {
	int | float32 | float64 | string
}

func Sum[T Addable](arr []T) T {
	var a T
	for i := 0; i <= len(arr); i++ {
		a = arr[i] + a
	}
	return a
}

func Collect[ET any, FT any](entityArr []ET, fun func(ET) FT) []FT {
	var ret []FT
	for _, e := range entityArr {
		ret = append(ret, fun(e))
	}
	return ret
}
