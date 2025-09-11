package featureflags

type FlagEvaluator interface {
	Evaluate(flag string) bool
	EvaluateWithOrg(flag string) bool
	EvaluateWithCustomVariables(flag string, variables map[string]interface{}) (bool, error)
	EvaluateWithOrgAndCustomVariables(flag string, variables map[string]interface{}) (bool, error)
}
