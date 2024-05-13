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

	cfg := parser.ParseStartCommand()
	log.Info("Считанная конфигурация", slog.Any("cfg", cfg))

	server := chi.New(log, cfg)

	server.Start()

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	sign := <-exitChan
	log.Info("Программа завершила работу", slog.Any("сигнал", sign))
}
