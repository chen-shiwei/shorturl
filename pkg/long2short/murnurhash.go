package long2short

import (
	"strconv"

	"github.com/spaolacci/murmur3"
)

func LongUrl2ShortUrl(url string) (string, error) {
	var hasher = murmur3.New32()
	_, err := hasher.Write([]byte(url))
	if err != nil {
		return "", err
	}
	n := hasher.Sum32()

	short := strconv.FormatInt(int64(n), 32)

	writeMap(short, url)

	return short, nil
}
