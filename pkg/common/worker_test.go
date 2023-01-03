package common

import (
	"net"
	"testing"
	"time"

	"github.com/slackhq/nebula/cert"
)

func Test_VerifyValidCertificatShouldSucceed(t *testing.T) {

	validUntil := time.Now()
	validUntil = validUntil.AddDate(0, 0, 1)

	nc := buildNebulaCert(time.Date(1, 0, 0, 1, 0, 0, 0, time.UTC), validUntil)

	pemContent, err := nc.MarshalToPEM()
	if err != nil {
		t.Fatal(err)
	}

	validationResult, err := ValidateCertificate(pemContent)
	if err != nil {
		t.Fatal(err)
	}

	if validationResult.Expired {
		t.Errorf("Cert is not expired; got true; expected false")
	}
}

func Test_VerifyExpiredCertificatShouldSucceed(t *testing.T) {

	validUntil := time.Now()

	nc := buildNebulaCert(time.Date(1, 0, 0, 1, 0, 0, 0, time.UTC), validUntil)

	pemContent, err := nc.MarshalToPEM()
	if err != nil {
		t.Fatal(err)
	}

	validationResult, err := ValidateCertificate(pemContent)
	if err != nil {
		t.Fatal(err)
	}

	if !validationResult.Expired {
		t.Errorf("Cert is expired; got false; expected true")
	}
}

func buildNebulaCert(notBefore time.Time, notAfter time.Time) cert.NebulaCertificate {

	time.Local = time.UTC
	pubKey := []byte("1234567890abcedfghij1234567890ab")

	return cert.NebulaCertificate{
		Details: cert.NebulaCertificateDetails{
			Name: "testing",
			Ips: []*net.IPNet{
				{IP: net.ParseIP("10.1.1.1"), Mask: net.IPMask(net.ParseIP("255.255.255.0"))},
				{IP: net.ParseIP("10.1.1.2"), Mask: net.IPMask(net.ParseIP("255.255.0.0"))},
				{IP: net.ParseIP("10.1.1.3"), Mask: net.IPMask(net.ParseIP("255.0.255.0"))},
			},
			Subnets: []*net.IPNet{
				{IP: net.ParseIP("9.1.1.1"), Mask: net.IPMask(net.ParseIP("255.0.255.0"))},
				{IP: net.ParseIP("9.1.1.2"), Mask: net.IPMask(net.ParseIP("255.255.255.0"))},
				{IP: net.ParseIP("9.1.1.3"), Mask: net.IPMask(net.ParseIP("255.255.0.0"))},
			},
			Groups:    []string{"test-group1", "test-group2", "test-group3"},
			NotBefore: notBefore, //time.Date(1, 0, 0, 1, 0, 0, 0, time.UTC),
			NotAfter:  notAfter,  //time.Date(1, 0, 0, 2, 0, 0, 0, time.UTC),
			PublicKey: pubKey,
			IsCA:      false,
			Issuer:    "1234567890abcedfghij1234567890ab",
		},
		Signature: []byte("1234567890abcedfghij1234567890ab"),
	}
}
