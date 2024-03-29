package opentelemetry

import (
	"github.com/go-web/web"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const instrumentationName = "github.com/onefiter/go-web/web/middlewares/opentelemetry"

type MiddlewareBuilder struct {
	Tracer trace.Tracer
}

func (m MiddlewareBuilder) Build() web.Middleware {
	if m.Tracer == nil {
		m.Tracer = otel.GetTracerProvider().Tracer(instrumentationName)
	}

	return func(next web.HandleFunc) web.HandleFunc {
		return func(ctx *web.Context) {
			reqCtx := ctx.Req.Context()
			reqCtx = otel.GetTextMapPropagator().Extract(reqCtx, propagation.HeaderCarrier(ctx.Req.Header))

			reqCtx, span := m.Tracer.Start(reqCtx, "unknown")
			defer span.End()

			span.SetAttributes(attribute.String("http.method", ctx.Req.Method))
			span.SetAttributes(attribute.String("http.url", ctx.Req.URL.String()))
			span.SetAttributes(attribute.String("http.scheme", ctx.Req.URL.Scheme))
			span.SetAttributes(attribute.String("http.host", ctx.Req.Host))

			// 你这里还可以继续加

			ctx.Req = ctx.Req.WithContext(reqCtx)
			// ctx.Ctx = reqCtx

			// 直接调用下一步
			next(ctx)
			// 这个是只有执行完 next 才可能有值
			span.SetName(ctx.MatchedRoute)

			// 把响应码加上去
			span.SetAttributes(attribute.Int("http.status", ctx.RespStatusCode))
		}
	}
}
