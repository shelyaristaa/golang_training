package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type article struct {
	ID int
	Title string
	Content string
}

func articles(w http.ResponseWriter, r *http.Request)  {
	var data = []article{
		article{1, "judul pertama", "isi pertama"},
		article{2, "judul pertama", "isi kedua"},
	}

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//marshal untuk merubah dari struct menjadi json
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
	return
}

type starWarsPeople struct{
	Name string `json:"name"`
	Height string `json:"height"`
	Mass string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender string `json:"gender"`
	Films []string `json:"films"`
}

func main() {
	response, _ := http.Get("http://swapi.co/api/people/1")
  	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	
	fmt.Println("raw data: ", string(responseData))
	var People starWarsPeople
	json.Unmarshal(responseData, &People)
	fmt.Println("------Print Field------")
	fmt.Println("Name: ", People.Name)
	fmt.Println("Height: ", People.Height)
	fmt.Println("Mass: ", People.Mass)
	fmt.Println("Hair Color: ", People.HairColor)
	fmt.Println("Film: ", People.Films[0])
	// http.HandleFunc("/articles", articles)
	// fmt.Println("starting web server at http://localhost:8082")
	// http.ListenAndServe(":8082", nil)
}