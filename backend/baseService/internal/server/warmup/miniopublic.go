package warmup

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

func InitMinioPublicDirectory() {
	entrypoint := "http://minio:9000"
	username := "root"
	password := "rootroot"
	bucketNameList := []string{"shortvideo"}

	// 定义可能的 mc 路径
	possiblePaths := []string{
		"./mc",              // 相对路径
		"/usr/local/bin/mc", // 常用的绝对路径
		"/home/gsk/.gvm/pkgsets/go1.22.8/global/bin/mc",
	}
	// 选择一个存在的路径
	var mcPath string
	for _, path := range possiblePaths {
		if fileExists(path) {
			mcPath = path
			break
		}
	}
	// 如果没有找到有效路径，记录错误并退出
	if mcPath == "" {
		log.Fatal("mc tool not found in any known path")
	}
	fmt.Println("os command params", mcPath, "config", "host", "add", "minio", entrypoint, username, password)
	if output, err := exec.Command(mcPath, "config", "host", "add", "minio", entrypoint, username, password).CombinedOutput(); err != nil {
		log.Errorf("mc config host add err: %v, output: %s", err, output)
		if strings.Contains(err.Error(), "502 Bad Gateway") {
			if output, err = exec.Command(mcPath, "config", "host", "add", "minio", "http://127.0.0.1:9000", username, password).CombinedOutput(); err != nil {
				log.Errorf("try `http://127.0.0.1:9000`, mc config host add err: %v, output: %s", err, output)
				panic(err)
			}
		}

	}
	for _, bucketName := range bucketNameList {
		if output, err := exec.Command(mcPath, "anonymous", "set", "public", "minio/"+bucketName+"/public").CombinedOutput(); err != nil {
			log.Errorf("mc anonymous set err: %v, output: %s", err, output)
			panic(err)
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
