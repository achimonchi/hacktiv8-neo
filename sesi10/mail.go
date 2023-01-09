package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"sesi10/config"

	"github.com/crewjam/saml/samlsp"
)

func main() {
	sp, err := samlMiddleware()
	if err != nil {
		panic(err)
	}
	http.Handle("/saml/", sp)

	http.Handle("/mail", sp.RequireAccount(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			name := samlsp.AttributeFromContext(r.Context(), "displayName")
			w.Write([]byte(fmt.Sprintf("Welcome : %s", name)))
		}),
	))
	http.Handle("/hello", sp.RequireAccount(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		}),
	))

	port := fmt.Sprintf(":%s", "5555")
	http.ListenAndServe(port, nil)
}

func samlMiddleware() (*samlsp.Middleware, error) {
	keyPair, err := tls.LoadX509KeyPair(config.SamlCertificatePath, config.SamlPrivateKeyPath)
	if err != nil {
		return nil, err
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		return nil, err
	}

	idpMedatadaURL, err := url.Parse(config.SamlIDP)
	if err != nil {
		return nil, err
	}

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMedatadaURL)
	if err != nil {
		return nil, err
	}

	rootURL, err := url.Parse(config.WebServerRootURL)
	if err != nil {
		return nil, err
	}

	middle, err := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})
	if err != nil {
		return nil, err
	}
	return middle, nil

}
