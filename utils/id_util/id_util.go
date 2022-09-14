package id_util

import (
	"fmt"
	"strconv"

	idworker "github.com/gitstliu/go-id-worker"
)

var currWorker *idworker.IdWorker

func init()  {
	currWorker = &idworker.IdWorker{}
	currWorker.InitIdWorker(2,1)
}

func GenID() string{
	newID, err := currWorker.NextId()
	if err == nil {
		fmt.Println(newID)
	}
	return strconv.FormatInt(newID,10)
}