package concat

import "strings"

func concat(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}

	sb := strings.Builder{}
	sb.Grow(total)
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}
