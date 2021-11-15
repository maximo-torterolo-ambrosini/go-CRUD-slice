package main

import (
	"course-crud/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var courses models.Courses

func main() {
	courses = append(courses, models.Course{Id: generateId(), Name: "¿Cómo funciona la tecnología?: 1. Computadoras y Smartphones", Price: 0})
	e := echo.New()
	e.GET("/", getCourses)
	e.POST("/", PostCourse)

	e.GET("/:id", getCourse)
	e.PUT("/:id", putCourse)
	e.DELETE("/:id", deleteCourse)
	e.Logger.Fatal(e.Start(":8080"))
}

func generateId() int {
	id := rand.New(rand.NewSource(time.Now().UnixNano()))
	for idExist(id.Intn(10000)) {
		id = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return id.Intn(10000)
}

func idExist(id int) bool {
	for _, course := range courses {
		if course.Id == id {
			return true
		}
	}
	return false
}

func PostCourse(c echo.Context) error {
	body := models.Course{}
	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err)
	}

	if body.Name == "" {
		return c.String(http.StatusNotAcceptable, "Invalid body")
	}

	body.Id = generateId()

	courses = append(courses, body)
	return c.JSON(http.StatusCreated, body)
}

func getCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, courses)
}

func getCourse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusNotFound, "404")
	}
	for _, course := range courses {
		if course.Id == id {
			return c.JSON(http.StatusOK, course)
		}
	}
	return c.JSON(http.StatusNotFound, "404")
}

func deleteCourse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusNotFound, "404")
	}
	for i, _ := range courses {
		if courses[i].Id == id {
			courses = append(courses[:i], courses[i+1:]...)
			return c.JSON(http.StatusOK, courses)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func putCourse(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusNotFound, "404")
	}

	bodyCourse := models.Course{}

	err = c.Bind(&bodyCourse)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err)
	}
	for i, _ := range courses {
		if courses[i].Id == id {
			courses[i].Name = bodyCourse.Name
			courses[i].Price = bodyCourse.Price
			return c.JSON(http.StatusOK, courses[i])
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}
