package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	//"bytes"
	"net/http"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

type Place struct {
	Id   int
	Name string
	Ctry string
	Desc string
	Lat  string
	Lon  string
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/place", CreatePlace)
		v1.GET("/place/:id", GetPlace)
		v1.GET("/places", GetAllPlaces)
		v1.DELETE("/place", DeletePlace)
		v1.GET("/health", HealthCheck)
		v1.GET("/url/:id", GetPlaceUrl)
	}
	return router
}

func CreatePlace(c *gin.Context) {
	// var buffer bytes.Buffer
	var newPlace Place

	fmt.Println(newPlace.Name)
	//stmt, err := Init().Prepare("insert into places (name, country, description, latitude, longitude) values(?,?,?,?,?);")

	//if err != nil {
	//	fmt.Print(err.Error())
	//}
	//_, err = stmt.Exec(name, country, description, latitude, longitude)
    
	//if err != nil {
	//	fmt.Print(err.Error())
	//}

	//buffer.WriteString(name)
	//buffer.WriteString(" ")
	//defer stmt.Close()
	//placename := buffer.String()

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created"),
	})
}

func GetAllPlaces(c *gin.Context) {
	var (
		place  Place
		places []Place
	)
	rows, err := Init().Query("select * from places;")
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
		place  Place
		result gin.H
	)
	id := c.Param("id")
	err := Init().QueryRow("select * from places where id = ?;", id).Scan(&place.Id, &place.Name, &place.Ctry, &place.Desc, &place.Lat, &place.Lon)

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
	stmt, err := Init().Prepare("delete from places where id= ?;")

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
		place  Place
		result gin.H
	)
	id := c.Param("id")
	err := Init().QueryRow("select * from places where id = ?;", id).Scan(&place.Id, &place.Name, &place.Ctry, &place.Desc, &place.Lat, &place.Lon)

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
	
	Url := "https://maps.google.com/maps?q=" + place.Lat + "," + place.Lon + "&t=k"
	result = gin.H {
		"result": Url,
	}

	c.IndentedJSON(http.StatusOK, result)
}

func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": 200, "data": "testing api", "alive": true})
}

func Init() *sql.DB {
	mysqlHost, provided := os.LookupEnv("MARIADB_HOST")
	if !provided {
		mysqlHost = "localhost"
	}
	mysqlUser, provided := os.LookupEnv("MARIADB_USER")
	if !provided {
		log.Fatalf("Environment variable %s is not set", "MARIADB_USER")
	}
	mysqlPass, provided := os.LookupEnv("MARIADB_PASS")
	if !provided {
		log.Fatalf("Environment variable %s is not set", "MARIADB_PASS")
	}
	config := mysql.Config{
		User:                 mysqlUser,
		Passwd:               mysqlPass,
		Net:                  "tcp",
		Addr:                 mysqlHost,
		AllowNativePasswords: true,
		DBName:               "placesdb",
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	checkErr(err)

	err = db.Ping()
	checkErr(err)
	fmt.Printf("Connection successfully")

	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
