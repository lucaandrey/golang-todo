package models

// `todos belongs user, userId is the foreing key`
type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        User   `json:"user"`
	UserId      uint   `json:"userId"`
}

func (b *Todo) TableName() string {
	return "todo"
}
