package files

import (
	"archive/dbconection"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	File struct {
		ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
		Name        string `json:"name"`
		Description string `json:"description"`
		CategoryId  string `json:"category_id"`
		Title       string `json:"title"`
		Path        string `json:"path"`
		Size        int64  `json:"size"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	Files []File
)

func (f *File) Rest() echo.Map {

	resp := echo.Map{
		"id":          f.ID,
		"name":        f.Name,
		"dexcription": f.Description,
		"category_id": f.CategoryId,
		"title":       f.Title,
		"path":        f.Path,
		"size":        f.Size,
		"created_at":  f.CreatedAt,
		"updated_at":  f.UpdatedAt,
	}
	return resp
}

func (f *Files) Rest() []echo.Map {
	resp := make([]echo.Map, 0)

	for _, r := range *f {
		resp = append(resp, r.Rest())
	}

	return resp
}

func SaveFile(id int64, file *File) (*File, error) {

	if id == 0 {
		dbconection.DB.Create(file)

		return file, nil
	}
	dbconection.DB.Save(file)

	return file, nil
}

func LoadFileById(id int64) (*File, error) {

	file := new(File)
	if err := dbconection.DB.Where("id = ?", id).First(file).Error; err != nil {

		return nil, err
	}

	return file, nil
}

func LoadFiles() (Files, error) {

	files := make(Files, 0)

	dbconection.DB.Find(&files)

	return files, nil
}

func DeleteFile(id int64) error {

	file := new(File)
	if err := dbconection.DB.Where("id = ?", id).First(file).Error; err != nil {
		return errors.New("user not found")
	}

	dbconection.DB.Delete(file)

	return nil
}

func SearchFile(keyword string) (Files, error) {

	files := make(Files, 0)

	dbconection.DB.Where("category_id = ?", keyword).Or("name = ?", keyword).Or("title = ?", keyword).Find(&files)

	return files, nil
}
