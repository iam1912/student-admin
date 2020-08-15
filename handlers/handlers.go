package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iam1912/student-admin/model"
)

var (
	students = model.NewStuSlice()
)

func LoginHandler(c *gin.Context) {
	var user model.Accout
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	if user.Name == "xjh" && user.Password == "1900" {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

func IndexHandler(c *gin.Context) {
	count := model.Counts()
	pages := int(math.Ceil(float64(count) / 10.000))
	page := c.Query("page")
	intpage, _ := strconv.Atoi(page)

	stu := students.List(intpage)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"data":   stu,
		"counts": pages,
		"page":   page,
	})
}

func SearchHandler(c *gin.Context) {
	id := c.PostForm("id")
	intid, _ := strconv.Atoi(id)
	err := students.FindIndex(intid)
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/index")
	} else {
		stu := students.Search(intid)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"data":   stu,
			"counts": "1",
			"page":   "1",
		})
	}
}

func AddHandler(c *gin.Context) {
	var stu model.Student
	if err := c.ShouldBind(&stu); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err := students.Add(stu)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteHandler(c *gin.Context) {
	id := c.PostForm("delid")
	intid, _ := strconv.Atoi(id)
	err := students.Delete(intid)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func PreviewHandler(c *gin.Context) {
	id := c.Query("id")
	intid, _ := strconv.Atoi(id)
	stu := students.Search(intid)
	c.HTML(http.StatusOK, "preview.html", gin.H{
		"data": stu,
	})
}

func EditHandler(c *gin.Context) {
	var stu model.Student
	if err := c.ShouldBind(&stu); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	students.Modify(stu)
	c.JSON(http.StatusOK, nil)
}
