package integer

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Ordered interface {
	Integer | Float | ~string
}

func Min[T Integer](arr ...T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}
	min := arr[0]
	for _, num := range arr {
		if min > num {
			min = num
		}
	}
	return min
}

func Max[T Integer](arr ...T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}
	max := arr[0]
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

func Pow[T Integer](n, p T) T {
	if n == 0 {
		return 0
	}

	var result T = 1
	for ; p > 0; p-- {
		result *= n
	}

	return result
}
