package models

import "context"

type StopList struct {
	ctx context.Context
}

func NewStopList(ctx context.Context) *StopList {
	return &StopList{ctx: ctx}
}
