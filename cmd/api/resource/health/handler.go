package health

import "net/http"

// Read godoc
//
// @summary Read Health
// @description Read Health
// @tags health
// @success 200
// @router /../livez [get]
func Read(w http.ResponseWriter, r *http.Request) {}
