// --------------------------------------------
// Author: Qadeer Hussain (C00270632@setu.ie)
// Created on 14/10/2024
// Modified by: Qadeer Hussain
// Issues:
// --------------------------------------------
package main

func producer() {

}

func consumer() {

}
func main() {
	numProd := 5
	numCons := 3

	for i := 0; i < numProd; i++ {
		go producer()
	}
	for i := 0; i < numCons; i++ {
		go consumer()
	}
}
