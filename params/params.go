package params

import (
	"fmt"

	"github.com/spiegel-im-spiegel/philo/ecode"
)

// Params holds tunable configuration values for a Dining Philosophers simulation.
type Params struct {
	NumPhilos int // Total number of philosophers (and forks). Must be >= 2.
	TimeDie   int // Maximum time in milliseconds a philosopher can go without eating before being considered dead.
	TimeEat   int // Time in milliseconds a philosopher spends eating (holding both forks).
	TimeSleep int // Time in milliseconds a philosopher sleeps after eating.
	TimeThink int // Optional target time in milliseconds a philosopher thinks after sleeping.
	NumEat    int // Optional target number of meals each philosopher must consume; if == 0 the simulation runs until a death occurs.
}

// Validate verifies that all Params fields contain acceptable values.
// It returns an error if any constraint is violated:
//   - NumPhilos must be at least 2 (cannot run with fewer than two philosophers)
//   - TimeDie, TimeEat, and TimeSleep must all be positive (greater than 0)
//   - TimeThink must be zero or positive
//   - NumEat must be zero or positive (zero can indicate unlimited eating)
//
// On success (all fields valid) it returns nil.
func (p *Params) Validate() error {
	if p.NumPhilos < 2 {
		return fmt.Errorf("invalid number of philosophers: %d (%w)", p.NumPhilos, ecode.ErrInvalidParameter)
	}
	if p.TimeDie <= 0 {
		return fmt.Errorf("invalid TimeDie: %d (%w)", p.TimeDie, ecode.ErrInvalidParameter)
	}
	if p.TimeEat <= 0 {
		return fmt.Errorf("invalid TimeEat: %d (%w)", p.TimeEat, ecode.ErrInvalidParameter)
	}
	if p.TimeSleep <= 0 {
		return fmt.Errorf("invalid TimeSleep: %d (%w)", p.TimeSleep, ecode.ErrInvalidParameter)
	}
	if p.TimeThink < 0 {
		return fmt.Errorf("invalid TimeThink: %d (%w)", p.TimeThink, ecode.ErrInvalidParameter)
	}
	if p.NumEat < 0 {
		return fmt.Errorf("invalid NumEat: %d (%w)", p.NumEat, ecode.ErrInvalidParameter)
	}
	return nil
}

// String returns a human-readable summary of the Params configuration,
// listing the number of philosophers, each lifecycle duration (die, eat,
// sleep, think in milliseconds), and the required number of meals.
// Example:
//
//	Philosophers: 5, Die: 800ms, Eat: 200ms, Sleep: 200ms, Think: 0ms, Eat Count: 3
func (p *Params) String() string {
	return fmt.Sprintf("Philosophers: %d, Die: %dms, Eat: %dms, Sleep: %dms, Think: %dms, Eat Count: %d",
		p.NumPhilos, p.TimeDie, p.TimeEat, p.TimeSleep, p.TimeThink, p.NumEat)
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
