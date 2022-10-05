package misskey

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func apiPost(jsonByte []byte, apiEndpoint string) ([]byte, error) {

	req, err := http.NewRequest(
		"POST",
		instanceInfo.Host+"/api/"+apiEndpoint,
		bytes.NewBuffer(jsonByte),
	)
	if err != nil {
		return jsonByte, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return jsonByte, err
	}

	resJsonByte, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode, string(resJsonByte))
		os.Exit(1)
	}

	defer resp.Body.Close()

	return resJsonByte, err

}
