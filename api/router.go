package api

import (
	"net/http"

	"QUICK-Template/logger"
	"QUICK-Template/models"

	_ "QUICK-Template/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ory/graceful"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title quicktmp
// @version 1.0
// @description This is a test case for quicktmp
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

type (
	// Custom HTTP handler. Error is returned to handle errors easier with better style.
	HandlerFunc func(http.ResponseWriter, *http.Request) error
	Config      struct {
		Port string
	}
	Handler struct {
		config  Config
		service models.Service
		logger  *logger.Logger
		router  *chi.Mux
		writer  writer
	}
)

func New(c Config, l *logger.Logger, s models.Service) *Handler {
	h := new(Handler)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/swagger/*", httpSwagger.WrapHandler)
	router.Route("/users", h.usersRouter)
	router.Route("/wallets", h.walletsRouter)

	h.logger = l
	h.config = c
	h.service = s
	h.router = router
	h.writer = newWriter()

	return h
}

// converts custom HTTP handler to net/http handler.
func (h *Handler) convert(fn HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := fn(rw, r); err != nil {
			if catch := h.writer.WriteError(rw, r, err); catch != nil {
				http.Error(rw,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}
	})
}

// Starts HTTP server with graceful shutdown.
func (h *Handler) Serve() error {
	server := graceful.WithDefaults(&http.Server{
		Addr:    ":" + h.config.Port,
		Handler: h.router,
	})

	return graceful.Graceful(server.ListenAndServe, server.Shutdown)
}
