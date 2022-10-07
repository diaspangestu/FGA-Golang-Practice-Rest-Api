package controllers

import (
	"github.com/diaspangestu/FGA-Golang-Practice-Rest-Api/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// to get one data with (id)
func (idb *Controller) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.Masterdb.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data in person
func (idb *Controller) GetAllPerson(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.Masterdb.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to database
func (idb *Controller) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.First_Name = first_name
	person.Last_Name = last_name

	idb.Masterdb.Create(&person)
	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

// update data with (id) as query
func (idb *Controller) UpdatePerson(c *gin.Context) {
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	id := c.Param("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	err := idb.Masterdb.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	newPerson.First_Name = first_name
	newPerson.Last_Name = last_name
	err = idb.Masterdb.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "Update failed",
		}
	} else {
		result = gin.H{
			"result": "Successfully updated data",
		}
	}

	c.JSON(http.StatusOK, result)
}

// delete data with (id)
func (idb *Controller) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.Masterdb.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	err = idb.Masterdb.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "Delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfuly",
		}
	}

	c.JSON(http.StatusOK, result)
}
