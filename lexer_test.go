package gosh

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lexer", func() {
	It("parses empty strings", func() {
		_, items := lex("")
		var result []item
		for item := range items {
			result = append(result, item)
		}
		Expect(result).To(HaveLen(1))
		Expect(result[0].typ).To(Equal(itemEOF))
	})

	It("parses basic commands", func() {
		_, items := lex("/bin/ls -la")
		var result []item
		for item := range items {
			result = append(result, item)
		}
		Expect(result).To(HaveLen(4))
		Expect(result[0].typ).To(Equal(itemIdentifier))
		Expect(result[0].val).To(Equal("/bin/ls"))
		Expect(result[1].typ).To(Equal(itemSpace))
		Expect(result[1].val).To(Equal(" "))
		Expect(result[2].typ).To(Equal(itemIdentifier))
		Expect(result[2].val).To(Equal("-la"))
		Expect(result[3].typ).To(Equal(itemEOF))
	})

	It("parses basic quotes", func() {
		_, items := lex("echo \"string argument\"")
		var result []item
		for item := range items {
			result = append(result, item)
		}
		Expect(result).To(HaveLen(4))
		Expect(result[0].typ).To(Equal(itemIdentifier))
		Expect(result[0].val).To(Equal("echo"))
		Expect(result[1].typ).To(Equal(itemSpace))
		Expect(result[1].val).To(Equal(" "))
		Expect(result[2].typ).To(Equal(itemString))
		Expect(result[2].val).To(Equal("\"string argument\""))
		Expect(result[3].typ).To(Equal(itemEOF))
	})

	It("parses single quotes", func() {
		_, items := lex("echo 'string argument'")
		var result []item
		for item := range items {
			result = append(result, item)
		}
		Expect(result).To(HaveLen(4))
		Expect(result[0].typ).To(Equal(itemIdentifier))
		Expect(result[0].val).To(Equal("echo"))
		Expect(result[1].typ).To(Equal(itemSpace))
		Expect(result[1].val).To(Equal(" "))
		Expect(result[2].typ).To(Equal(itemString))
		Expect(result[2].val).To(Equal("'string argument'"))
		Expect(result[3].typ).To(Equal(itemEOF))
	})
})
