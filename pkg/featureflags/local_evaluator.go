package featureflags

type LocalEvaluator struct {
	flags map[string]bool
}

func NewLocalEvaluator() *LocalEvaluator {
	return &LocalEvaluator{
		flags: map[string]bool{
			IacAttachCustomFrameworks:  true,
			IacAttachDefaultFrameworks: true,
			IacDisableKicsRule:         false,
		},
	}
}

func (l LocalEvaluator) Evaluate(flag string) bool {
	return l.flags[flag]
}

func (l LocalEvaluator) EvaluateWithOrg(flag string) bool {
	return l.flags[flag]
}

func (l LocalEvaluator) EvaluateWithCustomVariables(flag string, variables map[string]interface{}) (bool, error) {
	return l.flags[flag], nil
}

func (l LocalEvaluator) EvaluateWithOrgAndCustomVariables(flag string, variables map[string]interface{}) (bool, error) {
	return l.flags[flag], nil
}
