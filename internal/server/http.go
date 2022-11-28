package server

import (
	"context"
	"fmt"
	net "net/http"

	v1 "real_world/api/real_world/v1"
	"real_world/internal/conf"
	"real_world/internal/service"
	"real_world/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, rw *service.RealWorldService, logger log.Logger, jwt *conf.JWT) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(jwt.Secret)).Match(NewSkipListMatcher()).Build(),
		),
		http.Filter(
			// Filter demo
			func(next net.Handler) net.Handler {
				return net.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Println("route filter in")
					next.ServeHTTP(w, r)
					fmt.Println("route filter out")
				})
			},
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRealWorldHTTPServer(srv, rw)
	return srv
}

// NewSkipListMatcher jwt中间件放行方法
func NewSkipListMatcher() selector.MatchFunc {
	skipList := make(map[string]struct{})
	// gRPC path 的拼接规则为 /包名.服务名/方法名 详情见官网
	skipList["realworld.v1.RealWorld/Login"] = struct{}{}
	skipList["/realworld.v1.RealWorld/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipList[operation]; ok {
			return false
		}
		return true
	}
}
