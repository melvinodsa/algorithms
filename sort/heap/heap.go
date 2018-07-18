package heap

//HeapSort sorts a given integer array in ascending order
func HeapSort(array []int) {
	/*
	 * We will first build the heap
	 * Then one by one we will swaps the root and last element
	 */
	//Building the heap
	buildHeap(array)

	//swaping the root and last element
	for length := len(array); length > 1; length-- {
		swapRoot(array, length)
	}
}

//buildHeap builds the heap structure of an integer array
func buildHeap(array []int) {
	/*
	 * We will run heapify for the entire array
	 */
	for i := len(array) / 2; i >= 0; i-- {
		heapify(array, i, len(array))
	}
}

//swapRoot swaps the root and last element of the given array followed by a heapify from root element
func swapRoot(array []int, length int) {
	/*
	 * We will first swap the first and the last elements
	 * Then runa heapify
	 */
	//Swapping the root and the last element
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]

	//Running a heapify after the element swaps
	heapify(array, 0, lastIndex)
}

//heapify heapifies an integer array
func heapify(array []int, root, length int) {
	/*
	 * First we will check whether the parent element is greater than the children
	 * If not we will swap the parent and child node and run a heapify on the same
	 */
	//Checking if the root element is greater than its child nodes
	var max = root
	var l, r = (root * 2) + 1, (root * 2) + 2

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		//If not swap between the root and the child nodes
		array[root], array[max] = array[max], array[root]
		heapify(array, max, length)
	}
}
