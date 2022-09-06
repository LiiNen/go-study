h3. 특이사항
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