package mojang

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

const MOJANG_PROFILE_ENDPOINT = "https://sessionserver.mojang.com/session/minecraft/profile/"

var cache = make(map[string]string)

func FetchPlayerName(uuid string) (string, error) {
	if name, ok := cache[uuid]; ok {
		return name, nil
	}

	u, err := url.Parse(MOJANG_PROFILE_ENDPOINT)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, uuid)

	res, err := http.Get(MOJANG_PROFILE_ENDPOINT + uuid)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	profile := new(struct {
		Name string
	})

	if err := json.Unmarshal(data, profile); err != nil {
		return "", err
	}

	cache[uuid] = profile.Name
	time.Sleep(500 * time.Millisecond)

	return profile.Name, nil
}
