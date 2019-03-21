package csv

import (
	"testing"

	"github.com/go-tamate/tamate"
	"github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	drivers := tamate.Drivers()
	d, has := drivers[driverName]
	assert.EqualValues(t, &csvDriver{}, d)
	assert.True(t, has)
}

func Test_Open(t *testing.T) {
	var (
		dsn = ""
	)

	ds, err := tamate.Open(driverName, dsn)
	defer func() {
		err := ds.Close()
		assert.NoError(t, err)
	}()
	assert.NoError(t, err)
}
