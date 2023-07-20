package files

import "errors"

type FileForm struct {
	Description string `json:"description"`
	CategoryId  string `json:"category_id"`
	Title       string `json:"title"`
}

func (f *FileForm) Validate() error {

	if len(f.Title) < 3 {
		return errors.New(" title must be at least 3 characters")
	}
	return nil
}
