package practice3

import(
	"fmt"
	"sort"
	"strings"
)

type User struct {
	Id int
	Name string
	Age int
}

// Задача 1.
// Напишите функцию, которая с помощью цикла находит 
// наибольшее число в массиве. Например, при заданном 
// массиве [1, 5545, 13, 789] функция должна вернуть 
// значение 5545.
func Odin() int {
	chisl := []int{1, 10, 13, 0, 78, -13, -9, -1, 0, -34, 105, 104}
	mx := chisl[0]
	for _, name := range chisl {
		if name > mx {
			mx = name
		}
	}
	return mx
}

// Задача 2.
// Напишите функцию, которая сортирует исходный массив и 
// возвращает новый массив, состоящий только из положительных 
// элементов. Необходимо использовать конструкции, описанные 
// в этом и предыдущем блоке.
func Dva() []int {
	dev := []int{1, -1, 1, 10, -10, -21, 23, 32, 10, -20, 28}
	sort.Slice(dev, func(i, j int) bool {
		return dev[i] < dev[j]
	})
	dev2 := []int{}
	for i, name := range dev {
		if name <= 0 {
			dev2 = dev[i+1:]
		} else {
			break
		}
	}
	
	return dev2
}

// Задача 3.
// Напишите функцию, которая возвращает массив ключей 
// переданной в него мапы. Считается, что ключи и значения 
// типа string.
func Tri() []string {
	m := map[string] string{"one": "Artem", "dva": "Anton", 
	"tri": "Sergei"}
	arr := []string{}
	for i := range m {
		arr = append(arr, i)
	}

	return arr
}

// Задача 4.
// Напишите функцию, которая считает и возвращает количество 
// цифр в строке.
func Che(s string) int {
	c := 0
	for _, i := range s {
		if i >= 48 && i <= 57 {
			c++
		}
	}

	return c
}

// Задача 5.
// Необходимо написать функцию, преобразующие данные из массива 
// выше в набор данных без специальных символов (запятые, точки) 
// и лишних пробелов. Следует учитывать, что имя может быть 
// указано с фамилией. Из полученного набора данных должно 
// получиться 9 отдельных пользователей.
func Piat(Us []string) []string {
	var res string
	var a [] string
	res = strings.Join(Us, " ")
	res = strings.ReplaceAll(res, ",", "")
	res = strings.ReplaceAll(res, ".", "")
	res = strings.ReplaceAll(res, "  ", " ")
	a = strings.Split(res, " ")

	return a
}

// Задача 6.
// Напишите функцию, выводящую в консоль имена всех 
// пользователей в наборе структур. Ещё функция должна считать 
// средний возраст пользователей и возвращать его. Необходимо 
// создать структуру “User”, в которой должны быть поля 
// id, name, age.
func shest() string {
	users := []User {
		{Id: 1, Name: "Artem", Age: 17},
		{Id: 2, Name: "Den", Age: 19},
		{Id: 3, Name: "Stepa", Age: 44},
		{Id: 4, Name: "Roma", Age: 4},
	}
	st := &strings.Builder{}
	sr := 0
	for i, name := range users {
		st.WriteString(name.Name)
		sr += name.Age
		if i == len(users) - 1 {
			break
		}
		st.WriteString(", ")
	}
	sr = sr / len(users)
	st.WriteString(".")
	st.WriteString(fmt.Sprintf(" Средний возраст равен: %d", sr))

	return st.String()
}