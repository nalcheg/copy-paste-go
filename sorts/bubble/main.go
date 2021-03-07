package main

func main() {

}

func BubbleSort(arr []int) []int {
	l := len(arr)
	swapped := true
	for swapped == true {
		swapped = false
		for i := 0; i < l-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}

	return arr
}
