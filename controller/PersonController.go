package controller

import (
	"challange3_17Mar/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person []model.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error

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

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []model.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
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

func (idb *InDB) CreatePerson(c *gin.Context) {

	// Mendeklarasikan variabel
	var (
		person model.Person
		result gin.H
	)

	// Mengambil data dari form
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	// Memasukkan data ke dalam struct
	person.FirstName = first_name
	person.LastName = last_name

	// Memasukkan data ke dalam database
	idb.DB.Create(&person)

	// Mengembalikan respons
	result = gin.H{
		"result": person,
	}

	// Mengembalikan respons dalam bentuk JSON
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	var (
		person    model.Person
		newPerson model.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data tidak ditemukan",
		}
	}

	newPerson.FirstName = first_name
	newPerson.LastName = last_name
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "Update data gagal",
		}
	} else {
		result = gin.H{
			"result": "Data berhasil diupdate",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person model.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data tidak ditemukan",
		}
	}

	err = idb.DB.Delete(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Hapus data gagal",
		}
	} else {
		result = gin.H{
			"result": "Data berhasil dihapus",
		}
	}

	c.JSON(http.StatusOK, result)
}
