package main

import (
	"fmt"
	"time"
)

/*
A structure that stores status of the new routine
*/
type Future struct {
	status string
}

/*
Main function to start testing the future functions
*/
func main() {
	fmt.Println("Hello Universe")
	futurepointer := setup_routine(print_something)
	time.Sleep(1 * time.Second)
	fmt.Println(" can it be cancelled:", routine_cancel(futurepointer))
	fmt.Println("\nis add_done_callback done:", add_done_callback(print_something_new, futurepointer))
	time.Sleep(3 * time.Second)
	fmt.Println(" is routine cancelled:", is_routine_cancelled(futurepointer))
	time.Sleep(3 * time.Second)
	fmt.Println("is the routine running:", is_routine_running(futurepointer))
	time.Sleep(11 * time.Second)
	fmt.Println("is the routine done:", is_routine_done(futurepointer))
	fmt.Println("is add_done_callback done:", add_done_callback(print_something_new, futurepointer))
	time.Sleep(5 * time.Second)
}

/*
Method to setup and start routines, channels and pointers
*/
func setup_routine(printsomething func()) *Future {
	futureStatus := Future{"notstarted"}
	futurepointer := &futureStatus
	started := make(chan bool)
	ended := make(chan bool)
	cancel := make(chan bool)
	go start_go_routine(printsomething, started, ended, cancel)
	go update_start_status(started, futurepointer)
	go update_end_status(ended, futurepointer)
	return futurepointer
}

/*
Method to read the channel data and update routine status
*/
func update_start_status(started chan bool, pointer *Future) {
	<-started
	pointer.status = "started"
}

/*
Method to read the channel data and update routine status
*/
func update_end_status(ended chan bool, pointer *Future) {
	<-ended
	pointer.status = "ended"
}

/*
Method to run the actual function
*/
func start_go_routine(actual_function func(), started chan bool, ended chan bool, cancel chan bool) {
	started <- true
	select {
	case <-cancel:
		fmt.Printf("routine cancelled")
	default:
		actual_function()
	}
	ended <- true
}

/*
A dummy function
*/
func print_something() {
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
		time.Sleep(time.Second)
	}
}

/*
Method to cancel the routine
*/
func routine_cancel(pointer *Future) bool {
	if pointer.status == "notstarted" {
		pointer.status = "cancelled"
		return true
	} else {
		return false
	}
}

/*
Method to check whether a routine is cancelled?
*/
func is_routine_cancelled(pointer *Future) bool {
	if pointer.status == "cancelled" {
		return true
	} else {
		return false
	}
}

/*
Method to check whether a routine is running?
*/

func is_routine_running(pointer *Future) bool {
	if pointer.status == "started" {
		return true
	} else {
		return false
	}
}

/*
Method to check whether a routine is completed?
*/
func is_routine_done(pointer *Future) bool {
	if pointer.status == "ended" {
		return true
	} else {
		return false
	}
}

/*
Method to add new function after the completion of the first function
*/
func add_done_callback(newfunc func(), pointer *Future) bool {
	if pointer.status == "ended" || pointer.status == "cancelled" {
		fmt.Printf("Started new func")
		newfunc()
		return true
	} else {
		fmt.Printf("unable to start new func")
		return false
	}
}

/*
Another dummy method
*/
func print_something_new() {
	for a := 0; a < 5; a++ {
		fmt.Printf("new value of a: %d\n", a)
		time.Sleep(time.Second)
	}
}
