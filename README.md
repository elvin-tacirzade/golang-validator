This is a golang request data validation package with simple rules. The package is not completely finished and continues to be developed.
# Installation
```
go get -u github.com/elvin-tacirzade/golang-validator
```
# Usage
```
package main

import (
	"encoding/json"
	"log"
	"net/http"
	validator "github.com/elvin-tacirzade/golang-validator"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string][]string{
			"name":             {"required"},
			"email":            {"required", "email"},
			"logo":             {"image:png,jpg,jpeg", "min:100", "max:1024"},
			"age":              {"required", "numeric"},
			"city":             {"required", "string", "min:3"},
			"salary":           {"required", "numeric", "min:1000", "max:2000"},
			"url":              {"required", "url"},
			"password":         {"required", "min:8"},
			"password_confirm": {"required", "same:password"},
		}
		messages := validator.New(r, data)
		_ = json.NewEncoder(w).Encode(messages)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
# Validation Rules
* **`required`** - The field must be present in the input data and not empty.
* **`email`** - The field must have a valid email.
* **`numeric`** - The field must be entirely numeric characters.
* **`string`** - The field must not consist entirely of numeric characters.
* **`url`** - The field must be a valid URL.
* **`image`** - The field must be a image and verify a file mime type.
* **`min`** - This field can be of 3 different types
    *  `image` -   The field must have a minimum of kilobytes for photo.
    *  `numeric` - The field must have a minimum of numeric for numeric.
    *  `string` - The field must have a minimum of characters for string.
* **`max`** - This field can be of 3 different types
    *  `image` -   The field must have a maximum of kilobytes for photo.
    *  `numeric` - The field must have a maximum of numeric for numeric.
    *  `string` - The field must have a maximum of characters for string.
* **`same`** - The field must be the same as any other field.