package figma

import "time"

const (
	// BaseURL - базовый URL API Figma
	BaseURL = "https://api.figma.com"

	// UserAgent - User-Agent по умолчанию для запросов к API
	UserAgent = "Figma-Go-Client-v0.1.0"

	// Timeout - таймаут для запросов к API
	Timeout = 10 * time.Second
)
