package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	log "github.com/sirupsen/logrus"
	"way-jasy-cron/user/ecode"
	"way-jasy-cron/user/internal/model/ent"
)

func (svc *Service) Register(ctx context.Context, user *ent.User) error{
	if err := svc.DecodePwd(ctx, user); err != nil {
		log.Error("decode pwd err:", err)
		return err
	}
	if err := svc.ent.Register(ctx, user); err != nil {
		log.Error("method: Register#service. Register err:", err)
		return err
	}
	log.Info("Register success!")
	return nil
}

func (svc *Service) GetPublicKey(ctx context.Context, username string) (string, error) {
	key, err := svc.ent.GetPublicKey(ctx, username)
	if err != nil {
		log.Error("GetPublicKey err:",err)
	}
	return key, err
}

func (svc *Service) GenerateRSA(ctx context.Context, username string) (string, error) {
	u, err := svc.ent.QueryUser(ctx, username)
	if err != nil {
		log.Error("method: Register#service. QueryUser err:", err)
		return "", err
	}
	if u != nil {
		return "", ecode.UserExist
	}
	pubKey, priKey, err := GenRsaKey(1024)
	if err != nil {
		log.Error("GenRsaKey err:", err)
		return "", err
	}
	if err := svc.ent.SaveRsaKey(ctx, username, pubKey, priKey); err != nil {
		log.Error("SaveRsaKey err:",err)
		return "", err
	}
	return pubKey, err
}

func GenRsaKey(bits int) (publicKeyStr, privateKeyStr string, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	bufferPrivate := new(bytes.Buffer)
	err = pem.Encode(bufferPrivate, block)
	if err != nil {
		return
	}
	privateKeyStr = bufferPrivate.String()
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, block)
	if err != nil {
		return
	}
	publicKeyStr = bufferPublic.String()
	return
}

func (svc *Service) DecodePwd(ctx context.Context, user *ent.User) error{
	u, err := svc.ent.QueryUser(ctx, user.Username)
	if err != nil {
		log.Error("method: DecodePwd#user, QueryUser err:",err)
		return err
	}
	block, _ := pem.Decode([]byte(u.PrivateKey))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	decodeString, _ := base64.StdEncoding.DecodeString(user.Password)
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader,privateKey,decodeString)
	user.Password = string(plainText)
	return err
}

