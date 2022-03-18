package main

import (
	"fmt"
	"gorx/pkg/gorx"
	"time"
)

func main() {
	fmt.Println("little RXJS style observable demo using the new generics")
	// create an observable
	name := gorx.NewObservable("Max")
	// subscribe to the observable and print the resulting value to stdout
	name.Subscribe(func(value string) {
		fmt.Println("subscription one read value:", value)
	})
	// subscribe to the observable and print the resulting value to stdout
	name.Subscribe(func(value string) {
		fmt.Println("subscription two read value:", value)
	})
	time.Sleep(time.Second)
	name.Set("Anna")
	time.Sleep(time.Second)
	name.Set("Lisa")
	time.Sleep(time.Second)
	name.Set("Ben")
	time.Sleep(time.Second)

	/*
		little RXJS Style observable demo using the new generics
		subscription two read value: Max
		subscription one read value: Max
		subscription two read value: Anna
		subscription one read value: Anna
		subscription one read value: Lisa
		subscription two read value: Lisa
		subscription two read value: Ben
		subscription one read value: Ben
	*/
}
