
### 서문

템플릿 패키지는 텍스트를 출력하는 데이터 기반 템플릿을 구현한다.

HTML 템플릿을 만들고 싶으면 html/template를 봐야한다. html/template 패키지는 자동으로 특정 공격으로 부터 안전한 템플릿을 생성한다.

템플릿들은 자료구조에 적용시킴으로 실행된다. 템플릿에 있는 표기들은 실행을 제어하고 표시되는 값을 얻기 위한 자료 구조의 요소를 의미한다. 템플릿을 실행시키면, 자료 구조를 읽고, '.'으로 표기된('점'이라고 읽는다) 커서를 진행과정을 따라서 구조의 현제 값에 위치시킨다. 

템플릿에 입력된 텍스트는 UTF-8로 변환된 텍스트의 어떤 형태라고 가능하다. **"액션"** -- 데이터 평가 혹은 제어 구조 --는 "{{", 과 "}}"으로 구분된다. **액션 밖의 모든 텍스트는 변하지 않고 출력된다.**

분석이 끝나면, 템플릿은 병렬적으로 실행될 수 있다, 혹시 병렬 실행이 Writer를 공유한다고해도 결과는 끼워질수 있다. 

여기 간단한 예시는 "17 items are made of wool"을 출력한다. 

 ```go
 type Inventory struct{
	 Material string
	 Count uint
}
sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
if err != nil {panic(err)}
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil {panic(err)}

```

더 복잡한 예시는 아래에서 볼 수 있다. 

### 텍스트와 공백

기본적으로, 액션들 사이의 텍스트는 템플릿이 실행될 때, 문자 그대로 복사된다. 예를 들어서 위의 예시의 문자열 "items are made of"은 프로그램이 실행될 때 그대로 출력된다. 

그러나, 템플릿 소스 코드의 형식을 정돈하는 것을 도와주기 위해서 간단한 기능을 제공한다. 만일 액션의 왼쪽 구별자에  마이너스 표식과 공백이 따라오면 앞에 있는 텍스트에 붙어 있는 공백이 모두 잘려나간다. 비슷하게, 만일 오른쪽 구별자가 공백과 마이너스 표시보다 뒤에 있으면 따라오는 공백은 뒤에 오는 텍스트에서 잘려나간다. 이 자르기 마커에서는 공백을 반드시 필요로 한다. "{{- 3}}" 은 "{{3}}"과 동일 하지만 바로 선행되는 텍스트의 공백을 제거한다. 반면 "{{-3}}"은 -3을 포함하는 액션으로 분석된다.

예를 들어서,  "{{23 -}} < {{- 45}}"을 템플릿으로 실행하면 생성되는 결과는 "23<45" 이다.

이 자르기를 위해서는, 공백의 정의는 Go와 동일해야 한다.

**정리**
1. 액션에 포함되지 않은 텍스트는 그대로 생성된다.
2. 액션에서 마이너스 표시는 공백을 자르는 기능을 한다.
3. 액션의 왼쪽 구별자 다음에 마이너스와 공백을 입력하면 앞에 있는 텍스트의 공백이 잘려진다.
4. 액션의 오른쪽 구별자 앞에 공백과 마이너스를 입력하면 뒤에 있는 텍스트의 공백이 잘려진다.

### 액션

아래는 액션 리스트이다. "Arguments" 와 "piplines"는 평가된 데이터이다.  아래의 각각의 섹션에서 세부적인 정의가 있다. 

 ```go
{{/* 코멘트 */}}
코멘트는 무시되며 줄바꿈이 가능하다. 하지만 네스트할 수 없다. 

{{pipline}}
기본적인 텍스트 형식의 표현(fmt.Print로 출력되는 것과 동일하다). pipeline의 값은 출력 값에 복사된다

{{if pipeline}} T1 {{end}}
만일 pipeline의 값이 없다면, 아무런 것도 생성되지 않는다.
pipline의 값이 있다면, T1이 실행된다. false, 0, nil, interface 값,길이가 0인 배열, 슬라이스, 맵, 문자열은 모두 값이 비어 있는 것으로 정의된다.
점은 영향을 받지 않는다.
{{if pipeline}} T1 {{else}} T0 {{end}}
만일 pipeline의 값이 비어 있다면 T0가 실행된다. 그렇지 않다면 T1이 실행된다. 점은 영향을 받지 않는다. 

{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
if-else 체인를 간단하게 나타내기 위해서,  else 액션에서 바로 if를 포함할 수 있다.

{{range pipeline}} T1 {{end}}
pipeline의 값은 반드시 배열,슬라이스, 맵, 혹은 채널 이어야한다.
만일 pipeline의 길이가 0이면 결과는 없다. 
그렇지 않으면, 점은 배열, 슬라이스, 맵의 연속된 값으로 설정된다. 

{{range pipeline}} T1 {{else}} T0{{end}}
pipeline의 값은 배열, 슬라이스, 맵, 채널이다. 만일 길이가 0이면, T0가 실행된다. 
그렇지 않다면, T1은 연속된 pipline의 요소이다. 

{{break}}
가장 안쪽의 {{range pipeline}} 루프가 조기에 종료된다. 현제의 반복을 중단하고 나머지 반복도 통과한다. 

{{continue}}
가장 안쪽의 {{range pipeline}} 이 멈춘다. 그리고 다음 루프를 실행한다.

{{template "name"}}
특정한 이름의 템플릿이 빈 데이터와 실행된다. 

{{template "name" pipeline}}
특정한 이름의 템플릿이 점을 pipeline으로 설정하고 실행된다.

{{block "name" pipeline}} T1 {{end}}
a block is shorthand for defining a template
{{define "name"}} T1 {{end}}
를 한 이후에 
{{template "name" pipeline}}

{{with pipeline }} T1 {{end}}
만일 pipeline의 값이 비어 있다면 결과가 없다. 
그렇지 않으면, 점은 pipeline의 값으로 설정되고 T1이 실행된다.

{{with pipeline}} T1 {{else}} T0 {{end}}
만일 pipeline의 값이 비어 있다면, T0가 실행된다. 
그렇지 않으면, 점은 pipeline으로 설정된고 T1이 실행된다.

```

