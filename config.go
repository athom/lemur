package lemur

import (
	"strconv"
	"strings"
)

var version = 2.0

func getPostfixURL() string {
	versionStr := strconv.FormatFloat(version, 'f', 1, 32)
	return "api.mailchimp.com/" + versionStr + "/"
}

func getBaseURL(apiKey string) (r string) {
	return `https://` + getDataCenterViaApiKey(apiKey) + `.` + getPostfixURL()
}

func getDataCenterViaApiKey(apiKey string) string {
	ss := strings.Split(apiKey, `-`)
	if len(ss) != 2 {
		panic("wrong api key")
	}
	return ss[1]
}
