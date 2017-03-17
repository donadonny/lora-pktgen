package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hokaccha/go-prettyjson"
)

//dont do this, see above edit
func prettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}

func GetPrettyJson(body string) string {
	pj, err := prettyjson.Format([]byte(body))
	if err != nil {
		return err.Error()
	} else {
		return string(pj)
	}
}

func GetHttpRespBody(rsp *http.Response, err error) string {
	if err != nil {
		return err.Error()
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}
