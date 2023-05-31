## Go Doc Comments

고의 패키지를 이용해서 문서를 만들수 있다.

### Packages

패키지를 설명하는 코멘트를 사용하여 패키지를 소개한는 것이 좋다. 

### Commands

패키지 커멘트와 달리 프로그램의 동작을 설명한다. 첫 문장은 프로그램의 이름으로 시작하는 것이 관례이다. 또한  `semantic linefeed`, 마침표 마다 줄바꿈을 하면 git diff 를 읽기 편해진다. 

블록 코멘트 에서 들여 쓰기를 하면 미리 설정된 형식의 문자로 출력이 된다. 

### Types

타입의 코멘트는 해당하는 타입의 인스턴스가 무엇을 나타내고 무엇을 제공하는지 설명해야한다. 

 ```
 // A Reader serves content from a ZIP archive
 type Reader struct{
 ...
 }
```


타입이 여러개의 고루틴을 받을 수 있으면 명시한다. 

```
// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines,
// except for configuration methods, such as Longest.
type Regexp struct {
    ...
}

```

zero-value 의 의미가 특이하면 명시한다.

```
package bytes

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
    ...
}
```

스트럭트의 필드 값이 추출(export)이 될 수 있으면 설명을 한다.

```
package io

// A LimitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0.
type LimitedReader struct {
    R   Reader // underlying reader
    N   int64  // max bytes remaining
}
```

타입의 코멘트는 선언된 심볼을 지명하는 완전한 문장으로 시작해야한다. 예를 들어서  "Buffer is ....", "Reader is ...".


### Funcs 

함수의 코멘트는 함수가 무엇을 리턴하는지 설명해야한다. 혹은 부수효과를 만드는 함수면 어떤 기능을 하는지 설명한다. 

함수의 결과가 여러 개면, 결과에 이름을 정해주는 것이 설명하기 좋다.
```
package io

// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the total number of bytes
// written and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
func Copy(dst Writer, src Reader) (n int64, err error) {
    ...
}
```

