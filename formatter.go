package throughput

import (
	"math"
	"strconv"
	"time"
)

// https://en.wikipedia.org/wiki/Measuring_network_throughput
// https://en.wikipedia.org/wiki/Data_rate_units

const (
	kilo = float64(1000)
	mega = float64(1000) * kilo
	giga = float64(1000) * mega
	tera = float64(1000) * giga

	kibi = float64(1024)
	mebi = float64(1024) * kibi
	gibi = float64(1024) * mebi
	tebi = float64(1024) * gibi

	none = -1
)

type rate struct {
	multiple float64
	next     float64
	symbol   string
}

var bitPrefixesDecimal = []rate{
	rate{multiple: 1, next: kilo, symbol: "bit/s"},
	rate{multiple: kilo, next: mega, symbol: "kbit/s"},
	rate{multiple: mega, next: giga, symbol: "Mbit/s"},
	rate{multiple: giga, next: tera, symbol: "Gbit/s"},
	rate{multiple: tera, next: none, symbol: "Tbit/s"},
}

var bitPrefixesBinary = []rate{
	rate{multiple: 1, next: kibi, symbol: "bit/s"},
	rate{multiple: kibi, next: mebi, symbol: "Kibit/s"},
	rate{multiple: mebi, next: gibi, symbol: "Mibit/s"},
	rate{multiple: gibi, next: tebi, symbol: "Gibit/s"},
	rate{multiple: tebi, next: none, symbol: "Tibit/s"},
}

var bytePrefixesDecimal = []rate{
	rate{multiple: 1, next: kilo, symbol: "B/s"},
	rate{multiple: kilo, next: mega, symbol: "kB/s"},
	rate{multiple: mega, next: giga, symbol: "MB/s"},
	rate{multiple: giga, next: tera, symbol: "GB/s"},
	rate{multiple: tera, next: none, symbol: "TB/s"},
}

var bytePrefixesBinary = []rate{
	rate{multiple: 1, next: kibi, symbol: "B/s"},
	rate{multiple: kibi, next: mebi, symbol: "KiB/s"},
	rate{multiple: mebi, next: gibi, symbol: "MiB/s"},
	rate{multiple: gibi, next: tebi, symbol: "GiB/s"},
	rate{multiple: tebi, next: none, symbol: "TiB/s"},
}

// FormatBitsDecimal outputs n measured in bit per second (bit/s), kilobits per
// second (kbit/s), megabits per second (Mbit/s) or gigabits per second (Gbit/s)
// whichever is best represented. 1 kbit is defined as 1000 bits.
func FormatBitsDecimal(n uint64, period time.Duration, decimals int) string {
	return format(n, period, decimals, bitPrefixesDecimal)
}

// FormatBitsBinary outputs n measured in bit per second (bit/s), kibibit per
// second (Kibit/s), mebibit per second (Mibit/s) or gibibit per second (Gibit)
// whichever is best represented. 1 kibibit is defined as 1024 bits.
func FormatBitsBinary(n uint64, period time.Duration, decimals int) string {
	return format(n, period, decimals, bitPrefixesBinary)
}

// FormatBytesDecimal outputs n measured in byte per second (B/s), kilobyte per
// second (kB/s), megabyte per second (MB/s) or gigabyte per second (GB/s)
// whichever is best represented. 1 kilobyte is defined as 1000 bytes.
func FormatBytesDecimal(n uint64, period time.Duration, decimals int) string {
	return format(n, period, decimals, bytePrefixesDecimal)
}

// FormatBytesBinary outputs n measured in byte per second (B/s), kibibyte per
// second (KiB/s), mebibyte per second (MiB/s) or gibibyte per second (GiB)
// whichever is best represented. 1 kibibyte is defined as 1024 bytes.
func FormatBytesBinary(n uint64, period time.Duration, decimals int) string {
	return format(n, period, decimals, bytePrefixesBinary)
}

func format(n uint64, period time.Duration, decimals int, rates []rate) string {
	var r rate
	bps := (float64(n) / float64(period)) * float64(time.Second)
	for _, cr := range rates {
		if bps < cr.next || cr.next == none {
			r = cr
			break
		}
	}
	measure := (bps / r.multiple)
	measure *= math.Pow10(decimals)
	measure = math.Trunc(measure)
	measure /= math.Pow10(decimals)
	return strconv.FormatFloat(measure, 'f', decimals, 64) + " " + r.symbol
}
