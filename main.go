package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		defer f.Close()

		var str string
		buff := make([]byte, 8)

		for {
			n, err := f.Read(buff)

			if err == io.EOF {
				if str != "" {
					out <- str
				}
				break
			}

			if err != nil {
				return
			}

			data := buff[:n]

			for {
				i := bytes.IndexByte(data, '\n')
				if i == -1 {
					str += string(data)
					break
				}

				str += string(data[:i])
				out <- str
				str = ""
				data = data[i+1:]
			}
		}
	}()

	return out
}

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := getLinesChannel(file)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}

