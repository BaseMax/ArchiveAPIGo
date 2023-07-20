package apis

import (
	"archive/service/files"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// GET /files: Get a list of all files with optional sorting and filtering parameters.
func GetFiles(c echo.Context) error {

	fs := make(files.Files, 0)
	fs, err := files.LoadFiles()
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while loading files"})
	}

	return c.JSON(http.StatusOK, echo.Map{"file": fs.Rest(), "message": "files loaded successfully"})
}

// GET /files/{id}: Get details of a specific file by its ID.
func GetFileById(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	file := new(files.File)
	file, err = files.LoadFileById(intId)
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while loading file"})
	}

	return c.JSON(http.StatusOK, echo.Map{"file": file.Rest(), "message": "file loaded successfully"})
}

// POST /files: Upload a new file with optional caption, description, and category assignment.
func UploadFile(c echo.Context) error {

	var (
		isSuccess bool = true
		fileByte  []byte
		fileType  string
		fileName  string
		fileSize  int64
		filePath  string
	)

	fileForm := new(files.FileForm)
	title := c.QueryParam("title")
	description := c.QueryParam("description")
	categoryId := c.QueryParam("categoryId")

	fileForm.Title = title
	fileForm.CategoryId = categoryId
	fileForm.Description = description

	err := fileForm.Validate()
	if err != nil {

		return c.JSON(400, echo.Map{"error": "a problem occurred while validating"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		isSuccess = false
	} else {
		src, err := file.Open()
		if err != nil {
			isSuccess = false
		} else {
			fileByte, _ = ioutil.ReadAll(src)
			fileType = http.DetectContentType(fileByte)

			if fileType == "application/pdf" {
				fileName = strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
				filePath = "upload/" + fileName
			} else {

				return c.JSON(400, echo.Map{"error": "Invalid file type just upload pdf type"})
			}

			err := ioutil.WriteFile(filePath, fileByte, 0777)
			if err != nil {
				isSuccess = false
			} else {
				fileSize = file.Size
			}
		}
		defer src.Close()
	}

	f := new(files.File)

	if isSuccess {
		f.Name = fileName
		f.Description = fileForm.Description
		f.CategoryId = fileForm.CategoryId
		f.Size = fileSize
		f.Path = filePath
		f.Title = fileForm.Title
		f.CreatedAt = time.Now()
	}

	f, err = files.SaveFile(0, f)
	if err != nil {

		return c.JSON(500, echo.Map{"error": " a problem occurred while saving the file"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "file uploaded successfully", "file": f.Rest()})
}

// PUT /files/{id}: Update the details of a specific file by its ID.
func UpdateFile(c echo.Context) error {

	var (
		err       error
		isSuccess bool = true
		fileByte  []byte
		fileType  string
		fileName  string
		fileSize  int64
		filePath  string
	)

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	oldFile := new(files.File)
	oldFile, err = files.LoadFileById(intId)
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while loading old file"})
	}

	fileForm := new(files.FileForm)
	title := c.QueryParam("title")
	description := c.QueryParam("description")
	categoryId := c.QueryParam("categoryId")

	fileForm.Title = title
	fileForm.CategoryId = categoryId
	fileForm.Description = description

	err = fileForm.Validate()
	if err != nil {

		return c.JSON(400, echo.Map{"error": "a problem occurred while validating"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		isSuccess = false
	} else {
		src, err := file.Open()
		if err != nil {
			isSuccess = false
		} else {
			fileByte, _ = ioutil.ReadAll(src)
			fileType = http.DetectContentType(fileByte)

			if fileType == "application/pdf" {
				fileName = strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
				filePath = "upload/" + fileName
			} else {

				return c.JSON(400, echo.Map{"error": "Invalid file type just upload pdf type"})
			}

			err := ioutil.WriteFile(filePath, fileByte, 0777)
			if err != nil {
				isSuccess = false
			} else {
				fileSize = file.Size
			}
		}
		defer src.Close()
	}

	f := new(files.File)

	if isSuccess {
		f.ID = oldFile.ID
		f.Name = fileName

		if len(fileForm.Description) == 0 {
			f.Description = oldFile.Description
		} else {
			f.Description = fileForm.Description
		}

		if len(fileForm.CategoryId) == 0 {
			f.CategoryId = oldFile.CategoryId
		} else {
			f.CategoryId = fileForm.CategoryId
		}

		if len(fileForm.Title) == 0 {
			f.Title = oldFile.Title
		} else {
			f.Title = fileForm.Title
		}

		f.Size = fileSize
		f.Path = filePath
		f.CreatedAt = oldFile.CreatedAt
		f.UpdatedAt = time.Now()
	}

	f, err = files.SaveFile(intId, f)
	if err != nil {

		return c.JSON(500, echo.Map{"error": " a problem occurred while saving the file"})
	}

	return c.JSON(http.StatusOK, echo.Map{"file": f.Rest()})
}

// DELETE /files/{id}: Delete a specific file by its ID.
func DeleteFile(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	err = files.DeleteFile(intId)
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while deleting file"})
	}

	return c.JSON(http.StatusOK, echo.Map{"id": id, "message": "delete file successfully"})
}

// GET /files/{id}/public-link: Generate a public link for a specific file by its ID.
func GetPublicLink(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	f := new(files.File)
	f, err = files.LoadFileById(intId)
	if err != nil {

		return c.JSON(400, echo.Map{"error": "a problem occurred while loading file", "message": "id is not a valid "})
	}

	publicLink := fmt.Sprintf("http://localhost:8080/f/files/%d", f.ID)

	return c.JSON(http.StatusOK, echo.Map{"publicLink": publicLink})
}

// GET /search?q={keyword}: Search for files based on a keyword.
func SearchFile(c echo.Context) error {

	keyword := c.Param("keyword")
	fmt.Println(keyword)
	file := make(files.Files, 0)
	file, err := files.SearchFile(keyword)
	if err != nil {

		return c.JSON(400, echo.Map{"error": "not matched any file"})
	}
	return c.JSON(http.StatusOK, echo.Map{"file": file.Rest()})
}
