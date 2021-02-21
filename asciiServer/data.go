package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
)

func connect()*sql.DB{
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/ascii")
	if err != nil{
		panic(err.Error())
	}
	return db
}

func changeLocation(location Location) bool{
	db := connect()
	fmt.Println(db)
	db.Close()
	return true
}

func createUser(user User) bool{
	db := connect()
	fmt.Println(db)
	db.Close()
	return true
}

func login(user User)bool{
	db := connect()
	fmt.Println(db)
	db.Close()
	return true
}

func uploadPhoto(photo Photo)bool{
	db := connect()
	insert,err := db.Query("INSERT INTO photos (latitude,longitude,name) VALUES (?,?,?)",photo.Latitude,photo.Longitude,photo.Name)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	defer db.Close()
	return true
}

func getAllPhotosWithinSquare(zoom int,latitude float64,longitude float64) []Photo{
	photos := make([]Photo,0)
	circumference := math.Abs(math.Cos(latitude))*12756*1000*math.Pi
	r := 71*(math.Pow(2, float64(21-zoom)))
	radius := r/circumference*180
	lowerLat := latitude-radius
	upperLat := latitude+radius
	lowerLon := longitude-radius
	upperLon := longitude+radius
	db := connect()
	results,_ := db.Query("SELECT latitude,longitude,name FROM photos WHERE latitude>? AND latitude<? AND longitude<? AND longitude>?",lowerLat,upperLat,upperLon,lowerLon)
	for results.Next(){
		var photo Photo
		_ = results.Scan(&photo.Latitude,&photo.Longitude,&photo.Name)
		photos = append(photos,photo)
	}
	return photos
}

