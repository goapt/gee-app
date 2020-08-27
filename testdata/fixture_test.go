package testdata

import (
	"fmt"
	"testing"

	"github.com/goapt/dbunit"
	"github.com/ilibs/gosql/v2"

	"app/testutil"
)

// the connect development environment database
func db(dbname string) *gosql.DB {
	d, err := gosql.Open("mysql", fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=%s", dbname, "Asia%2FShanghai"))
	if err != nil {
		panic(err)
	}
	return d
}

// dump users testdata
func TestFixtures(t *testing.T) {
	if testing.Short() {
		t.Skip("dump testdata,skip")
	}

	_, err := dbunit.Dump(db("test"), testutil.Fixture("users"), "select * from users limit 10")
	if err != nil {
		panic(err)
	}
}
