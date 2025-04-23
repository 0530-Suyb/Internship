package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

}

func deepseekCall() {
	url := "https://api.deepseek.com/chat/completions"
	method := "POST"

	payload := strings.NewReader(`{
  "messages": [
    {
      "content": "You are a helpful assistant",
      "role": "system"
    },
    {
      "content": "Hi",
      "role": "user"
    }
  ],
  "model": "deepseek-chat",
  "frequency_penalty": 0,
  "max_tokens": 2048,
  "presence_penalty": 0,
  "response_format": {
    "type": "text"
  },
  "stop": null,
  "stream": false,
  "stream_options": null,
  "temperature": 1,
  "top_p": 1,
  "tools": null,
  "tool_choice": "none",
  "logprobs": false,
  "top_logprobs": null
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer sk-7b0f2df52b9349bf8e5cc46424e4ef0a")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func silliconCall() {
	url := "https://api.siliconflow.cn/v1/chat/completions"

	payload := strings.NewReader("{\n  \"model\": \"deepseek-ai/DeepSeek-V3\",\n  \"messages\": [\n    {\n      \"role\": \"user\",\n      \"content\": \"中国大模型行业2025年将会迎来哪些机遇和挑战？\"\n    }\n  ],\n  \"stream\": false,\n  \"max_tokens\": 512,\n  \"stop\": [\n    \"null\"\n  ],\n  \"temperature\": 0.7,\n  \"top_p\": 0.7,\n  \"top_k\": 50,\n  \"frequency_penalty\": 0.5,\n  \"n\": 1,\n  \"response_format\": {\n    \"type\": \"text\"\n  },\n  \"tools\": [\n    {\n      \"type\": \"function\",\n      \"function\": {\n        \"description\": \"<string>\",\n        \"name\": \"<string>\",\n        \"parameters\": {},\n        \"strict\": false\n      }\n    }\n  ]\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "Bearer sk-iixjviygojyctldwourbxgdcobtmqrhsfkfdhkyvaamqwqgy")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
