package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите пример:")
	input, _ := reader.ReadString('\n') //вводим пример.
	input = strings.TrimSpace(input)    //Очищает все пустоты.
	c := checkInt(string(input[0]))
	if c == false {
		fmt.Println("Вывод ошибки: формат строки не соответствует правилам ввода.")
		return
	}

	switch {
	case strings.Contains(input, "+"): //Ищем знак в примере.
		arr := strings.Split(input, " + ") //Превращаем пример в срез.
		if _, err := checkLen(arr); err != nil {
			fmt.Println(err)
			return
		}
		arr[0] = strings.ReplaceAll(arr[0], `"`, "")
		arr[1] = strings.ReplaceAll(arr[1], `"`, "")
		if _, err := checkLenArgs(arr[0], arr[1]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(`"%s%s"`, arr[0], arr[1])
	case strings.Contains(input, "-"): //Ищем знак в примере.
		arr := strings.Split(input, " - ") //Превращаем пример в срез.
		if _, err := checkLen(arr); err != nil {
			fmt.Println(err)
			return
		}
		arr[0] = strings.ReplaceAll(arr[0], `"`, "")
		arr[1] = strings.ReplaceAll(arr[1], `"`, "")
		if _, err := checkLenArgs(arr[0], arr[1]); err != nil {
			fmt.Println(err)
			return
		}
		if strings.Contains(arr[0], arr[1]) == true {
			result = strings.ReplaceAll(arr[0], arr[1], "")
			fmt.Printf(`"%s"`, result)
			return
		}
		fmt.Printf(`"%s"`, arr[0])
	case strings.Contains(input, "*"): //Ищем знак в примере.
		arr := strings.Split(input, " * ") //Превращаем пример в срез.
		if _, err := checkLen(arr); err != nil {
			fmt.Println(err)
			return
		}
		arr[0] = strings.ReplaceAll(arr[0], `"`, "")
		if _, err := checkLenArgs(arr[0], arr[1]); err != nil {
			fmt.Println(err)
			return
		}
		a1, err := strconv.Atoi(arr[1]) //Конвертируем первое число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат второго выражения неверный.", err)
			return
		}
		if _, err := checkArguments(a1); err != nil {
			fmt.Println(err)
			return
		}
		result = strings.Repeat(arr[0], a1)
		if len(result) > 40 {
			result = result[:40]
			fmt.Printf(`"%s..."`, result)
			return
		}
		fmt.Printf(`"%s"`, result)
	case strings.Contains(input, "/"): //Ищем знак в примере.
		arr := strings.Split(input, " / ") //Превращаем пример в срез.
		if _, err := checkLen(arr); err != nil {
			fmt.Println(err)
			return
		}
		arr[0] = strings.ReplaceAll(arr[0], `"`, "")
		a1, err := strconv.Atoi(arr[1]) //Конвертируем первое число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат второго выражения неверный.", err)
			return
		}
		if _, err := checkArguments(a1); err != nil {
			fmt.Println(err)
			return
		}
		a1 = len(arr[0]) / a1
		for i := 0; len(arr[0]) > a1; i++ {
			arr[0] = arr[0][:len(arr[0])-1]
		}
		fmt.Printf(`"%s"`, arr[0])
	default:
		fmt.Println("Вывод ошибки: формат строки неверный.")
	}
}

// Проверяем , не является ли первый аргумент выражения строкой.
func checkInt(a string) bool {
	_, err := strconv.Atoi(a)
	if err != nil {
		return true
	}
	return false
}

// Проверяем диапазон от 1 до 10 включительно.
func checkArguments(a int) (bool, error) {
	if a >= 1 && a <= 10 {
		return true, nil
	}
	return false, errors.New("Вывод ошибки, числа не могут быть меньше 1 и больше 10.")
}

// Проверяем размер строки  до 10 включительно.
func checkLenArgs(a, b string) (bool, error) {
	if len(a) <= 10 && len(b) <= 10 {
		return true, nil
	}
	return false, errors.New("Вывод ошибки, длина аргумента не могут быть меньше больше 10.")
}

// Проверяем кол-во аргументов в строке.
func checkLen(a []string) (bool, error) {
	if len(a) > 2 {
		return false, errors.New("Вывод ошибки, аргументов в строке больше двух.")
	}
	return true, nil
}
