package fasthttp

import (
	"bufio"
	"bytes"
	"net"
	"testing"
)

func BenchmarkParseIPv4(b *testing.B) {
	ipStr := []byte("123.145.167.189")
	b.RunParallel(func(pb *testing.PB) {
		var ip net.IP
		var err error
		for pb.Next() {
			ip, err = ParseIPv4(ip, ipStr)
			if err != nil {
				b.Fatalf("unexpected error: %s", err)
			}
		}
	})
}

func BenchmarkAppendIPv4(b *testing.B) {
	ip := net.ParseIP("123.145.167.189")
	b.RunParallel(func(pb *testing.PB) {
		var buf []byte
		for pb.Next() {
			buf = AppendIPv4(buf[:0], ip)
		}
	})
}

func BenchmarkInt2HexByte(b *testing.B) {
	buf := []int{1, 0xf, 2, 0xd, 3, 0xe, 4, 0xa, 5, 0xb, 6, 0xc, 7, 0xf, 0, 0xf, 6, 0xd, 9, 8, 4, 0x5}
	b.RunParallel(func(pb *testing.PB) {
		var n int
		for pb.Next() {
			for _, n = range buf {
				int2hexbyte(n)
			}
		}
	})
}

func BenchmarkHexByte2Int(b *testing.B) {
	buf := []byte("0A1B2c3d4E5F6C7a8D9ab7cd03ef")
	b.RunParallel(func(pb *testing.PB) {
		var c byte
		for pb.Next() {
			for _, c = range buf {
				hexbyte2int(c)
			}
		}
	})
}

func BenchmarkWriteHexInt(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var w bytes.Buffer
		bw := bufio.NewWriter(&w)
		i := 0
		for pb.Next() {
			writeHexInt(bw, i)
			i++
			if i > 0x7fffffff {
				i = 0
			}
			w.Reset()
			bw.Reset(&w)
		}
	})
}

func BenchmarkParseUint(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		buf := []byte("1234567")
		for pb.Next() {
			n, err := ParseUint(buf)
			if err != nil {
				b.Fatalf("unexpected error: %s", err)
			}
			if n != 1234567 {
				b.Fatalf("unexpected result: %d. Expecting %s", n, buf)
			}
		}
	})
}

func BenchmarkAppendUint(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var buf []byte
		i := 0
		for pb.Next() {
			buf = AppendUint(buf[:0], i)
			i++
			if i > 0x7fffffff {
				i = 0
			}
		}
	})
}

func BenchmarkLowercaseBytesNoop(b *testing.B) {
	src := []byte("foobarbaz_lowercased_all")
	b.RunParallel(func(pb *testing.PB) {
		s := make([]byte, len(src))
		for pb.Next() {
			copy(s, src)
			lowercaseBytes(s)
		}
	})
}

func BenchmarkLowercaseBytesAll(b *testing.B) {
	src := []byte("FOOBARBAZ_UPPERCASED_ALL")
	b.RunParallel(func(pb *testing.PB) {
		s := make([]byte, len(src))
		for pb.Next() {
			copy(s, src)
			lowercaseBytes(s)
		}
	})
}

func BenchmarkLowercaseBytesMixed(b *testing.B) {
	src := []byte("Foobarbaz_Uppercased_Mix")
	b.RunParallel(func(pb *testing.PB) {
		s := make([]byte, len(src))
		for pb.Next() {
			copy(s, src)
			lowercaseBytes(s)
		}
	})
}

func BenchmarkEqualBytesStrEq(b *testing.B) {
	s := "foobarbaraz"
	bs := []byte(s)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if !EqualBytesStr(bs, s) {
				b.Fatalf("unexpected result: %q != %q", bs, s)
			}
		}
	})
}

func BenchmarkEqualBytesStrNe(b *testing.B) {
	s := "foobarbaraz"
	bs := []byte(s)
	bs[len(s)-1] = 'a'
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if EqualBytesStr(bs, s) {
				b.Fatalf("unexpected result: %q = %q", bs, s)
			}
		}
	})
}

func BenchmarkAppendBytesStr(b *testing.B) {
	s := "foobarbazbaraz"
	b.RunParallel(func(pb *testing.PB) {
		var dst []byte
		for pb.Next() {
			dst = AppendBytesStr(dst[:0], s)
		}
	})
}
