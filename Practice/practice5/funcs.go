package practice5

import (
	"Practice/practice1"
	"Practice/practice3"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func menu() string {
	return `
	<a href="/">Главная</a>
	<a href="/about">О сайте</a>
	<a href="/func1">Соединение строк</a>
	<a href="/func2">Кол-во цифр в строке</a>
	`
}

func index(w http.ResponseWriter, r *http.Request) {
	res := menu() +
		`<p>Это главная страница сайта,<br/>
		Песики лучше котят!<br/><br/>
		___________$$_$$$$$_____________________________<br/>
		___________$$_____$$$$$$________________________<br/>
		__________$$_________$$_$_______________________<br/>
		_________$$______________$______________________<br/>
		_________$$_$___________$$______________________<br/>
		____________$$$$_________$______________________<br/>
		$$$________$$0$$________$$_____________$$$$$____<br/>
		$$$$$$$$$$$$$$$$________$$____________$$_$__$___<br/>
		$$$_________$$$$$$$$$$$$$$____________$____$$___<br/>
		$$$_______$$$$$$$_$$___$______________$$$___$___<br/>
		__$$$$$$$$$$$___$$$$____________________$$$$____<br/>
		___________$$___$$$$_________$$__$$______$$_____<br/>
		____________$$__$$$$_______$$__$$$_$$____$$_____<br/>
		__________$$$$_$$$$$$$_____$$_______$$$_$$$_____<br/>
		_________$$__$$$_____$$___$$$__________$$$______<br/>
		_________$$__$$$_____$$$$$$$$$$_________$$______<br/>
		__________$$$$$$___$$$___$___$$$_________$$_____<br/>
		____________$$_$$$$$$____$$___$$_________$$_____<br/>
		___________$$_____________$___$$_________$______<br/>
		___________$______________$$___$_______$$_______<br/>
		__________$_______________$___$______$$$$$______<br/>
		___________$$____________$$_$$$$__$$$$$_$$______<br/>
		____________$$_________$_$$$__$$$$$______$______<br/>
		_______$$___$$$_$$$$$$_$$______$$$$_____$$______<br/>
		_____$$$_$$$$$_$$_$$$$$_________$_$_____$$______<br/>
		____$$__$$$$$_$$$$_$$___________$$$$___$$_______<br/>
		___$$__$$____$$$$$_$$____________$$$$__$________<br/>
		___$$__$_________$$$$_____________$$$$_$________<br/>
		____$$$___________$_$______________$$$$_$_______<br/>
		__________________$$$_____________$$_$_$________<br/>
		__________________$_$_________$$$$$$$$$$________<br/>
		______________$$$$_$$_________$_$$$___$_________<br/>
		___________$$$____$$__________$$$____$$_________<br/>
		___________$$$$$$$$____________$$$$$$___________<br/>
		_____________$$$___$$___________________________<br/>
		</p>`

	fmt.Fprint(w, res)
}

func about(w http.ResponseWriter, r *http.Request) {
	res := menu() +
		`<p>Это простенький сайт, запущенный на Гоу,<br/>
	тут даже что-то есть.<br/>
	Тут нет пропоганды собачек,<br/>
	Они просто лучше<br/>
	</p>`

	fmt.Fprint(w, res)
}

func functia1(w http.ResponseWriter, r *http.Request) {

	res := menu() +
		`<p></br></p>` +
		practice1.Soed()

	fmt.Fprint(w, res)
}

func functia2(w http.ResponseWriter, r *http.Request) {
	res := menu() +
		`<p>
		</br>
		localhost:9776/func2/ВПИШИТЕ ЗДЕСЬ СВОЮ СТРОКУ ДЛЯ ПРОВЕРКИ
		</p>`

	fmt.Fprint(w, res)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inp := vars["inp"]
	response := fmt.Sprintf("%s", inp)
	res := menu() +
		`<p>
		</br>
		Количество чисел в строке равно:
		</p>` +
		fmt.Sprintf("%v", practice3.Che(response))

	fmt.Fprint(w, res)
}
