package main

import (
	"fmt"
	"time"
)

/*
A structure that stores status of the new routine
*/
type Future struct {
	status          string
	cancel          chan bool
	futureCallReady chan bool
}

/*
Main function to start testing the future functions
*/
func main() {
	fmt.Println("Hello Universe")
	futurepointer := setup_routine(print_something)
	// fmt.Println("\ncan it be cancelled:", routine_cancel(futurepointer))
	time.Sleep(3 * time.Second)
	fmt.Println("\nis add_done_callback done:", add_done_callback(print_something_new, futurepointer))
	fmt.Println("\nis the 1st routine running:", is_routine_running(futurepointer))
	fmt.Println("\nis 1st routine cancelled:", is_routine_cancelled(futurepointer))
	time.Sleep(3 * time.Second)
	fmt.Println("\nis the 1st routine done:", is_routine_done(futurepointer))
	time.Sleep(5 * time.Second)
	fmt.Println("\nis add_done_callback done:", add_done_callback(print_something_new, futurepointer))
}

/*
Method to setup and start routines, channels and pointers
*/
func setup_routine(printsomething func()) *Future {
	started := make(chan bool)
	ended := make(chan bool)
	cancel := make(chan bool)
	futureCallReady := make(chan bool)
	futureStatus := Future{"notstarted", cancel, futureCallReady}
	futurepointer := &futureStatus
	go start_go_routine(printsomething, started, ended, cancel, futureCallReady)
	go update_start_status(started, futurepointer)
	go update_end_status(ended, futurepointer)
	return futurepointer
}

/*
Method to read the channel data and update routine status
*/
func update_start_status(started chan bool, pointer *Future) {
	startedorcancelled := <-started
	if startedorcancelled == true {
		pointer.status = "started"
	}
}

/*
Method to read the channel data and update routine status
*/
func update_end_status(ended chan bool, pointer *Future) {
	endedorcancelled := <-ended
	if endedorcancelled == true {
		pointer.status = "ended"
	}
}

/*
Method to run the actual function
*/
func start_go_routine(actual_function func(), started chan bool,
	ended chan bool, cancel chan bool, futureCallReady chan bool) {
	select {
	case iscancelledValue := <-cancel:
		close(started)
		fmt.Printf("routine cancelled\n", iscancelledValue)
		close(ended)
		futureCallReady <- true
	default:
		started <- true
		fmt.Printf("routine started\n")
		close(cancel)
		actual_function()
		ended <- true
		futureCallReady <- true
	}
}

/*
A dummy function
*/
func print_something() {
	for a := 0; a < 10; a++ {
		fmt.Printf("\nvalue of a: %d", a)
		time.Sleep(time.Second)
	}
}

/*
Method to cancel the routine
*/
func routine_cancel(pointer *Future) bool {
	if pointer.status == "notstarted" {
		fmt.Printf("\n trying to cancel func\n")
		pointer.cancel <- true
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
	} else {
		fmt.Printf("Started new func")
		<-pointer.futureCallReady
		go newfunc()
	}
	return true
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

/*
Method to read result from a process

func read_result(timeout int){
	time.Sleep(time.Duration(timeout)*time.Second)

}
*/
