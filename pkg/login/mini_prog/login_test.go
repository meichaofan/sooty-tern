package mini_prog_test

import (
	"fmt"
	"sooty-tern/pkg/login/mini_prog"
	"testing"
)

func TestCode2Session(t *testing.T) {
	res, err := mini_prog.Code2Session()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("openId: %s\n", res.OpenID)
	fmt.Printf("sessionKey: %s\n", res.SessionKey)
	fmt.Printf("unionId: %s\n", res.UnionID)
}
