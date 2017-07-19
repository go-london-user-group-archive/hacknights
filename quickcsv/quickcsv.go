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

	//buf := make([]byte, 0, 1024)

	var output [][]byte

	inQuote := false

	column := []byte{}

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
				//	output = append(output, column)
				//	column = nil
			} else {
				inQuote = true
			}
		case ch == ',' && !inQuote:
			output = append(output, column)
			column = []byte{}
		case ch == eor:
			if inQuote {
				panic("code this bit")
			}
			output = append(output, column)
			cb(output)
			column = []byte{}
			output = nil
		default:
			column = append(column, ch)
			// add to buffer
		}
	}
}
