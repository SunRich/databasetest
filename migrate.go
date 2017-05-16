package databasetest

import (
	"os/exec"
	"fmt"
	"os"
)

func InitDatabase(path, driver, conn string)  {
	dir:=os.Getenv("GOPATH") + "/src/" + path
	resetDatabase(dir,driver, conn)
	migDatabase(dir,driver, conn)
}
func resetDatabase(dir, driver, conn string) {
	cmd := exec.Command("bash", "-c", "bee migrate reset -driver="+driver+" -conn='"+conn+"'")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
func migDatabase(dir, driver, conn string)  {
	cmd := exec.Command("bash", "-c", "bee migrate -driver="+driver+" -conn='"+conn+"'")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
