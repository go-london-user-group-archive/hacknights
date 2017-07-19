package quickcsv

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleUnquoted(t *testing.T) {
	assert := assert.New(t)

	in := `a,b,c
dasd,e,,f
1,2,3
`
	r := strings.NewReader(in)

	var output [][][]byte

	err := Parse(r, ',', '\n', func(data [][]byte) bool {
		output = append(output, data)
		return true
	})
	assert.NoError(err)

	assert.Equal([][][]byte{
		[][]byte{[]byte("a"), []byte("b"), []byte("c")},
		[][]byte{[]byte("dasd"), []byte("e"), []byte(""), []byte("f")},
		[][]byte{[]byte("1"), []byte("2"), []byte("3")},
	}, output)
}

func TestSimpleQuoted(t *testing.T) {
	assert := assert.New(t)

	in := `"a","b","c"
"dasd","e,","f"
"1","2","3"
`
	r := strings.NewReader(in)

	var output [][][]byte

	err := Parse(r, ',', '\n', func(data [][]byte) bool {
		output = append(output, data)
		return true
	})
	assert.NoError(err)

	assert.Equal([][][]byte{
		[][]byte{[]byte("a"), []byte("b"), []byte("c")},
		[][]byte{[]byte("dasd"), []byte("e,"), []byte("f")},
		[][]byte{[]byte("1"), []byte("2"), []byte("3")},
	}, output)
}
