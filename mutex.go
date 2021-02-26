// Copyright (c) 2020 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mu provides mutex helpers.
package mu

import (
	"sync"
)

// Mutex extended with helper methods.
type Mutex struct {
	sync.Mutex
}

// Guard invokes f while keeping the mutex locked.
func (m *Mutex) Guard(f func()) {
	m.Lock()
	defer m.Unlock()
	f()
}

// GuardBool invokes f while keeping the mutex locked.  The return value is
// passed through.
func (m *Mutex) GuardBool(f func() bool) bool {
	m.Lock()
	defer m.Unlock()
	return f()
}

// GuardError invokes f while keeping the mutex locked.  The return value is
// passed through.
func (m *Mutex) GuardError(f func() error) error {
	m.Lock()
	defer m.Unlock()
	return f()
}

// GuardBoolError invokes f while keeping the mutex locked.  The return values
// are passed through.
func (m *Mutex) GuardBoolError(f func() (bool, error)) (bool, error) {
	m.Lock()
	defer m.Unlock()
	return f()
}

// UnlockGuard unlocks the mutex, invokes f, and locks the mutex again.
func (m *Mutex) UnlockGuard(f func()) {
	m.Unlock()
	defer m.Lock()
	f()
}

// UnlockGuardBool unlocks the mutex, invokes f, and locks the mutex again.
// The return value is passed through.
func (m *Mutex) UnlockGuardBool(f func() bool) bool {
	m.Unlock()
	defer m.Lock()
	return f()
}

// UnlockGuardError unlocks the mutex, invokes f, and locks the mutex again.
// The return value is passed through.
func (m *Mutex) UnlockGuardError(f func() error) error {
	m.Unlock()
	defer m.Lock()
	return f()
}

// UnlockGuardBoolError unlocks the mutex, invokes f, and locks the mutex
// again.  The return values are passed through.
func (m *Mutex) UnlockGuardBoolError(f func() (bool, error)) (bool, error) {
	m.Unlock()
	defer m.Lock()
	return f()
}
