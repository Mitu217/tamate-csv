package csv

import (
	"testing"

	"github.com/Mitu217/tamate"
	"github.com/stretchr/testify/assert"
)

func Test_Open(t *testing.T) {
	_, err := tamate.Open(driverName, "")
	assert.NoError(t, err)
}
