package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func joinPath(rootDir, fileName string) string {
	if strings.Index(fileName, ".csv") > -1 {
		return filepath.Join(rootDir, fileName)
	}
	return filepath.Join(rootDir, fileName+".csv")
}

func createFile(path, data string) error {
	r := strings.NewReader(data)
	values, err := read(r)
	if err != nil {
		return err
	}
	return writeToFile(path, values)
}

func deleteFile(path string) error {
	return os.Remove(path)
}

func readFromFile(path string) ([][]string, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := r.Close()
		if cerr == nil {
			return
		}
		err = fmt.Errorf("Failed to close: %v, the original error was %v", cerr, err)
	}()
	return read(r)
}

func read(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)

	// Set FieldsPerRecord to a negative to avoid "wrong number of fields in line" error
	// https://golang.org/pkg/encoding/csv/
	reader.FieldsPerRecord = -1

	values, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return values, err
}

func writeToFile(path string, values [][]string) error {
	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer func() {
		cerr := w.Close()
		if cerr == nil {
			return
		}
		err = fmt.Errorf("Failed to close: %v, the original error was %v", cerr, err)
	}()
	return write(w, values)
}

func write(w io.Writer, values [][]string) error {
	return csv.NewWriter(w).WriteAll(values)
}
