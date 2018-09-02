package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	// var
	FABRIC       = "fabric"
	CLICONTAINER = "cli"

	// dir
	CONTRACTFATHERDIR = "/opt/data/contracts"
)

func StringEncryption(str string) string {
	md5Hash := md5.New()
	return hex.EncodeToString(md5Hash.Sum([]byte(str)))
}
