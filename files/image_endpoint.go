package files

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alridev/figma"
	m "github.com/alridev/figma/models"
)

func GetImageWithAdditionalParams(api *figma.FigmaAPI, key string, scale float64, format string, svg_outline_text int, svg_include_id int, svg_include_node_id int, svg_simplify_stroke int, contents_only int, use_absolute_bounds int, version string, ids ...string) (m.Images, error) {
	// Рендерит изображения из файла.
	//
	// Если ошибок нет, "images" будет заполнен картой соответствий ID узлов URL-адресам
	// отрендеренных изображений, а "status" будет опущен. Срок хранения изображений - 30 дней.
	// Можно экспортировать изображения размером до 32 мегапикселей. Более крупные изображения
	// будут уменьшены.
	//
	// Важно: карта изображений может содержать null значения. Это указывает на то, что рендеринг
	// конкретного узла не удался. Это может быть связано с тем, что ID узла не существует или
	// по другим причинам, например, узел не имеет компонентов для рендеринга. Например, невидимый
	// узел или узел с 0% прозрачности не может быть отрендерен. Гарантируется, что любой узел,
	// запрошенный для рендеринга, будет представлен в этой карте, независимо от того, успешно ли
	// выполнен рендеринг.
	//
	// Для рендеринга нескольких изображений из одного файла используйте параметр ids для указания
	// нескольких ID узлов.
	//
	// Параметры:
	// !key - ключ файла из URL: https://www.figma.com/:file_type/:file_key/:file_name
	// ?scale - масштаб изображения (от 0.01 до 4)
	// ?format - формат изображения (png, svg, pdf, jpg)
	// ?svg_outline_text - включить текстовые элементы в SVG (default: false)
	// ?svg_include_id - включить ID в SVG (default: false)
	// ?svg_include_node_id - включить ID узла в SVG (default: false)
	// ?svg_simplify_stroke - упростить контуры в SVG (default: true)
	// ?contents_only - включить только содержимое изображения (default: true)
	// ?use_absolute_bounds - использовать абсолютные границы (default: false)
	// ?version - версия файла
	// !ids - список ID узлов для рендеринга

	var image m.Images

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/images/%s", api.BaseURL, key), nil)
	if err != nil {
		return image, err
	}

	if len(ids) > 0 {
		api.AddQuery(request, "ids", formatIds(ids...))
	} else {
		return image, fmt.Errorf("ids are required")
	}
	if version != "" {
		api.AddQuery(request, "version", version)
	}
	if scale != 0 {
		if scale < 0.1 || scale > 4.0 {
			return image, fmt.Errorf("scale must be between 0.1 and 4.0")
		}
		api.AddQuery(request, "scale", strconv.FormatFloat(scale, 'f', -1, 64))
	}
	if format != "" {
		api.AddQuery(request, "format", string(format))
	}
	if svg_outline_text != -1 {
		api.AddQuery(request, "svg_outline_text", strconv.FormatBool(svg_outline_text == 1))
	}
	if svg_include_id != -1 {
		api.AddQuery(request, "svg_include_id", strconv.FormatBool(svg_include_id == 1))
	}
	if svg_include_node_id != -1 {
		api.AddQuery(request, "svg_include_node_id", strconv.FormatBool(svg_include_node_id == 1))
	}
	if svg_simplify_stroke != -1 {
		api.AddQuery(request, "svg_simplify_stroke", strconv.FormatBool(svg_simplify_stroke == 1))
	}
	if contents_only != -1 {
		api.AddQuery(request, "contents_only", strconv.FormatBool(contents_only == 1))
	}
	if use_absolute_bounds != -1 {
		api.AddQuery(request, "use_absolute_bounds", strconv.FormatBool(use_absolute_bounds == 1))
	}

	response, err := api.GetResponse(request)
	if err != nil {
		return image, err
	}

	err = api.ParseJsonResponse(response, &image)
	if err != nil {
		return image, err
	}

	return image, nil
}

func GetImage(api *figma.FigmaAPI, key string, ids ...string) (m.Images, error) {
	// Рендерит изображения из файла.
	//
	// Если ошибок нет, "images" будет заполнен картой соответствий ID узлов URL-адресам
	// отрендеренных изображений, а "status" будет опущен. Срок хранения изображений - 30 дней.
	// Можно экспортировать изображения размером до 32 мегапикселей. Более крупные изображения
	// будут уменьшены.
	//
	// Важно: карта изображений может содержать null значения. Это указывает на то, что рендеринг
	// конкретного узла не удался. Это может быть связано с тем, что ID узла не существует или
	// по другим причинам, например, узел не имеет компонентов для рендеринга. Например, невидимый
	// узел или узел с 0% прозрачности не может быть отрендерен. Гарантируется, что любой узел,
	// запрошенный для рендеринга, будет представлен в этой карте, независимо от того,
	// успешно ли выполнен рендеринг.
	//
	// Для рендеринга нескольких изображений из одного файла используйте параметр ids для указания
	// нескольких ID узлов.
	//
	// Параметры:
	// !key - ключ файла из URL: https://www.figma.com/:file_type/:file_key/:file_name
	// !ids - список ID узлов для рендеринга

	return GetImageWithAdditionalParams(api, key, 1, "png", -1, -1, -1, -1, -1, -1, "", ids...)
}
