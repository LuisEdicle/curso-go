package utils

import "encoding/base64"


func EncodeString(s string) string{
	encode := base64.StdEncoding.EncodeToString([]byte(s))
	return encode
}

func DecodeString(s string) (string, error){
	decode, err := base64.StdEncoding.DecodeString(s)
	if err != nil{
		return "", err
	}
	return string(decode), nil
}