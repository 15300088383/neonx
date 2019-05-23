package service

import (
	"context"

	"github.com/geekymedic/neonx/logger"

	"github.com/geekymedic/neonx"
	"google.golang.org/grpc/metadata"
)

var (
	stateName = "neonx.service.State"
)

type State struct {
	*logger.Logger
	ctx context.Context
	*neonx.Session
}

func (m *State) Context() context.Context {
	return m.ctx
}

func NewState(ctx context.Context) *State {

	v := ctx.Value(stateName)

	state, ok := v.(*State)

	if ok {
		return state
	}

	var (
		session    = &neonx.Session{}
		md, exists = metadata.FromIncomingContext(ctx)
		value      = func(name string, x *string) {

			data := md.Get(name)

			if len(data) > 0 {
				*x = data[0]
			}

		}
	)

	if exists {
		for name, ref := range session.Keys() {
			value(name, ref)
		}
	}

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.New(
			session.Encode(),
		),
	)

	state = &State{
		Session: session,
		Logger:  logger.NewLogger(session),
	}
	ctx = context.WithValue(ctx, stateName, state)

	state.ctx = ctx

	return state

}
