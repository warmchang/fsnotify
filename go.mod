module github.com/fsnotify/fsnotify

go 1.16

require (
	github.com/power-devops/ahafs v0.1.4 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956
)

retract (
	v1.5.3 // Published an incorrect branch accidentally https://github.com/fsnotify/fsnotify/issues/445
	v1.5.0 // Contains symlink regression https://github.com/fsnotify/fsnotify/pull/394
)
