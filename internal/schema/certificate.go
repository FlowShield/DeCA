package schema

import (
	"github.com/cloudslit/newca/pkg/util/json"
	"time"
)

type Certificate struct {
	SerialNumber           string                 `json:"serial_number"`
	AuthorityKeyIdentifier string                 `json:"authority_key_identifier"`
	CertPem                string                 `json:"cert_pem"`
	NotBefore              time.Time              `json:"not_before"`
	NotAfter               time.Time              `json:"not_after"`
	MetaData               map[string]interface{} `json:"meta_data"`
}

func (a *Certificate) String() string {
	return json.MarshalToString(a)
}
