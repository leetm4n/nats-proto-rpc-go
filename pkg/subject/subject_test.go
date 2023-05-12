package subject

import "testing"

var subjectPrefix string = "app.us-east1"
var subjectPrefixEmptyString string = ""

func TestGetSubject(t *testing.T) {
	type args struct {
		serviceName   string
		methodName    string
		subjectPrefix *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should get without subject prefix given",
			args: args{
				serviceName:   "service",
				methodName:    "method",
				subjectPrefix: nil,
			},
			want: "service.rpc.method",
		},
		{
			name: "should get with empty subject prefix string given",
			args: args{
				serviceName:   "service",
				methodName:    "method",
				subjectPrefix: &subjectPrefixEmptyString,
			},
			want: "service.rpc.method",
		},
		{
			name: "should get with subject prefix given",
			args: args{
				serviceName:   "service",
				methodName:    "method",
				subjectPrefix: &subjectPrefix,
			},
			want: "app.us-east1.service.rpc.method",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSubject(tt.args.serviceName, tt.args.methodName, tt.args.subjectPrefix); got != tt.want {
				t.Errorf("GetSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}
