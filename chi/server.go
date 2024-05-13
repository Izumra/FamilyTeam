package chi

import (
	"fmt"
	"log/slog"
	"net/http"
	"syscall"

	"github.com/Izumra/FamilyTeam/utils/parser"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type server struct {
	log    *slog.Logger
	cfg    *parser.Config
	server *chi.Mux
}

func New(log *slog.Logger, cfg *parser.Config) *server {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	files, err := parser.ParseArchive(cfg.FilePath, cfg.Extension)
	if err != nil {
		panic("Произошла ошибка пр парсинге архива - " + err.Error())
	}

	html := fmt.Sprintf("Extension: %s", cfg.Extension)
	for _, v := range files {
		html += "\n" + v
	}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	})

	return &server{
		log,
		cfg,
		mux,
	}
}

func (s *server) Start() {
	op := "chi/server.Start"
	log := s.log.With("op", op)

	addr := fmt.Sprintf(":%d", s.cfg.Port)
	err := http.ListenAndServe(addr, s.server)
	if err != nil {
		log.Error("Не удалось запустить сервер по причине", slog.Any("cause", err))

		err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		if err != nil {
			syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		}
	}
}
