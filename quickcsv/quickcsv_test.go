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

	var output [][]string

	err := Parse(r, ',', '\n', func(data [][]byte) bool {
		var row []string
		for _, cell := range data {
			row = append(row, string(cell))
		}
		output = append(output, row)
		return true
	})
	assert.NoError(err)

	assert.Equal([][]string{
		[]string{"a", "b", "c"},
		[]string{"dasd", "e", "", "f"},
		[]string{"1", "2", "3"},
	}, output)
}

func TestSimpleQuoted(t *testing.T) {
	assert := assert.New(t)

	in := `"a","b","c"
"dasd","e,","f"
"1","2","3"
`
	r := strings.NewReader(in)

	var output [][]string

	err := Parse(r, ',', '\n', func(data [][]byte) bool {
		var row []string
		for _, cell := range data {
			row = append(row, string(cell))
		}
		output = append(output, row)
		return true
	})
	assert.NoError(err)

	assert.Equal([][]string{
		[]string{"a", "b", "c"},
		[]string{"dasd", "e,", "f"},
		[]string{"1", "2", "3"},
	}, output)
}
