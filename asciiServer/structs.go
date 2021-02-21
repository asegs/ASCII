package main

type Location struct {
	Latitude float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Zoom int `json:"zoom,omitempty"`
	Inverse bool `json:"inverse"`
}

type AddressLocation struct {
	Address string `json:"address,omitempty"`
	Zoom int `json:"zoom,omitempty"`
	Inverse bool `json:"inverse"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Photo struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name string `json:"name"`
}

type Response struct {
	Body string `json:"body"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Photos []Photo `json:"photos"`
}

type Coords struct {
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Wrapper struct {
	Name string `json:"name"`
}
