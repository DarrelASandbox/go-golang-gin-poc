package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=127"` // >=1, <=127
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	//  xml:"title" form:"title" "validate": "email" binding:""
	Title       string `json:"title" binding:"min=2,max=255" validate:"is-cool"`
	Description string `json:"description" binding:"max=255"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
