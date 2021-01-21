package ent

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
	"way-jasy-cron/cron/internal/model/ent"
)

func TestJob(t *testing.T) {
	j := &ent.Job{}
	j.Stoppable = 0
	j.Creator = "wjb"
	j.Method = "GET"
	j.Header = "xx"
	j.Spec = "@every 5s"
	j.Comment = "vvv"
	open, err := ent.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/job?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Asia%2FShanghai&charset=utf8mb4,utf8")
	fmt.Println(time.Now())
	if err != nil {
		log.Fatal("1",err)
		return
	}
	//err = open.Schema.Create(context.Background())
	//if err != nil {
	//	log.Fatal("2",err)
	//	return
	//}
	_, err = open.Job.Create().
		SetName(j.Name).SetBody(j.Body).SetComment(j.Comment).SetSpec(j.Spec).SetHeader(j.Header).
		SetMethod(j.Method).SetCreator(j.Creator).SetStoppable(j.Stoppable).
		SetStatus(j.Status).SetUpdater(j.Updater).SetURL(j.URL).SetCtime(time.Now()).SetMtime(time.Now()).
		Save(context.TODO())
	if err != nil {
		log.Fatal("3",err)
		return
	}
}
