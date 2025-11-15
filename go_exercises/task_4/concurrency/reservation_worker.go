package concurrency

import "fmt"
import "time"
import "library_management/services"

type reservation struct {
	bookId		int
	memberId	int
}

var lib services.Library
var reserve_chan chan reservation
var reserved_chan chan reservation

func auto_cancel_reservation(r reservation) {
	// waits for five seconds
	ch := time.NewTimer(5 * time.Second)
	<- ch.C

	err := lib.UnReserveBook(r.bookId, r.memberId)
	if err == nil {
		fmt.Printf("Unreserved Book %d by %d\n", r.bookId, r.memberId)
	} else {
		fmt.Printf("%s\n",err.Error())
	}
}
func Init(worker_count int, _lib services.Library) {
	reserved_chan = make(chan reservation)
	reserve_chan = make(chan reservation)
	lib = _lib

	for range worker_count {
		go reservation_worker()
	}

	go func() {
		for reserved := range reserved_chan {
			go auto_cancel_reservation(reserved)
		}
	}()
}

func Reserve_book(bookId, memberId int) {
	reserve_chan <- reservation{bookId: bookId, memberId: memberId}
}

func reservation_worker(){
	for reserve := range reserve_chan {
		err := lib.ReserveBook(reserve.bookId, reserve.memberId)
		if err == nil {
			fmt.Printf("Book %d reverved by %d\n", reserve.bookId, reserve.memberId)
		} else {
			fmt.Printf("Unable to reserve Book %d by %d\n", reserve.bookId, reserve.memberId)
		}
		reserved_chan <- reserve
	}
}
