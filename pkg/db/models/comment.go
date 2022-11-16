package models

type Comment struct {
	ID      int64  `json:"id"`
	Comment string `json:"comment"`
	User *User `pg:"rel:has-one" json:"user"`
}
