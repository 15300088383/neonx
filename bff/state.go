package bff

import (
	"context"
	"net/http"

	"github.com/geekymedic/neonx/logger"

	"github.com/geekymedic/neonx/errors"
	"google.golang.org/grpc/metadata"

	"github.com/geekymedic/neonx"

	"github.com/gin-gonic/gin"
)

var (
	empty     = struct{}{}
	stateName = "neonx.bff.State"
)

type State struct {
	*neonx.Session
	*logger.Logger
	Gin *gin.Context
	ctx context.Context
}

func (m *State) httpJson(code int, v interface{}) {
	m.Gin.JSON(
		http.StatusOK,
		map[string]interface{}{"Code": code,
			"Message": GetMessage(code),
			"Data":    v,
		},
	)

}

func (m *State) Error(code int, err error) {
	if err != nil {
		m.LogError(err.Error())
	}
	m.httpJson(code, empty)
}

func (m *State) Success(v interface{}) {
	m.httpJson(
		CodeSuccess,
		v)
}

func (m *State) Context() context.Context {
	return m.ctx
}

func newSession(ctx *gin.Context) *neonx.Session {
	var (
		s = &neonx.Session{}
	)

	for name, ref := range s.Keys() {
		*ref = ctx.Query(name)
	}

	return s
}

func (m *State) BindJSON(v interface{}) error {

	err := m.Gin.BindJSON(v)

	if err != nil {
		return errors.By(err)
	}

	return nil
}

func newContext(ctx context.Context, session *neonx.Session) context.Context {

	return metadata.NewOutgoingContext(
		ctx,
		metadata.New(session.Encode()),
	)
}

func NewState(ctx *gin.Context) *State {

	x, exists := ctx.Get(stateName)

	if exists {
		return x.(*State)
	}

	var (
		context = context.Background()
		session = newSession(ctx)
		state   = &State{
			Gin:     ctx,
			Session: session,
			Logger:  logger.NewLogger(session),
			ctx:     newContext(context, session),
		}
	)

	ctx.Set(stateName, state)

	return state
}
