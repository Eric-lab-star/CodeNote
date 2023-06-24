템플릿 패키지와 chi 패키지를 이용해서 간단한 웹 사이트를 만들수 있다.

### 템플릿 만들기 

`html/template` 패키지를 이용하면 중복되는 형식을 간결하게 만들수 있고, 전체적인 HTML이 깨끗한 느낌을 준다. 

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
    <link rel="stylesheet" href="static/blog.css" />
  </head>
  <body>
    {{template "header"}}
    {{template "book" .}}
  </body>
</html>

```

먼저 blog.html이라는 파일에 위와 같이 작성해 주었다. 여기서 봐야하는 것은 css 파일의 위치와 body 안에 작성된 코드이다.  css 파일은 파일서버를 이용해서 웹에 공개를 해주는 것이고, body 안의 코드는 다른 템플릿을 실행시키는 코드이다. 템플릿은 여러개를 만들 수 있고 네스팅할 수 있다. body 안엔는 "header" 와 "book" 이 있는데 이것들은 네스팅 된 템플릿이다. 이런 템플릿을 만들기 위해서는 위에서 작성한 템플릿과는 조금 다른 형식의 템플릿이 필요하다. 아래의 코드를 먼저 살펴보자.

```html
{{define "book"}}
{{range $book := .}}
<div id="book">
  <div id="title">{{$book.Title}}</div>
  <div>{{$book.Review}}</div>
  <span id="count">Page Count: {{$book.PageCount}}</span>
  <div id="category">
    {{range $category := $book.Categories}}
    <div>- {{$category}}</div>
    {{end}}
  </div>
</div>
{{end}}
{{end}}
```

위의 코드는 book.html에 작성된 코드이다. 이 템플릿은 다른 템플릿에서 사용되는 템플릿이기 때문에 템플릿을 시작할 때, `{{define "book"}} {{end}}` 을 먼저 작성한 다음에 이 안에서 템플릿을 작성해야한다. 이렇게 하는 것은 "book" 템플릿을 만드는 역할을 한고 다른 템플릿, blog.html에서 `{{template "book"}}` 으로 이 템플릿을 읽을 수 있다. 그런데 book 템플릿에서는 데이터를 필요로 함으로  blog.html에서는 `{{template "book" . }}` 이라고 적어줌으로 데이터를 전달할 수 있다. 

header 템를릿 또한 동일한 방식으로 만들어 주면 된다. 

### 서버에서 템플릿 읽기

서버에서는 템플릿을 읽고 데이터를  템플릿에 넣어주고 웹에 보여준다. 이 작업을 하기 위해서 chi 패키지와  http 패키지, html/template 패키지를 사용한다. 

먼저 main.go 파일을 만들어 아래와 같이 작성한다.

```go
func main() {
	books := []book{LegacyAndTheDouble, Ghost, TheDeepEnd, WreckingBall, DogManAndCatKid}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler(books))
	http.ListenAndServe("localhost:8000", r)
}

func homeHandler(data []book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("blog.gotmpl", "book.gotmpl", "header.gotmpl")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}
```

변수 books는 임의로 만든 데이터이다. 이 데이터를 템플릿에 전달해 줄 것이다. 
변수 r 은 새로운 chi 라우터를 의미한다. 이 라우터를 `http.ListenAndServe` 의 두번째 변수로 전달을 해주면 localhost:8000 에서 이 라우터를 이용한다.  즉 `http://localhost:8000` 으로 가면  `r.Get("/", homeHandler(books))` 로 인해서  `homHandler(books)` 가 실행된다. 

다음으로  homeHandler은  템플릿을 분석하고 html의 공격에 안전한 파일로 만들어 데이터를 데이터를 넣어서 보여주는 역할을 한다. `r.Get` 이 `http.HandlerFunc` 타입을 인수로 받기 때문에 클로우저를 이용해서 데이터를 넘겨주었다. 이런 방식을 자주 사용된다. 

템플릿을 분석하는 것은 `template.ParseFiles` 를 이용하고, 데이터를 넣어서 실행하는 것은 `template.Execute` 를 이용한다. 

이렇게 하면 템플릿을 만들 수있다. 하지만 css 를 추가하려면 파일서버를 만들어야하는 작업이 필요하다. 

### 파일서버 만들어 CSS 적용하기

먼저 main 함수에 `fileServer(r, "/static")` 을 추가하자. 

```go
func main() {
	books := []book{LegacyAndTheDouble, Ghost, TheDeepEnd, WreckingBall, DogManAndCatKid}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler(books))
	fileServer(r, "/static") // 파일서버 추가
	http.ListenAndServe("localhost:8000", r)
}
```

그 다음  `fileServer` 함수를 만들어 주어야한다. 

 ```go
 func fileServer(r chi.Router, path string) {
	// url에 {}*가 있는지 확인후 거르기
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}
	//url의 마지막 글자가 "/"가 아닌경우 "/"를 붙여서 리다이렉트 시킴
	if path != "/" && path[len(path)-1] != '/' {
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

 이런 과정을 거치면 css 파일을 성공적으로  HTTP에 전송할 수 있다. 

### 완성된 모습
![[템플릿 프로젝트.png]]