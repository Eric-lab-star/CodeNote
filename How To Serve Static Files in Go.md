고에서는 쉽게 http 서버를 만들수 있다. 

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", HomeHandler())
	http.ListenAndServe("localhost:4000", nil)
}

type homeHandler func(w http.ResponseWriter, r *http.Request)

func (f homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func HomeHandler() homeHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home")
	}
}

```

`http.Handle`의 정의는 아래와 같다.
`func Handle(pattern string, handler Handler)`

첫 번째 인수는 문자열, 두 번째 인수는 `Handler` 타입을 전달해야 한다.

`http.Handler` 는 인터페이스 타입으로 
`ServeHTTP(ResponseWriter, *Request)`
를 메소드로 갖는 타입이다.

 `Handler` 를 만족시켜 주기 위해서 `homeHandler` 함수 타입을 만들고 `ServeHTTP` 메소드도 구현해 주었다. 다음으로 `HomeHandler`는 `homeHandler`를 리턴하도록하여 `http.Handle` 의 두 번째 인수로 넣어주었다. 

  이렇게 `ServeHTTP` 를 메소드로 가지는 함수 타입을 매번 만들기는 쉽지 않고, 매번 만들면 코드가 더러워 질수 있다. 그래서 `http` 패키지에서는 `homeHandler` 의 기능을 하는 타입을 제공한다. 

### 간결한 방법

 ```go
func main() {
	http.Handle("/", http.HandlerFunc(HomeHandler))
	http.ListenAndServe("localhost:4000", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home ")
}

```

`http.HandlerFunc`을 이용하면 코드가 더욱 간결해진다.
하지만 여기서 더 간단하게 `http.HandlerFunc` 또한 줄일 수 있다. ㅓ

### 더욱 간결한 방법

```go
func main(){
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe("localhost:4000",nil)
}
```

위의 코드와 같이 `http.Handle` 대신에 `http.HandleFunc`을 이용하면 `http.HandlerFunc` 으로 타입을 변환하지 않아도 된다. `http.HandleFunc`가 내부에서 동일한 과정을 수행하기 때문이다.

그래서 고에서는 `http.HandleFunc`가 가장 간편하기 때문에 `http.HandlFunc`이 주로 사용된다. 간혹 `http.RedirectHandler`와 같이 `http.Handler` 타입을 반환하는 경우가 있는데 이럴 때에는 `ServeHTTP` 를 호출하지 않고 전달을 하면 된다.



```go
http.HandleFunc("/", http.RedirectHandler("/home", 301).ServeHTTP)
```

### Handle vs Handler

또한 `Handle` 과 `Handler`가 한 글자 차이나기 때문에 헷갈리는데 고에서는 메소드가 하나인 인터페이스는 er로 끝나게 하는 것을 좋아하니, `Handler`는 타입, `Handle`은 함수라고 기억하면 편리하다.



### 파일서버 만들기 

"파일 서버"는 css, js 그리고 이미지 파일과 같이 미리 서버에 저장되어 있는 파일을 전달해주는  역할을 한다. 

 ```go
 package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe("localhost:4000", nil)

}

```

`http.FileServer`는 서버를 http로 내보내는 역활을 한다. 
위의 예시에서는 "static" 이라는 폴더(디렉토리)를 HTTP로 내보낸다. 하지만 `http.Dir`로 만들어서 `http.FileServer`에 전달을 해야한다. `http.Dir`는 string 타입지만 `Open` 메소드를 가지고 있기 때문에, `http.FileSystem` 인터페이스를 만족시키기 때문이다.

대부분의 경우에서 "/" 의 위치에서 css 파일이나 js 파일을 여는 경우는 없다. "static" 이라는 별개의 폴더를 루트로 하는 경우가 일반적이다. 파일 서버를 "/" 이외의 다른 라우터를 통해서 전달을 하려면 `http.StripPrefix` 를 이용해야한다. 

```go
func main() {

	http.Handle("/folder/", http.StripPrefix("/folder", http.FileServer(http.Dir("static"))))
	http.ListenAndServe("localhost:4000", nil)
}

```

위의 코드에서는 "/folder" 라는 라우트에서 파일을 전달한다. 이때 반드시 마지막에도 "/"를 적어야한다. "/folder/" 와 같이 적어야 "/folder/home.css" 와 같은 파일을 열수있다. 

### Chi - router FileServer

고에서 좋은 라우터를 제공하지만 여러가지 편의사항을 위해서 다른 패키지에서 라우터를 사용하기도 한다. 그중에서 chi 라우터를 이용해서 파일서버를 만드는 방법을 알아보자. 

고에서 파일서버를 만드는 방법과 동일하지만 chi 에서는 context를 이용한다.

```go

func fileServer(r chi.Router, path string) {
	// url에 {}*가 있는지 확인후 거르기
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}
	//url의 마지막 글자가 "/"가 아닌경우 "/"를 붙여서 리다이렉트 시킴
	if path != "/" && path[len(path)-1] != '/' {
		fmt.Println(path)
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filesDir := http.Dir(filepath.Join(wd, "static"))

	path += "*"
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})
}
```

먼저 나오는 두개의 if 제어문은 올바른 url 를 만들어주는 역할을 하고 마지막에 나오는 `r.Get` 함수를 살펴보면 고에서 기본으로 제공하는 `http` 패키지의 파일서버와 매우 유사하다는 것을 알수있다. 

[context pkg in go](https://go.dev/blog/context)

차이점은 context를 이용했다는 것인데 context를 이용하면 request-scope 변수를 설정할 수 있고 고루틴을 효과적으로 관리를 해서 메모리 누수를 막을 수 있다는 장점이 있다. 

*추가정보*
path의 마지막이 "/"고 끝나지 않을 때,  r.Get()을 두 번사용한다. 이것은  path에 / 를 붙여서 올바르게 리다이레그 시켜주기 위함이다. 