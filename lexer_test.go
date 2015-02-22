package gosh

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type lexTest struct {
	name  string
	input string
	items []item
}

var (
	tEOF   = item{itemEOF, 0, ""}
	tSpace = item{itemSpace, 0, " "}
)

var lexTests = []lexTest{
	{"empty", "", []item{tEOF}},
	{"basic commands", "/bin/ls -la", []item{
		{itemIdentifier, 0, "/bin/ls"},
		tSpace,
		{itemIdentifier, 0, "-la"},
		tEOF,
		},
	},
	{"basic quoted", "echo \"string argument\"", []item{
		{itemIdentifier, 0, "echo"},
		tSpace,
		{itemString, 0, "\"string argument\""},
		tEOF,
		},
	},
	{"single quoted", "echo 'string argument'", []item{
		{itemIdentifier, 0, "echo"},
		tSpace,
		{itemString, 0, "'string argument'"},
		tEOF,
		},
	},
}

var _ = Describe("Lexer", func() {
	Describe("parses case:", func() {
		for _, test := range lexTests {
			It(test.name, func() {
				items := collect(&test)
				Expect(items).To(HaveLen(len(test.items)))
				for i := range items {
					Expect(items[i].typ).To(Equal(test.items[i].typ))
					Expect(items[i].val).To(Equal(test.items[i].val))
				}
			})
		}
	})
})

// collect gathers the emitted items into a slice.
func collect(t *lexTest) (items []item) {
	_, ch := lex(t.input)
	for item := range ch {
		items = append(items, item)
		if item.typ == itemEOF {
			break
		}
	}
	return
}
