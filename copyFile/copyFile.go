package CopyFile

import (
	"io/ioutil"
	"log"
)

func Copy(src string, dst string) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(dst, data, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
