package handlers

import "net/http"

/* TODO: implement a CORS middleware handler, as described
in https://drstearns.github.io/tutorials/cors/ that responds
with the following headers to all requests:

  Access-Control-Allow-Origin: *
  Access-Control-Allow-Methods: GET, PUT, POST, PATCH, DELETE
  Access-Control-Allow-Headers: Content-Type, Authorization
  Access-Control-Expose-Headers: Authorization
  Access-Control-Max-Age: 600
*/

// CorsMiddleWare adds the given headers when ServeHTTP is called
type CorsMiddleWare struct {
	Handler http.Handler
}

func (c *CorsMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set(headerAccessControlAllowOrigin, allowOriginValue)
		w.Header().Set(headerAccessControlAllowMethods, allowMethodsValue)
		w.Header().Set(headerAccessControlAllowHeaders, allowHeadersValue)
		w.Header().Set(headerAccessControlExposeHeaders, exposeHeadersValue)
		w.Header().Set(headerAccessControlMaxAge, controlMaxAgeValue)
	}
	c.Handler.ServeHTTP(w, r)
}
