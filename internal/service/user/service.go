package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha/facade/pkg/database"
)

type Service struct{}

func (Service) GetBatch(ctx context.Context, from int64, limit uint64) (us []User, err error) {
	db, c, err := database.Get(ctx)()
	defer c()

	q, a, err := sq.
		Select("*").
		From(`"user"`).
		Where(sq.Gt{"id": from}).
		Limit(limit).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	err = db.SelectContext(ctx, &us, q, a...)

	return
}

func (Service) Login(ctx context.Context, user int64) (isNew bool, err error) {

}