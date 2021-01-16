package test

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"net"
	"testing"
)

func TestShell(t *testing.T) {
	var (
		client *ssh.Client
		session *ssh.Session
		err error
	)
	addr := "106.52.6.144:22"
	if client, err = ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{ssh.Password("#Wjb123456")},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}); err != nil {
		log.Error(err)
		return
	}

	defer client.Close()
	if session, err = client.NewSession(); err != nil {
		log.Error(err)
		return
	}
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	session.Stdout = bufOut
	session.Stderr = bufErr
	//执行命令
	if err = session.Run("mkdir test"); err != nil {
		log.Error(err)
		return
	}
}
