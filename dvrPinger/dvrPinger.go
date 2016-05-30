package dvrPinger

import (
	"errors"
	"fmt"
	"github.com/ozgurOdun/dvrWeb/dvrDbOps"
	"os/exec"
	"regexp"
	"time"
)

const TIMEOUT = 60

func StartCheckTimer() {
	ticker := time.NewTicker(TIMEOUT * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			CheckDvrStatus()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func checkDvrStatus() {
	dvrs, count := dvrDbOps.GetAllDvr()
	if count > 0 {
		var i int64
		var s string
		var err error
		var flag bool
		for i = 0; i < count; i++ {
			s, err = ping(dvrs[i].IpAddress, false)
			if err != nil {
				fmt.Println(err, dvrs[i].IpAddress)
				if err == errors.New("timeout") {
					fmt.Println("Deleting dvr with ip address:", dvrs[i].IpAddress)
					if flag = dvrDbOps.DeleteDvrByIpAddr(dvrs[i].IpAddress); flag == true {
						fmt.Println("Dvr is deleted successfuly")
					} else {
						fmt.Println("Dvr cannot be deleted!!")
					}
				}
			} else {
				fmt.Println(s, dvrs[i].IpAddress)
			}
		}
	}
}

func ping(host string, ipv6 bool) (string, error) {
	matched, err := regexp.Match(`^[\w._:-]+$`, []byte(host))
	if err != nil {
		return "", err
	}
	if !matched {
		return "", errors.New("invalid host/IP")
	}
	six := ""
	if ipv6 {
		six = "6"
	}
	// -c: packet count, -w: timeout in seconds
	out, err := exec.Command("ping"+six, "-c", "1", "-w", "3", "--", host).Output()
	if err != nil {
		errs := fmt.Sprintf("%s", err)
		if errs == "exit status 1" {
			return "", errors.New("timeout")
		}
		if errs == "exit status 2" {
			return "", errors.New("unknown host")
		}
		return "", err
	}
	r, err := regexp.Compile(`\d+ bytes from .*`)
	if err != nil {
		return "", err
	}
	line := r.Find(out)
	if line == nil {
		return "", errors.New("cannot parse ping output")
	}
	return string(line), nil
}
