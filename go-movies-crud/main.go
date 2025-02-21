package main
import(
	"fmt"  // using for printing out stuff
	"log"  // log out any errors for when connecting to server (for example)
	"encoding/json" //need for encoding data to JSON for when we sent to POSTMAN for testing
	"math/rand" //create a new id 
	"net/http" //allows to create a server in GoLand
	"strconv"  //string conversion b/c id that is created in math/rand will need to be converted to a string
	"github.com/gorilla/mux" //github library we installed 
)

//No database, just structs (object like in Javascript) and slices!
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"` //* is a pointer to the Director struct, every movie has 1 director
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//slice
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // r is request user sent and mux is the package we installed from the github
	for index, item := range movies {
		if item.ID == params["id"]{
			//           replace this movie -> w/ rest in list (index + 1) and ...
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}
	//now return all exsistng movies
	json.NewEncoder(w).Encode(movies) //sending back all the movies
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//to get movie we want in goland we just need to loop through the exsisting movies in slice
	// find that one and encode into json and just return (to send it)!
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item) //dneding back just one
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // decoded from json
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // to get a new movie ID
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

//delete that movie then add a new movie (NOT THE WAY TO DO THIS IF YOU HAD AN ACTUAL DATABASE)
func updateMovie(w http.ResponseWriter, r *http.Request){
	
	//set json content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)

	//loop over the movies, range
	for index, item := range movies{
		if item.ID == params ["id"]{
			//delete the movie with the ID that we have sent
			movies = append(movies[:index], movies[index + 1:]...)
			//add the new movie - the movie that we sent in the body of postman
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
	
}



// a
func main(){
	r := mux.NewRouter() //function from mux directory from gorilla 

	//fill the slice with two example movies!
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Tropic Thunder", Director : &Director{Firstname: "Ben", Lastname: "Stiller"}})
	movies = append(movies, Movie{ID: "2", Isbn: "433822", Title: "Cat in the Hat", Director: &Director{Firstname: "Dr.", Lastname: "Seuss"}})
	
	//routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/[id]", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/[id]", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/[id]", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
