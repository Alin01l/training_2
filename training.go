package main

import (
	"fmt"
	"slices"
)

func main() {
	marks()
}

// Кофе машина
func chooseCoffee() {

	fmt.Println("Выберите кофе")
	fmt.Print("1 - Американо")
	fmt.Print("2 - Капучино")
	fmt.Print("3 - Латте")
	fmt.Print("4 - Эспрессо")
	fmt.Print("5 - Выход")

	var choise int

	for {
		fmt.Scan(&choise)

		switch choise {
		case 1:
			fmt.Println("Вы выбрали Американо")
		case 2:
			fmt.Println("Вы выбрали Капучино")
		case 3:
			fmt.Println("Вы выбрали Латте")
		case 4:
			fmt.Println("Вы выбрали Эспрессо")
		case 5:
			fmt.Println("Вы вышли из программы")
			return
		default:
			fmt.Println("Такого варианта нет")
		}
	}
}

// Проверка возраста
func yerspld() {
	year := 2026

	var bd int
	fmt.Println("Введи год своего рождения")
	fmt.Scan(&bd)
	age := year - bd
	fmt.Println("Тебе", age, "лет")
	if age > 40 {
		fmt.Println("Ебать ты старый)")
	} else if age >= 18 && age <= 25 {
		fmt.Println("Так то нихуя себе возраст)")
	} else if age < 18 {
		fmt.Println("Ты че еблан, ты же еще ребенок)")
	}
}

// Проверка на четное и нечетное число
func chetnechet() {
	var num int
	fmt.Println("Введите любое число")
	fmt.Scan(&num)

	if num == 0 {
		fmt.Println("число равно нулю")
	} else if num > 0 {
		fmt.Println("число положительное")
	} else if num > 0 {
		fmt.Println("число отрицательное")
	}
}

// Сравнение размеров членов
func dick() {
	var a string
	fmt.Println("Как тебя зовут?")
	fmt.Scan(&a)

	var aa int
	fmt.Println("Сколько у тебя СМ?")
	fmt.Scan(&aa)

	var b string
	fmt.Println("Как тебя зовут?")
	fmt.Scan(&b)

	var bb int
	fmt.Println("Сколько у тебя СМ?")
	fmt.Scan(&bb)

	if aa > bb {
		fmt.Println("У", a, "член больше чем у", b)
	} else if aa < bb {
		fmt.Println("У", b, "член больше чем у", a)
	} else if aa == bb {
		fmt.Println("У", a, "и", b, "члены одинакового размера")
	}
}

// Калькулятор
func math() {
	var a int
	var b int
	var choise int

	fmt.Println("Введите значние а")
	fmt.Scan(&a)
	fmt.Println("Введите значние b")
	fmt.Scan(&b)

	for {
		fmt.Println("Выберите что хотите узанть: 1 - Периметр  2 - площадь")
		fmt.Scan(&choise)

		switch choise {

		case 1:
			result := 2 * (a + b)
			fmt.Println("ваш периметр состовляет:", result)
			return

		case 2:

			result := a * b
			fmt.Println("ваша площадь состовляет:", result)
			return

		default:
			fmt.Println("Такого варианта нет")
		}
	}
}

// 1234567
func marks() {

	var mark []float64
	var input float64
	fmt.Println("Введите оценку ученику Для выхода из программы нажмите (0)")

	for {
		_, err := fmt.Scan(&input)
		if input == 0 {
			break
		}
		if err != nil {
			fmt.Println("Введите другое значение")
			continue
		}
		if input < 1 {
			fmt.Println("Оценка не может быть меньше 0")
		}

		mark = append(mark, input)
	}
	fmt.Println("Список оценок", mark)

	var sum float64
	var middle float64

	for i := 0; i < len(mark); i++ {
		sum += mark[i]
	}
	middle = sum / float64(len(mark))
	fmt.Println("Средняя оценка:", middle)

	for i := 0; i < len(mark); i++ {
		if mark[i] > middle {
			fmt.Println("Оценкa выше среднего", mark[i])
		}
	}
	maxMark := slices.Max(mark)
	fmt.Println("Самая высокая оценка", maxMark)

	minMark := slices.Min(mark)
	fmt.Println("Самая низкая оценка", minMark)
}
