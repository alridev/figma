package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/alridev/figma"
	"github.com/alridev/figma/fonts"
	m "github.com/alridev/figma/models"
	"github.com/stretchr/testify/assert"
)

func TestFindFont(t *testing.T) {
	font, err := fonts.FindFont("Fira Code", "FiraCode-Regular", 400)
	assert.NoError(t, err)
	assert.NotEmpty(t, font)

	font, err = fonts.FindFont("Inter", "Inter-Regular", 400)
	assert.NoError(t, err)
	assert.NotEmpty(t, font)
}

func TestParseImagesJSON(t *testing.T) {
	// Читаем тестовый файл
	data, err := os.ReadFile(filepath.Join("responses", "images.json"))
	assert.NoError(t, err)

	// Структура для парсинга
	var response m.Images

	// Парсим JSON
	err = json.Unmarshal(data, &response)
	assert.NoError(t, err)

	// Проверяем результаты
	assert.Empty(t, response.Err)
	assert.NotEmpty(t, response.Images)
	assert.Contains(t, response.Images, "1:14")
}

func TestParseFileNodesJSON(t *testing.T) {
	// Читаем тестовый файл
	data, err := os.ReadFile(filepath.Join("responses", "files-nodes.json"))
	assert.NoError(t, err)

	// Структура для парсинга
	var response m.FileNodes
	// Парсим JSON
	err = figma.ParseJsonFromBytes(data, &response)
	assert.NoError(t, err)

	// Проверяем основные поля
	assert.Equal(t, "Untitled", response.Name)
	assert.NotEmpty(t, response.ThumbnailUrl)
	assert.Equal(t, "owner", response.Role)
	assert.Equal(t, "figma", response.EditorType)

	// Проверяем наличие узлов
	assert.NotEmpty(t, response.Nodes)
	assert.Contains(t, response.Nodes, "1:14")

	// Проверяем структуру узла
	node := response.Nodes["1:14"]

	frameNode, ok := node.Document.GetActualInstance().(*m.FrameNode)
	assert.True(t, ok, "Expected FrameNode type")
	assert.Equal(t, "screen", frameNode.Name)
	assert.Equal(t, "FRAME", string(frameNode.Type))

	headNode, ok := frameNode.Children[1].GetActualInstance().(*m.TextNode)
	assert.True(t, ok, "Expected TextNode type")
	assert.Equal(t, "$head", headNode.Name)
	assert.Equal(t, "TEXT", string(headNode.Type))

	fill, ok := headNode.Fills[0].GetActualInstance().(*m.SolidPaint)
	assert.True(t, ok, "Expected SolidPaint type")
	assert.Equal(t, "SOLID", string(fill.Type))

	// Проверяем наличие дочерних узлов
	assert.NotEmpty(t, frameNode.Children)
}

func TestParseImageFillsJSON(t *testing.T) {
	// Читаем тестовый файл
	data, err := os.ReadFile(filepath.Join("responses", "images-fills.json"))
	assert.NoError(t, err)

	// Структура для парсинга
	var response m.ImageFills

	// Парсим JSON
	err = figma.ParseJsonFromBytes(data, &response)
	assert.NoError(t, err)

	// Проверяем результаты
	assert.Empty(t, response.Error)
	assert.NotEmpty(t, response.Meta.Images)
}

func TestParseCompleteFileJSON(t *testing.T) {
	// Читаем тестовый файл
	data, err := os.ReadFile(filepath.Join("responses", "files.json"))
	assert.NoError(t, err)

	// Структура для парсинга
	var response m.File

	// Парсим JSON
	err = figma.ParseJsonFromBytes(data, &response)
	assert.NoError(t, err)

	// Проверяем основные поля файла
	assert.Equal(t, "Untitled", response.Name)
	assert.Equal(t, "owner", response.Role)
	assert.Equal(t, "figma", response.EditorType)
	assert.NotEmpty(t, response.Version)
	assert.NotEmpty(t, response.LastModified)
	assert.NotEmpty(t, response.ThumbnailUrl)

	// Проверяем структуру документа
	assert.Equal(t, "Document", response.Document.Name)
	assert.NotEmpty(t, response.Document.Children)

	// Проверяем Canvas (Page 1)
	canvas := response.Document.Children[0]
	canvasNode, ok := canvas.GetActualInstance().(*m.CanvasNode)
	assert.True(t, ok, "Expected CanvasNode type")
	assert.Equal(t, "Page 1", canvasNode.Name)
	assert.NotEmpty(t, canvasNode.Children)

	// Проверяем Frame (screen)
	frame := canvasNode.Children[0]
	frameNode, ok := frame.GetActualInstance().(*m.FrameNode)
	assert.True(t, ok, "Expected FrameNode type")
	assert.Equal(t, "screen", frameNode.Name)
	assert.NotEmpty(t, frameNode.Children)

	// Проверяем наличие всех элементов во frame
	elementNames := []string{"$image", "$head", "$qrcode", "$time", "$desc"}
	for i, name := range elementNames {
		elementNode := frameNode.Children[i]
		actualNode := elementNode.GetActualInstance()
		if name == "$qrcode" {
			assert.Equal(t, name, actualNode.(*m.FrameNode).Name)
		} else if name == "$image" {
			assert.Equal(t, name, actualNode.(*m.RectangleNode).Name)
		} else {
			assert.Equal(t, name, actualNode.(*m.TextNode).Name)
		}
	}

	// Проверяем детали изображения
	image := frameNode.Children[0]
	rectangleNode, ok := image.GetActualInstance().(*m.RectangleNode)
	assert.True(t, ok, "Expected RectangleNode type")
	assert.Equal(t, "$image", rectangleNode.Name)
}
