package main

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, temp := range temperatures {
		frequency[temp]++
	}

	for temp, count := range frequency {
		println(temp, "occurs", count, "times")
	}
}
