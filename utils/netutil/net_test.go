package netutil

import (
	"net"
	"net/http"
	"testing"

	"github.com/x-module/utils/utils/internal"
)

func TestGetInternalIp(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetInternalIp")

	internalIp := GetInternalIp()
	ip := net.ParseIP(internalIp)
	assert.IsNotNil(ip)
}

func TestGetRequestPublicIp(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetPublicIpInfo")

	ip := "36.112.24.10"

	request := http.Request{
		Method: "GET",
		Header: http.Header{
			"X-Forwarded-For": {ip},
		},
	}
	publicIp := GetRequestPublicIp(&request)
	assert.Equal(publicIp, ip)

	request = http.Request{
		Method: "GET",
		Header: http.Header{
			"X-Real-Ip": {ip},
		},
	}
	publicIp = GetRequestPublicIp(&request)
	assert.Equal(publicIp, ip)
}

func TestGetPublicIpInfo(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetPublicIpInfo")

	publicIpInfo, err := GetPublicIpInfo()
	assert.IsNil(err)

	t.Logf("public ip info is: %+v \n", *publicIpInfo)
}

func TestIsPublicIP(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsPublicIP")

	ips := []net.IP{
		net.ParseIP("127.0.0.1"),
		net.ParseIP("192.168.0.1"),
		net.ParseIP("10.91.210.131"),
		net.ParseIP("172.20.16.1"),
		net.ParseIP("36.112.24.10"),
	}

	expected := []bool{false, false, false, false, true}

	for i := 0; i < len(ips); i++ {
		actual := IsPublicIP(ips[i])
		assert.Equal(expected[i], actual)
	}
}

func TestIsInternalIP(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsInternalIP")

	ips := []net.IP{
		net.ParseIP("127.0.0.1"),
		net.ParseIP("192.168.0.1"),
		net.ParseIP("10.91.210.131"),
		net.ParseIP("172.20.16.1"),
		net.ParseIP("36.112.24.10"),
	}

	expected := []bool{true, true, true, true, false}

	for i := 0; i < len(ips); i++ {
		actual := IsInternalIP(ips[i])
		assert.Equal(expected[i], actual)
	}
}

func TestGetIps(t *testing.T) {
	ips := GetIps()
	t.Log(ips)
}

func TestGetMacAddrs(t *testing.T) {
	macAddrs := GetMacAddrs()
	t.Log(macAddrs)
}

func TestEncodeUrl(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsInternalIP")

	urlAddr := "http://www.lancet.com?a=1&b=[2]"
	encodedUrl, err := EncodeUrl(urlAddr)
	if err != nil {
		t.Log(err)
	}

	expected := "http://www.lancet.com?a=1&b=%5B2%5D"
	assert.Equal(expected, encodedUrl)
}
