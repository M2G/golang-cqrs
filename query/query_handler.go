package query

type Handler interface {
	Execute(args []string) (interface{}, error)
}

type QueryHandler[T any] struct {
	QueryName string
	query     func(args []string) (T, error)
}

func NewHandler[T any](queryName string, query func(args []string) (T, error)) *QueryHandler[T] {
	return &QueryHandler[T]{
		QueryName: queryName,
		query:     query,
	}
}

func (qh *QueryHandler[T]) Execute(args []string) (interface{}, error) {
	return qh.query(args)
}
