package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		// канал для счётчика
		ch      = make(chan struct{}, 1)
		counter = 1
		// канал для сигнала остановки приложения
		listener = make(chan os.Signal, 1)
	)

	// регистрируем listener на ожидание сигналов terminate (kill), interrupt (ctrl+c)
	signal.Notify(listener, syscall.SIGTERM, syscall.SIGINT)

	// создаём пустой контекст
	ctx := context.Background()
	// получаем контекст с функцией завершения
	ctx, cancel := context.WithCancel(ctx)

	// запускаем 1000 горутин
	for i := 1; i <= 1000; i++ {
		ch <- struct{}{}

		go func() {
			<-ch
			// увеличиваем счётчик на 1
			counter++
		}()
	}
	close(ch)
	fmt.Println("counter: ", counter)
	fmt.Println("Exit: Ctrl+C")

	// дожидаемся первого события в listener
	select {
	case s := <-listener:
		log.Println("received: ", s)
		// таймаут на 1 секунду
		time.Sleep(1 * time.Second)
		cancel()
	}
}
