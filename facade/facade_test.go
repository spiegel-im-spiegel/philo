package facade

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
)

// helper to build *rwi.RWI with separate out/err buffers
func newTestUI() (*rwi.RWI, *bytes.Buffer, *bytes.Buffer) {
	var outBuf, errBuf bytes.Buffer
	ui := rwi.New(
		rwi.WithReader(strings.NewReader("")),
		rwi.WithWriter(&outBuf),
		rwi.WithErrorWriter(&errBuf),
	)
	return ui, &outBuf, &errBuf
}

func TestExecute(t *testing.T) {
	testCases := []struct {
		args   []string
		out    string
		outErr string
	}{
		{args: []string{}, out: "Philosophers: 5, Die: 200ms, Eat: 20ms, Sleep: 80ms, Think: 80ms, Eat Count: 3\n", outErr: ""},
		{args: []string{
			"--philosophers", "7",
			"-d", "300",
			"-e", "25",
			"-s", "90",
			"-t", "100",
			"-m", "5",
		}, out: "Philosophers: 7, Die: 300ms, Eat: 25ms, Sleep: 90ms, Think: 100ms, Eat Count: 5\n", outErr: ""},
	}

	for _, tc := range testCases {
		ui, out, errOut := newTestUI()
		exit := Execute(ui, tc.args)
		if exit != exitcode.Normal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Normal)
		}
		if out.String() != tc.out {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), tc.out)
		}
		if errOut.String() != tc.outErr {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), tc.outErr)
		}
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
