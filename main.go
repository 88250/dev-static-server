// Dev Static Server - HTTP server for dev.
// Copyright (c) 2019-present, b3log.org
//
// Dev Static Server is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

func main() {
	portArg := flag.String("p", "9090", "port, default is 9090")
	flag.Parse()

	ifaces, err := net.Interfaces()
	if nil != err {
		fmt.Println("can't get interfaces")
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if nil == err {
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP.To4()
				case *net.IPAddr:
					ip = v.IP.To4()
				}
				if nil != ip {
					fmt.Println(ip.String())
				}
			}
		}
	}

	port := *portArg
	fmt.Println(":" + port)
	http.Handle("/", cors(http.FileServer(http.Dir(`.`))))
	panic(http.ListenAndServe(":"+port, nil))
}

func cors(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
