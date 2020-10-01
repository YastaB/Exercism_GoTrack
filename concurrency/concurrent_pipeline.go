package main

func gen(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range values {
			ch <- val
		}
		close(ch)
	}()
	return ch
}

func sqrt(ch <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		for val := range ch {
			outCh <- val * val
		}
		close(outCh)
	}()
	return outCh
}

func consume(ch <- chan int) []int {
	var arr []int
	for val := range ch {
		arr = append(arr, val)
	}
	return arr
}
