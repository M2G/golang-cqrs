package query

type Query[T any] struct {
	QueryId string
	Query   func(args ...string) (T, error)
	Args    []string
}

func NewQuery[T any](queryId string, query func(args ...string) (T, error), args ...string) *Query[T] {
	return &Query[T]{
		QueryId: queryId,
		Query:   query,
		Args:    args,
	}
}
