package remember

import "fmt"

func GenericsTest[K comparable, V float64 | int64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}

	return sum
}

func Remember(name string) string {
	return fmt.Sprintf("Remember %s perseverance and discipline beats talent", name)
}
