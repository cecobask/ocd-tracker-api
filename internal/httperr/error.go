package httperr

import (
	"github.com/cecobask/ocd-tracker-api/internal/log"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
)

func Unauthorised(w http.ResponseWriter, r *http.Request, slug string, err error) {
	httpRespondWithError(w, r, slug, err, "unauthorised", http.StatusUnauthorized)
}

func httpRespondWithError(w http.ResponseWriter, r *http.Request, slug string, err error, message string, status int) {
	logger := log.LoggerFromContext(r.Context())
	logger.Warn(message, zap.String("error-slug", slug), zap.Int("status", status), zap.Error(err))
	resp := ErrorResponse{slug, status}
	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}