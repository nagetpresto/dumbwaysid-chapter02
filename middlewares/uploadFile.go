package middleware

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
)

// function with handler argument and handler output (func dalam func)
func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	// return func to process the upload file
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// getting the uploaded file from the form
        file, handler, err := r.FormFile("input-image")
		if err != nil {
			http.Error(w, "Error Retrieving the File: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
        fmt.Printf("Uploaded File: %+v\n", handler.Filename)

		// save the uploaded file to directory using ioutil
        tempFile, err := ioutil.TempFile("uploads", "image-*"+handler.Filename)
		if err != nil {
			http.Error(w, "Path upload error: "+err.Error(), http.StatusBadRequest)
			return
		}
        defer tempFile.Close()

		// read the handler.filename in [bytes] using ioutil.ReadAll
        fileBytes, err := ioutil.ReadAll(file)
        if err != nil {
            fmt.Println(err)
        }
		// storing to the temFile Var
        tempFile.Write(fileBytes)

		// chunking name
        data := tempFile.Name()
		fmt.Println(data)
        filename := data[8:]

		// creating context of the upload name
		// sets the key-value pair of "dataFile" and the filename
        ctx := context.WithValue(r.Context(), "dataFile", filename)

		// passing the name in request context (as argument) to be used (accessible) in other function
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}