package query

import "uuid"

type Query struct {
	QueryId   string
	QueryName string
	Args      []string
}

func NewQuery(queryName string, args ...string) *Query {
	return &Query{
		QueryId:   uuid.New().String(),
		QueryName: queryName,
		Args:      args,
	}
}
