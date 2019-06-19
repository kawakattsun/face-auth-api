package domains

import (
	"encoding/json"
)

type RequestPostFace struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

func GetRequestPostFace(body string) (*RequestPostFace, error) {
	var r *RequestPostFace
	err := json.Unmarshal([]byte(body), &r)
	return r, err
}
