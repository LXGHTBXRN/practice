package main

import "fmt"

func main() {
	//Целочисленные типы
	var i int = 42
	var i8 int8 = -8
	var i16 int16 = 32000
	var i32 int32 = 400000000
	var i64 int64 = 8000000000
	//Беззнаковые целочисленные типы
	var u uint = 100
	var u8 uint8 = 255
	var u16 uint16 = 65000
	var u32 uint32 = 4000000000
	var u64 uint64 = 18000000000000000000
	//Числа с плавающей точкой
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653
	//Комплексные типы
	var c64 complex64 = complex(1, 2)
	var c128 complex128 = complex(3, 4)
	//Логические типы
	var b bool = true
	//Строка
	var s string = "Hello, Go!"

	fmt.Println("int:", i)
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)

	fmt.Println("uint:", u)
	fmt.Println("uint8:", u8)
	fmt.Println("uint16:", u16)
	fmt.Println("uint32:", u32)
	fmt.Println("uint64:", u64)

	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)

	fmt.Println("complex64:", c64)
	fmt.Println("complex128:", c128)

	fmt.Println("bool:", b)
	fmt.Println("string:", s)
}
