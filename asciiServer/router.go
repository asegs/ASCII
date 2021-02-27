package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)


func locationUpdateHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var location Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil{
		fmt.Println(err.Error())
	}
	_ = changeLocation(location)
	photoName := saveImageOfLocation(location)
	_ = json.NewEncoder(w).Encode(Response{
		Body:      handler(photoName,location.Inverse,".png"),
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Photos: getAllPhotosWithinSquare(location.Zoom,location.Latitude,location.Longitude),
	})
	deleteFile("E:\\Go\\asciiServer\\images\\"+photoName+".png")
}

func addressUpdateHandler(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var location AddressLocation
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil{
		fmt.Println(err.Error())
	}
	photoName := saveImageOfAddress(location)
	coords := getAddressCoords(location.Address)
	_ = json.NewEncoder(w).Encode(Response{
		Body:      handler(photoName,location.Inverse,".png"),
		Latitude:  coords.Latitude,
		Longitude: coords.Longitude,
		Photos: getAllPhotosWithinSquare(location.Zoom,coords.Latitude,coords.Longitude),
	})
	deleteFile("E:\\Go\\asciiServer\\images\\"+photoName+".png")
}

func userCreationHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		fmt.Println(err.Error())
	}
	success := createUser(user)
	_ = json.NewEncoder(w).Encode(success)
}

func userLoginHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		fmt.Println(err.Error())
	}
	success := login(user)
	_ = json.NewEncoder(w).Encode(success)
}

func savePhotoHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	r.ParseMultipartForm(5*1024*1024)
	file, _, _ := r.FormFile("file")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")
	extension := r.FormValue("extension")
	if extension != ".png" && extension != ".jpeg" && extension != ".jpg"{
		w.WriteHeader(403)
		return
	}
	name := randomBase64String(96)
	if extension == ".jpg"{
		extension = ".jpeg"
	}

	f, _ := os.Create("E:\\Go\\asciiServer\\images\\"+name+extension)
	defer f.Close()
	io.Copy(f, file)
	var photo Photo
	FLongitude,_ := strconv.ParseFloat(longitude,64)
	FLatitude,_ := strconv.ParseFloat(latitude,64)
	photo.Longitude = FLongitude
	photo.Latitude = FLatitude
	photo.Name = name
	success := uploadPhoto(photo)
	str := handler(name,false,extension)
	fmt.Println(name)
	Write("E:\\Go\\asciiServer\\textfiles\\"+name+".txt",str)
	_ = json.NewEncoder(w).Encode(success)
	deleteFile("E:\\Go\\asciiServer\\images\\"+name+".png")
	defer file.Close()
}

func viewPhoto(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var name Wrapper
	_ = json.NewDecoder(r.Body).Decode(&name)
	_ = json.NewEncoder(w).Encode( ReadToString("E:\\Go\\asciiServer\\textfiles\\"+name.Name+".txt"))
}




func main() {
	r := mux.NewRouter()
	r.HandleFunc("/changelocation",locationUpdateHandler).Methods("PUT")
	r.HandleFunc("/changeaddress",addressUpdateHandler).Methods("PUT")
	r.HandleFunc("/createuser",userCreationHandler).Methods("POST")
	r.HandleFunc("/login",userLoginHandler).Methods("PUT")
	r.HandleFunc("/uploadphoto",savePhotoHandler).Methods("POST")
	r.HandleFunc("/viewphoto",viewPhoto).Methods("PUT")
	log.Println("http server started on :8888")
	err := http.ListenAndServe(":8888",r)
	if err != nil{
		log.Fatal("ListenAndServe: ",err)
	}




}
