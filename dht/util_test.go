package dht

import (
	"fmt"
	"testing"
)

func TestInt2Bytes(t *testing.T) {
	cases := []struct {
		in  uint64
		out []byte
	}{
		{0, []byte{0}},
		{1, []byte{1}},
		{256, []byte{1, 0}},
		{22129, []byte{86, 113}},
	}

	for _, c := range cases {
		r := int2bytes(c.in)
		if len(r) != len(c.out) {
			t.Fail()
		}

		for i, v := range r {
			if v != c.out[i] {
				t.Fail()
			}
		}
	}
}

func TestBytes2Int(t *testing.T) {
	cases := []struct {
		in  []byte
		out uint64
	}{
		{[]byte{0}, 0},
		{[]byte{1}, 1},
		{[]byte{1, 0}, 256},
		{[]byte{86, 113}, 22129},
	}

	for _, c := range cases {
		if bytes2int(c.in) != c.out {
			t.Fail()
		}
	}
}

func TestDecodeCompactIPPortInfo(t *testing.T) {
	cases := []struct {
		in  string
		out struct {
			ip   string
			port int
		}
	}{
		{"123456", struct {
			ip   string
			port int
		}{"49.50.51.52", 13622}},
		{"abcdef", struct {
			ip   string
			port int
		}{"97.98.99.100", 25958}},
	}

	for _, item := range cases {
		ip, port, err := decodeCompactIPPortInfo(item.in)
		if err != nil || ip.String() != item.out.ip || port != item.out.port {
			t.Fail()
		}
	}
}

func TestEncodeCompactIPPortInfo(t *testing.T) {
	cases := []struct {
		in struct {
			ip   []byte
			port int
		}
		out string
	}{
		{struct {
			ip   []byte
			port int
		}{[]byte{49, 50, 51, 52}, 13622}, "123456"},
		{struct {
			ip   []byte
			port int
		}{[]byte{97, 98, 99, 100}, 25958}, "abcdef"},
	}

	for _, item := range cases {
		info, err := encodeCompactIPPortInfo(item.in.ip, item.in.port)
		if err != nil || info != item.out {
			t.Fail()
		}
	}
}

func TestGetLocalIPs(t *testing.T) {
	ips := getLocalIPs()
	if len(ips) == 0 {
		t.Fail()
	}

	fmt.Println(ips)
}

func TestGetRemoteIP(t *testing.T) {
	ip, err := getRemoteIP()
	if err != nil {
		t.Fail()
	}

	fmt.Println(ip)
}

func TestGenAddress(t *testing.T) {
	cases := []struct {
		ip      string
		port    int
		address string
	}{
		{"127.0.0.1", 80, "127.0.0.1:80"},
		{"127.0.0.1", 0, "127.0.0.1:0"},
	}

	for _, c := range cases {
		if addr := genAddress(c.ip, c.port); addr != c.address {
			t.Fail()
		}
	}
}
