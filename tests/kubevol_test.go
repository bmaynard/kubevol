package tests

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/bmaynard/kubevol/cmd"
)

func Test_RedisConfigMapExists(t *testing.T) {
	cmd := cmd.NewKubevolApp()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"configmap"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	matched, err := regexp.MatchString("redis-config", string(out))
	if !matched {
		t.Fatalf("expected \"%s\" got \"%s\"", "redis-config", string(out))
	}
}
