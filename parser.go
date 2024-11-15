// parser.go
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Config struct {
	Period   string // YYYYMM
	FilePath string
}

func ParseArgs() (*Config, error) {
	if len(os.Args) != 3 {
		return nil, errors.New("usage: program YYYYMM filepath")
	}

	period := os.Args[1]
	filePath := os.Args[2]

	// Validate period format (YYYYMM)
	if err := validatePeriod(period); err != nil {
		return nil, fmt.Errorf("invalid period format: %v", err)
	}

	// Validate file path
	if err := validateFilePath(filePath); err != nil {
		return nil, fmt.Errorf("invalid file path: %v", err)
	}

	return &Config{
		Period:   period,
		FilePath: filePath,
	}, nil
}

func validatePeriod(period string) error {
	// Check basic format
	matched, err := regexp.MatchString(`^\d{6}$`, period)
	if err != nil || !matched {
		return errors.New("period must be in YYYYMM format")
	}

	// Parse as actual date using time.Parse
	_, err = time.Parse("200601", period)
	if err != nil {
		return errors.New("invalid year/month combination")
	}

	return nil
}

func validateFilePath(path string) error {
	// Check if file exists
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("file does not exist")
		}
		return err
	}

	// Check if it's a regular file (not a directory)
	if info.IsDir() {
		return errors.New("path points to a directory, not a file")
	}

	// Check file extension
	ext := filepath.Ext(path)
	if ext != ".csv" {
		return errors.New("file must be a CSV file")
	}

	return nil
}
