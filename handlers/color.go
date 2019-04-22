package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yashdalfthegray/color-info/utils"

	"github.com/yashdalfthegray/color-info/models"
)

// ColorHandler responds with an RGB object or an error when you POST
// with a hex color string
func ColorHandler(w http.ResponseWriter, req *http.Request) {
	var b models.ColorRequestBody

	jsonEnc := json.NewEncoder(w)
	jsonDec := json.NewDecoder(req.Body)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if decErr := jsonDec.Decode(&b); decErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if !utils.ValidateColorString(b.Color) {
			w.WriteHeader(http.StatusBadRequest)
			jsonEnc.Encode(models.NewErrorResponse("Malformed color string"))
			return
		}

		rgbColor, convertErr := models.HexCode(b.Color).ToRGB()
		if convertErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			jsonEnc.Encode(models.NewErrorResponse("Could not convert given color"))
			return
		}
		w.WriteHeader(http.StatusOK)

		encErr := json.NewEncoder(w).Encode(models.NewColorResponse(rgbColor))
		if encErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			jsonEnc.Encode(models.NewErrorResponse(encErr.Error()))
		}
	}
}
