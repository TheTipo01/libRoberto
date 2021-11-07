package libroberto

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// GetWikipedia returns the first paragraph of the give article
func GetWikipedia(link string) string {
	// Gets article title
	tmp := strings.Split(link, "/")
	title := tmp[len(tmp)-1]

	// Gets wikipedia language
	language := strings.Split(strings.TrimPrefix(link, "https://"), ".")[0]

	res, err := http.Get("https://" + language + ".wikipedia.org/w/api.php?action=query&format=json&titles=" + title + "&prop=extracts&exintro&explaintext")
	if err != nil || http.StatusOK != res.StatusCode {
		return ""
	}

	body, _ := ioutil.ReadAll(res.Body)

	out := string(body)
	_ = res.Body.Close()

	out, _ = strconv.Unquote("\"" + strings.TrimSuffix(strings.Split(out, "\"extract\":\"")[1], "\"}}}}") + "\"")

	return out
}