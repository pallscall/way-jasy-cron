package test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestRSA(t *testing.T) {
	pri := `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC4OXNBCFdcdyaAxpQCHaFJvCd6qTb5Og0mZCHviqcXah2IrkLj
19/reW5bEOJeD9rWgpzV4XXRW/GnP2PmjonuOuI13/S8b77zAhTDzFQbKU4vhV6D
CQlu1mtka1bnNfocgKyom9WXKeU0GHGkANzIAHaLrHRFVTid7GO3p0PejQIDAQAB
AoGAVxIYLjGCENrj4NN6WvGkLJR4WFon+G1+j5meDHiM9t0ZrmMmjEeYttPC29nE
E88MkHxVIiWYqOX69iBY2DygZaGoDdUi8N3ulFqTygPsAMWaRlnjoDQtOKFmYzsW
RhIY5fqF9ZLNA9EJY//yx1BPiaPX5UO6RCdu/0Wu0em5BiUCQQDJqG5Kevyy4cQd
l77urK2oO9AaMHCH23BwcOp3+PAj28YKygrqvLvoaL8N8hOPS2zyOfSbR3kBrVAu
X499geZjAkEA6d5Z5JAJ9t7VcStd6ZRyOgLgcj8a6VYnVjFQ3vtoK6gii9Y8EAEf
klaau8/Hj3RMHB3zrS8Uh42HgdCv/2sCTwJBAKUfHIgyIn7a7HoKyhWnIV8S8viu
10X3Qh9f5i2skf3atFQbUksZlYfdVSu3H5sC+Mdy+z62jHf6ESCXsOflGN8CQGWU
sNL4RXu3WlpnjckRM3RDH55ADr5fL4LetNPFu3+K7kZy7W++LSpw95CwNWv4Bb8I
u3jzxCUEsmEOdgOpRT0CQQCBnr2pEN8crm061iVhXncU6W5UKvX6mY9llvhym3/2
qGmpPeedQPzsg+f5tdrT1PsztnmtOeEueT8aIJdx7k38
-----END RSA PRIVATE KEY-----
`

	pwd := `N18b44JcSvn4hp1lUKXGrs1JAzktbDatrxNKMkJk4BlQqAFmeQwyvi1x7XvoO4slpxZyJls6LeXaJf3CuJu2Puk2bV6Nlz4hq9HQC0oabZ/PReyQ+TfjFt2DoPnjS7qM3awoYhB9YBxKhnPS6fXVW3VeosYFqfYgI0g8TYWhHRw=`
	block, _ := pem.Decode([]byte(pri))
	fmt.Println(block)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	fmt.Println(privateKey)
	decodeString, _ := base64.StdEncoding.DecodeString(pwd)
	plainText,_:= rsa.DecryptPKCS1v15(rand.Reader,privateKey,[]byte(decodeString))
	fmt.Println(string(plainText))
}


func GenRsaKey(bits int) (publicKeyStr, privateKeyStr string, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	bufferPrivate := new(bytes.Buffer)
	err = pem.Encode(bufferPrivate, block)
	if err != nil {
		return
	}
	privateKeyStr = bufferPrivate.String()
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, block)
	if err != nil {
		return
	}
	publicKeyStr = bufferPublic.String()
	log.Debug("-------------公钥----------------")
	fmt.Printf("%s\n",publicKeyStr)
	fmt.Printf("%d\n",len(publicKeyStr))
	log.Debug("--------------私钥---------------")
	fmt.Printf("%s\n",privateKeyStr)
	return
}
