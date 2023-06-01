package main

import "github.com/shine-bright-team/LAAS/v2/initialize"

func main() {
	initialize.LookForEnv()
	initialize.DbSetUp()
}
