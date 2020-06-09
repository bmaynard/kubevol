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
	cmd.SetArgs([]string{"--namespace", "kubevol-test-run", "configmap"})
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

func Test_RedisConfigMapOutdated(t *testing.T) {
	cmd := cmd.NewKubevolApp()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--namespace", "kubevol-test-run", "configmap"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	matched, err := regexp.MatchString("Yes", string(out))
	if !matched {
		t.Fatalf("expected \"%s\" got \"%s\"", "Yes", string(out))
	}
}

func Test_RedisSecretExists(t *testing.T) {
	cmd := cmd.NewKubevolApp()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--namespace", "kubevol-test-run", "secret"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	matched, err := regexp.MatchString("redis-secret", string(out))
	if !matched {
		t.Fatalf("expected \"%s\" got \"%s\"", "redis-secret", string(out))
	}
}

func Test_RedisSecretOutdated(t *testing.T) {
	cmd := cmd.NewKubevolApp()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--namespace", "kubevol-test-run", "secret"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	matched, err := regexp.MatchString("Yes", string(out))
	if !matched {
		t.Fatalf("expected \"%s\" got \"%s\"", "Yes", string(out))
	}
}
