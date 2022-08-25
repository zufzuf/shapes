package util

import (
	"context"
	"net/http"
	"shapes/libs/logger"

	"github.com/rotisserie/eris"
	"github.com/unrolled/render"
	"go.uber.org/zap"
)

type CTXValue string

const (
	CTXTrackerID = CTXValue("CTX.Tracker.ID")
)

func GetTracker(ctx context.Context) string {
	v, _ := ctx.Value(CTXTrackerID).(string)
	return v
}

var (
	Render = render.New()
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
	Err     any    `json:"error"`
}

type Error struct {
	TrackerID string `json:"tracker_id"`
	Cause     any    `json:"cause,omitempty"`
}

func HTTPResponse(rw http.ResponseWriter, code int, message string, payload any) error {
	return Render.JSON(rw, code, Response{
		Code:    code,
		Message: message,
		Payload: payload,
	})
}

func ErrHTTPResponse(ctx context.Context, rw http.ResponseWriter, err error) {
	var (
		trackerId = GetTracker(ctx)
		code      = http.StatusInternalServerError
		unpack    = eris.Unpack(err)
	)

	logger.Log.With(
		zap.String("tracker_id", trackerId),
		zap.Any("error", eris.ToJSON(err, true)),
	).Error(unpack.ErrRoot.Msg)

	logger.Console.With(
		zap.String("tracker_id", trackerId),
	).Error(unpack.ErrRoot.Msg)

	Render.JSON(rw, code, Response{
		Code:    code,
		Message: unpack.ErrRoot.Msg,
		Err: Error{
			TrackerID: trackerId,
		},
	})
}

func ErrorHTTPResponse(rw http.ResponseWriter, code int, message string, err any) error {
	return Render.JSON(rw, code, Response{
		Code:    code,
		Message: message,
		Err:     err,
	})
}
