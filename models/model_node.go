package models

// Node - struct for Node
type Node map[string]interface{}

func (n Node) GetActualInstance() interface{} {
	nodeType, ok := n["type"].(string)
	if !ok {
		return nil
	}

	switch nodeType {
	case "FRAME":
		var frameNode FrameNode
		MapToStruct(n, &frameNode)
		return &frameNode
	case "DOCUMENT":
		var docNode DocumentNode
		MapToStruct(n, &docNode)
		return &docNode
	case "TEXT":
		var textNode TextNode
		MapToStruct(n, &textNode)
		return &textNode
	case "RECTANGLE":
		var rectNode RectangleNode
		MapToStruct(n, &rectNode)
		return &rectNode
	case "COMPONENT":
		var compNode ComponentNode
		MapToStruct(n, &compNode)
		return &compNode
	case "COMPONENT_SET":
		var compSetNode ComponentSetNode
		MapToStruct(n, &compSetNode)
		return &compSetNode
	case "INSTANCE":
		var instNode InstanceNode
		MapToStruct(n, &instNode)
		return &instNode
	case "BOOLEAN_OPERATION":
		var boolNode BooleanOperationNode
		MapToStruct(n, &boolNode)
		return &boolNode
	case "VECTOR":
		var vecNode VectorNode
		MapToStruct(n, &vecNode)
		return &vecNode
	case "STAR":
		var starNode StarNode
		MapToStruct(n, &starNode)
		return &starNode
	case "LINE":
		var lineNode LineNode
		MapToStruct(n, &lineNode)
		return &lineNode
	case "ELLIPSE":
		var ellipseNode EllipseNode
		MapToStruct(n, &ellipseNode)
		return &ellipseNode
	case "REGULAR_POLYGON":
		var polyNode RegularPolygonNode
		MapToStruct(n, &polyNode)
		return &polyNode
	case "GROUP":
		var groupNode GroupNode
		MapToStruct(n, &groupNode)
		return &groupNode
	case "SLICE":
		var sliceNode SliceNode
		MapToStruct(n, &sliceNode)
		return &sliceNode
	case "CONNECTOR":
		var connNode ConnectorNode
		MapToStruct(n, &connNode)
		return &connNode
	case "SECTION":
		var sectNode SectionNode
		MapToStruct(n, &sectNode)
		return &sectNode
	case "STICKY":
		var stickyNode StickyNode
		MapToStruct(n, &stickyNode)
		return &stickyNode
	case "WASHI_TAPE":
		var washiNode WashiTapeNode
		MapToStruct(n, &washiNode)
		return &washiNode
	case "SHAPE_WITH_TEXT":
		var shapeTextNode ShapeWithTextNode
		MapToStruct(n, &shapeTextNode)
		return &shapeTextNode
	case "TABLE":
		var tableNode TableNode
		MapToStruct(n, &tableNode)
		return &tableNode
	case "TABLE_CELL":
		var cellNode TableCellNode
		MapToStruct(n, &cellNode)
		return &cellNode
	case "LINK_UNFURL":
		var linkNode LinkUnfurlNode
		MapToStruct(n, &linkNode)
		return &linkNode
	case "EMBED":
		var embedNode EmbedNode
		MapToStruct(n, &embedNode)
		return &embedNode
	case "WIDGET":
		var widgetNode WidgetNode
		MapToStruct(n, &widgetNode)
		return &widgetNode
	case "CANVAS":
		var canvasNode CanvasNode
		MapToStruct(n, &canvasNode)
		return &canvasNode
	default:
		return nil
	}
}
