package throughput

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		f        func(uint64, time.Duration, int) string
		n        uint64
		period   time.Duration
		decimals int
		expected string
	}{
		{FormatBitsDecimal, 0, time.Second, 0, "0 bit/s"},
		{FormatBitsDecimal, 0, time.Second, 1, "0.0 bit/s"},
		{FormatBitsDecimal, 1, time.Second * 2, 1, "0.5 bit/s"},
		{FormatBitsDecimal, 125, time.Second * 2, 1, "62.5 bit/s"},
		{FormatBitsDecimal, 999, time.Second, 0, "999 bit/s"},
		{FormatBitsDecimal, 1000, time.Second, 0, "1 kbit/s"},
		{FormatBitsDecimal, 3, time.Millisecond, 2, "3.00 kbit/s"},
		{FormatBitsDecimal, 2 * ((1000 * 1000) - 1), 2 * time.Second, 2, "999.99 kbit/s"},
		{FormatBitsDecimal, 1000 * 1000, time.Second, 0, "1 Mbit/s"},
		{FormatBitsDecimal, 1000 * 1000 * 124, time.Second * 42, 1, "2.9 Mbit/s"},
		{FormatBitsDecimal, 1000 * 1000 * 1000, time.Second * 3, 4, "333.3333 Mbit/s"},
		{FormatBitsDecimal, (1000 * 1000 * 1000) - 1, time.Second, 0, "999 Mbit/s"},
		{FormatBitsDecimal, 1000 * 1000 * 1000, time.Second, 0, "1 Gbit/s"},
		{FormatBitsDecimal, (1000 * 1000 * 1000 * 1000) - 1, time.Second, 1, "999.9 Gbit/s"},
		{FormatBitsDecimal, 1000 * 1000 * 1000 * 1000, time.Second, 0, "1 Tbit/s"},
		{FormatBitsDecimal, 1<<64 - 1, time.Second, 3, "18446744.073 Tbit/s"},
		{FormatBitsBinary, 256, time.Second / 2, 0, "512 bit/s"},
		{FormatBitsBinary, 251 * 1024, time.Second * 10, 1, "25.1 Kibit/s"},
		{FormatBitsBinary, 256 * 1024 * 1024, time.Second / 2, 0, "512 Mibit/s"},
		{FormatBitsBinary, 256 * 1024 * 1024 * 1024, time.Second, 2, "256.00 Gibit/s"},
		{FormatBitsBinary, 256 * 1024 * 1024 * 1024 * 1024, time.Second, 2, "256.00 Tibit/s"},
		{FormatBytesDecimal, 125, time.Second * 2, 1, "62.5 B/s"},
		{FormatBytesDecimal, 2 * ((1000 * 1000) - 1), 2 * time.Second, 2, "999.99 kB/s"},
		{FormatBytesDecimal, 1000 * 1000 * 124, time.Second * 42, 1, "2.9 MB/s"},
		{FormatBytesDecimal, (1000 * 1000 * 1000 * 1000) - 1, time.Second, 1, "999.9 GB/s"},
		{FormatBytesDecimal, 1000 * 1000 * 1000 * 1000, time.Second, 0, "1 TB/s"},
		{FormatBytesBinary, 256, time.Second / 2, 0, "512 B/s"},
		{FormatBytesBinary, 251 * 1024, time.Second * 10, 1, "25.1 KiB/s"},
		{FormatBytesBinary, 256 * 1024 * 1024, time.Second / 2, 0, "512 MiB/s"},
		{FormatBytesBinary, 256 * 1024 * 1024 * 1024, time.Second, 2, "256.00 GiB/s"},
		{FormatBytesBinary, 256 * 1024 * 1024 * 1024 * 1024, time.Second, 2, "256.00 TiB/s"},
	}

	for i, testcase := range tests {
		got := testcase.f(testcase.n, testcase.period, testcase.decimals)
		if got != testcase.expected {
			t.Errorf("[%d] link throughput: expected %s got %s", i, testcase.expected, got)
		}
	}
}

func BenchmarkFormatBitsDecimal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FormatBitsDecimal(1000*1000*1000, time.Second*3, 4)
	}
}
