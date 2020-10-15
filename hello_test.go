package main

import (
	"fmt"
	"testing"
	"time"
)

/*
Testcase to run everything
*/
// func Test_main(t *testing.T) {
// 	fmt.Printf("Hello Universe")
// 	result := make(chan string)
// 	futurepointer := setup_routine(print_something, result)
// 	go add_done_callback(print_something_new, futurepointer)
// 	fmt.Println("is the 1st routine running:  ", is_routine_running(futurepointer))
// 	fmt.Println("is 1st routine cancelled:  ", is_routine_done(futurepointer))
// 	fmt.Println("final result:  ", read_result(15, result))
// 	fmt.Println("is 1st routine cancelled:  ", is_routine_cancelled(futurepointer))
// 	fmt.Println("is the 1st routine done:  ", is_routine_done(futurepointer))
// 	time.Sleep(10 * time.Second)
// }

/*
Testcase to setup a go routine
*/
func Test_setup_routine(t *testing.T) {
	result := make(chan string)
	setup_routine(print_something, result)
	time.Sleep(7 * time.Second)
}

/*
Testcase to cancel a started  process
*/
func Test_cancel_function(t *testing.T) {
	result := make(chan string)
	futurepointer := setup_routine(print_something, result)
	routine_cancel(futurepointer)
}

/*
Testcase to check if a process is cancelled
*/
func Test_is_routine_cancelled(t *testing.T) {
	result := make(chan string)
	futurepointer := setup_routine(print_something, result)
	status := is_routine_cancelled(futurepointer)
	if status != false {
		t.Errorf("error in function")
	}
}

/*
Testcase to check if a process is running
*/
func Test_is_routine_running(t *testing.T) {
	result := make(chan string)
	futurepointer := setup_routine(print_something, result)
	time.Sleep(3 * time.Second)
	status := is_routine_running(futurepointer)
	fmt.Println("\n", status)
}

/*
Testcase to check if a process is done
*/
func Test_is_routine_done(t *testing.T) {
	result := make(chan string)
	futurepointer := setup_routine(print_something, result)
	time.Sleep(7 * time.Second)
	status := is_routine_done(futurepointer)
	fmt.Println("\n", status)
}

/*
Testcase to check the result
*/
func Test_to_read_result(t *testing.T) {
	result := make(chan string)
	setup_routine(print_something, result)
	resp := read_result(7, result)
	fmt.Println("\n", resp)
}

/*
Testcase to add a new function when done with initial function
*/
func Test_add_done_callback(t *testing.T) {
	result := make(chan string)
	futurepointer := setup_routine(print_something, result)
	go add_done_callback(print_something_new, futurepointer)
	time.Sleep(10 * time.Second)
}
