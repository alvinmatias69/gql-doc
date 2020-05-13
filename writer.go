package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func write(data, path string) error {
	if path == "." {
		fmt.Println(data)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	n, err := w.WriteString(data)
	if err != nil {
		return err
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	log.Printf("success write %d bytes in %s\n", n, path)
	return nil
}
