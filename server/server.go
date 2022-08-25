package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"shapes/db"
	"shapes/handler"
	"shapes/libs/logger"
	"shapes/repository"
	AppMiddleware "shapes/server/middleware"
	"shapes/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func NewHTTPServer() *HTTPServer {
	db := db.Init()
	logger.New()

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"*"},
		AllowCredentials:   true,
		MaxAge:             60,
		OptionsPassthrough: false,
		Debug:              false,
	}))
	r.Use(AppMiddleware.Tracker)

	repoArea := &repository.AreaRepository{
		DB: db,
	}

	srv := &service.AreaService{
		Repo: repoArea,
	}

	server := &HTTPServer{
		Router:      r,
		DB:          db,
		AreaHandler: &handler.AreaHTTPHandler{Service: srv},
	}

	server.routes()

	return server
}

type HTTPServer struct {
	Router *chi.Mux
	DB     *gorm.DB

	AreaHandler AreaHandler
}

func (hs *HTTPServer) Run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(
		ctx,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	port, ok := os.LookupEnv("API_PORT")
	if !ok {
		port = "3000"
	}

	server := http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           hs.Router,
		IdleTimeout:       0,
		WriteTimeout:      5 * time.Second,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		log.Printf("start area api")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start / shutdown area api, err : \n%+v\n", err)
		}
	}()

	<-ctx.Done()

	shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("server shutdown")

	if err := server.Shutdown(shutdown); err != nil {
		log.Fatalf("shutdown area api, err : \n%+v\n", err)
	}

	log.Printf("server shutdown properly")

	if err := db.Close(hs.DB); err != nil {
		log.Fatal("unable close db connection")
	}

	return nil
}

type AreaHandler interface {
	Insert(rw http.ResponseWriter, r *http.Request)
}

func (hs *HTTPServer) routes() {
	hs.Router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("area api 🔥"))
	})
	hs.Router.Post("/area", hs.AreaHandler.Insert)
}
