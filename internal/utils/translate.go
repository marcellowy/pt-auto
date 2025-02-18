package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/marcellowy/go-common/gogf/vconfig"
	"github.com/marcellowy/go-common/gogf/vlog"
	"github.com/marcellowy/go-common/tools"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type TranslateResult struct {
	Trans              string  `json:"trans"`
	SourceLanguageCode string  `json:"source_language_code"`
	SourceLanguage     string  `json:"source_language"`
	TrustLevel         float64 `json:"trust_level"`
}

type TranslateConfig struct {
	Key  string `json:"key"`
	Host string `json:"host"`
}

func TranslateTextOld(ctx context.Context, config *TranslateConfig, text string) (targetText string, err error) {

	postUrl := "https://google-translate113.p.rapidapi.com/api/v1/translator/text"

	payload := strings.NewReader("{\"from\":\"auto\",\"to\":\"en\",\"text\":\"" + text + "\"}")

	var req *http.Request
	if req, err = http.NewRequest("POST", postUrl, payload); err != nil {
		vlog.Error(ctx, err)
		return "", err
	}

	req.Header.Add("x-rapidapi-key", config.Key)
	req.Header.Add("x-rapidapi-host", config.Host)
	req.Header.Add("Content-Type", "application/json")

	var res *http.Response
	http.DefaultClient.Timeout = time.Second * 10

	proxyEnabled := vconfig.Get("tools.proxy.enabled", false).Bool()
	proxyAddr := vconfig.Get("tools.proxy.addr", "").String()

	if proxyEnabled {
		var proxy *url.URL
		if proxy, err = url.Parse(proxyAddr); err != nil {
			vlog.Error(ctx, err, proxyAddr)
			return
		}
		http.DefaultClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}

	if res, err = http.DefaultClient.Do(req); err != nil {
		vlog.Error(ctx, err)
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", res.Status)
		vlog.Error(ctx, err)
		return "", err
	}

	defer tools.Close(res.Body)
	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		vlog.Error(ctx, err)
		return
	}

	var tr = TranslateResult{}
	if err = json.Unmarshal(body, &tr); err != nil {
		vlog.Info(ctx, "response body:", body)
		vlog.Error(ctx, err)
		return "", err
	}

	vlog.Info(ctx, "translate text:", tr.Trans)

	targetText = tr.Trans

	return
}

func TranslateText(ctx context.Context, config *TranslateConfig, text string, proxyEnabled bool, proxyAddr string) (targetText string, err error) {

	postUrl := "https://google-translate113.p.rapidapi.com/api/v1/translator/text"

	payload := "{\"from\":\"auto\",\"to\":\"en\",\"text\":\"" + text + "\"}"

	client := g.Client()
	client.SetTimeout(time.Second * 10)
	client.SetHeader("x-rapidapi-key", config.Key)
	client.SetHeader("x-rapidapi-host", config.Host)
	client.SetHeader("Content-Type", "application/json")

	if err = SetProxy(ctx, client, proxyEnabled, proxyAddr); err != nil {
		vlog.Error(ctx, err)
		return
	}

	var response *gclient.Response
	if response, err = client.Post(ctx, postUrl, payload); err != nil {
		vlog.Error(ctx, err)
		return
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", response.Status)
		vlog.Error(ctx, err)
		return "", err
	}

	defer tools.Close(response.Body)
	var body []byte
	body, err = io.ReadAll(response.Body)
	if err != nil {
		vlog.Error(ctx, err)
		return
	}

	var tr = TranslateResult{}
	if err = json.Unmarshal(body, &tr); err != nil {
		vlog.Info(ctx, "response body:", body)
		vlog.Error(ctx, err)
		return "", err
	}

	vlog.Info(ctx, "translate text:", tr.Trans, "trust level:", tr.TrustLevel)

	targetText = tr.Trans

	return
}
