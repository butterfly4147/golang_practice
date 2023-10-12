package main

import "context"

// interfacer
type DBCommon interface {
	Insert(ctx context.Context, instance interface{})
}
