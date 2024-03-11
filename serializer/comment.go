package serializer

import "4-TikTok/model"

type Comment struct {
	ID        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	VideoId   int64  `json:"video_id"`
	LikeCount uint   `json:"like_count"`
	Comment   string `json:"comment"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

func BuildComment(item model.Comment) Comment {
	return Comment{
		ID:        item.ID,
		UserId:    item.UserId,
		VideoId:   item.VideoId,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
		DeletedAt: item.DeletedAt.Unix(),
	}
}
func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}
