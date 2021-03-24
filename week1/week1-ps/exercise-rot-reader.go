package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(storedBytes []byte) (int, error) {
	storedBeginIndx := 0

	UPPERCASE_LETTERS := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWERCASE_LETTERS := "abcdefghijklmnopqrstuvwxyz"

	for {
		var bytesAtATime = make([]byte, 8)
		nBytesRead, err := reader.r.Read(bytesAtATime)

		if err == io.EOF {
			break
		}

		for i := 0; i < nBytesRead; i++ {
			ch := bytesAtATime[i]

			if UPPERCASE_LETTERS[0] <= ch && ch <= UPPERCASE_LETTERS[25] {
				chInc := ch + 13
				if UPPERCASE_LETTERS[25] < chInc {
					chInc = chInc % UPPERCASE_LETTERS[25]
					storedBytes[storedBeginIndx] = UPPERCASE_LETTERS[0] + chInc - 1
				} else {
					storedBytes[storedBeginIndx] = chInc
				}
			} else if LOWERCASE_LETTERS[0] <= ch && ch <= LOWERCASE_LETTERS[25] {
				chInc := ch + 13
				if LOWERCASE_LETTERS[25] < chInc {
					chInc = chInc % LOWERCASE_LETTERS[25]
					storedBytes[storedBeginIndx] = LOWERCASE_LETTERS[0] + chInc - 1
				} else {
					storedBytes[storedBeginIndx] = chInc
				}
			} else {
				storedBytes[storedBeginIndx] = ch
			}

			storedBeginIndx++
		}
	}

	return storedBeginIndx + 1, io.EOF
}

func main() {
	s := strings.NewReader("ALbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
