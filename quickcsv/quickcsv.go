package quickcsv

import (
	"bufio"
	"io"
)

type Callback func([][]byte) bool

func Parse(
	r io.Reader,
	sep byte,
	eor byte,
	cb Callback,
) error {
	br := bufio.NewReader(r)

	buf := make([]byte, 0, 1024)

	output := make([][]byte, 0, 64)

	inQuote := false

	var lastIndex int
	var index int

	for {
		ch, err := br.ReadByte()
		if err != nil {
			if err == io.EOF {
				if inQuote {
					//error here
				}
				return nil
			}
			return err
		}
		switch {
		case ch == '"':
			if inQuote {
				// end of field
				inQuote = false
			} else {
				inQuote = true
			}
		case ch == ',' && !inQuote:
			output = append(output, buf[lastIndex:index])
			lastIndex = index
		case ch == eor:
			if inQuote {
				panic("code this bit")
			}
			output = append(output, buf[lastIndex:index])
			//lastIndex = index
			cb(output)
			index = 0
			lastIndex = 0
			output = output[0:0]
			buf = buf[0:0]
		default:
			buf = append(buf, ch)
			index++
		}
	}
}
