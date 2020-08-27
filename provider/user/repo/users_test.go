package repo

import (
	"testing"

	"github.com/goapt/dbunit"
	"github.com/goapt/test"
	"github.com/ilibs/gosql/v2"
	"github.com/stretchr/testify/assert"

	"app/provider/user/cache"
	"app/testutil"
)

func TestNewUsers(t *testing.T) {
	db := &gosql.DB{}
	redis := test.NewRedis()
	assert.Equal(t, NewUsers(db, redis), &Users{Base: Base{db: db}, userRedis: cache.NewUsers(redis)})
}

func TestUsers_GetUser(t *testing.T) {
	dbunit.Run(t, testutil.Schema(), func(t *testing.T, db *gosql.DB) {
		type args struct {
			userId int
		}
		tests := []struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{
			{
				name: "exists",
				args: args{
					userId: 1,
				},
				want:    "test",
				wantErr: false,
			},
			{
				name: "not exists",
				args: args{
					userId: 333,
				},
				want:    "hello",
				wantErr: true,
			},
		}
		u := NewUsers(db, test.NewRedis())
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := u.GetUser(tt.args.userId)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && got.UserName != tt.want {
					t.Errorf("GetUser() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
