# chapter2 variable

## 변수 선언

`int age;`

## 변수 초기화

`int age = 0`

## 리터럴

숫자를 할당할 때, 0을 앞에 입력하면 8진수로 해석을 하고,
0x 를 하면 16진법, 0b 를 하면 2진법으로 해석을 한다.

long 타입을 선언할 때는 L 접미사를 표기해야한다.
float 타입을 선언할 때는 f 접미사를 표기해야한다.

## 상수 선언

상수를 선언 할 때는 final를 붙인다.
`final int MAX_SPEED`

## printf

이진법을 문자열로 출력할때
`Intger.toBinaryString()`

문자를 숫자로 출력할 때는 형변환을 해야한다.
`(int) A`

지시자 앞에 숫자를 입력하면 공백을 만들수 있다.
`%10s`

## 진수 바꾸기

2진수를 8진수로 변환하려면, 2진수를 뒤에서 부터 3자리씩 끊어서 그에 새당하는 8진수로 바꾸면 된다.

```java

class Hello {
	public static void main(String[] args) {

		int bin = 0b1010101100;
		System.out.printf("%d%n", bin);
		System.out.printf("%x%n", bin); //16진수
		System.out.printf("%o%n", bin); //8진수
	}
}

```


## 정수형간의 형변환을

큰 타입에서 작은 타입으로 형변환을 하면 값손실이 발생한다.
