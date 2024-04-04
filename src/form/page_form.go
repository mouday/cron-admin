package form

type PageForm struct {
	Page   int `json:"page"`
	Size   int `json:"size"`
	Status int `json:"status"`
}

func (pageForm PageForm) PageOffset() int {
	return (pageForm.Page - 1) * pageForm.Size
}
