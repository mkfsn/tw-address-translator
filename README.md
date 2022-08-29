# Taiwan address translator

This repository contains three components:

- Go library: the library to translate a Taiwan address from Chinese to English
- Web app: https://tw-address-translator.vercel.app
- API: the api to translate a Taiwan address from Chinese to English

## Go library

### Installation

```
go get github.com/mkfsn/tw-address-translator
```

### How to use

```go
package main

import (
	"log"
	
	"github.com/mkfsn/tw-address-translator/pkg/translator"
)

func main() {
	t := translator.New()
	
	englishAddress, err := t.Translate("/* Valid Chinese address */")
	if err != nil {
		log.Fatalln(err)
	}
	
	log.Println(englishAddress)
}
```


## Web app

[https://tw-address-translator.vercel.app](https://tw-address-translator.vercel.app)

## API

TODO

# References

## Data source

- [中華郵政全球資訊網 ▶︎ 下載專區 ▶︎ 下載項目一覽](https://www.post.gov.tw/post/internet/Download/all_list.jsp?ID=2201#dl_txt_A09)
  - [6.1 縣市鄉鎮中英對照Excel檔(漢語拼音)](https://www.post.gov.tw/post/download/county_h_10706.xls)
  - [6.3 村里文字巷中英對照Excel檔 106/02(漢語拼音)](https://www.post.gov.tw/post/download/Village_H_10602.xls)
  - [6.5 路街中英對照Excel檔 111/07(漢語拼音)](https://www.post.gov.tw/post/download/6.5_CEROAD11107.xlsx)
