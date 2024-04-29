package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	writer io.Writer
	err    error
}

func main() {
	err := proverbs("writer.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}

	_, sw.err = fmt.Fprintln(sw.writer, s)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()

	sw := safeWriter{writer: f}

	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")

	return sw.err
}
