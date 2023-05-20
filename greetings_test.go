package greetings

import (
    "testing"
    "regexp"
)

func TestHelloName(t *testing.T) {
    name := "skullArms"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("skullArms")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("skullArms") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

