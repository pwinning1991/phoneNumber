package main

func normalize(phone string) string {
	var buf bytes.Buffer

	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRun(ch)
		}
	}
	return buf.String()
}
