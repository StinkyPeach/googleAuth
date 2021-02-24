package auth

import (
	"testing"
)

func TestGoogleAuth_VerifyCode(t *testing.T) {
	authenticator := NewGoogleAuth()
	secret := authenticator.GetSecret()
	code, err := authenticator.GetCode(secret)
	if err != nil {
		t.Error(err)
	}

	url := authenticator.GetQrcodeUrl("gotest", secret)
	t.Log(secret, code, url)
	isOk, err := authenticator.VerifyCode(secret, code)
	if err != nil {
		t.Error(err)
	}

	if isOk == true {
		t.Log("verify successful")
	}else {
		t.Log("verify failed")
	}

}