package main

import "sync"

/*func main() {

	odd := make(chan int)
	even := make(chan int)
	result := make(chan int)

	go func() {

		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(odd)
		close(even)
	}()

	go func() {
		for {
			oddVal, oddOk := <-odd
			evenVal, evenOk := <-even
			if oddOk {
				result <- oddVal
			}
			if evenOk {
				result <- evenVal
			}

			if !oddOk && !evenOk {
				close(result)
				return
			}
		}
	}()

	for v := range result {
		println(v)
	}

}*/

// more clear answer
func main() {
    
	wg := &sync.WaitGroup{}
    odd := make(chan int)
	even := make(chan int)
    result := make(chan int)

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i:=1;i<=10;i+=2 {
			odd <- i
		}
		close(odd)
	}()

	go func() {
		defer wg.Done()
		for i:=2;i<=10;i+=2 {
			even <- i
		}
		close(even)
	}()
    
	go func() {
	defer wg.Done()
	oddOpen,evenOpen := true,true
	for oddOpen || evenOpen{
        select {
		case v1,ok := <-odd:
			if !ok {
	            oddOpen = false
				continue
			}
			result <- v1
        case v2,ok := <-even:
			if !ok {
				evenOpen = false
				continue
			}
			result <- v2
		}
	}
	close(result)
}()

  for res := range result {
	println(res)
  }
  wg.Wait()

}
