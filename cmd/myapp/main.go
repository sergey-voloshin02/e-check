package main

import (
    "net/http"
    "github.com/sergey-voloshin02/e-check/internal/app"
) 
 
func main() {
   r := app.SetupRouter()
 
    http.ListenAndServe(":8081", r)
}  
  