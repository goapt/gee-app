package repo

import (
	"github.com/ilibs/gosql/v2"
)

type BuilderFunc func(b *gosql.Builder)

type Base struct {
	db *gosql.DB
}

func (b *Base) Create(m gosql.IModel) (int64, error) {
	id, err := b.db.Model(m).Create()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b *Base) Update(m gosql.IModel) (int64, error) {
	aff, err := b.db.Model(m).Update()
	if err != nil {
		return 0, err
	}
	return aff, nil
}

func (b *Base) Find(m gosql.IModel) error {
	err := b.db.Model(m).Get()

	if err != nil {
		return err
	}

	return nil
}

func (b *Base) FindAll(m interface{}, fn ...BuilderFunc) error {
	builder := b.db.Model(m)

	if len(fn) > 0 {
		for _, f := range fn {
			f(builder)
		}
	}

	err := builder.All()
	if err != nil {
		return err
	}
	return nil
}
