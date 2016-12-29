package golexer

import (
	"fmt"
	"reflect"
	"strconv"
)

type Token struct {
	value   string
	raw     string
	matcher TokenMatcher
	line    int // 行
	index   int // 列
}

func (self *Token) Line() int {
	if self == nil {
		return 0
	}

	return self.line
}

func (self *Token) Index() int {
	if self == nil {
		return 0
	}

	return self.index
}

func (self *Token) MatcherID() int {

	if self == nil || self.matcher == nil {
		return 0
	}

	return self.matcher.ID()
}

func (self Token) Value() string {
	if self.matcher == nil {
		return ""
	}

	return self.value
}

func (self Token) Raw() string {
	if self.matcher == nil {
		return ""
	}

	return self.raw
}

func (self *Token) ToFloat32() float32 {
	v, err := strconv.ParseFloat(self.value, 32)

	if err != nil {
		return 0
	}

	return float32(v)
}

func (self *Token) ToInt32() int32 {
	v, err := strconv.ParseInt(self.value, 10, 32)

	if err != nil {
		return 0
	}

	return int32(v)
}

func (self *Token) ToInt() int {
	v, err := strconv.ParseInt(self.value, 10, 32)

	if err != nil {
		return 0
	}

	return int(v)
}

func (self *Token) ToInt64() int64 {
	v, err := strconv.ParseInt(self.value, 10, 64)

	if err != nil {
		return 0
	}

	return v
}

func (self *Token) MatcherName() string {
	if self == nil || self.matcher == nil {
		return ""
	}

	return reflect.TypeOf(self.matcher).Elem().Name()
}

func (self *Token) MatcherString() string {
	if self == nil || self.matcher == nil {
		return ""
	}

	return self.matcher.String()
}

func (self *Token) String() string {

	if self == nil {
		return ""
	}

	return fmt.Sprintf("line: %d id:%d matcher: %s  value:%s", self.line, self.MatcherID(), self.MatcherName(), self.value)
}

func NewToken(m TokenMatcher, tz *Tokenizer, v string, raw string) *Token {

	if raw == "" {
		raw = v
	}

	self := &Token{
		value:   v,
		raw:     raw,
		matcher: m,
	}

	if tz != nil {
		self.line = tz.Line()
		self.index = tz.Index()
	}

	return self
}
