package page

func makeURLWithAttributes(origin string, params map[string]string) string {
	var paramPart string

	for key, value := range params {
		if value != "" {
			paramPart = paramPart + key + "=" + value + "&"
		}
	}
	return "/" + origin + "?" + paramPart
}
