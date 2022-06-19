package util

import (
	"regexp"
)

type UrlReplaceUtil struct {
	UrlStr string
}

const (
	SkuRegexpV1 = `(?m)((?P<key>wareId=)(?P<value>\d+)[&].*)` // "(?<=wareId=).*?(?=&)"
	SkuRegexpV2 = `(?m)((?P<key>sku=)(?P<value>\d+)[&].*)`    // (?<=sku=).*?(?=&)
	SkuRegexpV3 = `(?m)((?P<key>/)(?P<value>\d+).*)`          // (?<=\\/)\\d*(?=\\.)
)

func NewUrlReplaceUtil(urlStr string) *UrlReplaceUtil {
	return &UrlReplaceUtil{UrlStr: urlStr}
}

// GetUrl twpUrl提取
func (ur *UrlReplaceUtil) GetUrl() ([]string, error) {
	reg, err := regexp.Compile(`(http|https)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	if err != nil {
		return nil, err
	}
	return reg.FindAllString(ur.UrlStr, -1), nil
}

// GetSkuId skuId提取
func (ur *UrlReplaceUtil) GetSkuId() (string, error) {
	reg, err := regexp.Compile("(?m)(?P<key>(抢购|抢购:|抢购：|下单|下单:|下单：))(?P<value>[0-9]+)")
	if err != nil {
		return "", err
	}
	nameIndex := reg.SubexpIndex("value") // nameIndex = 3
	allIndexes := reg.FindAllStringSubmatch(ur.UrlStr, -1)
	for _, loc := range allIndexes {
		return loc[nameIndex], nil
	}
	return "", err
}

// GetSkuIdByCustom skuId提取
func (ur *UrlReplaceUtil) GetSkuIdByCustom(exp string, tag string) (string, error) {
	reg, err := regexp.Compile(exp)
	if err != nil {
		return "", err
	}
	nameIndex := reg.SubexpIndex(tag) // nameIndex = 3
	allIndexes := reg.FindAllStringSubmatch(ur.UrlStr, -1)
	for _, loc := range allIndexes {
		return loc[nameIndex], nil
	}
	return "", err
}
