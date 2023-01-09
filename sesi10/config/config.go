package config

import "fmt"

var (
	SamlCertificatePath = "./svc.cert"
	SamlPrivateKeyPath  = "./svc.key"
	SamlIDP             = "https://samltest.id/saml/idp"

	WebServerPort    = "4444"
	WebServerRootURL = fmt.Sprintf("http://localhost:%s", WebServerPort)
)
