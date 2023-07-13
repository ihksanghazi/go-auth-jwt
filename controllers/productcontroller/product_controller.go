package productcontroller

import (
	"net/http"

	"github.com/ihksanghazi/go-auth-jwt/helpers"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":1,
			"name":"Laptop",
			"stock":1000,
		},
		{
			"id":2,
			"name":"Komputer",
			"stock":1000,
		},
		{
			"id":3,
			"name":"Gadget",
			"stock":1000,
		},
	}

	helpers.ResponseJSON(w,http.StatusOK,data)
}