package parser

import (
	"archive/zip"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

type Config struct {
	Port      int
	FilePath  string
	Extension string
}

func ParseStartCommand() (*Config, error) {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 80, "port of the server")
	flag.StringVar(&cfg.FilePath, "file", "", "path to file")
	flag.StringVar(&cfg.Extension, "ext", "", "extension of the files")

	flag.Parse()

	if cfg.Port < 1 || cfg.Port > 65535 {
		return nil, fmt.Errorf("Порт сервера должен находиться в диапазоне от 1 до 65535")
	}

	return &cfg, nil
}

func ParseArchive(filePath string, extension string) ([]string, error) {

	reader, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, fmt.Errorf("Произошла ошибка при парсинге архива: открытии архива - %w", err)
	}
	defer reader.Close()

	xss := regexp.MustCompile(`<script>.*</script>`)
	var titles []string
	for _, file := range reader.File {
		title := file.Name
		if strings.HasSuffix(title, extension) && xss.FindStringSubmatch(title) == nil {
			titles = append(titles, title)
		}
	}

	return titles, nil
}
