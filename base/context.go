package base

import (
	"context"
	"errors"
	"sync"
	"time"
)

type BaseContext interface {
	context.Context
	SetError(err error)
	GetError() error
	Set(key string, value any)
	GetString(key any) string
	Get(any) (any, bool)
	SetRequestId(string)
	GetRequestId() string
	SetUserId(any)
	GetUserId() any
}

type Context struct {
	m         *sync.Map
	RequestId *string
	UserId    *any
	err       *error
}

func (y Context) Deadline() (deadline time.Time, ok bool) {
	//TODO implement me
	return
}

func (y Context) Done() <-chan struct{} {
	//TODO implement me
	return nil
}

func (y Context) Err() error {
	return *y.err
}

func (y Context) Value(key any) any {
	//TODO implement me
	return nil
}

func (y Context) SetError(err error) {
	if y.err == nil {
		y.err = &err
	} else {
		ter := errors.Join(*y.err, err)
		y.err = &ter
	}
}
func (y Context) GetError() error {
	return *y.err
}
func (y Context) Set(key string, v any) {
	if y.m == nil {
		y.m = new(sync.Map)
	}
	y.m.Store(key, v)
}

func (y Context) GetString(key any) (value string) {
	val, ok := y.Get(key)
	if ok && val != nil {
		value, _ = val.(string)
	}
	return
}

func (y Context) Get(key any) (value any, exists bool) {
	value, exists = y.m.Load(key)
	return
}

func (y Context) SetRequestId(requestId string) {
	y.RequestId = &requestId
}
func (y Context) GetRequestId() string {
	return *y.RequestId
}

func (y Context) SetUserId(userId any) {
	y.UserId = &userId
}

func (y Context) GetUserId() any {
	return *y.UserId
}
