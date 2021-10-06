package context

type AppContext interface {
	RunService() error
}

type appCtx struct {
	dbCtx *DbCtx
	routeContext *RouteContext
}

func NewAppContext() *appCtx {
	dbConn := NewDbContext()
	r := NewRouteContext()
	return &appCtx{dbCtx: dbConn, routeContext: r}
}

func (a *appCtx) RunService() error {
	a.routeContext.RunRouteContext(a.dbCtx)

	if err := a.routeContext.Run(); err != nil {
		return err
	}

	return nil
}