package main

import (
	"context"
	"errors"
	"flag"
	"github.com/golang/groupcache"
	"log"
	"net/http"
	"strings"
)

var DBStore = map[string]string{
	"red":   "红色",
	"green": "绿色",
	"blue":  "蓝色",
}

var Group = groupcache.NewGroup("foobar", 64<<20, groupcache.GetterFunc(
	func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
		log.Println("读DB", key)
		v, ok := DBStore[key]
		if !ok {
			return errors.New("color not found")
		}
		dest.SetString(v)
		return nil
	},
))

func main() {
	addr := flag.String("addr", ":8080", "server address")
	peers := flag.String("pool", "http://localhost:8080", "server pool list")
	flag.Parse()
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		color := r.FormValue("name")
		log.Printf("获取 name=%s", color)
		var b []byte
		err := Group.Get(context.Background(), color, groupcache.AllocatingByteSliceSink(&b))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Write(b)
		w.Write([]byte{'\n'})
	})
	p := strings.Split(*peers, ",")
	pool := groupcache.NewHTTPPool(p[0])
	pool.Set(p...) // 设置服务器列表
	http.ListenAndServe(*addr, nil)
}
