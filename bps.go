package bytefmt

import (
	"fmt"
	"strings"
)

const (
	BITSPS     = 1.0
	KILOBITSPS = 1000 * BITSPS
	MEGABITSPS = 1000 * KILOBITSPS
	GIGABITSPS = 1000 * MEGABITSPS
	TERABITSPS = 1000 * GIGABITSPS
)

// BPSSize returns a human-readable bits/sec string of the form 10Mbps,
// 12.5Kbps, and so forth.  The following units are available:
//	Tbps: Terabit/sec
//	Gbps: Gigabit/sec
//	Mbps: Megabit/sec
//	Kbps: Kilobit/sec
//	bps: Bits/sec
// The unit that results in the smallest number greater than or equal to 1 is
// always chosen.
func BPSSize(bps float64) string {
	unit := ""

	switch {
	case bps >= TERABITSPS:
		unit = "Tbps"
		bps = bps / TERABITSPS
	case bps >= GIGABITSPS:
		unit = "Gbps"
		bps = bps / GIGABITSPS
	case bps >= MEGABITSPS:
		unit = "Mbps"
		bps = bps / MEGABITSPS
	case bps >= KILOBITSPS:
		unit = "Kbps"
		bps = bps / KILOBITSPS
	case bps >= BITSPS:
		unit = "bps"
	case bps == 0:
		return "0"
	}

	stringValue := fmt.Sprintf("%.1f", bps)
	stringValue = strings.TrimSuffix(stringValue, ".0")

	return fmt.Sprintf("%s%s", stringValue, unit)
}
