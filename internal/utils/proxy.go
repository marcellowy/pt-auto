package utils

import (
	"context"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/marcellowy/go-common/gogf/vlog"
)

// SetProxy set proxy
func SetProxy(ctx context.Context, client *gclient.Client, proxyEnabled bool, proxyAddr string) (err error) {

	if !proxyEnabled {
		vlog.Info(ctx, "skip proxy")
		return
	}
	client.SetProxy(proxyAddr)
	//var proxy *url.URL
	//if proxy, err = url.Parse(proxyAddr); err != nil {
	//	vlog.Error(ctx, err, proxyAddr)
	//	return
	//}
	//client.Transport = &http.Transport{
	//	Proxy: http.ProxyURL(proxy),
	//}
	return nil
}
