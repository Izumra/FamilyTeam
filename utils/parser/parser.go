package parser

import (
	"archive/zip"
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	Port      int
	FilePath  string
	Extension string
}

func ParseStartCommand() *Config {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 80, "port of the server")
	flag.StringVar(&cfg.FilePath, "file", "", "path to file")
	flag.StringVar(&cfg.Extension, "ext", "", "extension of the files")

	flag.Parse()

	return &cfg
}

func ParseArchive(filePath string, extension string) ([]string, error) {

	reader, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, fmt.Errorf("Произошла ошибка при парсинге архива - " + err.Error())
	}

	var titles []string
	for i := range reader.File {
		title := reader.File[i].Name
		if strings.HasSuffix(title, extension) {
			titles = append(titles, title)
		}
	}

	return titles, nil
}
