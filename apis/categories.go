package apis

import (
	"archive/service/categories"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// GET /categories: Get a list of all categories.
func GetCategories(c echo.Context) error {

	var err error

	ct := make(categories.Categories, 0)
	ct, err = categories.LoadCategoryList()
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while loading categories"})
	}

	return c.JSON(http.StatusOK, echo.Map{"categories": ct.Rest(), "message": "loaded categories successfully"})
}

// GET /categories/{id}: Get details of a specific category by its ID.
func GetCategoryByID(c echo.Context) error {

	var err error

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	ct := new(categories.Category)
	ct, err = categories.LoadCategoryById(intId)
	if err != nil {

		return c.JSON(400, echo.Map{"error": "a problem occurred while loading category"})
	}

	return c.JSON(http.StatusOK, echo.Map{"category": ct.Rest(), "message": "loaded category successfully"})
}

// POST /categories: Create a new category.
func CreateCategory(c echo.Context) error {

	ctForm := new(categories.CategoryForm)
	err := c.Bind(ctForm)
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while binding category form"})
	}

	err = ctForm.Validate()
	if err != nil {

		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	ct := new(categories.Category)
	ct.Name = ctForm.Name
	ct.Description = ctForm.Description
	ct.CreatedAt = time.Now()

	ct, err = categories.SaveCategory(0, ct)
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while saving the category"})
	}

	return c.JSON(http.StatusOK, echo.Map{"category": ct.Rest(), "message": "category create successfully"})
}

// PUT /categories/{id}: Update the details of a specific category by its ID.
func UpdateCategory(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	ctForm := new(categories.CategoryForm)
	err = c.Bind(ctForm)
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while binding category form"})
	}

	err = ctForm.Validate()
	if err != nil {

		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	ct := new(categories.Category)
	ct, err = categories.LoadCategoryById(intId)
	if err != nil {

		return c.JSON(400, echo.Map{"error": "a problem occurred while loading category"})
	}

	ct.Name = ctForm.Name
	ct.Description = ctForm.Description
	ct.UpdatedAt = time.Now()

	ct, err = categories.SaveCategory(intId, ct)
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while updating the category"})
	}

	return c.JSON(http.StatusOK, echo.Map{"category": ct.Rest(), "message": "category updated successfully"})
}

// DELETE /categories/{id}: Delete a specific category by its ID.
func DeleteCategory(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	err = categories.DeleteCategory(intId)
	if err != nil {

		return c.JSON(500, echo.Map{"error": "a problem occurred while deleting category"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Category deleted successfully"})
}
