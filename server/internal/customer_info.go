package internal

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

type CustomerInfo map[string]string

func (u CustomerInfo) Empty() bool {
	return len(u) == 0
}

func GenerateCustomerID(u CustomerInfo) string {
	u[CustomerFieldKeyID] = "" // remove id

	b, _ := json.Marshal(u)
	sum := sha1.Sum(b)
	return hex.EncodeToString(sum[:])
}
