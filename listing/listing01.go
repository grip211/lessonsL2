package main

import "fmt"

//Что выведет программа? Объяснить вывод программы.

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

// Программа выведет [77,78,79]
/*
 вот пример: Нарезка выполняется путем указания полуоткрытого диапазона с двумя индексами, разделенными двоеточием.
	Например, выражение b[1:4]создает срез, включающий элементы с 1 по 3 b(индексы результирующего среза будут от 0 до 2).

			b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
			// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
	Начальный и конечный индексы выражения слайса необязательны; по умолчанию они равны нулю и длине среза соответственно:

			// b[:2] == []byte{'g', 'o'}
			// b[2:] == []byte{'l', 'a', 'n', 'g'}
			// b[:] == b
*/
