package main

// 메인 패키지에 속한 코드임을 알려줌 (main은 프로그램의 시작점) -> go 언어는 패키지 선언으로 시작되어야 함

// 표준 입출력을 다루는 내장 패키지
import "fmt"

func Add(a int, b int) int { // 함수 정의
	return a + b
}

// 재귀호출: 함수 안에서 자기 자신 함수를 다시 호출하는 것
func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func UploadFile() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println("Hello Go World")
	c := Add(3, 6) // 함수 호출하여 변수에 저장
	fmt.Println(c)

	printNo(3)

	/* 1. 초기문; 조건문
	조건문 검사 전에 초기문을 넣을 수 있음. 초기문은 검사에 사용할 변수를 초기화할 때 주로 사용. */
	if filename, success := UploadFile(); success {
		// 초기문에서 선언한 변수의 범위는 if문 안으로 한정됨.
		fmt.Println("Upload Success", filename)
	} else {
		fmt.Println("Failed to upload")
	}

	/* 2. switch 문
	값에 따라 다른 로직 사용할 떄.
	*/

	// 복잡한 if else문을 보기 좋게 정리할 수 있음.
	day := 3

	switch day {
	case 1:
		fmt.Println("첫째 날입니다.")
	case 2:
		fmt.Println("둘째 날입니다.")
	}

	// 하나의 case는 하나 이상의 값을 검사할 수 있음. 각 값은 쉼표로 구분.

}

func main() {
	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("Chile")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}
