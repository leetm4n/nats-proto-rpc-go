package correlation

import (
	"context"
	"reflect"
	"testing"
)

func TestContextWithCorrelationID(t *testing.T) {
	type args struct {
		ctx       context.Context
		requestID string
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "should create context with correlation id",
			args: args{
				ctx:       context.Background(),
				requestID: "id",
			},
			want: context.WithValue(context.Background(), CorrelationIDKey, "id"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContextWithCorrelationID(tt.args.ctx, tt.args.requestID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContextWithCorrelationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCorrelationIDFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should get correlation id from context",
			args: args{
				ctx: context.WithValue(context.Background(), CorrelationIDKey, "id"),
			},
			want: "id",
		},
		{
			name: "should get empty correlation id from context if id not set within context",
			args: args{
				ctx: context.Background(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CorrelationIDFromContext(tt.args.ctx); got != tt.want {
				t.Errorf("CorrelationIDFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
