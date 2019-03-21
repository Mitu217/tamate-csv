package csv

import (
	"context"
	"log"
	"testing"

	"github.com/go-tamate/tamate"
	"github.com/go-tamate/tamate/driver"
	"github.com/stretchr/testify/assert"
)

func Test_GetSchema(t *testing.T) {
	var (
		rootDir  = "./"
		fileName = "getSchema"
		testData = `
			(id),name,age
		`
	)
	path := joinPath(rootDir, fileName)
	err := createFile(path, testData)
	assert.NoError(t, err)
	defer func() {
		cerr := deleteFile(path)
		assert.NoError(t, cerr)
	}()

	ds, err := tamate.Open(driverName, rootDir)
	if assert.NoError(t, err) {
		ctx := context.Background()
		schema, err := ds.GetSchema(ctx, fileName)
		if assert.NoError(t, err) {
			columns := schema.Columns
			assert.Equal(t, driver.ColumnTypeString, columns[0].Type)
			assert.Equal(t, "id", columns[0].Name)
			assert.Equal(t, 0, columns[0].OrdinalPosition)

			assert.Equal(t, driver.ColumnTypeString, columns[1].Type)
			assert.Equal(t, "name", columns[1].Name)
			assert.Equal(t, 1, columns[1].OrdinalPosition)

			assert.Equal(t, driver.ColumnTypeString, columns[2].Type)
			assert.Equal(t, "age", columns[2].Name)
			assert.Equal(t, 2, columns[2].OrdinalPosition)
		}
	}
}

func Test_SetSchema(t *testing.T) {
	var (
		rootDir    = "./"
		fileName   = "setSchema"
		beforeData = `
			(id),name,age
		`
		afterData = `
			(id),name,from
		`
	)
	path := joinPath(rootDir, fileName)
	err := createFile(path, beforeData)
	assert.NoError(t, err)
	defer func() {
		cerr := deleteFile(path)
		assert.NoError(t, cerr)
	}()

	log.Println(afterData)
}

func Test_GetRows(t *testing.T) {
	var (
		rootDir  = "./"
		fileName = "getRows"
		testData = `
			(id),name,age
			1,hana,16
		`
	)
	path := joinPath(rootDir, fileName)
	err := createFile(path, testData)
	assert.NoError(t, err)
	defer func() {
		cerr := deleteFile(path)
		assert.NoError(t, cerr)
	}()
}

func Test_SetRows(t *testing.T) {
	var (
		rootDir    = "./"
		fileName   = "setRows"
		beforeData = `
			(id),name,age
			1,hana,16
		`
		afterData = `
			(id),name,age
			1,tamate,15
		`
	)
	path := joinPath(rootDir, fileName)
	err := createFile(path, beforeData)
	assert.NoError(t, err)
	defer func() {
		cerr := deleteFile(path)
		assert.NoError(t, cerr)
	}()

	log.Println(afterData)
}
