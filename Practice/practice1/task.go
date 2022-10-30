package practice1

import (
	"fmt"
	"math"
)

	
// #2.1
// Напишите функцию, возвращающую число, 
// введённое в консоль.
func Ret() float64 {
	var ab float64
	fmt.Scan(&ab)
	return ab
}

// #2.2
// Напишите функцию, соединяющую две строки в одну.
func Soed() string{
	s1:="sdsdfsdfdfsf"
	s2:="ooosdfsf"
	s3:=s1+s2
	return s3
}

// #2.3
// Напишите функцию, выводящую две переменные в консоль, 
// одна из переменных объявляется в функции и 
// задаётся с консоли.
func Hz(p1 int){
	p2:=0
	fmt.Scan(&p2)
	fmt.Println(p1,p2)
}

// #3.1
// Напишите функцию, реализующую возведение числа 
// в третью степень для больших чисел включительно.
func Stepen(){
	var ch float64
	fmt.Scan(&ch)
	ch3:=math.Pow(ch,3)
	fmt.Println(ch3)
}

// #3.2
// Реализуйте функцию, способную вывести значение формулы 
// x^2+4x-10. Для реализации возведения в степень 
// используйте пакет math и функцию Pow(число, степень).
func Formul(){
	xf:=0
	fmt.Scan(&xf)
	fmt.Println(xf*xf+4*xf-10)
}

// #3.3
// Напишите функцию, возвращающую площадь эллипса. 
// Необходимо учитывать дробные числа.
func Elips(){
	var ae float32
	var be float32
	fmt.Scan(&ae)
	fmt.Scan(&be)
	S:=math.Pi*ae*be
	fmt.Println(S)
}

// #3.4
// Напишите функцию для вычисления выражения.
// Для реализации необходимо использовать пакет math, 
// для которого необходимо изучить необходимые функции
// самостоятельно.
func Vichislit(){
	var xv float64
	fmt.Scan(&xv) 
	res:=(math.Sin(math.Pow(xv,3))) + (math.Sqrt(xv + (math.Pow(xv,4)))) * (math.Log2(xv))
	fmt.Println(res)
}