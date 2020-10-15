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
	futureCallReady chan bool
}

/*
Main function to start testing the future functions
*/
func main() {
	fmt.Printf("Hello Universe")
	result := make(chan string)
	futurepointer := setup_routine(print_something, result)
	go add_done_callback(print_something_new, futurepointer)
	// fmt.Println("\ncan it be cancelled:", routine_cancel(futurepointer))
	time.Sleep(3 * time.Second)
	fmt.Println("is the 1st routine running:  ", is_routine_running(futurepointer))
	fmt.Println("is 1st routine cancelled:  ", is_routine_done(futurepointer))
	fmt.Println("final result:  ", read_result(15, result))
	fmt.Println("is 1st routine cancelled:  ", is_routine_cancelled(futurepointer))
	fmt.Println("is the 1st routine done:  ", is_routine_done(futurepointer))
	time.Sleep(15 * time.Second)
	// go add_done_callback(print_something_new, futurepointer)
}

/*
Method to setup and start routines, channels and pointers
*/
func setup_routine(printsomething func() string, result chan string) *Future {
	futureCallReady := make(chan bool)
	futureStatus := Future{"notstarted", futureCallReady}
	futurepointer := &futureStatus
	go start_go_routine(printsomething, result, futurepointer, futureCallReady)
	go check_readiness_future(futureCallReady)
	go check_future_Result(result)
	return futurepointer
}

/*
Method to check readiness of adding new func
*/
func check_future_Result(result chan string) {
	resp := <-result
	result <- resp
}

/*
Method to check readiness of adding new func
*/
func check_readiness_future(futureCallReady chan bool) {
	<-futureCallReady
	futureCallReady <- true
}

/*
Method to run the actual function
*/
func start_go_routine(actual_function func() string, result chan string,
	future *Future, futureCallReady chan bool) {
	if future.status == "cancelled" {
		fmt.Printf("\nroutine cancelled")
		close(result)
		futureCallReady <- true
	} else {
		future.status = "started"
		fmt.Printf("\nroutine started")
		resp := actual_function()
		future.status = "ended"
		futureCallReady <- true
		result <- resp
	}
}

/*
A dummy function
*/
func print_something() string {
	for a := 0; a < 10; a++ {
		fmt.Printf("\nvalue of a: %d", a)
		time.Sleep(time.Second)
	}
	return "I am done"
}

/*
Method to cancel the routine
*/
func routine_cancel(pointer *Future) bool {
	if pointer.status == "notstarted" {
		fmt.Printf("\n trying to cancel func")
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
		fmt.Printf("\nStarted new func")
		newfunc()
	} else {
		fmt.Printf("\nStarted new func")
		<-pointer.futureCallReady
		newfunc()
	}
	return true
}

/*
Another dummy method
*/
func print_something_new() {
	for a := 0; a < 5; a++ {
		fmt.Printf("\nnew value of b: %d", a)
		time.Sleep(time.Second)
	}
}

/*
Method to read result
*/
func read_result(timeout int, result chan string) string {
	time.Sleep(time.Duration(timeout) * time.Second)
	select {
	case resp := <-result:
		return resp
	default:
		panic("no result")
	}
}
