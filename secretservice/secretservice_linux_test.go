package secretservice

import (
	"testing"

	"github.com/docker/docker-credential-helpers/credentials"
)

func TestSecretServiceHelper(t *testing.T) {
	t.Skip("test requires gnome-keyring but travis CI doesn't have it")

	creds := &credentials.Credentials{
		ServerURL: "https://foobar.docker.io:2376/v1",
		Username:  "foobar",
		Password:  "foobarbaz",
	}

	helper := New()
	if err := helper.Add(creds); err != nil {
		t.Fatal(err)
	}

	username, password, err := helper.Get(creds.ServerURL)
	if err != nil {
		t.Fatal(err)
	}

	if username != "foobar" {
		t.Fatalf("expected %s, got %s\n", "foobar", username)
	}

	if password != "foobarbaz" {
		t.Fatalf("expected %s, got %s\n", "foobarbaz", password)
	}

	if err := helper.Delete(creds.ServerURL); err != nil {
		t.Fatal(err)
	}
}

func TestMissingCredentials(t *testing.T) {
	t.Skip("test requires gnome-keyring but travis CI doesn't have it")

	helper := New()
	_, _, err := helper.Get("https://adsfasdf.wrewerwer.com/asdfsdddd")
	if err != credentials.ErrCredentialsNotFound {
		t.Fatalf("exptected ErrCredentialsNotFound, got %v", err)
	}
}
