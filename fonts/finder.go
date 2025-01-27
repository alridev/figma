package fonts

import (
	"fmt"
	"net/http"
)

func findInFigma(fontPostScriptName string) string {
	urls := []string{
		"https://static.figma.com/font/%s_1",
		"https://static.figma.com/font/%s_2",
		"https://static.figma.com/font/%s_3",
	}

	for _, url := range urls {
		// Создаем HEAD запрос для проверки существования файла
		req, err := http.NewRequest("HEAD", fmt.Sprintf(url, fontPostScriptName), nil)
		if err != nil {
			continue
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
		client := &http.Client{}
		response, err := client.Do(req)
		if err == nil && response.StatusCode == 200 {
			return fmt.Sprintf(url, fontPostScriptName)
		}
	}
	return ""
}

func FindFont(fontFamily string, fontPostScriptName string, fontWeight int) (string, error) {
	fromFigma := findInFigma(fontPostScriptName)
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
