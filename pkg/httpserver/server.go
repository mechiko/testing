// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"mime"
	"net/http"
	"time"

	httpV1 "testing/internal/controller/http/v1"
	"testing/pkg"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(e *echo.Echo, opts ...Option) *Server {
	// HTTP Server
	mime.AddExtensionType(".js", "application/javascript")
	// e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// e.Logger.SetOutput(os.Stdout)
	// 	e.Logger.SetOutput(ioutil.Discard)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Pre(middleware.Rewrite(map[string]string{
		"/spa/requestrests/*":   "/",
		"/spa/requestclients/*": "/",
		"/test":                 "/",
	}))

	// Prometheus metrics
	// p := prometheus.NewPrometheus("echo", nil)
	// p.Use(e)

	httpServer := &http.Server{
		Handler:      e,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	e.GET("/shutdown", func(c echo.Context) error {
		// e.Shutdown(context.Background())
		c.String(http.StatusOK, "Shutdown in proccess")
		return s.Shutdown()
	})

	assetHandler := http.FileServer(pkg.GetFileSystem())
	e.GET("/*", echo.WrapHandler(assetHandler))

	// изучаем структуру приложения в части контроллера http
	httpV1.NewRouterSimple(e)

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
