package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outChan := make(chan int)

	go func() {
		defer close(outChan)

		select {
		case val := <-firstChan:
			outChan <- val * val
		case val := <-secondChan:
			outChan <- val * 3
		case <-stopChan:
			return
		}
	}()

	return outChan
}

func main() {
	first := make(chan int)
	second := make(chan int)
	stop := make(chan struct{})

	result := calculator(first, second, stop)

	go func() {
		first <- 5
		close(first)
	}()

	if res, ok := <-result; ok {
		fmt.Printf("Квадрат числа: %d\n", res)
	}

	first2 := make(chan int)
	second2 := make(chan int)
	stop2 := make(chan struct{})

	result2 := calculator(first2, second2, stop2)

	go func() {
		second2 <- 7
		close(second2)
	}()

	if res, ok := <-result2; ok {
		fmt.Printf("Значение, умноженное на 3: %d\n", res)
	}

	first3 := make(chan int)
	second3 := make(chan int)
	stop3 := make(chan struct{})

	result3 := calculator(first3, second3, stop3)

	go func() {
		close(stop3)
	}()

	if res, ok := <-result3; ok {
		fmt.Printf("Результат: %d\n", res)
	} else {
		fmt.Println("Работа функции завершена без отправки значений.")
	}
}
