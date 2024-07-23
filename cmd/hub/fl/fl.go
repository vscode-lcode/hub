package fl

import (
	"errors"
	"runtime"
	"time"
	"unsafe"

	"golang.org/x/net/webdav"
)

type LockSystem struct {
	webdav.LockSystem
}

func New(fl webdav.LockSystem) *LockSystem {
	return &LockSystem{
		LockSystem: fl,
	}
}

func (fl *LockSystem) Confirm(params ConfirmParams, reply *uintptr) error {
	release, err := fl.LockSystem.Confirm(params.Now, params.Name0, params.Name0, params.Conditions...)
	if err != nil {
		return err
	}

	var pinner runtime.Pinner
	callback := func() {
		defer pinner.Unpin()
		release()
	}
	pinner.Pin(&callback)

	ptr := uintptr(unsafe.Pointer(&callback))
	*reply = ptr

	return nil
}

var ErrNotCallback = errors.New("the uintptr is not callback")

func (fl *LockSystem) ConfirmCallback(ptr uintptr, reply *bool) error {
	// 这里可能 panic, 但我不知道如何防止
	callback := *(*func())(unsafe.Pointer(ptr))
	callback()
	*reply = true
	return nil
}

type ConfirmParams struct {
	Now        time.Time
	Name0      string
	Name1      string
	Conditions []webdav.Condition
}

func (fl *LockSystem) Create(params CreateParams, reply *string) error {
	token, err := fl.LockSystem.Create(params.Now, params.Details)
	if err != nil {
		return err
	}
	*reply = token
	return nil
}

type CreateParams struct {
	Now     time.Time
	Details webdav.LockDetails
}

func (fl *LockSystem) Refresh(params RefreshParams, reply *webdav.LockDetails) error {
	ld, err := fl.LockSystem.Refresh(params.Now, params.Token, params.Duration)
	if err != nil {
		return err
	}
	*reply = ld
	return nil
}

type RefreshParams struct {
	Now      time.Time
	Token    string
	Duration time.Duration
}

func (fl *LockSystem) Unlock(params UnlockParams, reply *bool) error {
	err := fl.LockSystem.Unlock(params.Now, params.Token)
	if err != nil {
		return err
	}
	*reply = true
	return nil
}

type UnlockParams struct {
	Now   time.Time
	Token string
}
