package main

func counter(ch <-chan bool) *int {
	count := 0
	go func() {
		for {
			if <-ch {
				count++
			}
		}
	}()
	return &count
}
