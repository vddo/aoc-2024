package importdata

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Import(path string) (*[]string, error) {
	data := make([]string, 0, 20)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	firstTime, rowLen := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if firstTime == 0 {
			rowLen = len(scanner.Text())
			firstTime++
		}

		row := scanner.Text()
		if scanner.Err() != nil {
			return nil, fmt.Errorf("reading file: %w", err)
		}
		if len(row) == 0 {
			continue
		}
		if len(row) != rowLen {
			return nil, errors.New("row are not of equal length")
		}

		data = append(data, row)
	}

	return &data, nil
}
