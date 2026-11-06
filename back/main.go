package main

import ()

type Notes struct {
	ID          int     `json:"id"`
	UserID      int     `json:"userid"`
	Title       string  `json:"title"`
	Content     string  `json:"content"`
	Createdtime string  `json:"createdtime"`
	Updatetime  string  `json:"updatetime"`
	Cos         float64 `json:"cos"`
	Embedding   string  `json:"embedding"`
}

func main() {
	r := api()
	r.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}
