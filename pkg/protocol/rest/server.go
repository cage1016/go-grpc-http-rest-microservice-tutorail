package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"time"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cage1016/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	"github.com/cage1016/go-grpc-http-rest-microservice-tutorial/pkg/logger"
	"github.com/cage1016/go-grpc-http-rest-microservice-tutorial/pkg/protocol/rest/middleware"
	"github.com/cage1016/go-grpc-http-rest-microservice-tutorial/pkg/ui/data/swagger"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, httpPort, SwaggerDir string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, "localhost:"+grpcPort, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			logger.Log.Fatal("fail to serve swagger.json", zap.String("Not Found:", r.URL.Path))
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join(SwaggerDir, p)

		logger.Log.Info("Serving swagger-file:", zap.String(p, ""))

		http.ServeFile(w, r, p)
	})
	serveSwaggerUI(mux)

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: middleware.AddRequestID(middleware.AddLogger(logger.Log, mux)),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
