package downloader

import (
	"io/ioutil"
	"net/http"
	"os"
)


//Download download media file with given url
func Download(url string) error {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	bs, err := ioutil.ReadAll(res.Body)


	f, err := os.OpenFile("tmp/sample.mp4", os.O_RDONLY|os.O_WRONLY, 0755)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(bs)

	return err
}