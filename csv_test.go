package csv

import (
	"bytes"
	"encoding/csv"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JoinPath_Extention(t *testing.T) {
	assert.Equal(t, "join-path.csv", joinPath(".", "join-path"))
	assert.Equal(t, "join-path.csv", joinPath(".", "join-path.csv"))
}

func Test_JoinPath_Relation(t *testing.T) {
	assert.Equal(t, "join-path.csv", joinPath(".", "join-path"))
	assert.Equal(t, "join-path.csv", joinPath("./", "join-path"))
	assert.Equal(t, "../join-path.csv", joinPath("../", "join-path"))
}

func Test_JoinPath_Absolute(t *testing.T) {
	assert.Equal(t, "/join-path.csv", joinPath("/", "join-path"))
}

func Test_CreateFile(t *testing.T) {
	var (
		rootDir  = "."
		fileName = "create"
		data     = ""
		path     = joinPath(rootDir, fileName)
	)

	// create
	assert.NoError(t, createFile(path, data))

	// check exist
	_, err := os.Stat(path)
	assert.NoError(t, err)

	// delete
	assert.NoError(t, os.Remove(path))
}

func Test_DeleteFile(t *testing.T) {
	var (
		rootDir  = "."
		fileName = "delete"
		path     = joinPath(rootDir, fileName)
	)

	// create
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	assert.NoError(t, err)
	defer file.Close()

	// delete
	assert.NoError(t, deleteFile(path))

	// check exist
	_, err = os.Stat(path)
	assert.Error(t, err)
}

func Test_Read(t *testing.T) {
	var (
		data = "1,2,3,4"
	)

	values, err := read(strings.NewReader(data))
	if assert.NoError(t, err) {
		assert.Equal(t, "1", values[0][0])
		assert.Equal(t, "2", values[0][1])
		assert.Equal(t, "3", values[0][2])
		assert.Equal(t, "4", values[0][3])
	}
}

func Test_Read_From_File(t *testing.T) {
	var (
		rootDir  = "."
		fileName = "read"
		data     = "1,2,3,4"
		path     = joinPath(rootDir, fileName)
	)

	// create & append data
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if assert.NoError(t, err) {
		defer f.Close()
		f.WriteString(data)
	}

	// read
	values, err := readFromFile(path)
	if assert.NoError(t, err) {
		assert.Equal(t, "1", values[0][0])
		assert.Equal(t, "2", values[0][1])
		assert.Equal(t, "3", values[0][2])
		assert.Equal(t, "4", values[0][3])
	}

	// delete
	assert.NoError(t, os.Remove(path))
}

func Test_Write(t *testing.T) {
	var (
		values = [][]string{
			[]string{"1", "2", "3", "4"},
		}
	)
	buf := &bytes.Buffer{}
	write(buf, values)
	assert.Equal(t, "1,2,3,4\n", buf.String())
}

func Test_Write_To_File(t *testing.T) {
	var (
		rootDir  = "."
		fileName = "write"
		path     = joinPath(rootDir, fileName)
		values   = [][]string{
			[]string{"1", "2", "3", "4"},
		}
	)
	// write
	assert.NoError(t, writeToFile(path, values))

	// check
	f, err := os.Open(path)
	if assert.NoError(t, err) {
		defer f.Close()
		rows, err := csv.NewReader(f).ReadAll()
		if assert.NoError(t, err) {
			assert.Equal(t, values, rows)
		}
	}

	// delete
	assert.NoError(t, os.Remove(path))
}
