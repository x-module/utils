package main

// import (
// 	"fmt"
// 	"study/utils"
// )
//
// func main() {
// 	response, err := utils.RequestUtils.SetCookies(map[string]string{
// 		"name":     "124",
// 		"password": "1234",
// 	}).SetHeaders(map[string]string{
// 		"auth": "super man",
// 		"sign": "sign the request",
// 	}).Debug().SetTimeout(1).Json().Get("http://127.0.0.1:9090", map[string]interface{}{
// 		"username": "username",
// 		"password": "password",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(response.Content())
//
// }
