package main

import (
    "net/http"
    "github.com/e-check/internal/app"
)

   
func main() {
   r := app.SetupRouter()
 
    http.ListenAndServe(":8081", r)
}  
  