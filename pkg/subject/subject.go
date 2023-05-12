package subject

import "fmt"

// GetSubject returns the following format:
// ${subjectPrefix}.${serviceName}.rpc.${methodName}
func GetSubject(serviceName, methodName string, subjectPrefix *string) string {
	if subjectPrefix == nil || *subjectPrefix == "" {
		return fmt.Sprintf("%s.rpc.%s", serviceName, methodName)
	}
	return fmt.Sprintf("%s.%s.rpc.%s", *subjectPrefix, serviceName, methodName)
}
