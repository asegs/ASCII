package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"github.com/Jeffail/gabs"
)

func randomBase64String(l int) string {
	buff := make([]byte, int(math.Round(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l]
}

func deleteFile(filename string){
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}

func saveImageOfLocation(location Location) string{
	url := "https://maps.googleapis.com/maps/api/staticmap?center="+fmt.Sprintf("%f", location.Latitude)+","+fmt.Sprintf("%f", location.Longitude)+"&zoom="+fmt.Sprintf("%d", location.Zoom)+"&size=620x620&maptype=satellite&&key=KEY"
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	name := randomBase64String(128)
	file, err := os.Create("E:\\Go\\asciiServer\\images\\"+name+".png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func saveImageOfAddress(location AddressLocation) string{
	url := "https://maps.googleapis.com/maps/api/staticmap?center=\""+strings.ReplaceAll(location.Address," ","+")+"\"&zoom="+fmt.Sprintf("%d", location.Zoom)+"&size=620x620&maptype=satellite&&key=KEY"
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	name := randomBase64String(128)
	file, err := os.Create("E:\\Go\\asciiServer\\images\\"+name+".png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	var buf bytes.Buffer
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	if int64(int(capacity)) == capacity {
		buf.Grow(int(capacity))
	}
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var n int64 = bytes.MinRead

	if fi, err := f.Stat(); err == nil {
		if size := fi.Size() + bytes.MinRead; size > n {
			n = size
		}
	}
	return readAll(f, n)
}

func ReadToString(filename string) string {
	s,err:=ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(s)
}



func getAddressCoords(address string)Coords{
	url := "https://maps.googleapis.com/maps/api/geocode/json?address="+strings.ReplaceAll(address," ","+")+"&key=AIzaSyDlpGn52XAyQUyy9va5NWITy2DKZcO7CJ4"
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	jsonParsed,_ := gabs.ParseJSON(body)
	lat := jsonParsed.Path("results.0.geometry.location.lat").String()
	lng := jsonParsed.Path("results.0.geometry.location.lng").String()
	latF,_ := strconv.ParseFloat(lat,64)
	lonF,_ := strconv.ParseFloat(lng,64)
	return Coords{
		Latitude:  latF,
		Longitude: lonF,
	}
}
