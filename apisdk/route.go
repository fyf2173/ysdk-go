package apisdk

type Route struct {
	Group   string
	Method  string
	Path    string
	Desc    string
	RT      string
	Handler interface{}
}

type OptRoute func(*Route)

var (
	OptGin  = OptRouteType("gin")
	OptEcho = OptRouteType("echo")
	OptMux  = OptRouteType("mux")
)

func OptDesc(desc string) OptRoute {
	return func(r *Route) {
		r.Desc = desc
	}
}

func OptRouteType(rt string) OptRoute {
	return func(r *Route) {
		r.RT = rt
	}
}

func OptHandler(handler interface{}) OptRoute {
	return func(r *Route) {
		r.Handler = handler
	}
}

func NewRoute(method, path string, opts ...OptRoute) *Route {
	r := &Route{Method: method, Path: path}
	for _, f := range opts {
		f(r)
	}
	return r
}
