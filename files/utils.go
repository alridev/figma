package files

import (
	"strings"
)

func formatIds(ids ...string) string {
	return strings.Join(ids, ",")
}
