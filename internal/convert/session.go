package convert

import (
	"encoding/json"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

func UUID(r *http.Request) (*model.Session, error) {
	var uuid string
	if err := json.NewDecoder(r.Body).Decode(&uuid); err != nil {
		return nil, err
	}

	return &model.Session{
		UUID: uuid,
	}, nil
}
