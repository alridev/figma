package files

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alridev/figma"
	m "github.com/alridev/figma/models"
)

func GetFileNodesWithAdditionalParams(api *figma.FigmaAPI, key string, version string, depth int, geometry string, plugin_data string, ids ...string) (m.FileNodes, error) {
	// Возвращает узлы, на которые ссылаются ids, в виде JSON объекта. Узлы извлекаются из файла Figma,
	// на который ссылается key.
	//
	// ID узла и ключ файла можно получить из URL: https://www.figma.com/:file_type/:file_key/:file_name?node-id=:id
	//
	// Атрибуты name, lastModified, thumbnailUrl, editorType и version являются метаданными указанного файла.
	//
	// Поле linkAccess описывает уровень доступа к файлу по ссылке. Существует 5 типов разрешений:
	// "inherit" - наследует разрешения проекта (по умолчанию)
	// "view" - просмотр
	// "edit" - редактирование
	// "org_view" - просмотр только для пользователей организации
	// "org_edit" - редактирование только для пользователей организации
	//
	// Атрибут document содержит узел типа DOCUMENT.
	// В components содержится маппинг ID узлов на метаданные компонентов.
	// В styles содержится маппинг ID стилей на их метаданные.
	//
	// По умолчанию векторные данные не возвращаются. Для их получения используйте параметр geometry=paths.
	//
	// Важно: карта узлов может содержать null значения, если указанный ID не существует в файле.
	//
	// Параметры:
	// key - ключ файла из URL: https://www.figma.com/:file_type/:file_key/:file_name
	// version - версия файла
	// depth - глубина дерева документа
	// geometry - формат геометрии (paths или bounds)
	// plugin_data - данные плагина
	// ids - список ID узлов для получения

	var fileNodes m.FileNodes

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/files/%s/nodes", api.BaseURL, key), nil)
	if err != nil {
		return fileNodes, err
	}
	if len(ids) > 0 {
		api.AddQuery(request, "ids", formatIds(ids...))
	} else {
		return fileNodes, fmt.Errorf("ids are required")
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

	response, err := api.GetResponse(request)
	if err != nil {
		return fileNodes, err
	}

	err = api.ParseJsonResponse(response, &fileNodes)
	if err != nil {
		return fileNodes, err
	}

	return fileNodes, nil
}

func GetFileNodes(api *figma.FigmaAPI, key string, ids ...string) (m.FileNodes, error) {
	// Возвращает узлы, на которые ссылаются :ids, в виде JSON объекта. Узлы извлекаются из файла Figma,
	// на который ссылается :key.
	//
	// ID узла и ключ файла можно получить из URL: https://www.figma.com/:file_type/:file_key/:file_name?node-id=:id
	//
	// Атрибуты name, lastModified, thumbnailUrl, editorType и version являются метаданными указанного файла.
	//
	// Поле linkAccess описывает уровень доступа к файлу по ссылке. Существует 5 типов разрешений:
	// "inherit" - наследует разрешения проекта (по умолчанию)
	// "view" - просмотр
	// "edit" - редактирование
	// "org_view" - просмотр только для пользователей организации
	// "org_edit" - редактирование только для пользователей организации
	//
	// Атрибут document содержит узел типа DOCUMENT.
	// В components содержится маппинг ID узлов на метаданные компонентов.
	// В styles содержится маппинг ID стилей на их метаданные.
	//
	// По умолчанию векторные данные не возвращаются. Для их получения используйте параметр geometry=paths.
	//
	// Важно: карта узлов может содержать null значения, если указанный ID не существует в файле.
	//
	// Параметры:
	// key - ключ файла из URL: https://www.figma.com/:file_type/:file_key/:file_name
	// ids - список ID узлов для получения

	return GetFileNodesWithAdditionalParams(api, key, "", 0, "", "", ids...)
}
