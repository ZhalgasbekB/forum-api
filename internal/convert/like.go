package convert

import (
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

func NewConvertVote(r *http.Request, session *model.Session) (*model.LikePost, error) {
	// numIdPost, err := ConvertDatePost(r.URL.Path)
	// if err != nil {
	// 	return nil, err
	// }
	// vote, err := strconv.Atoi(r.FormValue("vote"))
	// if err != nil {
	// 	return nil, err
	// }
	// voteModel := &model.Vote{}
	// if err := json.NewDecoder(r.Body).Decode(voteModel); err != nil {
	// 	return nil, err
	// }
	// return &model.Vote{
	// 	UserId: session.UserID,
	// 	PostId: numIdPost,
	// 	Vote:   vote,
	// }, nil
	return nil, nil
}