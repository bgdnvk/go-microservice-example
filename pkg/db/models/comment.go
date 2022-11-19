package models

import "github.com/go-pg/pg/v10"

type Comment struct {
	ID      int64  `json:"id"`
	Comment string `json:"comment"`
	UserID  int64  `json:"user_id"`
	User    *User  `pg:"rel:has-one" json:"user"`
}

func CreateComment(db *pg.DB, req *Comment) (*Comment, error) {
	_, err := db.Model(req).Insert()
	if err != nil {
		return nil, err
	}

	comment := &Comment{}

	err = db.Model(comment).
		Relation("User").
		Where("comment.id = ?", req.ID).
		Select()

	return comment, err
}

func GetComment(db *pg.DB, commentID string) (*Comment, error) {
	comment := &Comment{}

	err := db.Model(comment).
		Relation("User").
		Where("comment.id = ?", commentID).
		Select()

	return comment, err
}

func GetComments(db *pg.DB) ([]*Comment, error) {
	comments := make([]*Comment, 0)

	err := db.Model(&comments).
		Relation("User").
		Select()

	return comments, err
}
