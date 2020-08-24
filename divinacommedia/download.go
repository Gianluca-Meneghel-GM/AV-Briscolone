package divinacommedia

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DownloadFromGithub() error {
	shouldDownload := false

	_, statErr := os.Stat("etag.txt")
	if os.IsNotExist(statErr) {
		fmt.Println("ETAG not found")
		shouldDownload = true
	} else {
		fmt.Println("ETAG found")
		headResp, err := http.DefaultClient.Head(`https://raw.githubusercontent.com/genez/dante/master/dante.txt`)
		if err != nil {
			log.Println(err.Error())
			return errors.New("Could not fetch headers for divinacommedia txt")
		}
		defer headResp.Body.Close()

		currentETag := headResp.Header.Get("ETag")
		fmt.Printf("CURRENT ETAG: %s\n", currentETag)

		oldETag, err := ioutil.ReadFile("etag.txt")
		if err != nil {
			log.Println(err.Error())
			return errors.New("Could not read etag txt")
		}

		fmt.Printf("OLD ETAG: %s\n", string(oldETag))
		if currentETag != string(oldETag) {
			shouldDownload = true
			fmt.Println("ETAGS are different")
		}
	}

	if shouldDownload {
		fmt.Println("downloading file...")
		resp, err := http.DefaultClient.Get(`https://raw.githubusercontent.com/genez/dante/master/dante.txt`)
		if err != nil {
			log.Println(err.Error())
			return errors.New("Could not download divinacommedia txt")
		}
		defer resp.Body.Close()

		err = ioutil.WriteFile("etag.txt", []byte(resp.Header.Get("ETag")), os.ModePerm)
		if err != nil {
			log.Println(err.Error())
			return errors.New("Could write etag txt file")
		}

		buff, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			return errors.New("Could not read HTTP response body")
		}

		err = ioutil.WriteFile("divinacommedia.txt", buff, os.ModePerm)
		if err != nil {
			log.Println(err.Error())
			return errors.New("Could not write file divinacommedia txt")
		}

		fmt.Println("...done")
	}

	return nil
}
