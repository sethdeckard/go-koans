package go_koans

import (
	"bytes"
	"io"
	"log"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		if _, err := io.Copy(out, in); err != nil {
			log.Fatal(err)
		}

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		if _, err := io.CopyN(out, in, 5); err != nil {
			log.Fatal(err)
		}

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
