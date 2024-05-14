package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Izumra/FamilyTeam/chi"
	"github.com/Izumra/FamilyTeam/utils/logger"
	"github.com/Izumra/FamilyTeam/utils/parser"
)

func main() {
	log := logger.New()

	cfg, err := parser.ParseStartCommand()
	if err != nil {
		panic("Ошибка при парсинге строки конфигурации - " + err.Error())
	}
	log.Info("Считанная конфигурация", slog.String("path", cfg.FilePath), slog.String("extension", cfg.Extension), slog.Int("port", cfg.Port))

	server, err := chi.New(log, cfg)
	if err != nil {
		panic("Ошибка при инициализации сервера - " + err.Error())
	}

	chanServ := make(chan error)
	go func() {
		chanServ <- server.Start()
	}()

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sign := <-exitChan:
		log.Info("Программа завершила работу", slog.Any("сигнал", sign))
	case err := <-chanServ:
		log.Error("Ошибка при запуске сервера", slog.Any("ошибка", err))
	}
}
