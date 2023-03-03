package agtapi

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/vela-ssoc/backend-common/pubody"
	"github.com/vela-ssoc/vela-minion/logic/inout"
	"github.com/vela-ssoc/vela-minion/logic/regist"
	"github.com/xgfone/ship/v5"
)

func FM() regist.Router {
	return new(fmCtrl)
}

type fmCtrl struct{}

func (fc *fmCtrl) Route(arr, _ *ship.RouteGroupBuilder) {
	arr.Route("/fm").GET(fc.Browser)
}

// Browser 文件目录浏览
func (fc *fmCtrl) Browser(c *ship.Context) error {
	var r inout.FM
	if err := c.BindQuery(&r); err != nil {
		return err
	}

	name := r.Path
	if name == "" {
		name = "/"
	}
	abs, err := filepath.Abs(name)
	if err != nil {
		return err
	}
	open, err := os.Open(abs)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer open.Close()

	if stat, _ := open.Stat(); stat != nil && !stat.IsDir() {
		fnm, mtime := stat.Name(), stat.ModTime()
		res, req := c.ResponseWriter(), c.Request()
		http.ServeContent(res, req, fnm, mtime, open)
		return nil
	}

	infos, err := open.Readdir(r.Limit)
	if err != nil {
		return err
	}
	sep := string(os.PathSeparator)
	ret := &pubody.Folder{
		Abs:       abs,
		Separator: sep,
		Items:     make(pubody.FileItems, 0, 100),
	}

	var me error
	match := r.Match
	for _, info := range infos {
		nm := info.Name()
		if me == nil && match != "" {
			var matched bool
			if matched, me = filepath.Match(match, nm); me == nil && !matched {
				continue
			}
		}
		dir := info.IsDir()
		item := &pubody.FileItem{
			Path:  filepath.Join(abs, nm),
			Name:  nm,
			Size:  info.Size(),
			Mtime: info.ModTime(),
			Dir:   dir,
			Mode:  info.Mode().String(),
		}
		if !dir {
			item.Ext = filepath.Ext(nm)
		}
		ret.Items = append(ret.Items, item)
	}

	return c.JSON(http.StatusOK, ret)
}