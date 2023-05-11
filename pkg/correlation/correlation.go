package correlation

import "context"

type ContextKey string

const CorrelationIDKey ContextKey = "nats-rpc-correlation-id"
const RequestIDHeaderKey = "x-request-id"

func ContextWithCorrelationID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, CorrelationIDKey, requestID)
}

func CorrelationIDFromContext(ctx context.Context) string {
	requestID, ok := ctx.Value(CorrelationIDKey).(string)
	if !ok {
		return ""
	}

	return requestID
}
