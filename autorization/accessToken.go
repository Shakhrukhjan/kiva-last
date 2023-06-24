package autorization

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func AccessToken() string {
	url := "https://auth.k1.kiva.org/oauth/token"
	method := "POST"
	payload := strings.NewReader("grant_type=client_credentials&scope=create%3Aloan_draft%20read%3Aloans&audience=https%3A%2F%2Fpartner-api.k1.kiva.org&client_id=gykq4zLNGpYBm134u9yA5anE2dl36ZPMj&client_secret=h1tTQT_562jlw55qdR_sMEwfoJ_u8luluwDOx%2BHP1CR231LPfWMn29gan6EqC8MQnk")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//fmt.Println(string(body))
	err = ioutil.WriteFile("./token.json", body, 0666)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fileContent, err := os.Open("token.json")

	if err != nil {
		log.Fatal(err)
		return ""
	}

	defer fileContent.Close()
	byteResult, _ := ioutil.ReadAll(fileContent)
	var ress map[string]string
	json.Unmarshal([]byte(byteResult), &ress)
	return ress["access_token"]
}
