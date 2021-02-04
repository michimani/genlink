package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

func Generate(url, genType, outDir string) (string, error) {
	err := validate(url, genType)
	if err != nil {
		return "", err
	}

	p, err := Fetch(url)
	if err != nil {
		return "", err
	}

	switch genType {
	case "md":
		return p.md(), nil
	case "html":
		return p.html(false), nil
	case "html-bl":
		return p.html(true), nil
	case "qr":
		return p.qr(outDir)
	}

	return "", errors.New("Invalid generate type.")
}

func validate(url, genType string) error {
	if url == "" {
		return errors.New("flag -u: expected one argument")
	}

	if genType == "" {
		return errors.New("flag -t: expected one argument")
	}

	return nil
}

func (p *Page) md() string {
	return fmt.Sprintf("[%s](%s)", p.Title, p.Url)
}

func (p *Page) html(blank bool) string {
	target := ""
	if blank {
		target = " target=\"_blank\" rel=\"noopener\""
	}
	return fmt.Sprintf("<a href=\"%s\"%s>%s</a>", p.Url, target, p.Title)
}

func (p *Page) qr(outDir string) (string, error) {
	cd, _ := os.Getwd()
	if outDir == "" {
		outDir = cd
	}

	if f, err := os.Stat(outDir); os.IsNotExist(err) || !f.IsDir() {
		return "", err
	}

	r := sha256.Sum256([]byte(p.Url + strconv.FormatInt(time.Now().UnixNano(), 10)))
	fname := fmt.Sprintf("%s%s%s.png", outDir, string(os.PathSeparator), hex.EncodeToString(r[:]))

	err := qrcode.WriteFile(p.Url, qrcode.Medium, 256, fname)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("QR code has been created.\n%s", fname), nil
}
