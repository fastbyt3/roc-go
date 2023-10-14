// DO NOT EDIT! Code generated by generate_enums script from roc-toolkit
// roc-toolkit git tag: v0.2.5-11-g14d642e9, commit: 14d642e9

package roc

// Resampler backend.
//
// Affects speed and quality. Some backends may be disabled at build time.
//
//go:generate stringer -type ResamplerBackend -trimprefix ResamplerBackend -output resampler_backend_string.go
type ResamplerBackend int

const (
	// Default backend.
	//
	// Depends on what was enabled at build time.
	ResamplerBackendDefault ResamplerBackend = 0

	// Slow built-in resampler.
	//
	// Always available.
	ResamplerBackendBuiltin ResamplerBackend = 1

	// Fast good-quality resampler from SpeexDSP.
	//
	// May be disabled at build time.
	ResamplerBackendSpeex ResamplerBackend = 2
)