package facade

import (
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/philo/params"
)

// getParams extracts execution parameters for the philosophers simulation from the provided
// Cobra command flags, validates them, and returns a populated *params.Params.
//
// Expected integer flags:
//
//	--philosophers : Number of philosophers (required, > 0)
//	--to-die       : Time (ms) before a philosopher dies without eating (required, > 0)
//	--to-eat       : Time (ms) a philosopher spends eating (required, > 0)
//	--to-sleep     : Time (ms) a philosopher spends sleeping (required, > 0)
//	--to-think     : (Optional) Time (ms) spent thinking; may be zero or validated by Params.Validate
//	--must-eat     : (Optional) Number of times each philosopher must eat; non-positive may mean “unbounded”
//
// Behavior:
//  1. Sequentially reads each flag, returning the first retrieval error encountered.
//  2. Populates a params.Params instance with the retrieved values.
//  3. Invokes Params.Validate() to enforce semantic constraints; any validation failure
//     is returned as an error.
//  4. On success, returns the fully validated *params.Params.
//
// Errors:
//   - Flag lookup failures (e.g., missing or malformed flag values)
//   - Validation errors from Params.Validate()
//
// Returns:
//
//	*params.Params on success, or an error if flag extraction or validation fails.
func getParams(cmd *cobra.Command) (*params.Params, error) {
	prms := &params.Params{}
	var err error
	if prms.NumPhilos, err = cmd.Flags().GetInt("philosophers"); err != nil {
		return nil, err
	}
	if prms.TimeDie, err = cmd.Flags().GetInt("to-die"); err != nil {
		return nil, err
	}
	if prms.TimeEat, err = cmd.Flags().GetInt("to-eat"); err != nil {
		return nil, err
	}
	if prms.TimeSleep, err = cmd.Flags().GetInt("to-sleep"); err != nil {
		return nil, err
	}
	if prms.TimeThink, err = cmd.Flags().GetInt("to-think"); err != nil {
		return nil, err
	}
	if prms.NumEat, err = cmd.Flags().GetInt("must-eat"); err != nil {
		return nil, err
	}
	if err = prms.Validate(); err != nil {
		return nil, err
	}
	return prms, nil
}

/* MIT License
 *
 * Copyright 2025 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
