/*
Предложите реализацию примера так, чтобы аварийная остановка программы не выполнилась
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	a := func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}

	go func() {
		defer a()
		panic("A-A-A!!!")
	}()

	time.Sleep(time.Second)
}
