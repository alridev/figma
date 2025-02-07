package fonts

import (
	"fmt"
	"net/http"
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

func findInFigma(fontFamily string, fontPostScriptName string) string {
	// Проверяем кэш
	if cachedURL, ok := fontCache.Load(fontPostScriptName); ok {
		return cachedURL.(string)
	}
	if cachedURL, ok := fontCache.Load(fontFamily); ok {
		return cachedURL.(string)
	}

	urls := []string{
		"https://static.figma.com/font/%s_1",
		"https://static.figma.com/font/%s_2",
		"https://static.figma.com/font/%s_3",
		"https://static.figma.com/font/%s_wght_1",
		"https://static.figma.com/font/%s_wght_2",
		"https://static.figma.com/font/%s_wght_3",
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
		go checkURL(url, fontFamily)
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

func FindFont(fontFamily string, fontPostScriptName string, fontWeight int) (string, error) {
	fromFigma := findInFigma(fontFamily, fontPostScriptName)
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
