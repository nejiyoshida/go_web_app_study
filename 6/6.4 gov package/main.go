package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// 引数が空インタフェースなのでなんでも受けられる
func store(data interface{}, filename string) {
	// 可変バッファ。ReadとWriteのメソッドを持つ。
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer) // バッファをエンコーダに渡す。エンコードを実行するとバッファに入れられる
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600) // バッファの内容をファイルに書き込む
	if err != nil {
		panic(err)
	}

}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename) // 生データを読み込んだ
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw) // バッファに生データを渡した
	dec := gob.NewDecoder(buffer)  // バッファをデコードするためにセットした。
	err = dec.Decode(data)         // dataに生データをデコードして入れた
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello hoge", Author: "tanaka"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1") // 子の入れ物postReadにデータをいれてね、なので参照渡し
	fmt.Println(postRead)
}
