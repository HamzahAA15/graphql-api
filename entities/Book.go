package entities

type Book struct {
	Id        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Author    string `json:"email" form:"email"`
	Publisher string `json:"password" form:"password"`
	Year      int    `json:"year" form:"year"`
}
