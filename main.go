//go:generate bash versioninfo.sh

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
)

const IDI_APPLICATION = 32512 // application icon index.

func main() {
	runtime.LockOSThread()

	info, ok := debug.ReadBuildInfo()
	if !ok {
		os.Exit(1)
	}
	version := info.Main.Version

	packageInfo := getPackageInfo()

	wnd := ui.NewMain(
		ui.OptsMain().
			ClassIconId(IDI_APPLICATION).
			Title(fmt.Sprintf("Example Go Windows App %s", version)).
			Style(co.WS_TILEDWINDOW).
			Size(ui.Dpi(680, 280)),
	)

	ui.NewStatic(
		wnd,
		ui.OptsStatic().
			Text(packageInfo).
			Position(ui.Dpi(10, 10)),
	)

	os.Exit(wnd.RunAsMain())
}

func getPackageInfo() string {
	var sb strings.Builder

	id, err := win.GetCurrentPackageId()
	if wErr, ok := err.(co.ERROR); ok && wErr == co.ERROR_APPMODEL_ERROR_NO_PACKAGE {
		return "This process does not have a package identity."
	}
	if err == nil {
		var arch string
		switch id.ProcessorArchitecture {
		case co.PROCESSOR_ARCH_ENUM_X86:
			arch = "X86"
		case co.PROCESSOR_ARCH_ENUM_X64:
			arch = "X64"
		case co.PROCESSOR_ARCH_ENUM_X86ONARM64:
			arch = "X86ONARM64"
		case co.PROCESSOR_ARCH_ENUM_ARM:
			arch = "ARM"
		case co.PROCESSOR_ARCH_ENUM_ARM64:
			arch = "ARM64"
		default:
			arch = "unknown"
		}
		fmt.Fprintf(&sb, "Package Processor Architecture: %s\n", arch)
		fmt.Fprintf(&sb, "Package Name: %s\n", id.Name())
		fmt.Fprintf(&sb, "Package Version: %d.%d.%d.%d\n", id.Version.Major, id.Version.Minor, id.Version.Build, id.Version.Revision)
		fmt.Fprintf(&sb, "Package Publisher: %s\n", id.Publisher())
		fmt.Fprintf(&sb, "Package PublisherId: %s\n", id.PublisherId())
		fmt.Fprintf(&sb, "Package ResourceId: %s\n", id.ResourceId())
	} else {
		fmt.Fprintf(&sb, "Package: win.GetCurrentPackageId error: %s\n", err)
	}

	familyName, err := win.GetCurrentPackageFamilyName()
	if err == nil {
		fmt.Fprintf(&sb, "Package Family Name: %s\n", familyName)
	} else {
		fmt.Fprintf(&sb, "Package Family Name: win.GetCurrentPackageFamilyName error: %s\n", err)
	}

	fullName, err := win.GetCurrentPackageFullName()
	if err == nil {
		fmt.Fprintf(&sb, "Package Full Name: %s\n", fullName)
	} else {
		fmt.Fprintf(&sb, "Package Full Name: win.GetCurrentPackageFullName error: %s\n", err)
	}

	path, err := win.GetCurrentPackagePath()
	if err == nil {
		fmt.Fprintf(&sb, "Package Path: %s\n", path)
	} else {
		fmt.Fprintf(&sb, "Package Path: win.GetCurrentPackagePath error: %s\n", err)
	}

	installPath, err := win.GetCurrentPackagePath2(co.PACKAGE_PATH_TYPE_Install)
	if err == nil {
		fmt.Fprintf(&sb, "Package Install Path: %s\n", installPath)
	} else {
		fmt.Fprintf(&sb, "Package Install Path: win.GetCurrentPackagePath2 error: %s\n", err)
	}

	mutablePath, err := win.GetCurrentPackagePath2(co.PACKAGE_PATH_TYPE_Mutable)
	if err == nil {
		fmt.Fprintf(&sb, "Package Mutable Path: %s\n", mutablePath)
	} else {
		fmt.Fprintf(&sb, "Package Mutable Path: win.GetCurrentPackagePath2 error: %s\n", err)
	}

	effectivePath, err := win.GetCurrentPackagePath2(co.PACKAGE_PATH_TYPE_Effective)
	if err == nil {
		fmt.Fprintf(&sb, "Package Effective Path: %s\n", effectivePath)
	} else {
		fmt.Fprintf(&sb, "Package Effective Path: win.GetCurrentPackagePath2 error: %s\n", err)
	}

	machineExternalPath, err := win.GetCurrentPackagePath2(co.PACKAGE_PATH_TYPE_MachineExternal)
	if err == nil {
		fmt.Fprintf(&sb, "Package Machine External Path: %s\n", machineExternalPath)
	} else {
		fmt.Fprintf(&sb, "Package Machine External Path: win.GetCurrentPackagePath2 error: %s\n", err)
	}

	userExternalPath, err := win.GetCurrentPackagePath2(co.PACKAGE_PATH_TYPE_UserExternal)
	if err == nil {
		fmt.Fprintf(&sb, "Package User External Path: %s\n", userExternalPath)
	} else {
		fmt.Fprintf(&sb, "Package User External Path: win.GetCurrentPackagePath2 error: %s\n", err)
	}

	effectiveExternalPath, err := win.GetCurrentPackagePath2(co.PACKAGE_PATH_TYPE_EffectiveExternal)
	if err == nil {
		fmt.Fprintf(&sb, "Package Effective External Path: %s\n", effectiveExternalPath)
	} else {
		fmt.Fprintf(&sb, "Package Effective External Path: win.GetCurrentPackagePath2 error: %s\n", err)
	}

	return sb.String()
}
