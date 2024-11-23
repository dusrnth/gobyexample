package main

import "fmt"

func main() {

	// 기본 int 타입
	// int 는 플랫폼에 따라 크기가 32비트 또는 64비트
	var a int = 42
	fmt.Println("a: ", a)

	// 고정 크기 타입 - int32
	// 크기가 항상 32비트 (4바이트)
	var b int32 = 42
	fmt.Println("b: ", b)

	// 고정 크기 타입 - int64
	// 크기가 항상 64비트 (8바이트)
	var c int64 = 42
	fmt.Println("c: ", c)

	// 4. 타입 변환 예제
	// int와 int32는 서로 다른 타입이므로 명시적으로 변환해야 합니다.
	var d int = int(b)     // int32 → int
	var e int64 = int64(a) // int → int64
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
}
