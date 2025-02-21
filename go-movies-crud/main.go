package main
import(
	"fmd"  // using for printing out stiff
	"log"  // log out any errors for when connecting to server (for example)
	"encoding/json" //need for encoding data to JSON for when we sent to POSTMAN for testing
	"math/rand" //create a new id 
	"net/http" //allows to create a server in GoLand
	"strconv"  //string conversion b/c id that is created in math/rand will need to be converted to a string
	"github.com/gorilla/mux" //github library we installed 
)

//No database, just structs and slices!
