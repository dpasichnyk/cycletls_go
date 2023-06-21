package cycletls_test

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"testing"

	cycletls "github.com/dpasichnyk/cycletls_go"
)

func SimpleFileWriteTest(t *testing.T) {

	client := cycletls.Init()
	response, err := client.Do("http://httpbin.org/image/jpeg", cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
	}, "GET")
	if err != nil {
		log.Print("Request Failed: " + err.Error())
	}
	// Decode Base64
	dec, err := base64.StdEncoding.DecodeString(response.Body)
	if err != nil {
		panic(err)
	}
	//create file to write
	f, err := os.Create("test.jpeg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//write b64 to file
	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
}

func WriteFile(Body string, Filepath string) {
	// Decode Base64
	dec, err := base64.StdEncoding.DecodeString(Body)
	if err != nil {
		panic(err)
	}
	//create file to write
	f, err := os.Create(Filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//write b64 to file
	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

}

func CompareFiles(filepath1 string, filepath2 string) bool {
	f1, err1 := ioutil.ReadFile(filepath1)

	if err1 != nil {
		log.Fatal(err1)
	}

	f2, err2 := ioutil.ReadFile(filepath2)

	if err2 != nil {
		log.Fatal(err2)
	}

	return bytes.Equal(f1, f2)
}
func GetRequest(url string, client cycletls.CycleTLS) cycletls.Response {
	resp, err := client.Do(url, cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
	}, "GET")
	if err != nil {
		log.Print("Request Failed: " + err.Error())
	}
	return resp
}
