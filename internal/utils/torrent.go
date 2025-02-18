package utils

import (
	"context"
	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/marcellowy/go-common/gogf/vlog"
	"github.com/marcellowy/go-common/tools"
	"os"
)

// CreateTorrent create torrent file
func CreateTorrent(ctx context.Context, root, file string) (err error) {
	var (
		mi      = metainfo.MetaInfo{}
		private = true
		f       *os.File
	)
	info := metainfo.Info{
		Private: &private,
	}
	if err = info.BuildFromFilePath(root); err != nil {
		vlog.Error(ctx, err)
		return
	}
	if mi.InfoBytes, err = bencode.Marshal(info); err != nil {
		vlog.Error(ctx, err)
		return
	}
	if f, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm); err != nil {
		vlog.Error(ctx, err)
		return
	}
	defer tools.Close(f)
	if err = mi.Write(f); err != nil {
		vlog.Error(ctx, err)
		return
	}
	return
}
