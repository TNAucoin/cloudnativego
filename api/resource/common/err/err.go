package err

import "net/http"

var (
	RespDBDataInsertFailure = []byte(`{"error":"db data insert failure"`)
	RespDBDataAccessFailure = []byte(`{"error":"db data access failure"`)
	RespDBDataUpdateFailure = []byte(`{"error":"db data update failure"`)
	RespDBDataRemoveFailure = []byte(`{"error":"db data remove failure"`)
	RespJSONEncodeFailure   = []byte(`{"error":"json encode failure"`)
	RespJSONDecodeFailure   = []byte(`{"error":"json decode failure"`)
	RespInvalidURLParamID   = []byte(`{"error": "invalid url param-id"`)
)

func ServerError(w http.ResponseWriter, reps []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(reps)
}

func BadRequest(w http.ResponseWriter, reps []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(reps)
}

func ValidationErrors(w http.ResponseWriter, reps []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(reps)
}

type Error struct {
	Error error `json:"error"`
}

type Errors struct {
	Error []string `json:"errors"`
}
