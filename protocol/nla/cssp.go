package nla

import (
	"encoding/asn1"

	"github.com/tomatome/grdp/glog"
)

type CredSSPVersion int

const (
	CredSSPVersion2 CredSSPVersion = 2
	CredSSPVersion5 CredSSPVersion = 5
	CredSSPVersion6 CredSSPVersion = 6
)

type NegoToken struct {
	Data []byte `asn1:"explicit,tag:0"`
}

type TSRequest struct {
	Version    CredSSPVersion `asn1:"explicit,tag:0"`
	NegoTokens []NegoToken    `asn1:"optional,explicit,tag:1"`
	AuthInfo   []byte         `asn1:"optional,explicit,tag:2"`
	PubKeyAuth []byte         `asn1:"optional,explicit,tag:3"`
	ErrorCode  int            `asn1:"optional,explicit,tag:4"`
	Nonce      []byte         `asn1:"optional,explicit,tag:5"`
}

type TSCredentials struct {
	CredType    int    `asn1:"explicit,tag:0"`
	Credentials []byte `asn1:"explicit,tag:1"`
}

type TSPasswordCreds struct {
	DomainName []byte `asn1:"explicit,tag:0"`
	UserName   []byte `asn1:"explicit,tag:1"`
	Password   []byte `asn1:"explicit,tag:2"`
}

type TSCspDataDetail struct {
	KeySpec       int    `asn1:"explicit,tag:0"`
	CardName      string `asn1:"explicit,tag:1"`
	ReaderName    string `asn1:"explicit,tag:2"`
	ContainerName string `asn1:"explicit,tag:3"`
	CspName       string `asn1:"explicit,tag:4"`
}

type TSSmartCardCreds struct {
	Pin        string            `asn1:"explicit,tag:0"`
	CspData    []TSCspDataDetail `asn1:"explicit,tag:1"`
	UserHint   string            `asn1:"explicit,tag:2"`
	DomainHint string            `asn1:"explicit,tag:3"`
}

func EncodeDERTRequest(version CredSSPVersion, msgs []Message, authInfo []byte, pubKeyAuth []byte, nonce []byte) []byte {
	req := TSRequest{
		Version: version,
	}

	if len(msgs) > 0 {
		req.NegoTokens = make([]NegoToken, 0, len(msgs))
	}

	for _, msg := range msgs {
		token := NegoToken{msg.Serialize()}
		req.NegoTokens = append(req.NegoTokens, token)
	}

	if len(authInfo) > 0 {
		req.AuthInfo = authInfo
	}

	if len(pubKeyAuth) > 0 {
		req.PubKeyAuth = pubKeyAuth
	}

	if req.ErrorCode != 0 {
		req.ErrorCode = 0
	}

	if len(nonce) > 0 {
		req.Nonce = nonce
	}

	result, err := asn1.Marshal(req)
	if err != nil {
		glog.Error(err)
	}
	return result
}

func DecodeDERTRequest(s []byte) (*TSRequest, error) {
	treq := &TSRequest{}
	_, err := asn1.Unmarshal(s, treq)
	return treq, err
}

func EncodeDERTCredentials(domain, username, password []byte) []byte {
	tpas := TSPasswordCreds{domain, username, password}
	result, err := asn1.Marshal(tpas)
	if err != nil {
		glog.Error(err)
	}
	tcre := TSCredentials{1, result}
	result, err = asn1.Marshal(tcre)
	if err != nil {
		glog.Error(err)
	}
	return result
}

func DecodeDERTCredentials(s []byte) (*TSCredentials, error) {
	tcre := &TSCredentials{}
	_, err := asn1.Unmarshal(s, tcre)
	return tcre, err
}
