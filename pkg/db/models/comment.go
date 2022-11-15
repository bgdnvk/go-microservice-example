package models

type Comment struct {
	ID      int64  `json:"id"`
	Comment string `json:"comment"`
	//time
	User *User `pg:"rel-hasone" json:"user"`
}
