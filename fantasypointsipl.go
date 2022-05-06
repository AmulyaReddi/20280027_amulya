package main
import ("encoding/json"
    "log"
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
)
type Player struct {
    ID        string   `json:"ID,omitempty"`
    Name string   `json:"Name,omitempty"`
    Team  string   `json:"team,omitempty"`
}
type Employee struct {
    ID        string   `json:"ID,omitempty"`
    Score   *Score `json:"score,omitempty"`
}
type Score struct {
    Match  string `json:"Match,omitempty"`
    Runs string `json:"Runs,omitempty"`
    Wickets string `json:"Wickets,omitempty"`
}

type Fservice struct {
    PurpleCap      string   `json:"PurpleCap,omitempty"`
    OrangeCap      string   `json:"OrangeCap,omitempty"`
}
type Fscore struct {
    Name     string   `json:"Name,omitempty"`
    Fantasyscore      string   `json:"Fantasyscore,omitempty"`
}
var fservice []Fservice
var emp []Employee
var ply []Player
var fscore []Fscore

func Getcapholder(w http.ResponseWriter, req *http.Request) {
 json.NewEncoder(w).Encode(fservice)

}
func Getfancyscore(w http.ResponseWriter, req *http.Request) {
 json.NewEncoder(w).Encode(fscore)

}

//getplayer
func Getplayers(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(ply)
}
//createplayer
func createPlayer(w http.ResponseWriter, r *http.Request) {
var player Player
reqBody, err := ioutil.ReadAll(r.Body)
if err != nil {
fmt.Fprintf(w, "Kindly enter the  data")
}
json.Unmarshal(reqBody, &player)
ply= append(ply, player)
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(ply)
}

//getscore
func GetScore(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(emp)
}
func Createscore(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Employee
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    emp = append(emp, person)
    json.NewEncoder(w).Encode(emp)
}

func main() {
    router := mux.NewRouter()
    emp = append(emp, Employee{ID: "1",
        Score: &Score{Match: "1", Runs: "20",Wickets:"0"}})
   ply= append(ply,Player{ID:"1",Name:"warner",Team:"srh"})
   ply= append(ply,Player{ID:"2",Name:"kane",Team:"srh"})
fservice=append(fservice,Fservice{PurpleCap:"warner",OrangeCap:"kane"})
fscore=append(fscore,Fscore{Name:"warner",Fantasyscore:"100"})
    router.HandleFunc("/players/scores", GetScore).Methods("GET")
    router.HandleFunc("/capholder", Getcapholder).Methods("GET")
    router.HandleFunc("/players", Getplayers).Methods("GET")
router.HandleFunc("/fantasyscore",Getfancyscore).Methods("GET")
    router.HandleFunc("/player/{id}/score", Createscore).Methods("POST")
    router.HandleFunc("/player", createPlayer).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", router))
}