package main

import (
    "net/http"
    "e-check/routes/app"
)    

func main() { 
   r := routes.SetupRouter()
 
    http.ListenAndServe(":8081", r)
}  
  