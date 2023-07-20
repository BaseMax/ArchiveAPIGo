package categories

import "errors"

type CategoryForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *CategoryForm) Validate() error {

	if c.Name == "" {
		return errors.New("name must not be empty")
	}

	if c.Description == "" {
		return errors.New("description must not be empty")
	}

	return nil
}
