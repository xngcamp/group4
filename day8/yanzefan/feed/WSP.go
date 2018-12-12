// Code generated by wsp, DO NOT EDIT.

package main

import "net/http"
import "time"
import "camp/feed/controller/feed"
import "camp/feed/controller/user"
import "camp/feed/filter"

func init() {
	http.HandleFunc("/feed/add", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/add"}); !ok {
				return
			}
		}()
		c.Add(w, r)
	})

	http.HandleFunc("/feed/del", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/del"}); !ok {
				return
			}
		}()
		c.Del(w, r)
	})

	http.HandleFunc("/feed/get_all", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/get_all"}); !ok {
				return
			}
		}()
		c.GetAll(w, r)
	})

	http.HandleFunc("/feed/update", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(feed.Feed)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/feed/update"}); !ok {
				return
			}
		}()
		c.Update(w, r)
	})

	http.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(user.User)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/user/login"}); !ok {
				return
			}
		}()
		c.Login(w, r)
	})

	http.HandleFunc("/user/add_user", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(user.User)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/user/add_user"}); !ok {
				return
			}
		}()
		c.AddUser(w, r)
	})

	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(user.User)
		defer func() {
			e = recover()
			if ok := filter.Boss(w, r, map[string]interface{}{"__T__": t, "__C__": c, "__E__": e, "__P__": "/user/update"}); !ok {
				return
			}
		}()
		c.Update(w, r)
	})

}