package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func f(inputStream <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		var prev string
		first := true
		for el := range inputStream {
			if first || prev != el {
				out <- el
				prev = el
				first = false
			}
		}
	}()
	return out
}

func main() {
	in := make(chan string)

	go func() {
		defer close(in)
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Введите строку (пустая строка для завершения ввода): ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
				return
			}

			input = strings.TrimSpace(input)

			if input == "" {
				return
			}

			in <- input
		}
	}()

	answer := f(in)

	var outputs []string
	for j := range answer {
		outputs = append(outputs, j)
	}

	fmt.Println("\nРезультат обработки:")
	for _, line := range outputs {
		fmt.Println(line)
	}
}
