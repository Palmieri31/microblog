package handlers
 
import (
    "log"
    "net/http"
    "os"
	"github.com/gorilla/mux"
    "github.com/Palmieri31/microblog/middlew"
    "github.com/Palmieri31/microblog/routers"
    "github.com/rs/cors"
)
 
/*Manejadores: seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores(){
    router := mux.NewRouter()

    /*Por cada EndPoint vamos a tener un renglon de codigo que permita manejar la funcion correspondiente*/
    router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

    PORT := os.Getenv("PORT")
    if PORT == ""{
        PORT = "8080"
    }
    handler := cors.AllowAll().Handler(router)
    log.Fatal(http.ListenAndServe(":"+PORT,handler))
}
