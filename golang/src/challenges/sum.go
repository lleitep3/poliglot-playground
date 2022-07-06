package challenges

func Sum[T int64 | float64](numbers ...T) T {

	var total T

	for _, n := range numbers {
		total = total + n
	}

	return total
}
