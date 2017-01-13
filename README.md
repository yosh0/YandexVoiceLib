# YandexVoiceLib
Golang implementation of Yandex Voice Library (speech to text)

##Example
```
file := "file.wav"
topic := "q"
keyR := "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
lang := "en-US"
xml0 := yandexVoiceLib.Recognize(file, topic, keyR, lang)
```
