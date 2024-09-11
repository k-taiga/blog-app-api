package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	// http.ResponseWriterのinterfaceをフィールドとしてもつ
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeaderメソッドをオーバーライド
func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	// HandlerFunc型はServeHTTPメソッドを持っておりhttp.Handlerのinterfaceを満たす
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.RequestURI, req.Method)

		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Println("res:", rlw.code)
	})
}
