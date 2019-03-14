package csv

import (
	"context"

	"github.com/Mitu217/tamate"
	"github.com/Mitu217/tamate/driver"
)

const driverName = "csv"

type csvDriver struct{}

func (cd *csvDriver) Open(ctx context.Context, dsn string) (driver.Conn, error) {
	return newCSVConn(dsn, 0), nil
}

func init() {
	tamate.Register(driverName, &csvDriver{})
}
