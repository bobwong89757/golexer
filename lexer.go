package golexer

import (
	"errors"
)

type TokenMatcher interface {
	Match(*Tokenizer) (*Token, error)
}

type matcherMeta struct {
	m      TokenMatcher
	ignore bool
}

type tokenAndError struct {
	tk  *Token
	err error
}

type Lexer struct {
	matchers []matcherMeta

	comm chan tokenAndError
}

// 添加一个匹配器，如果结果匹配，返回token
func (self *Lexer) AddMatcher(m TokenMatcher) {
	self.matchers = append(self.matchers, matcherMeta{
		m:      m,
		ignore: false,
	})
}

// 添加一个匹配器，如果结果匹配，直接忽略匹配内容
func (self *Lexer) AddIgnoreMatcher(m TokenMatcher) {
	self.matchers = append(self.matchers, matcherMeta{
		m:      m,
		ignore: true,
	})
}

func (self *Lexer) Start(src string) {

	if self.comm != nil {
		close(self.comm)
	}

	self.comm = make(chan tokenAndError)

	go self.tokenWorker(src)
}

func (self *Lexer) Read() (*Token, error) {
	if self.comm == nil {
		return nil, errors.New("call 'Start' first")
	}

	te := <-self.comm

	return te.tk, te.err
}

func (self *Lexer) tokenWorker(src string) {

	tz := NewTokenizer(src)

	if len(self.matchers) > 0 {

		for !tz.EOF() {

			for _, mm := range self.matchers {

				token, err := mm.m.Match(tz)

				if err != nil {
					self.comm <- tokenAndError{nil, err}
					return
				}

				if token == nil {
					continue
				}

				if mm.ignore {
					break
				}

				self.comm <- tokenAndError{token, nil}

				// 重新从matcher开始检查
				break

			}
		}
	}

	self.comm <- tokenAndError{nil, nil}
}

func NewLexer() *Lexer {

	return &Lexer{}

}