package helper

import (
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

func MergePDF(files []string) (out string, err error) {
	log.Println("files is:", files) // [./upload/1.pdf ./upload/5.pdf]
	fileVars := strings.Join(files, " ")
	log.Println("file vars:", fileVars) // ./upload/1.pdf ./upload/5.pdf

	// ready for out combined file
	rand.Seed(int64(time.Now().UnixNano()))
	out = "./upload/" + RandString(10) + ".pdf"

	plainCmd := "cpdf -merge " + fileVars + " -o " + out
	log.Println("plaincmd:", plainCmd) // cpdf -merge ./upload/1.pdf ./upload/5.pdf -o out12.pdf
	sliceA := strings.Fields(plainCmd)
	cmd := exec.Command(sliceA[0], sliceA[1:]...)

	err = cmd.Run()
	if err != nil {
		log.Println("cmd run error:", err)
		return "", err
	}

	return out, err
}
