# Go

GO에서 제공하는 데이터 처리 도구를 사용하여 Web Scrapper를 만들면서 GO언어를 익혀보자. 도구를 익히면서 GO의 특징인 `multi-core`와 `병행성(concurrency)`의 매력을 느껴보자

# 01. THEORY

## #1.0 Main Package

- `main.go`라는 이름의 패키지를 만든다.
- 만약 나만의 프로젝트를 컴파일하고 싶다면 패키지 이름은 선택사항이 아니다. → 무조건 `main.go`라고 만들어야 한다는 말
- 프로젝트를 컴파일 하고 → 서버를 시작하고 → 웹 스크래핑을 하고 등...
- 목적에 따라 프로젝트 컴파일이 필요 없을 수는 있다. 공유를 위한 라이브러리를 만든다든지 오픈소스에 기여한다든지.. 이러한 경우에는 굳이 `main.go`를 사용하지 않아도 된다.

```go
package main
... // 컴파일을 위해선 이렇게 만들어준다.

package learning
... // 내가 만들고 있는 기능들의 묶음이라든가, 다른 사람이 사용(컴파일)할 수 있도록 하기 위한 네이밍
```

- `main.go`는 진입점(entry point)라서 go 컴파일러는 패키지의 이름이 main인 것부터 찾아낸다. 자동적으로 컴파일러는 `main.go`라는 이름의 패키지와 그 안에 있는 main function을 먼저 찾고 실행시킨다.

### Go에서 Hello World 프린트하기

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

## #1.1 Packages and Imports

- vs code go 익스텐션에서 제공하는 기능으로 main() {} 안에서 fmt를 입력하면 `import "fmt"`가 자동으로 위에 생성된다.
- "fmt"는 포맷팅을 해주는 패키지
- Println에서 P가 대문자인 이유. Go에서는 대문자로 시작하는 모듈은 `export`하겠다는 의미.

```go
// something 폴더 안에 something.go 패키지
func something

import "fmt"

func sayHello() {
	fmt.Println("Hello")
} // export 안 됨

func sayBye() {
	fmt.Println("Bye")
} // export 안 됨

func SayHello() {
	fmt.Println("Hello")
} // export 됨!
```

```go
package main

import (
	"fmt"
	"github.com/wooseopim/learngo/something" // vs code에서 자동 추가
)

func main() {
	fmt.Println("Hello World!")
	something.SayHello()
	// something.sayBye() 불가능
	// something.sayHello9) 불가능

}
```

- fmt 패키지에 대해서 자세히 알고 싶으면 fmt를 `Ctrl+클릭`한다. fmt 패키지를 통해서 우리가 사용할 수 있는 대부분의 모듈들은 대문자로 시작하는 함수임을 알 수 있다.

## #1.2 Variables and Constants

- Go의 상수 선언: `const`
  Go의 변수 선언: `var`

```go
package main

import "fmt"

func main() {
	const name string = "nico" // 상수: 바꿀 수 없다.
	var name2 string = "nico"  // 변수: 바꿀 수 있다.
	name2 = "lynn"

	fmt.Println(name)
	fmt.Println(name2) // 바꿀 수 있다.
}
```

- 좀 더 SEXY한 방법을 사용해보자.
- 아래와 같은 축약형(`:=`)은 오로지 `func` 안에서만 사용 가능하고 `변수`에만 적용 가능하다는 것에 주의!
- 축약형 표현이 존재한다면 Go가 첫 번째 값을 기준으로 변수의 type을 찾아 정해줄 것이다.

```go
package main

import "fmt"

func main() {
	name := "nico" // 축약법: Go가 알아서 type을 찾아준다.
	name = "lynn"
	fmt.Println(name) // 결과: lynn

	name = false // 처음 정의된 type에 의해서 boolean인 false는 넣을 수 없다.
	fmt.Println(name) // 결과: Error
}
```

## #1.3 Functions 1

- 곱셈 함수 만들기
- 주의할 점: 매개변수 a와 b 각각의 type을 명시해주어야 한다.
- 물론 `multiply`라는 함수가 int type의 값을 return할 것이라고도 명시해줘야 한다. ⇒ 안하면 Error

```go
 package main

import "fmt"

func multiply(a int, b int) int {
	// 조금 더 간단하게 쓰고 싶다면 (a, b int)라고 써줘도 동작함
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}
```

- Go 함수의 특징: multiple return 값을 갖는다.

```go
package main

import (
	"fmt"
	"strings"
)

// Go 함수의 특징: multiple return 값을 갖는다.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("nico")
  // 두 return 값 중 하나만 받고 싶다면 "_"를 쓰면 된다.
	// totlaLength, _ := lenAndUpper
	fmt.Println(totalLength, upperName)
}
```

- 함수의 arguments를 여러 개 받고 싶다면?
- `...`을 함수의 argument와 type 사이에 넣어준다.

```go
package main

import (
	"fmt"
	"strings"
)

// 반복하는 함수
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}
```

## #1.4 Functions 2

- `naked return` 함수: return할 variable을 굳이 명시하지 않아도 된다.

```go
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name) // 새로 만드는 것이 아니기 때문에 :=를 쓰지 않음
	uppercase = strings.ToUpper(name) // 마찬가지로 uppercase는 이미 정의
	return // return 뒤에 뭐라고 (굳이) 써줄 필요가 없다.
}

func main() {
	totalLength, uppercase := lenAndUpper("nico")
	fmt.Println(totalLength, uppercase) // 4 NICO
}
```

- `defer`: 함수의 기능이 끝났을 때 추가적으로 무엇인가 동작하도록 할 수 있다.
- defer는 언제 사용되나?
  - 시스템을 디자인 한다고 생각했을 때, 이미지를 열거나 파일을 생성하고 나서 function의 동작이 끝났을 때 열었던 이미지를 닫던가 생성된 파일을 닫거나 혹은 삭제할 수 있도록 만들 수 있다.
  - 또는 함수의 동작이 끝난 뒤에 API로 요청을 보낸다거나 하는 로직을 만들 수도 있다.

```go
package main

import (
	"fmt"
	"strings"
)

// defer: lenAndUpper 함수가 return하고 난 뒤 실행될 동작을 정의해준다.
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println(("I'm done"))
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, uppercase := lenAndUpper("nico")
	fmt.Println(totalLength, uppercase)
  // 결과
	// I'm done
  // 4 NICO
}
```

## #1.5 for, range, ...args

## #1.6 If with a Twist

## #1.7 Switch

## #1.8 Pointers!

## #1.9 Arrays and Slices

## #1.10 Maps

## #1.11 Structs

