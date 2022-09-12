<h3>진행 0907</h3>

 - 함수 항상 call by value
 - package 밖에서 package 안에 있는 호출.
  대문자로 시작
 - a, b = b, a 가능
 - golang은 컴파일 단계도 거치고 강타입언어.. ide를 쓰는게 좋다 -> 컴파일 단계에서 문제가 제거된다.
 - 첫글자 대문자는 문법. 패키지간 호출
 - 함수에서 변수명 '지정'해서 리턴해줄 때는
 반환값이 '하나'더라도 괄호로 묶어줘야 한다.


<h4>~11장</h3>

 - 변수 선언 이후 사용하지 않는다면 build error가 발생한다.  
  warning이 아닌, error.
 - 실수 연산 주의  
  Nextafter(a, b): 1bit 차이무시용도.  
  a에서 b로 1bit 만큼 이동  
  math/big 이라는 패키지도 존재. 보다 정확한 연산용도. big.Float 타입
 - 변수 swap 가능  
  a, b = b, a
 - 명시적 변수 return -> 코드읽기 쉽지 않을까  
  함수 선언부에 return variable 선언, 함수 종료시 해당 변수들 return
 - enum 개념으로 iota  
  같은 const scope 내에서 iota가 사용될 때마다, 0, 1, 2, 3, ... 값으로 사용  
  iota는 같은 라인에서는 같은 값으로 취급
  
```
 const (
    a int = iota = 0 // 0으로 시작
    b int = iota = 1
    c int = iota * iota // 2 * 2 = 4
    d int = iota * iota * iota // 3 * 3 * 3 = 27
 )
```

 - javascript의 switch 구문은 case 안에 break를 적어주어야 하지만  
  golang은 그런거 안해줘도 된다.. 좋다..  
  fallthrough 쓰면 바로 다음 case문도 실행하지만 이걸 굳이 쓸까?
 - 다중 반복문에서 중첩된 반복문 한번에 탈출 가능 : Label  
  권장하지는 않는다고..


<h4>12장</h4>

- 배열 초기화 시에는 배열 크기와 함께 중괄호를 꼭 써주어야한다.

```
a := [5]int{1,2,3,4,5}
```
- 단 여기서 [5]는 [...]으로 갯수 생략, 뒤의 배열 갯수로 된다.
- 생략하는 것이기 때문에 아래와 같은 상황에서는 에러 발생  
동적할당 같은 개념이 아니라, 뒤의 갯수만큼 그대로 캐스팅되는 것으로 보임

```
var a [6]int = [...]int{1,2,3,4,5}
// cannot use [...]int{…} (value of type [5]int) as type [6]int in variable declaration
```
- python에서처럼 -1 인덱스 참조는 불가능하다
- C에서 변수 값을 리스트 크기로 선언하는 것은 동적할당만 가능하듯이, go도 마찬가지로 변수는 선언할 때 사용 불가
- 하고싶다면 len(a)-1 형태로 사용
- for loop의 iterate 같은 경우에는 python과 다르게 index, value 두개를 바탕으로 순회

```
for i := range a {
  fmt.Println(i)
}
// 0\n1\n2\n3\n4

for i, v := range a {
  fmt.Println(i, v)
}
// 0 1\n1 2\n2 3\n3 4\n4 5 
```

- 불필요한 경우 for _, v:= range a 식으로도 사용(blank identifier)  
i, v 라고 선언해두고 i를 호출하지 않으면 에러 발생. _ 로 사용  
값을 무시하겠다는 의미로 별도의 메모리 할당을 하지 않는 것으로 보임
- C에서는 포맷을 지정하고 포인터 개념이 들어가서  
다중배열에서 모두 탐색하지 않으면 포인터가 출력되지만  
go에서는 출력포맷을 명시하지 않게 되면 리스트를 출력함  
대충 C와 python의 중간정도 형태..


<h4>13장</h4>

- struct 초기화 시 여러 줄 초기화
- 마지막 줄 값 뒤에 꼭 쉼표를 달아주세요 라고 하기보단,  
마지막 값과 닫음 중괄호} 가 같은 줄에 위치하지 않으면 쉼표가 필요하다.
- 같은 줄에 위치할 때 쉼표를 써도 vscode에서 파일 저장 시에는 자동으로 마지막 쉼표를 없애주는데,  
vi 로 그냥 파일편집해서 쉼표를 남기게된다고 하더라도 잘 수행된다.  
아래 모든 경우가 정상적으로 수행

```
a := A{1,2,3}
a := A{1,2
      3}
a := A{1,2
      3,}
a := A{
  1, 2, 3,
}
```

- 당연히 인덱스는 안되고, 필드명으로 참조
- 구조체 안에 구조체가 들어간 경우에도 필드명 하나로 바로 참조가 가능
- 필드명이 겹치는 경우에는 가장 상위로 참조

- 메모리 패딩 잘하자..
- 구조체에서 가장 큰 크기의 변수에 맞추어서  
그 크기에 맞도록 몰아서 배치


<h4>14장</h4>

- 포인터 구조체 사용의 경우 golang에서는 그냥 바로 참조해도됨  
283페이지 아래
- 스택메모리 / 힙메모리 저장 위치는 코드 보고 파악


<h4>15~16장</h4>

- 백쿼트 (`) 이용하는건 js와 유사한 점
  - 여러 줄로 문자열 묶기 가능
  - 단순히 여러 줄로 묶는 것 뿐만 아니라 \n 이 들어가는 것
  - 들여쓰기의 경우 그대로 반영이 되어, 2번째 줄부터는 inline을 모두 없애야한다.
  - \n과 같이 escape 문자를 이용해도, 그대로 출력된다
- 단순히 여러 줄로 string을 표현하고 싶을 때, c와는 다르게 \ + {Enter}는 되지 않는다.
- 여러 줄로 표현하고 싶으면 + 를 통해 여러 줄에 걸쳐서 더하자
  - 단, + 기호는 string 오른쪽에 붙어야한다.

  ```
  "a" + 
  "b" // "ab"
  "a"
  + "b" // 불가
  ```

- 한글 등의 영어, 특수문자 이외 문자 사용시에는 string 이 아닌 rune을 이용해 파싱
- 변수를 통한 문자열 대입의 경우 헤더 정보를 복사하는 것으로,  
값만 복사되는 것이 아닌, 참조하는 메모리 주소까지 복사된다.
- string 자체는 불변. 중간을 임의로 변경할 수 없다  
이 불변원칙 때문에, 복제 이외에 string이 새로 생길 때마다 메모리가 계속 할당되어 메모리 낭비  
string.Builder 사용해서 합연산 하는것으로..

- package import 에서 underscore의 사용  
https://stackoverflow.com/questions/26972615/a-use-case-for-importing-with-blank-identifier-in-golang