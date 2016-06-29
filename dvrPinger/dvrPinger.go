package dvrPinger

import (
	"fmt"
	"github.com/ozgurOdun/dvrWeb/dvrDbOps"
	"time"
)

const TIMEOUT = 10

func StartCheckTimer() {
	fmt.Println("Timer started.")
	ticker := time.NewTicker(TIMEOUT * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			fmt.Println("ping")
			checkDvrStatus()
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
		for i = 0; i < count; i++ {
			if dvrs[i].Status == dvrDbOps.Alive {
				_ = dvrDbOps.UpdateDvrStatusByName(dvrs[i].Name, dvrDbOps.Dead)
			} else if dvrs[i].Status == dvrDbOps.Dead {
				_ = dvrDbOps.DeleteDvrByName(dvrs[i].Name)

			} else {
				fmt.Println("none")
			}

		}
	} else {
		fmt.Println("No dvrs found.")
	}
}
