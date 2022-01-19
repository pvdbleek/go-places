package controllers

import (
	"bytes"
	"fmt"
	"github.com/pvdbleek/go-places/db"
	"github.com/pvdbleek/go-places/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePlace(c *gin.Context) {
	var buffer bytes.Buffer
	Name := c.PostForm("name")
	Ctry := c.PostForm("country")
	Desc := c.PostForm("description")
	Lat := c.PostForm("latitude")
	Lon := c.PostForm("longitude")

	stmt, err := db.Init().Prepare("insert into places (name, country, description, latitude, longitude) values(?,?,?,?,?);")

	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(Name, Ctry, Desc, Lat, Lon)

	if err != nil {
		fmt.Print(err.Error())
	}

	buffer.WriteString(Name)
	buffer.WriteString(" ")
	defer stmt.Close()
	placename := buffer.String()

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created", placename),
	})
}

func GetAllPlaces(c *gin.Context) {
	var (
		place  models.Place
		places []models.Place
	)
	rows, err := db.Init().Query("select * from places;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&place.Id, &place.Name, &place.Ctry, &place.Desc, &place.Lat, &place.Lon)
		places = append(places, place)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.IndentedJSON(http.StatusOK, gin.H{
		"result": places,
	})
}

func GetPlace(c *gin.Context) {
	var (
		place  models.Place
		result gin.H
	)
	id := c.Param("id")
	err := db.Init().QueryRow("select * from places where id = ?;", id).Scan(&place.Id, &place.Name, &place.Ctry, &place.Desc, &place.Lat, &place.Lon)

	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": place,
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}

func DeletePlace(c *gin.Context) {
	id := c.Query("id")
	stmt, err := db.Init().Prepare("delete from places where id= ?;")

	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted place: %s", id),
	})
}

func GetPlaceUrl(c *gin.Context) {
	var (
		place  models.Place
		result gin.H
	)
	id := c.Param("id")
	err := db.Init().QueryRow("select * from places where id = ?;", id).Scan(&place.Id, &place.Name, &place.Ctry, &place.Desc, &place.Lat, &place.Lon)

	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": place,
		}
	}
	UrlLat := strconv.FormatFloat(place.Lat, 'f', -1, 64)
	UrlLon := strconv.FormatFloat(place.Lon, 'f', -1, 64)
	Url := "https://maps.google.com/maps?q=" + UrlLat + "," + UrlLon + "&t=k"
	result = gin.H {
		"result": Url,
	}

	c.PureJSON(http.StatusOK, result)
}
