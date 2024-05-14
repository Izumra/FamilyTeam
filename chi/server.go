package chi

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Izumra/FamilyTeam/utils/parser"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type server struct {
	log    *slog.Logger
	cfg    *parser.Config
	server *chi.Mux
}

func New(log *slog.Logger, cfg *parser.Config) (*server, error) {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	files, err := parser.ParseArchive(cfg.FilePath, cfg.Extension)
	if err != nil {
		return nil, fmt.Errorf("Произошла ошибка пр парсинге архива - %w", err)
	}

	html := fmt.Sprintf(`
	<html>
		<head></head>
		<body>
			<div>Extension: %s</div>
			<ul>`, cfg.Extension)
	for _, v := range files {
		html += fmt.Sprintf("<li>%s</li>", v)
	}
	html += `
			</ul>
		</body>
	</html>
	`

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	})

	return &server{
		log,
		cfg,
		mux,
	}, nil
}

func (s *server) Start() error {
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	err := http.ListenAndServe(addr, s.server)
	if err != nil {
		return fmt.Errorf("Не удалось запустить сервер по причине %w", err)
	}
	return nil
}
