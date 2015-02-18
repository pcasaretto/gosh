package gosh

import "testing"

func TestPlaceholder(t *testing.T) {
	_, items := lex("empty", "")
	item := <-items
	if item.typ != itemEOF {
		t.Errorf("Channel was not closed, received %+v", item)
	}
}
