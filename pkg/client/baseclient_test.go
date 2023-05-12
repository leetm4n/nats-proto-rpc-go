package client

import (
	"reflect"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/builtin"
)

func TestBaseClient_SetEncoder(t *testing.T) {
	type fields struct {
		encoder               nats.Encoder
		isClientValidationSet bool
		timeout               time.Duration
	}
	type args struct {
		encoder nats.Encoder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   nats.Encoder
	}{
		{
			name: "should set a nil encoder",
			fields: fields{
				encoder:               nil,
				isClientValidationSet: false,
				timeout:               time.Microsecond,
			},
			args: args{
				encoder: nil,
			},
			want: nil,
		},
		{
			name: "should set a non-nil encoder",
			fields: fields{
				encoder:               nil,
				isClientValidationSet: false,
				timeout:               time.Microsecond,
			},
			args: args{
				encoder: &builtin.DefaultEncoder{},
			},
			want: &builtin.DefaultEncoder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := &BaseClient{
				encoder:               tt.fields.encoder,
				isClientValidationSet: tt.fields.isClientValidationSet,
				timeout:               tt.fields.timeout,
			}
			base.SetEncoder(tt.args.encoder)

			if !reflect.DeepEqual(base.encoder, tt.want) {
				t.Errorf("base.encoder = %v, want %v", base.encoder, tt.want)
			}
		})
	}
}

func TestBaseClient_GetEncoder(t *testing.T) {
	type fields struct {
		encoder               nats.Encoder
		isClientValidationSet bool
		timeout               time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   nats.Encoder
	}{
		{
			name: "should get the nil encoder",
			fields: fields{
				encoder:               nil,
				isClientValidationSet: false,
				timeout:               time.Microsecond,
			},
			want: nil,
		},
		{
			name: "should get the non-nil encoder",
			fields: fields{
				encoder:               &builtin.DefaultEncoder{},
				isClientValidationSet: false,
				timeout:               time.Microsecond,
			},
			want: &builtin.DefaultEncoder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := &BaseClient{
				encoder:               tt.fields.encoder,
				isClientValidationSet: tt.fields.isClientValidationSet,
				timeout:               tt.fields.timeout,
			}
			if got := base.GetEncoder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseClient.GetEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseClient_SetClientValidation(t *testing.T) {
	type fields struct {
		encoder               nats.Encoder
		isClientValidationSet bool
		timeout               time.Duration
	}
	type args struct {
		isSet bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should set validation to true",
			args: args{
				isSet: true,
			},
			fields: fields{
				encoder:               nil,
				isClientValidationSet: false,
				timeout:               time.Microsecond,
			},
			want: true,
		},
		{
			name: "should set validation to false",
			args: args{
				isSet: false,
			},
			fields: fields{
				encoder:               nil,
				isClientValidationSet: true,
				timeout:               time.Microsecond,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := &BaseClient{
				encoder:               tt.fields.encoder,
				isClientValidationSet: tt.fields.isClientValidationSet,
				timeout:               tt.fields.timeout,
			}
			base.SetClientValidation(tt.args.isSet)

			if !reflect.DeepEqual(base.isClientValidationSet, tt.want) {
				t.Errorf("base.isClientValidationSet = %v, want %v", base.isClientValidationSet, tt.want)
			}
		})
	}
}

func TestBaseClient_IsClientValidationEnabled(t *testing.T) {
	type fields struct {
		encoder               nats.Encoder
		isClientValidationSet bool
		timeout               time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "should get true if client validation is set",
			fields: fields{
				encoder:               nil,
				isClientValidationSet: true,
				timeout:               time.Microsecond,
			},
			want: true,
		},
		{
			name: "should get false if client validation is not set",
			fields: fields{
				encoder:               nil,
				isClientValidationSet: false,
				timeout:               time.Microsecond,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := &BaseClient{
				encoder:               tt.fields.encoder,
				isClientValidationSet: tt.fields.isClientValidationSet,
				timeout:               tt.fields.timeout,
			}
			if got := base.IsClientValidationEnabled(); got != tt.want {
				t.Errorf("BaseClient.IsClientValidationEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseClient_SetTimeout(t *testing.T) {
	type fields struct {
		encoder               nats.Encoder
		isClientValidationSet bool
		timeout               time.Duration
	}
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		{
			name: "should set timeout",
			fields: fields{
				encoder:               nil,
				isClientValidationSet: false,
				timeout:               time.Hour,
			},
			args: args{
				timeout: time.Microsecond,
			},
			want: time.Microsecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := &BaseClient{
				encoder:               tt.fields.encoder,
				isClientValidationSet: tt.fields.isClientValidationSet,
				timeout:               tt.fields.timeout,
			}
			base.SetTimeout(tt.args.timeout)

			if !reflect.DeepEqual(base.timeout, tt.want) {
				t.Errorf("base.timeout = %v, want %v", base.timeout, tt.want)
			}
		})
	}
}

func TestBaseClient_GetTimeout(t *testing.T) {
	type fields struct {
		encoder               nats.Encoder
		isClientValidationSet bool
		timeout               time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "should get timeout set",
			fields: fields{
				encoder:               nil,
				isClientValidationSet: false,
				timeout:               time.Hour,
			},
			want: time.Hour,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := &BaseClient{
				encoder:               tt.fields.encoder,
				isClientValidationSet: tt.fields.isClientValidationSet,
				timeout:               tt.fields.timeout,
			}
			if got := base.GetTimeout(); got != tt.want {
				t.Errorf("BaseClient.GetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBaseClient(t *testing.T) {
	tests := []struct {
		name string
		want *BaseClient
	}{
		{
			name: "should create base client with defaults",
			want: &BaseClient{
				encoder:               &builtin.JsonEncoder{},
				isClientValidationSet: true,
				timeout:               DefaultTimeout,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBaseClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
