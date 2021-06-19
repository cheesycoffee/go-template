package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo"
	"pkg.agungdp.dev/candi/config/env"
)

// CustomEchoErrorHandler custom echo http error
func CustomEchoErrorHandler(err error, c echo.Context) {
	var message string
	code := http.StatusInternalServerError
	if err != nil {
		message = err.Error()
	}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if code == http.StatusNotFound {
			message = fmt.Sprintf(`Resource "%s %s" not found`, c.Request().Method, c.Request().URL.Path)
		}
	}
	NewHTTPResponse(code, message).JSON(c.Response())
}

// HTTPRoot http handler
func HTTPRoot(serviceName string, buildNumber string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]string{
			"message":   fmt.Sprintf("Service %s up and running", serviceName),
			"timestamp": time.Now().Format(time.RFC3339Nano),
		}
		if env.BaseEnv().BuildNumber != "" {
			payload["build_number"] = env.BaseEnv().BuildNumber
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	}
}

// HTTPMemstatsHandler calculate runtime statistic
func HTTPMemstatsHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	data := struct {
		NumGoroutine int         `json:"num_goroutine"`
		Memstats     interface{} `json:"memstats"`
	}{
		runtime.NumGoroutine(), m,
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
