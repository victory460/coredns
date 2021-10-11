package metrics

import (
	"runtime"

	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/miekg/dns"
)

// Recorder is a dnstest.Recorder specific to the metrics plugin.
type Recorder struct {
	*dnstest.Recorder
	// CallerN holds the string return value of the call to runtime.Caller(N)
	Caller1 string
	Caller2 string
	Caller3 string
}

// NewRecorder makes and returns a new Recorder.
func NewRecorder(w dns.ResponseWriter) *Recorder { return &Recorder{Recorder: dnstest.NewRecorder(w)} }

// WriteMsg records the status code and calls the
// underlying ResponseWriter's WriteMsg method.
func (r *Recorder) WriteMsg(res *dns.Msg) error {
	_, r.Caller1, _, _ = runtime.Caller(1)
	_, r.Caller2, _, _ = runtime.Caller(2)
	_, r.Caller3, _, _ = runtime.Caller(3)
	r.Len += res.Len()
	r.Msg = res
	return r.ResponseWriter.WriteMsg(res)
}
