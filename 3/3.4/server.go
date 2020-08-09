package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem") // certの方が証明書で、keyが秘密鍵

	//oreore()
}

// おれおれSSL証明書と秘密鍵生成
func oreore() {
	// bigは64bitとかに収まらないような値を使うときのパッケージ。Lshは Lsh sets z = x << n and returns z. らしいので、シフトして戻す
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max) // 証明書のシリアル暗号。本来は認証局（CA）に発行してもらうやつ
	subject := pkix.Name{
		Organization:       []string{"Hogehoge "},
		OrganizationalUnit: []string{"book"},
		CommonName:         "Go app hogehoge",
	}
	template := x509.Certificate{
		SerialNumber: serialNumber, // 発行してもらったシリアル番号を証明書に設定。
		Subject:      subject,      // XXXの証明書ですよ～ってのを設定
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // bitmap, 暗号化、デジタル署名
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // サーバ認証に使うよ
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},                           // このIPの時だけこの証明書を有効にする
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

}
