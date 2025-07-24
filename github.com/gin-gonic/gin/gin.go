package gin

import (
	"encoding/json"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)
type H map[string]interface{}

type route struct {
	method  string
	pattern string
	handler HandlerFunc
}

type Engine struct {
	routes []route
}

type RouterGroup struct {
	prefix string
	engine *Engine
}

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	params  map[string]string
}

func Default() *Engine { return New() }

func New() *Engine {
	return &Engine{}
}

func (e *Engine) addRoute(method, pattern string, h HandlerFunc) {
	e.routes = append(e.routes, route{method, pattern, h})
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	for _, rt := range e.routes {
		if r.Method != rt.method {
			continue
		}
		params, ok := match(rt.pattern, path)
		if ok {
			c := &Context{Writer: w, Request: r, params: params}
			rt.handler(c)
			return
		}
	}
	http.NotFound(w, r)
}

func (e *Engine) Run(addr ...string) error {
	a := ":8080"
	if len(addr) > 0 {
		a = addr[0]
	}
	return http.ListenAndServe(a, e)
}

func match(pattern, path string) (map[string]string, bool) {
	pParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")
	if len(pParts) != len(pathParts) {
		return nil, false
	}
	params := make(map[string]string)
	for i := range pParts {
		if strings.HasPrefix(pParts[i], ":") {
			params[pParts[i][1:]] = pathParts[i]
		} else if pParts[i] != pathParts[i] {
			return nil, false
		}
	}
	return params, true
}

func (e *Engine) GET(pattern string, h HandlerFunc)  { e.addRoute(http.MethodGet, pattern, h) }
func (e *Engine) POST(pattern string, h HandlerFunc) { e.addRoute(http.MethodPost, pattern, h) }

func (e *Engine) Group(prefix string) *RouterGroup {
	return &RouterGroup{prefix: prefix, engine: e}
}

func (e *Engine) Use(mw ...HandlerFunc) {}

func (g *RouterGroup) GET(path string, h HandlerFunc) {
	g.engine.addRoute(http.MethodGet, g.prefix+path, h)
}
func (g *RouterGroup) POST(path string, h HandlerFunc) {
	g.engine.addRoute(http.MethodPost, g.prefix+path, h)
}
func (g *RouterGroup) Group(path string) *RouterGroup {
	return &RouterGroup{prefix: g.prefix + path, engine: g.engine}
}
func (g *RouterGroup) Use(mw ...HandlerFunc) {}

func (c *Context) JSON(code int, obj interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(obj)
}

func (c *Context) ShouldBindJSON(obj interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(obj)
}

func (c *Context) Param(key string) string { return c.params[key] }

func (c *Context) GetHeader(key string) string { return c.Request.Header.Get(key) }

func (c *Context) AbortWithStatusJSON(code int, obj interface{}) {
	c.JSON(code, obj)
}

func (c *Context) Set(key string, value interface{}) {}
func (c *Context) Next()                             {}
