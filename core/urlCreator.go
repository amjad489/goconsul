package core

func getPrefix(prefix string) string {
	lastN := prefix[len(prefix)-1:]
	if lastN != "/" {
		return prefix + "/"
	} else {
		return lastN
	}

}
func getUrl(consulAddr string, prefix string) (string, string) {
	consulBaseUrl := "http://" + consulAddr + "/v1/kv/"
	if prefix == "all" {
		return consulBaseUrl, consulBaseUrl + "?keys"
	} else {
		prefixPath := getPrefix(prefix)
		return consulBaseUrl, consulBaseUrl + prefixPath + "?keys"
	}
}
