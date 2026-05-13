package Trimspace

func Trimspacemanual(s string) string {
	start := 0
	end := len(s) - 1

	for start <= end && (s[start] == ' ' || s[start] == '\n' || s[start] == '\t' || s[start] == '\r') {
		start++
	}
	for end >= start && (s[end] == ' ' || s[end] == '\n' || s[end] == '\t' || s[end] == '\r') {
		end--
	}
	return s[start : end+1]
}
