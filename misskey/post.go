package misskey

import (
	"bytes"
	"fmt"
	"net/http"
)

func apiPost(jsonByte []byte, apiEndpoint string) error {

	fmt.Println(jsonByte)

	fmt.Println("API")
	req, err := http.NewRequest(
		"POST",
		instanceInfo.Host+"/api/"+apiEndpoint,
		bytes.NewBuffer(jsonByte),
	)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return err
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	return err
}
