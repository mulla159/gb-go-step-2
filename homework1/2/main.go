/*
Дополните функцию из п.1 возвратом собственной ошибки в случае
возникновения панической ситуации. Собственная ошибка должна хранить
время обнаружения панической ситуации.
Критерием успешного выполнения задания является наличие обработки созданной
ошибки в функции main и вывод ее состояния в консоль
*/

package main

import (
	"fmt"
	"time"
)

type ErrorWithTime struct {
	Err  error
	Time time.Time
}

func New(e error) error {
	return &ErrorWithTime{e, time.Now()}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\ntime:%s\n", e.Err, e.Time)
}

func tmpFunc() {
	defer func() {
		if err := recover(); err != nil {
			err = New(fmt.Errorf("%v", err))
			panic(err)
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[3])
}

func main() {
	defer func() {
		e := recover()
		switch e.(type) {
		case *ErrorWithTime:
			fmt.Println("ErrorWithTime", e)
		default:
			fmt.Println("Other error", e)
		}
	}()

	fmt.Println("start")
	tmpFunc()
	fmt.Println("end")
}
