package util

import (
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlReplaceUtil_GetUrl(t *testing.T) {
	util := NewUrlReplaceUtil("https://item.jd.com/100024319032.html?cu=true&utm_source=kong&utm_medium=tuiguang&utm_campaign=t_2024968405_&utm_term=672cd15a5d2b4335a04052fd5faccb2e")
	urlList, err := util.GetUrl()
	assert.Nil(t, err)
	t.Log(urlList)
}

// https://item.jd.com/12414724882.html http://coupon.jd.com/ilink/couponSendFront/send_index.action?key=u3n5iaocnaw3xea5p2p18ab4fcec94ba&roleId=70553734&to=maijingtujydq.jd.com
func TestUrlReplaceUtil_GetUrl2(t *testing.T) {
	util := NewUrlReplaceUtil("https://item.jd.com/12414724882.html http://coupon.jd.com/ilink/couponSendFront/send_index.action?key=u3n5iaocnaw3xea5p2p18ab4fcec94ba&roleId=70553734&to=maijingtujydq.jd.com")
	urlList, err := util.GetUrl()
	assert.Nil(t, err)
	t.Log(urlList)
}

func TestUrlReplaceUtil_GetShortUrl(t *testing.T) {
	util := NewUrlReplaceUtil("https://u.jd.com/ld8HVGU")
	urlList, err := util.GetUrl()
	assert.Nil(t, err)
	t.Log(urlList)
}

func TestUrlReplaceUtil_GetSkuId(t *testing.T) {
	contents := []string{
		"下单100024319032",
		"下单:100024319032",
		"下单：100024319032",
		"抢购100024319032",
		"抢购:100024319032",
		"抢购：100024319032",
	}
	for _, v := range contents {
		util := NewUrlReplaceUtil(v)
		skuId, err := util.GetSkuId()
		assert.Nil(t, err)
		assert.Equal(t, skuId, "100024319032")
	}
}

func TestUrlReplaceUtil_GetSkuIdBatch(t *testing.T) {
	contents := "下单100024319032下单100024319038"
	reg := regexp.MustCompile("(?m)(?P<key>(抢购|抢购:|抢购：|下单|下单:|下单：))(?P<value>[0-9]+)")
	nameIndex := reg.SubexpIndex("value") // nameIndex = 3
	allIndexes := reg.FindAllStringSubmatch(contents, -1)
	log.Println(allIndexes)
	for _, loc := range allIndexes {
		log.Println(loc[nameIndex])
	}
}

func TestNewUrlReplaceUtil_GetSkuIdV1(t *testing.T) {
	contents := "https://item.m.jd.com/ware/view.action?ptag=138067.15.2&wareId=100024319032&pps=reclike.FO4O305:FOFO0049E9FC383O13O6:FOFOBO2423O1FO3OB23O7FFF5002807FO7O1749E9FC383655E0DC64F838D6&sid=&source=0&scene=0&jxsid=16546120743505638171&appCode=msc588d6d5"
	util := NewUrlReplaceUtil(contents)
	skuId, err := util.GetSkuIdByCustom(SkuRegexpV1, "value")
	log.Println(skuId, err)
	assert.Nil(t, err)
	assert.Equal(t, "100024319032", skuId)
}

func TestNewUrlReplaceUtil_GetSkuIdV2(t *testing.T) {
	contents := "https://wq.jd.com/item/view?ptag=138067.15.2&sku=100024319032&pps=reclike.FO4O305:FOFO0049E9FC383O13O6:FOFOBO2423O1FO3OB23O7FFF5002807FO7O1749E9FC383655E0DC64F838D6&sid=&source=0&scene=0&jxsid=16546120743505638171&appCode=msc588d6d5"
	util := NewUrlReplaceUtil(contents)
	skuId, err := util.GetSkuIdByCustom(SkuRegexpV2, "value")
	log.Println(skuId, err)
	assert.Nil(t, err)
	assert.Equal(t, "100024319032", skuId)
}

func TestNewUrlReplaceUtil_GetSkuIdV3(t *testing.T) {
	contents := "https://item.m.jd.com/product/100010969475.html?gx=RnE2kWBRPmbbn9RJ_t10XTBJ6jA&ad_od=share&utm_source=androidapp&utm_medium=appshare&utm_campaign=t_335139774&utm_term=CopyURL"
	util := NewUrlReplaceUtil(contents)
	skuId, err := util.GetSkuIdByCustom(SkuRegexpV3, "value")
	log.Println(skuId, err)
	assert.Nil(t, err)
	assert.Equal(t, "100010969475", skuId)
}
