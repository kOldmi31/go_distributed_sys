package cos418_hw1_1

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// HINT: use for loop over `nums`
	sum := 0
	//result of ranging over a channnel for each iteration - value from chan
	for n := range nums{
		sum += n
	}
	// Send sum to channel out
	out <- sum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
// HINT: use `readInts` and `sumWorkers`
// HINT: used buffered channels for splitting numbers between workers
func sum(num int, fileName string) int {
	file, err := os.Open(fileName) 		// For read access https://pkg.go.dev/os
	if err != nil {
	log.Fatal("Error in reading initial file: ", err)
	}
	nums := make(chan int, 50) 		//buffer for splitting numbers between workers
	out := make(chan int, 50)
	numbers, err2 := readInts(file) 	// create slice of integers - func below
	if err2 != nil {
	log.Fatal("Error readInts (str 39): ", err2)
	}
	for i := 0; i < num; i++ {
		go sumWorker(nums, out) 	// run goroutine, num - count of workers - variable
	}
	for _, j := range numbers{ 		// in parallel send numbers from readInts(file) slice to channel nums
		nums <- j
	}
	close(nums) 	// close channel nums!! otherwise deadlock
	result := 0
	defer file.Close() 		// also works without closing the file, but why not? wanna try defer
	return toResultFunc(num, out, result)
}

func toResultFunc(num int, out chan int, result int) int { 		// my func
	for i := 0; i < num; i++ { 		// for each worker send all values from all chan to result
		n := <- out
		result += n
	}
	return result
}
// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
