package context

import "github.com/foody-go-api/aws/s3"

type AppContext interface {
	Run()
}

type AppCtx struct {
	DbCtx        *DbCtx
	RouteContext *RouteContext
	s3Provider   s3.AwsS3Provider
}

func NewAppContext() *AppCtx {
	dbConn := NewDbContext()
	r := NewRouteContext()
	s3Provider := s3.NewS3Provider()
	s3Service := AwsS3Provider(s3Provider)
	return &AppCtx{DbCtx: dbConn, RouteContext: r, s3Provider: s3Service}
}

func (r *AppCtx) Run() {
	r.RoutesMapping()
}

func AwsS3Provider(upProvider s3.AwsS3Provider) s3.AwsS3Provider {
	return upProvider
}
