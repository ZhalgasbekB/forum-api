package convert

import (
	"encoding/json"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

func UUID(r *http.Request) (*model.Sessinon, error) {
	var uuid string
	if err := json.NewDecoder(r.Body).Decode(&uuid); err != nil {
		return nil, err
	}

	return &model.Sessinon{
		UUID: uuid,
	}, nil
}
