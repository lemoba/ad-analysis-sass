package server

import (
	"context"
	"net/url"
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type EndPointer interface {
	EndPoint() (*url.URL, error)
}
