package errutil

import (
	"database/sql"
	"database/sql/driver"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSystemError(t *testing.T) {
	assert.True(t, IsSystemError(sql.ErrConnDone))
	assert.True(t, IsSystemError(sql.ErrTxDone))
	assert.True(t, IsSystemError(driver.ErrBadConn))
	assert.True(t, IsSystemError(&net.DNSError{}))
}
