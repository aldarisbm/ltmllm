package main

import (
	"github.com/aldarisbm/ltmllm/config"
	"github.com/aldarisbm/ltmllm/llm"
)

func main() {
	conf := config.NewConfig()
	cb := llm.NewChatBot(conf)
	cb.NewPineconeClient()
	//emb := cb.CreateEmbeddings(context.Background(), "Hello, my name is Sasha")
	//cb.Chat()
	//fmt.Println(emb)
	//fmt.Println(emb.Model.String())
	//fmt.Println(emb.Usage)
}
