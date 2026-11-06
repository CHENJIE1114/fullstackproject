package main

import (
	"bytes"

	"encoding/json"
	"fmt"

	"net/http"
)

func rag(content string) (error, []float64) {

	map1 := map[string]any{
		"model":           "BAAI/bge-m3",
		"input":           []string{content},
		"encoding_format": "float",
	}
	payload, err := json.Marshal(map1)
	if err != nil {
		return err, nil
	}
	req, err := http.NewRequest("POST",
		"https://api.siliconflow.cn/v1/embeddings",
		bytes.NewReader(payload))
	if err != nil {
		return err, nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.SiliconFlowAPIKey)
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return err, nil
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("硅基流动请求失败"), nil
	}

	type A2 struct {
		Emb []float64 `json:"embedding"`
	}

	type A1 struct {
		Data []A2 `json:"data"`
	}
	var embb A1
	err = json.NewDecoder(response.Body).Decode(&embb)
	if err != nil {
		return err, nil
	}
	return nil, embb.Data[0].Emb

}
