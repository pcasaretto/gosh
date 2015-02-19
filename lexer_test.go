package gosh

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lexer", func() {
	It("should work", func()  {
		_, items := lex("empty", "")
		item := <-items
		Expect(item.typ).To(Equal(itemEOF))
	})
})
