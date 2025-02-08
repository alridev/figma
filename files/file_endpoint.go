package files

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alridev/figma"
	m "github.com/alridev/figma/models"
)

func GetFileWithAdditionalParams(api *figma.FigmaAPI, key string, version string, depth int, geometry string, plugin_data string, branch_data int, ids ...string) (m.File, error) {
	// Получаем документ по ключу файла с дополнительными параметрами
	// !key - ключ файла из URL: https://www.figma.com/:file_type/:file_key/:file_name
	// ?version - версия файла
	// ?depth - глубина дерева документа
	// ?geometry - формат геометрии (paths или bounds)
	// ?plugin_data - данные плагина
	// ?branch_data - включить данные о ветках
	// ?ids - список ID узлов для получения
	// Возвращает метаданные файла (name, lastModified, thumbnailUrl и т.д.) и сам документ
	// В components содержится маппинг ID узлов на метаданные компонентов

	var file m.File

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/files/%s", api.BaseURL, key), nil)
	if err != nil {
		return file, err
	}
	if len(ids) > 0 {
		api.AddQuery(request, "ids", formatIds(ids...))
	}
	if version != "" {
		api.AddQuery(request, "version", version)
	}
	if depth != 0 {
		api.AddQuery(request, "depth", strconv.Itoa(depth))
	}
	if geometry != "" {
		api.AddQuery(request, "geometry", geometry)
	}
	if plugin_data != "" {
		api.AddQuery(request, "plugin_data", plugin_data)
	}
	if branch_data != -1 {
		api.AddQuery(request, "branch_data", strconv.FormatBool(branch_data == 1))
	}

	response, err := api.GetResponse(request)
	if err != nil {
		return file, err
	}

	err = api.ParseJsonResponse(response, &file)
	if err != nil {
		return file, err
	}

	return file, nil
}

func GetFile(api *figma.FigmaAPI, key string, ids ...string) (m.File, error) {
	// Получаем документ по ключу файла
	// Ключ файла можно получить из URL: https://www.figma.com/:file_type/:file_key/:file_name
	// Возвращает метаданные файла (name, lastModified, thumbnailUrl и т.д.) и сам документ
	// В components содержится маппинг ID узлов на метаданные компонентов
	//
	// Параметры:
	// !key - ключ файла из URL: https://www.figma.com/:file_type/:file_key/:file_name
	// ?ids - список ID узлов для получения

	return GetFileWithAdditionalParams(api, key, "", 0, "", "", -1, ids...)
}
