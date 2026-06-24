package main

func generator(ip chan<- int) {
	for i := range 5 {
		ip <- i
	}
	close(ip)
}

func square(ip <-chan int, sq chan<- int) {
	for val := range ip {
		sq <- val * val
	}
	close(sq)
}

func evenSelector(sq <-chan int, evenResult chan<- int) {
	for val := range sq {
		if val%2 == 0 {
			evenResult <- val
		}
	}
	close(evenResult)
}

func main() {

	ip := make(chan int)
	sq := make(chan int)
	evenResult := make(chan int)

	go generator(ip)
	go square(ip, sq)
	go evenSelector(sq, evenResult)

	for value := range evenResult {
		println(value)
	}

}
