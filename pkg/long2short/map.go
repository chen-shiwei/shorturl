package long2short

var ShortMap = map[string]string{}

func GetShortMap(shortUrl string) string {
	return ShortMap[shortUrl]
}

func writeMap(short, long string) {
	ShortMap[short] = long
}
