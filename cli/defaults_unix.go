// +build aix dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris

package cli

import (
	"os"
	"path/filepath"
)

func ZvmDataDirPath(home string) string {
	zvm_path := os.Getenv("ZVM_PATH")
	if zvm_path == "" {
		zvm_path = os.Getenv("XDG_DATA_HOME")
		if zvm_path == "" {
			zvm_path = filepath.Join(home, ".local", "share")
		}
		zvm_path = filepath.Join(zvm_path, "zvm")
	}
	return zvm_path
}
func ZvmConfigDirPath(home string) string {
	zvm_path := os.Getenv("ZVM_PATH")
	if zvm_path == "" {
		zvm_path = os.Getenv("XDG_CONFIG_HOME")
		if zvm_path == "" {
			zvm_path = filepath.Join(home, ".config")
		}
		zvm_path = filepath.Join(zvm_path, "zvm")
	}
	return zvm_path
}
func ZvmStateDirPath(home string) string {
	zvm_path := os.Getenv("ZVM_PATH")
	if zvm_path == "" {
		zvm_path = os.Getenv("XDG_STATE_HOME")
		if zvm_path == "" {
			zvm_path = filepath.Join(home, ".local", "state")
		}
		zvm_path = filepath.Join(zvm_path, "zvm")
	}
	return zvm_path
}
func ZvmBinDirPath(home string) string {
	zvm_path := os.Getenv("ZVM_PATH")
	if zvm_path == "" {
		zvm_path = os.Getenv("XDG_BIN_HOME")
		if zvm_path == "" {
			zvm_path = filepath.Join(home, ".local", "bin")
		}
	}else {
		zvm_path = filepath.Join(zvm_path, "bin")
	}
	return zvm_path
}
func ZvmCacheDirPath(home string) string {
	zvm_path := os.Getenv("ZVM_PATH")
	if zvm_path == "" {
		zvm_path = os.Getenv("XDG_CACHE_HOME")
		if zvm_path == "" {
			zvm_path = filepath.Join(home, ".cache")
		}
		zvm_path = filepath.Join(zvm_path, "zvm")
	}
	return zvm_path
}
