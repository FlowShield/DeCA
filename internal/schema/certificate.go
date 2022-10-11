package schema

import (
	"time"

	"github.com/cloudslit/deca/pkg/util/json"
)

const (
	CrdtSnRevokePrefix = "certificate_sn_revoke_"
	CrdtSnCidPrefix    = "certificate_sn_cid_"
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

// CertificateStatus 证书状态信息
type CertificateRevoke struct {
	SerialNumber string    `json:"serial_number"`
	RevokeAt     time.Time `json:"revoke_at"`
}

func (a *CertificateRevoke) String() string {
	return json.MarshalToString(a)
}

// 吊销证书记录存储key
func SnRevokeKey(key string) string {
	return CrdtSnRevokePrefix + key
}

// 证书cid存储key
func SnCidKey(key string) string {
	return CrdtSnCidPrefix + key
}
