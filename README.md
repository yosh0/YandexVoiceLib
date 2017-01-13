# YandexVoiceLib
Golang implementation of Yandex Voice Library (speech to text)

##Example
```go
package main

import (
	"fmt"
	"github.com/yosh0/YandexVoiceLib"
)

func main() {
        file := "file.wav"
        topic := "queries|maps|buying|notes"
        keyR := "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
        lang := "en-US"
        xml0 := yandexVoiceLib.Recognize(file, topic, keyR, lang)
        fmt.Println(string(xml0))
}
```
