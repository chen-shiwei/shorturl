package long2short

import (
	"strconv"
	"testing"
)

func TestLongUrl2ShortUrl(t *testing.T) {

	n, err := LongUrl2ShortUrl("https://mp.weixin.qq.com/s/YTrBaERcyjvw7A0Fg2Iegw")
	if err != nil {
		t.Errorf("murmur fail:%s", err.Error())
		return
	}

	i, err := strconv.ParseInt(n, 32, 64)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(i)
}
