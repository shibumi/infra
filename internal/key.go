package internal

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

type SSHPublicKey struct {
	Bytes   []byte
	Comment string
}

func ReadSSHPublicKey(path string) (*SSHPublicKey, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var k SSHPublicKey
	_, k.Comment, _, _, err = ssh.ParseAuthorizedKey(raw)
	if err != nil {
		return nil, err
	}
	k.Bytes = raw
	return &k, nil
}
