package main

import (
    "log"
    "github.com/Palmieri31/microblog/handlers"
    "github.com/Palmieri31/microblog/bd"
)


func main(){
	if bd.ChequeoConnection() == 0 {
        log.Fatal("Sin conexión a la BD")
        return
    }
    handlers.Manejadores()

}