package subject

import "fmt"

// GetSubject requires serviceName, methodName and an optional prefix
type GetSubjectFn func(serviceName, methodName, prefix string) string

var _ GetSubjectFn = GetSubject

// Implements GetSubjectFn outputting the following
// ${prefix}.${serviceName}.rpc.${methodName}
func GetSubject(serviceName, methodName, prefix string) string {
	if prefix == "" {
		return fmt.Sprintf("%s.rpc.%s", serviceName, methodName)
	}
	return fmt.Sprintf("%s.%s.rpc.%s", prefix, serviceName, methodName)
}
