package main

import (
	"fmt"
	"sync"
)

func main() {
	var mtx sync.Mutex
	var wg sync.WaitGroup

	bisa := []string{"bisa1", "bisa2", "bisa3"}
	coba := []string{"coba1", "coba2", "coba3"}

	ch := make(chan []string)

	go func() {
		for {
			ch <- bisa
			ch <- coba
		}
	}()

	//GOROUTINE keduanya menampilkan secara acak
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(id int) {

			arr := <-ch

			fmt.Printf("%v %d\n", arr, id)
			wg.Done()
		}(i)
	}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(id int) {

			arr := <-ch

			fmt.Printf("%v %d\n", arr, id)
			wg.Done()
		}(i)
	}
	fmt.Println("GOROUTINE keduanya menampilkan secara acak")
	wg.Wait()

	//GOROUTINE keduanya menampilkan secara rapih
	//(gunakan mutex golang dengan fungsi lock, dan unlock)
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(id int) {
			mtx.Lock()
			defer mtx.Unlock()

			arr := <-ch

			fmt.Printf("%v %d\n", arr, id)
			wg.Done()
		}(i)
	}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(id int) {
			mtx.Lock()
			defer mtx.Unlock()

			arr := <-ch

			fmt.Printf("%v %d\n", arr, id)
			wg.Done()
		}(i)
	}
	fmt.Println("\nGOROUTINE keduanya menampilkan secara rapih")
	wg.Wait()

}
