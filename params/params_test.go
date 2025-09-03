package params

import (
	"errors"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/philo/ecode"
)

func TestParamsValidate_Valid(t *testing.T) {
	p := &Params{
		NumPhilos: 5,
		TimeDie:   800,
		TimeEat:   200,
		TimeSleep: 200,
		TimeThink: 100,
		NumEat:    0, // unlimited
	}
	if err := p.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestParamsValidate_Invalid(t *testing.T) {
	tests := []struct {
		name   string
		p      Params
		substr string
	}{
		{
			name:   "NumPhilos too small",
			p:      Params{NumPhilos: 1, TimeDie: 1, TimeEat: 1, TimeSleep: 1, TimeThink: 0, NumEat: 0},
			substr: "philosophers",
		},
		{
			name:   "TimeDie invalid",
			p:      Params{NumPhilos: 2, TimeDie: 0, TimeEat: 1, TimeSleep: 1, TimeThink: 0, NumEat: 0},
			substr: "TimeDie",
		},
		{
			name:   "TimeEat invalid",
			p:      Params{NumPhilos: 2, TimeDie: 1, TimeEat: 0, TimeSleep: 1, TimeThink: 0, NumEat: 0},
			substr: "TimeEat",
		},
		{
			name:   "TimeSleep invalid",
			p:      Params{NumPhilos: 2, TimeDie: 1, TimeEat: 1, TimeSleep: 0, TimeThink: 0, NumEat: 0},
			substr: "TimeSleep",
		},
		{
			name:   "TimeThink invalid",
			p:      Params{NumPhilos: 2, TimeDie: 1, TimeEat: 1, TimeSleep: 1, TimeThink: -1, NumEat: 0},
			substr: "TimeThink",
		},
		{
			name:   "NumEat negative",
			p:      Params{NumPhilos: 2, TimeDie: 1, TimeEat: 1, TimeSleep: 1, TimeThink: 0, NumEat: -1},
			substr: "NumEat",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.p.Validate()
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			if !errors.Is(err, ecode.ErrInvalidParameter) {
				t.Errorf("expected wrapped ecode.ErrInvalidParameter, got %v", err)
			}
			if !strings.Contains(err.Error(), tc.substr) {
				t.Errorf("expected error message to contain %q, got %q", tc.substr, err.Error())
			}
		})
	}
}

func TestParamsString(t *testing.T) {
	p := &Params{
		NumPhilos: 5,
		TimeDie:   800,
		TimeEat:   200,
		TimeSleep: 200,
		TimeThink: 150,
		NumEat:    3,
	}
	got := p.String()
	want := "Philosophers: 5, Die: 800ms, Eat: 200ms, Sleep: 200ms, Think: 150ms, Eat Count: 3"
	if got != want {
		t.Errorf("unexpected String() output.\n got: %q\nwant: %q", got, want)
	}
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
