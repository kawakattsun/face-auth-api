package domains

import (
	"encoding/json"
)

type RequestPostFace struct {
	Name string `json:"name"`
	body string `json:"body"`
}

func (r *RequestPostFace) GetRequestPostFace(body []byte) RequestPostFace, error {
	req := &r{}
	err := json.Unmarshal([]byte(body), &req)
	return req, err
}
