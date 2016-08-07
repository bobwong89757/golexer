package golexer

import (
	"fmt"
	"reflect"
	"strconv"
)

type Token struct {
	value   string
	matcher TokenMatcher
	line    int
}

func (self *Token) ToFloat32() float32 {
	v, err := strconv.ParseFloat(self.value, 32)

	if err != nil {
		return 0
	}

	return float32(v)
}

func (self *Token) MatcherName() string {
	return reflect.TypeOf(self.matcher).Elem().Name()
}

func (self *Token) String() string {
	return fmt.Sprintf("line: %d matcher: %s  value:%s", self.line, self.MatcherName(), self.value)
}

func NewToken(m TokenMatcher, tz *Tokenizer, v string) *Token {

	return &Token{
		value:   v,
		line:    tz.Line(),
		matcher: m,
	}
}