**헷갈리는 액션들 예시**

```go
{{with pipeline}} T1 {{end}}
{{with .Material}} {{.}} are made of {{end}}
.Material의 값이 존재하면  점은 .Material의 값으로 설정된다. 

{{if pipeline}} T1 {{end}}
{{if .Material}} {{.}} are made of {{end}}
.Material의 값이 있어도 점은 변경되지 않는다.

```

### 인자

인자는 간단한 값이다. 

- 변수 이름
	- 달러 표시가 앞에 있고 알파벳과 숫자로 이루어짐. 예를 들어 $piOver2
	- 변수의 값으로 평가된다.
- 데이터 필드의 이름
	- 구조체 이어야하고, 앞에 점을 붙여야 한다.
- 매개변수가 없는 매소드 이름
	- 하나 혹은 두개의 값을 리턴해야한다.
	- 리턴 값이 두개인 경우 한개는 에러 값이다.
	- 에러가 생긴 경우, 에러는 Execute의 값으로 전달된다. 
 - 매개변수가 없는 함수 이름
	 - 메소드와 동일하게 작동한다.
	 - 함수는 아래에서 설명한다.
 - 그룹을 만들기 위한 소괄호 ()


### Pipelines

파이프라인은 연속될 수도 있는 일련의 커맨드이다. 커맨드는 간단한 값이나 함수, 메소드이다. 

```go
Argument
	간단한 값
.Method [Argument...]
	마지막에 있는 메소드는 인수를 받을 수 있다.
funtionName [Argument...]
	결과는 함수를 이름과 호출하는 것과 동일하다.
```

**메소드 예시**

```go
package main

import (
	"os"
	"text/template"
)

type Invetory struct {
	Material string
	Count    uint
}

func (i *Invetory) Shout(name string) string {
	return name + "is made by master kim"
}
func main() {
	wools := Invetory{"wool", 12}


	tmpl, err := template.New("trivial").Parse("{{.Shout .Material}}\n")
	//.Shout는 Invetory의 메소드이다. 메소드가 마지막에 위치함으로 인수를 받을 수 있다. 
	// .Material은 .Shout의 인수이다. 
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, &wools)
	if err != nil {
		panic(err)
	}
}

```


### Variables

액션 안에 있는 파이프라인은 변수를 선언할 수 있다. 

```go
$variable := pipeline
```

range 에서 변수를 선언할 수 있다.

```go
range $index, $element := pipeline
```

변수의 스코프는 변수가 선언된 제어 구조의 "end" 까지이다. 만일 제어구조가 없다면 템플릿의 끝이 스코프이다. 변수는 다른 템플릿에 상속되지 않는다.

### Functions

실행이 되면, 함수를 두 곳에서 찾는다. 우선 템플릿에서 찾고, 그 다음에 전역 함수 맵에서 찾는다. 기본적으로, 템플릿에서는 어떠한 함수도 정의되어 있지 않는다. Funcs 메소드를 이용해서 추가를 해야한다. 

미리 정의된 전역 함수는 아래와 같다. 

```go

and
	인수에 대하여 불리안 AND를 실행한다.
	"and x y" 라고 하면 x가 참이면 y를 반환하고 그렇지 않으면 x를 반환한다.
call
	함수를 호출하는 기능을 한다.
	"call .X.Y 1 2"를 Go 표기법으로 나타내면 dot.X.Y(1,2)이다.
html
	탈출한 html 형태의 텍스트를 리턴한다.
index
	슬라이스, 배열, 맵에서 인수를 인덱스로 하는 값을 리턴한다.
	index  x 1 2 3 이라고 하면 x[1][2][3]과 동일하다.
slice
	문자열, 슬라이스, 배열에 대하여 슬라이스를 한다.
	slice x 1 2라고 하면 x[1:2]와 동일하다.
	slice x 2라고 하면 x[2:]와 동일하다.
len
	인수의 길이를 정수로 나타낸다.
not 
	인수의 부정을 리턴한다.
or
	전달받은 인수의 불리안 or를 실행한다.
	or x y 는 x 가 존재할 경우 x 그렇지 않으면 y를 반환한다.
print
	fmt.Sprint
printf
	fmt.Sprintf
urlquery
	URL 쿼리가 가능한 탈출한 텍스트를 반환한다.


```



