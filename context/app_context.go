package context

type AppContext interface {
	Run()
}

type AppCtx struct {
	DbCtx        *DbCtx
	RouteContext *RouteContext
}

func NewAppContext() *AppCtx {
	dbConn := NewDbContext()
	r := NewRouteContext()
	return &AppCtx{DbCtx: dbConn, RouteContext: r}
}

func (r *AppCtx) Run() {
	r.RoutesMapping()
}
