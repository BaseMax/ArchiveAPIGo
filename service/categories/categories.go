package categories

import (
	"archive/dbconection"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	Category struct {
		ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
		Name        string `json:"name"`
		Description string `json:"description"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	Categories []Category
)

func (c *Category) Rest() echo.Map {

	resp := echo.Map{
		"id":          c.ID,
		"name":        c.Name,
		"description": c.Description,
		"created_at":  c.CreatedAt,
		"updated_at":  c.UpdatedAt,
	}
	return resp
}

func (c *Categories) Rest() []echo.Map {
	resp := make([]echo.Map, 0)

	for _, r := range *c {
		resp = append(resp, r.Rest())
	}

	return resp
}

func LoadCategoryList() (Categories, error) {

	c := make(Categories, 0)

	dbconection.DB.Find(&c)

	return c, nil

}

func SaveCategory(id int64, category *Category) (*Category, error) {

	if id == 0 {
		dbconection.DB.Create(category)

		return category, nil
	}
	dbconection.DB.Save(category)

	return category, nil
}

func LoadCategoryById(id int64) (*Category, error) {

	c := new(Category)
	if err := dbconection.DB.Where("id = ?", id).First(c).Error; err != nil {

		return nil, err
	}

	return c, nil
}

func DeleteCategory(id int64) error {

	c := new(Category)
	if err := dbconection.DB.Where("id = ?", id).First(c).Error; err != nil {
		return errors.New("user not found")
	}

	dbconection.DB.Delete(c)

	return nil
}
