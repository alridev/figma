package fonts

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	client = &http.Client{
		Timeout: 5 * time.Second,
	}

	// Кэш для уже найденных шрифтов
	fontCache = sync.Map{}
)

func findInFigma(fontPostScriptName string) string {
	// Проверяем кэш
	if cachedURL, ok := fontCache.Load(fontPostScriptName); ok {
		return cachedURL.(string)
	}

	urls := []string{
		// base
		"https://static.figma.com/font/%s_1",
		"https://static.figma.com/font/%s_2",
		"https://static.figma.com/font/%s_3",
		"https://static.figma.com/font/%s_4",
		"https://static.figma.com/font/%s_5",
		"https://static.figma.com/font/%s_6",
		"https://static.figma.com/font/%s_7",
		"https://static.figma.com/font/%s_8",
		"https://static.figma.com/font/%s_9",
		"https://static.figma.com/font/%s_10",

		// wdth
		"https://static.figma.com/font/%s_wdth__1",
		"https://static.figma.com/font/%s_wdth__2",
		"https://static.figma.com/font/%s_wdth__3",
		"https://static.figma.com/font/%s_wdth__4",
		"https://static.figma.com/font/%s_wdth__5",
		"https://static.figma.com/font/%s_wdth__6",
		"https://static.figma.com/font/%s_wdth__7",
		"https://static.figma.com/font/%s_wdth__8",
		"https://static.figma.com/font/%s_wdth__9",
		"https://static.figma.com/font/%s_wdth__10",

		// wght
		"https://static.figma.com/font/%s_wght__1",
		"https://static.figma.com/font/%s_wght__2",
		"https://static.figma.com/font/%s_wght__3",
		"https://static.figma.com/font/%s_wght__4",
		"https://static.figma.com/font/%s_wght__5",
		"https://static.figma.com/font/%s_wght__6",
		"https://static.figma.com/font/%s_wght__7",
		"https://static.figma.com/font/%s_wght__8",
		"https://static.figma.com/font/%s_wght__9",
		"https://static.figma.com/font/%s_wght__10",

		// wdth_wght
		"https://static.figma.com/font/%s_wdth_wght__1",
		"https://static.figma.com/font/%s_wdth_wght__2",
		"https://static.figma.com/font/%s_wdth_wght__3",
		"https://static.figma.com/font/%s_wdth_wght__4",
		"https://static.figma.com/font/%s_wdth_wght__5",
		"https://static.figma.com/font/%s_wdth_wght__6",
		"https://static.figma.com/font/%s_wdth_wght__7",
		"https://static.figma.com/font/%s_wdth_wght__8",
		"https://static.figma.com/font/%s_wdth_wght__9",
		"https://static.figma.com/font/%s_wdth_wght__10",
	}

	// add urls from .env
	figmaUrls := os.Getenv("FIGMA_FONTS_URLS_LIST")
	if figmaUrls != "" {
		if os.Getenv("FIGMA_FONTS_REPLACE_URLS") == "true" {
			urls = strings.Split(figmaUrls, ",")
		} else {
			urlsEnv := strings.Split(figmaUrls, ",")
			urls = append(urls, urlsEnv...)
		}
	}

	// Канал для результатов
	results := make(chan string, len(urls)*2)
	var wg sync.WaitGroup

	// Функция для проверки URL
	checkURL := func(url, name string) {
		defer wg.Done()
		req, err := http.NewRequest("HEAD", fmt.Sprintf(url, name), nil)
		if err != nil {
			return
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

		response, err := client.Do(req)
		if err == nil && response.StatusCode == 200 {
			results <- fmt.Sprintf(url, name)
			return
		}
		if response != nil {
			response.Body.Close()
		}
	}

	// Запускаем горутины для проверки всех URL
	for _, url := range urls {
		wg.Add(2)
		go checkURL(url, fontPostScriptName)
	}

	// Ждем завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Получаем первый успешный результат
	for result := range results {
		fontCache.Store(fontPostScriptName, result)
		return result
	}

	return ""
}

func FindFont(fontPostScriptName string) (string, error) {
	fromFigma := findInFigma(fontPostScriptName)
	if fromFigma == "" {
		return "", fmt.Errorf("шрифт не найден в Figma")
	}
	return fromFigma, nil
}

/* Данную ссылку можно использовать в SVG
<style>
    @font-face {
        font-family: font;
        src: url(font.woff);
    }
    svg{
        font-family: font, fallBackFonts, sans-serif;
    }
</style>
*/
