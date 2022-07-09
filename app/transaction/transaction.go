package transaction

import "context"

type Transaction interface {
	DoInTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
