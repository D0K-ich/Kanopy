package movie

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetManga() {

	url := "https://mangalib.me"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}