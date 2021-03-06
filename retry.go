package common

//	MIT License
//
//	Copyright (c) Microsoft Corporation. All rights reserved.
//
//	Permission is hereby granted, free of charge, to any person obtaining a copy
//	of this software and associated documentation files (the "Software"), to deal
//	in the Software without restriction, including without limitation the rights
//	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//	copies of the Software, and to permit persons to whom the Software is
//	furnished to do so, subject to the following conditions:
//
//	The above copyright notice and this permission notice shall be included in all
//	copies or substantial portions of the Software.
//
//	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//	SOFTWARE

import (
	"time"
)

// Retryable represents an error which should be able to be retried
type Retryable string

// Error implementation for Retryable
func (r Retryable) Error() string {
	return string(r)
}

// Retry will attempt to retry an action a number of times if the action returns a retryable error
func Retry(times int, delay time.Duration, action func() (interface{}, error)) (interface{}, error) {
	var lastErr error
	for i := 0; i < times; i++ {
		item, err := action()
		if err != nil {
			if err, ok := err.(Retryable); ok {
				lastErr = err
				time.Sleep(delay)
				continue
			} else {
				return nil, err
			}
		}
		return item, nil
	}
	return nil, lastErr
}
