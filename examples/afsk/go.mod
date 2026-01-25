module tinygo.org/x/wireless/examples/afsk

go 1.25.3

replace tinygo.org/x/wireless => ../..

replace tinygo.org/x/wireless/examples/audio => ../audio

require (
	tinygo.org/x/drivers v0.34.1-0.20260107195827-c21cd39813be
	tinygo.org/x/wireless v0.0.0-20260121153201-f0d8647de68c
	tinygo.org/x/wireless/examples/audio v0.0.0-00010101000000-000000000000
)

require (
	github.com/ebitengine/oto/v3 v3.1.0 // indirect
	github.com/ebitengine/purego v0.7.1 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/gopxl/beep v1.4.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.40.0 // indirect
)
