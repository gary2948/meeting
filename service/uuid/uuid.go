package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenUUID() (string, error) {
	uuid := make([]byte, 10)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// TODO: verify the two lines implement RFC 4122 correctly

	uuid[2] = 0x80 // variant bits see page 5
	uuid[0] = 0x40 // version 4 Pseudo Random, see page 7
	s := strings.ToUpper(hex.EncodeToString(uuid))
	return fmt.Sprintf("%s-%s-%s-%s-%s", s[0:4], s[4:8], s[8:12], s[12:16], s[16:20]), nil

}

func GenUUID4() (string, error) {
	uuid := make([]byte, 2)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// TODO: verify the two lines implement RFC 4122 correctly
	//uuid[2] = 0x80 // variant bits see page 5
	//uuid[0] = 0x40 // version 4 Pseudo Random, see page 7
	s := strings.ToUpper(hex.EncodeToString(uuid))
	return fmt.Sprintf("%s", s[0:4]), nil

}
