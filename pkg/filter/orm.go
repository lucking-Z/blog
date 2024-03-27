package filter

import "context"

type Orm interface {
	BeginTx(ctx context.Context)
	Create(ctx context.Context, obj any) error
	Delete(ctx context.Context, filter Filter) error
	Update(ctx context.Context, filter Filter, obj any) error
	FindOne(ctx context.Context, filter Filter, obj any) error
	Find(ctx context.Context, filter Filter, objs any) error
}
