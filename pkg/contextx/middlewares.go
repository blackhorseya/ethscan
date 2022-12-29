package contextx

import (
	"context"

	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// AddContextxWitLoggerMiddleware add custom contextx middleware
func AddContextxWitLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(string(KeyCtx), BackgroundWithLogger(logger))

		c.Next()
	}
}

// StreamServerInterceptor serve caller to added Contextx into streaming
func StreamServerInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		newStream := grpc_middleware.WrapServerStream(stream)
		newStream.WrappedContext = context.WithValue(stream.Context(), KeyCtx, BackgroundWithLogger(logger))
		return handler(srv, newStream)
	}
}

// UnaryServerInterceptor serve caller to added Contextx into unary
func UnaryServerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := context.WithValue(c, KeyCtx, BackgroundWithLogger(logger))
		return handler(newCtx, req)
	}
}
