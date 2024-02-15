package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/***** Basic Logger  ****/
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing middlewareOne")
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	fmt.Println(r.Method, r.URL.Path, time.Since(start))
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))

}

/**** NewLogger constructs a new Logger middleware handler ****/
func NewLoggerMiddleware(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
