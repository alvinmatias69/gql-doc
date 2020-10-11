package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func write(data, path string) error {
	if len(path) == 0 {
		fmt.Println(data)
		return nil
	}

	log.Info().Msgf("Opening output file: %s", path)
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

	log.Info().Msgf("Writing %d lines into file", n)
	return w.Flush()
}
