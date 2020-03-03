package errorx

import (
	"database/sql"
	"errors"
	"testing"
)

func TestErrorx_Is(t *testing.T) {
	tests := []struct {
		name  string
		field Errorx
		err   error
		want  bool
	}{
		{"new error", Database(errors.New("new error")), ErrDatabase, true},
		{"sql no rows", Database(sql.ErrNoRows), sql.ErrNoRows, true},
		{"no error", Database(errors.New("new error")), ErrBusiness, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Is(tt.field, tt.err); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
