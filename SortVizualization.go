package main

import (
	"github.com/fogleman/gg"
	"math/rand"
	"sortingBenchmark/sort"
	"sortingBenchmark/util"
)

const (
	width  = 800
	height = 400
)

func createFolders() {
	util.CreateFolderIfNotExists("./bubble")
	util.CreateFolderIfNotExists("./cocktail")
	util.CreateFolderIfNotExists("./cycle")
	util.CreateFolderIfNotExists("./heap")
	util.CreateFolderIfNotExists("./insertion")
	util.CreateFolderIfNotExists("./radix")

}

func cleanImages() {
	util.CleanImages("./bubble")
	util.CleanImages("./cocktail")
	util.CleanImages("./cycle")
	util.CleanImages("./heap")
	util.CleanImages("./insertion")
	util.CleanImages("./radix")
}

func main() {

	//Size of the array
	size := 40

	cleanImages()
	createFolders()

	values := make([]int, size)
	for i := range values {
		values[i] = rand.Intn(height)
	}

	dc := gg.NewContext(width, height)

	values1 := make([]int, size)
	copy(values1, values)
	sort.InsertionSort(values1, dc)

	values2 := make([]int, size)
	copy(values2, values)
	sort.BubbleSort(values2, dc)

	values3 := make([]int, size)
	copy(values3, values)
	sort.CocktailShakerSort(values3, dc)

	values4 := make([]int, size)
	copy(values4, values)
	sort.CycleSort(values4, dc)

	values5 := make([]int, size)
	copy(values5, values)
	sort.HeapSort(values5, dc)

	values6 := make([]int, size)
	copy(values6, values)
	sort.RadixSort(values6, dc)

}
