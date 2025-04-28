package pkg

import (
	"encoding/json"
	"net/http"
)

func UnmarshalBody(r *http.Request, output any) error {
	return json.NewDecoder(r.Body).Decode(&output)
}
