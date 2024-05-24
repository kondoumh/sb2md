package cmd

import (
	"strings"
	"testing"
)

func TestMod(t *testing.T) {
	original :=`title
hoge
 aaa
 bbb
 ccc

table:foo
 head1	head2
 123	456`

	got := ToMd(strings.Split(original, "\n"), false)

  want :=
`title
=================
hoge
- aaa
- bbb
- ccc

foo

| head1 | head2 |
|:--|:--|
| 123 | 456 |
`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
