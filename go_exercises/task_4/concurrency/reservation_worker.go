package concurrency

import "fmt"
import "time"

type reservation struct {
	bookId		int
	memberId	int
}

var reserve_chan chan reservation
var reserved_chan chan reservation

func Init(worker_count int) {
	reserved_chan = make(chan reservation)
	reserve_chan = make(chan reservation)

	for range worker_count {
		go reservation_worker()
	}

	go func() {
		for reserved := range reserved_chan {
			fmt.Printf("Book %d reverved by %d\n", reserved.bookId, reserved.memberId)
		}
	}()
}

func Reserve_book(bookId, memberId int) {
	reserve_chan <- reservation{bookId: bookId, memberId: memberId}
}

func reservation_worker(){
	for reserve := range reserve_chan {
		fmt.Printf("Reserving book %d by %d\n", reserve.bookId, reserve.memberId)
		time.Sleep(2 * time.Second)
		reserved_chan <- reserve
	}
}
