package subject_test

import (
	"testing"

	"github.com/leetm4n/nats-proto-rpc-go/pkg/subject"
)

func TestGetSubject(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := subject.GetSubject(
				test.args.serviceName,
				test.args.methodName,
				test.args.subjectPrefix,
			); got != test.want {
				t.Errorf("subject.GetSubject() = %v, want %v", got, test.want)
			}
		})
	}
}
