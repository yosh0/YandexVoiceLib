package yandexVoiceLib

import (
	"os"
	"io"
	"fmt"
	"math"
	"time"
	"net/url"
	"strings"
	"net/http"
	"io/ioutil"
	"math/rand"
)

const (
	RECOGNIZE_SCHEME	= "https"
	RECOGNIZE_HOST		= "asr.yandex.net"
	RECOGNIZE_PATH		= "/asr_xml/"
)

func Recognize(file, topic, key, lang string) (body []byte) {
	tUuid := generateRandomSelection(0, 30, 64)
	uuid := strings.Join(tUuid, "")
	uuid = uuid[0:32]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	payload := io.MultiReader(f)

	var Url *url.URL
	Url, err = url.Parse(fmt.Sprintf("%s://%s", RECOGNIZE_SCHEME, RECOGNIZE_HOST))
	if err != nil {
		fmt.Println(err)
	}
	Url.Scheme = RECOGNIZE_SCHEME
	Url.Host = RECOGNIZE_HOST
	Url.Path = RECOGNIZE_PATH
	params := url.Values{}
	params.Add("key", key)
	params.Add("uuid", uuid)
	params.Add("topic", topic)
	params.Add("lang", lang)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest(http.MethodPost, Url.String(), payload)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "audio/x-wav")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return body
}

const (
	TOKENIZE_SCHEME	= "https"
	TOKENIZE_HOST	= "vins-markup.voicetech.yandex.net"
	TOKENIZE_PATH	= "/markup/0.x/"
)

func Tokenize(key, layers, text string) (body []byte) {
	var Url *url.URL
	Url, err := url.Parse(fmt.Sprintf("%s://%s", TOKENIZE_SCHEME, TOKENIZE_HOST))
	if err != nil {
		fmt.Println(err)
	}
	Url.Scheme = TOKENIZE_SCHEME
	Url.Host = TOKENIZE_HOST
	Url.Path = TOKENIZE_PATH
	params := url.Values{}
	params.Add("key", key)
	params.Add("layers", layers)
	params.Add("text", text)
	Url.RawQuery = params.Encode()
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest(http.MethodGet, Url.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
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
