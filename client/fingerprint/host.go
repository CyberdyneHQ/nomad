package fingerprint

import (
	"runtime"

	log "github.com/hashicorp/go-hclog"
	"github.com/shirou/gopsutil/v3/host"
)

// HostFingerprint is used to fingerprint the host
type HostFingerprint struct {
	StaticFingerprinter
	logger log.Logger
}

// NewHostFingerprint is used to create a Host fingerprint
func NewHostFingerprint(logger log.Logger) Fingerprint {
	f := &HostFingerprint{logger: logger.Named("host")}
	return f
}

func (f *HostFingerprint) Fingerprint(req *FingerprintRequest, resp *FingerprintResponse) error {
	hostInfo, err := host.Info()
	if err != nil {
		f.logger.Warn("error retrieving host information", "error", err)
		return err
	}

	resp.AddAttribute("os.name", hostInfo.Platform)
	resp.AddAttribute("os.version", hostInfo.PlatformVersion)

	resp.AddAttribute("kernel.name", runtime.GOOS)
	resp.AddAttribute("kernel.arch", hostInfo.KernelArch)
	resp.AddAttribute("kernel.version", hostInfo.KernelVersion)

	resp.AddAttribute("unique.hostname", hostInfo.Hostname)
	resp.Detected = true

	return nil
}
