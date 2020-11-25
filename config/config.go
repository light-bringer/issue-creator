package config

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	_Root string = filepath.Join(filepath.Dir(b), "../..")
)

const (
	PORT       = ":5000"
	APP_NAME   = "issue-creator"
	APP_SECRET = "6YJSuc50uJ18zj45"
	API_EXPIRY = "120"

	Log_FILE_PATH = "//Users/debapriya.das/Desktop/issue-creator/logs"
	LOG_FILE_NAME = "application.log"
)
