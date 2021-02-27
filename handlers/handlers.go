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
    router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
    router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
    router.HandleFunc("/modificarperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
    router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
    router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
    
    PORT := os.Getenv("PORT")
    if PORT == ""{
        PORT = "8080"
    }
    handler := cors.AllowAll().Handler(router)
    log.Fatal(http.ListenAndServe(":"+PORT,handler))
}
