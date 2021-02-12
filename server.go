package webify

import (
	"fmt"
	"net/http"
)




func NewServer(port string) {
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}