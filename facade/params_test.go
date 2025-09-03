package facade

import (
	"testing"

	"github.com/spf13/cobra"
)

// helper to build a command with selected flags defined
func newCmd(with map[string]int) *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	for k, v := range with {
		cmd.Flags().Int(k, v, k+" flag")
	}
	return cmd
}

func TestGetParamsSuccess(t *testing.T) {
	flags := map[string]int{
		"philosophers": 5,
		"to-die":       800,
		"to-eat":       200,
		"to-sleep":     200,
		"to-think":     50,
		"must-eat":     3,
	}
	cmd := newCmd(flags)

	prms, err := getParams(cmd)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if prms.NumPhilos != flags["philosophers"] ||
		prms.TimeDie != flags["to-die"] ||
		prms.TimeEat != flags["to-eat"] ||
		prms.TimeSleep != flags["to-sleep"] ||
		prms.TimeThink != flags["to-think"] ||
		prms.NumEat != flags["must-eat"] {
		t.Errorf("returned params mismatch: %+v", prms)
	}
}

func TestGetParamsMissingFlag(t *testing.T) {
	// Intentionally omit "philosophers"
	flags := map[string]int{
		"to-die":   800,
		"to-eat":   200,
		"to-sleep": 200,
	}
	cmd := newCmd(flags)

	if _, err := getParams(cmd); err == nil {
		t.Errorf("expected error for missing required flag 'philosophers'")
	}
}

func TestGetParamsValidateError(t *testing.T) {
	// Provide all flags so retrieval succeeds, but values likely invalid (zeros)
	flags := map[string]int{
		"philosophers": 0, // required > 0 per spec
		"to-die":       0,
		"to-eat":       0,
		"to-sleep":     0,
		"to-think":     0,
		"must-eat":     -1,
	}
	cmd := newCmd(flags)

	if _, err := getParams(cmd); err == nil {
		t.Errorf("expected validation error, got nil")
	}
}
