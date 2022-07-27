# Go에서 byte와 rune의 차이

## byte
* byte는 uint8의 alias이다. 
* 따라서 부호없는(unsigned) 정수를 0~255까지 저장할 수 있다. 

## rune
* rune은 int32의 alias이다.
* 0과 1이 32개 있는 32비트 자료형이다.
* 따라서 rune은 문자이며, UTF-8 문자 인코딩 스킴을 따른다.