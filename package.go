package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// see https://devblogs.microsoft.com/insidemsix/is-this-a-packaged-process/
// see https://learn.microsoft.com/en-us/windows/win32/api/appmodel/nf-appmodel-getcurrentpackagefullname
func GetCurrentPackageFullName() (string, error) {
	procGetCurrentPackageFullName := windows.NewLazySystemDLL("kernel32.dll").NewProc("GetCurrentPackageFullName")
	if procGetCurrentPackageFullName == nil {
		return "", fmt.Errorf("failed to load GetCurrentPackageFullName")
	}

	var length uint32
	ret, _, _ := procGetCurrentPackageFullName.Call(
		uintptr(unsafe.Pointer(&length)),
		uintptr(0),
	)

	errCode := windows.Errno(ret)
	if errCode != windows.ERROR_INSUFFICIENT_BUFFER {
		if errCode == windows.APPMODEL_ERROR_NO_PACKAGE {
			return "", fmt.Errorf("process has no package identity")
		}
		return "", fmt.Errorf("GetCurrentPackageFullName failed: %v", errCode)
	}

	buf := make([]uint16, length)

	ret, _, _ = procGetCurrentPackageFullName.Call(
		uintptr(unsafe.Pointer(&length)),
		uintptr(unsafe.Pointer(&buf[0])),
	)

	errCode = windows.Errno(ret)
	if errCode != windows.ERROR_SUCCESS {
		return "", fmt.Errorf("GetCurrentPackageFullName failed on second call: %v", errCode)
	}

	return syscall.UTF16ToString(buf), nil
}

// see https://learn.microsoft.com/en-us/windows/win32/api/appmodel/nf-appmodel-getcurrentpackagepath
func GetCurrentPackagePath() (string, error) {
	procGetCurrentPackagePath := windows.NewLazySystemDLL("kernel32.dll").NewProc("GetCurrentPackagePath")
	if procGetCurrentPackagePath == nil {
		return "", fmt.Errorf("failed to load GetCurrentPackagePath")
	}

	var length uint32
	ret, _, _ := procGetCurrentPackagePath.Call(
		uintptr(unsafe.Pointer(&length)),
		uintptr(0),
	)

	errCode := windows.Errno(ret)
	if errCode != windows.ERROR_INSUFFICIENT_BUFFER {
		if errCode == windows.APPMODEL_ERROR_NO_PACKAGE {
			return "", fmt.Errorf("process has no package identity")
		}
		return "", fmt.Errorf("GetCurrentPackagePath failed: %v", errCode)
	}

	buf := make([]uint16, length)

	ret, _, _ = procGetCurrentPackagePath.Call(
		uintptr(unsafe.Pointer(&length)),
		uintptr(unsafe.Pointer(&buf[0])),
	)

	errCode = windows.Errno(ret)
	if errCode != windows.ERROR_SUCCESS {
		return "", fmt.Errorf("GetCurrentPackagePath failed on second call: %v", errCode)
	}

	return syscall.UTF16ToString(buf), nil
}
