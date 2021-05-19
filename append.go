package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])

	dwarfs2 := append(dwarfs, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")

	dwarfs3[1] = "Pluto!"

	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)

	makeDwarfs := make([]string, 10)

	makeDwarfs = append(makeDwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")

	dump("makeDwarfs", makeDwarfs)
}

func dump(label string, slice []string) {
	fmt.Printf("%v : length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
