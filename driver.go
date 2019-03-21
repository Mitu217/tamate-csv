package csv

import (
	"context"

	"github.com/go-tamate/tamate"
	"github.com/go-tamate/tamate/driver"
)

const driverName = "csv"

type csvDriver struct{}

func (cd *csvDriver) Open(ctx context.Context, dsn string) (driver.Conn, error) {
	return newCSVConn(dsn, 0)
}

func init() {
	tamate.Register(driverName, &csvDriver{})
}
