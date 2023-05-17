package subject_test

import (
	"testing"

	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
)

func TestGetSubject(t *testing.T) {
	type args struct {
		serviceName   string
		methodName    string
		subjectPrefix string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should get with empty subject prefix string given",
			args: args{
				serviceName:   "service",
				methodName:    "method",
				subjectPrefix: "",
			},
			want: "service.rpc.method",
		},
		{
			name: "should get with subject prefix given",
			args: args{
				serviceName:   "service",
				methodName:    "method",
				subjectPrefix: "app.us-east1",
			},
			want: "app.us-east1.service.rpc.method",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subject.GetSubject(tt.args.serviceName, tt.args.methodName, tt.args.subjectPrefix); got != tt.want {
				t.Errorf("subject.GetSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}
