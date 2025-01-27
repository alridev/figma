package figma

type FigmaError struct {
	Status FigmaCodeError `json:"status"`
	Error  bool           `json:"error"`
}

type FigmaCodeError int

// Коды ошибок API Figma
const (
	BadRequest          FigmaCodeError = 400
	Forbidden           FigmaCodeError = 403
	NotFound            FigmaCodeError = 404
	RateLimit           FigmaCodeError = 429
	InternalServerError FigmaCodeError = 500
)

type FigmaErrorDescription struct {
	Code    FigmaCodeError `json:"code"`
	Message string         `json:"message"`
}

var FigmaErrorDescriptions = map[FigmaCodeError]string{
	BadRequest:          "Parameters are invalid or malformed. Please check the input formats. This error can also happen if the requested resources are too large to complete the request, which results in a timeout. Please reduce the number and size of objects requested.",
	Forbidden:           "The request was valid, but the server is refusing action. This can happen if the caller does not have the necessary permissions, or when making HTTP requests instead of HTTPS requests.",
	NotFound:            "The requested file or resource was not found.",
	RateLimit:           "In some cases API requests may be throttled or rate limited. Please wait a while before attempting the request again (typically a minute). Rate limiting is calculated on a per-user basis. If the caller is using an OAuth token, the rate limit is calculated based on the user associated with the token.",
	InternalServerError: "This most commonly occurs for very large image render requests, which may time out our server and return a 500. Please reduce the number and size of objects requested.",
}
