package util

import (
	"regexp"
	"strconv"
	"strings"
)

var RepGlobal *regexp.Regexp
var RepGlobalReplaceNumber *regexp.Regexp

/**
http://vod.cn-shanghai.aliyuncs.com/?AccessKeyId=fdfdfdfd&Acti替换为
http://vod.cn-shanghai.aliyuncs.com/?AccessKeyId=***&Acti
*/
func ReplaceAccessKeyId(requesturl string) string {
	if RepGlobal == nil {
		re3, _ := regexp.Compile("(^|\\?|&)" + "AccessKeyId" + "=([^&]*)(\\s|&|$)")
		RepGlobal = re3
	}

	rep := RepGlobal.ReplaceAllString(requesturl, "?AccessKeyId=***&")
	return rep
}

/**
  http://vod.cn-shanghai.aliyuncs.com/123/123?AccessKeyId=fdfdfdfd&Acti替换为
  http://vod.cn-shanghai.aliyuncs.com/_number_/_number_?AccessKeyId=***&Acti
*/

func ReplaceNumber(requesturl string) string {
	requesturlArray := strings.Split(requesturl, "/")
	result := ""
	for i := 0; i < len(requesturlArray); i++ {
		if IsNum(requesturlArray[i]) {
			result += "/_number_"
		} else {
			index := strings.Index(requesturlArray[i], "?")
			if index != -1 {
				numStr := requesturlArray[i][0:index]
				if IsNum(numStr) {
					result += "/_number_" + requesturlArray[i][index:len(requesturlArray[i])]
				}
			} else {
				result += "/" + requesturlArray[i]
			}
		}
	}

	return result[1:len(result)]
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
