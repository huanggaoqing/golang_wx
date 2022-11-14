package tools

import (
	"bufio"
	"io"
)

func ReadAll(r io.Reader) ([]byte, error) {
	read := bufio.NewReader(r)
	result := make([]byte, 0)
	for {
		b, readErr := read.ReadByte()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return nil, readErr
		}
		result = append(result, b)
	}
	return result, nil
}
