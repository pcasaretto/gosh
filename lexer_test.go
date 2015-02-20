package gosh

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lexer", func() {
	It("should parse empty strings successfuly", func() {
		_, items := lex("empty", "")
		var result []item
		for item := range items {
			result = append(result, item)
		}
		Expect(len(result)).To(Equal(1))
		Expect(result[0].typ).To(Equal(itemEOF))
	})
})
