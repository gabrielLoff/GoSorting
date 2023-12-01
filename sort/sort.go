package sort

import (
	"fmt"
	"github.com/fogleman/gg"
	"sortingBenchmark/draw"
)

func InsertionSort(arr []int, dc *gg.Context) {
	n := len(arr)
	count := 0
	name := "insertion"

	draw.DrawBars(dc, arr, 0, name, count, -1)

	for i := 1; i < n; i++ {

		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1

			count++
			draw.DrawBars(dc, arr, j, name, count, i)
		}
		arr[j+1] = key

		count++
		draw.DrawBars(dc, arr, i, name, count, j+1)
	}

	fmt.Println("insertion steps:", count)
}

func BubbleSort(values []int, dc *gg.Context) {
	n := len(values)
	count := 0
	name := "bubble"

	draw.DrawBars(dc, values, 0, name, count, -1)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
			count = i*n + j
			draw.DrawBars(dc, values, j, name, count, j+1)
		}
	}
	fmt.Println("bubble steps:", count)
}

func CocktailShakerSort(arr []int, dc *gg.Context) {
	n := len(arr)
	swapped := true
	count := 0
	name := "cocktail"

	draw.DrawBars(dc, arr, 0, name, count, -1)

	for swapped {
		swapped = false

		// Perform a bubble sort from left to right
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
			count++
			draw.DrawBars(dc, arr, i, name, count, i+1)
		}

		if !swapped {
			break
		}

		swapped = false

		// Perform a bubble sort from right to left
		for i := n - 1; i > 0; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
			count++
			draw.DrawBars(dc, arr, i, name, count, i-1)
		}
	}

	fmt.Println("cocktail steps:", count)
}

func CycleSort(arr []int, dc *gg.Context) {
	n := len(arr)
	count := 0
	name := "cycle"

	draw.DrawBars(dc, arr, -1, name, count, -1)

	for cycleStart := 0; cycleStart < n-1; cycleStart++ {
		item := arr[cycleStart]
		pos := cycleStart

		// Find the position where we put the element
		for i := cycleStart + 1; i < n; i++ {
			if arr[i] < item {
				pos++
			}
			count++
			draw.DrawBars(dc, arr, pos, name, count, i)
		}

		// If the element is already in the correct position
		if pos == cycleStart {
			continue
		}

		// Otherwise, put the element to the correct position
		for item == arr[pos] {
			pos++
		}
		arr[pos], item = item, arr[pos]

		count++
		draw.DrawBars(dc, arr, pos, name, count, cycleStart)

		// Rotate the rest of the cycle
		for pos != cycleStart {
			pos = cycleStart
			for i := cycleStart + 1; i < n; i++ {
				if arr[i] < item {
					pos++
				}
				count++
				draw.DrawBars(dc, arr, pos, name, count, i)
			}

			for item == arr[pos] {
				pos++
			}
			arr[pos], item = item, arr[pos]

			count++
			draw.DrawBars(dc, arr, pos, name, count, cycleStart)
		}
	}

	fmt.Println("cycle steps:", count)
}

func HeapSort(arr []int, dc *gg.Context) {
	n := len(arr)
	count := 0
	name := "heap"

	draw.DrawBars(dc, arr, -1, name, count, -1)

	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		count = heapify(arr, n, i, dc, count, name)
	}

	// One by one extract elements from the heap
	for i := n - 1; i > 0; i-- {
		// Move current root to the end
		arr[0], arr[i] = arr[i], arr[0]

		count++
		draw.DrawBars(dc, arr, 0, name, count, i)

		// Call max heapify on the reduced heap
		count = heapify(arr, i, 0, dc, count, name)
	}

	fmt.Println("heap steps:", count)
}

func heapify(arr []int, n, i int, dc *gg.Context, count int, name string) int {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	count++
	draw.DrawBars(dc, arr, left, name, count, right)

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected sub-tree
		return heapify(arr, n, largest, dc, count, name)
	}

	return count
}

// getMax finds the maximum number in a slice of integers
func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

// countingSort performs counting sort on a slice of integers based on a specific digit
func countingSort(arr []int, exp int, dc *gg.Context, name string, frame int) int {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	// Initialize count array
	for i := 0; i < n; i++ {
		count[(arr[i]/exp)%10]++
		frame++
		draw.DrawBars(dc, arr, i, name, frame, -1)
	}

	// Update count array to store the position of each element
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
		frame++
		draw.DrawBars(dc, arr, -1, name, frame, i)
	}

	// Copy the output array to the original array
	for i := 0; i < n; i++ {
		arr[i] = output[i]
		frame++
		draw.DrawBars(dc, arr, i, name, frame, -1)
	}

	return frame
}

// radixSort performs radix sort on a slice of integers
func RadixSort(arr []int, dc *gg.Context) {
	count := 0
	name := "radix"
	// Find the maximum number to know the number of digits
	max := getMax(arr)

	draw.DrawBars(dc, arr, -1, name, count, -1)

	// Perform counting sort for every digit
	for exp := 1; max/exp > 0; exp *= 10 {
		count = countingSort(arr, exp, dc, name, count)
		count++
		draw.DrawBars(dc, arr, -1, name, count, -1)
	}

	fmt.Println("radix steps:", count)
}
