package yandexVoiceLib

import (
	"os"
	"io"
	"fmt"
	"math"
	"time"
	"bytes"
	"net/url"
	"strings"
	"net/http"
	"io/ioutil"
	"math/rand"
)

func Recognize(file string, topic string, key string, lang string) (body []byte) {
	tUuid := generateRandomSelection(0, 30, 64)
	uuid := strings.Join(tUuid, "")
	uuid = uuid[0:32]
	f1, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	payload := io.MultiReader(f1)
	url := fmt.Sprintf("https://asr.yandex.net/asr_xml?key=%s&uuid=%s&topic=%s&lang=%s", key, uuid, topic, lang)
	rsp, err := http.NewRequest("POST", url, payload)
	rsp.Header.Set("Content-Type", "audio/x-wav")
	rq, err := http.DefaultClient.Do(rsp)
	if err != nil {
		fmt.Println(err)
	}
	defer rq.Body.Close()
	body, err = ioutil.ReadAll(rq.Body)
	return body
}

func Tokenize(key string, layers string, text string) (body []byte) {
	data := url.Values{}
	url := fmt.Sprintf("https://vins-markup.voicetech.yandex.net/markup/0.x/?key=%s&layers=%s&text=%s", key, layers, text)
	rsp, err := http.NewRequest("GET", url, bytes.NewBufferString(data.Encode()))
	rq, err := http.DefaultClient.Do(rsp)
	if err != nil {
		fmt.Println(err)
	}
	defer rq.Body.Close()
	body, err = ioutil.ReadAll(rq.Body)
	return body
}

func generateRandomSelection(min, max, c int) []string {
	var r []string
	if min > max {
		return r
	}
	c = minVal(int(max - min + 1), int(c))
	for len(r) < c {
		if max - len(r) == 0 {
			break
		} else {

		}
		v := random(min, max - len(r))
		for res, _ := range r {
			if res <= v {

			} else {
				break
			}
		}
		cVal := fmt.Sprintf("%x", v)
		r = append(r, string(cVal))
	}
	r = arrayShuffle(r)
	return r
}

func arrayShuffle(a []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j:= rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func minVal(x, y int) int {
	min := int(math.Min(float64(x), float64(y)))
	return min
}
