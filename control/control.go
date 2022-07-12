package control

import (
	"io/ioutil"
	"net/http"
)

// 主页面
func IndexView(w http.ResponseWriter, r *http.Request) {
	w.Write(LoadHtml("index.html"))
}

// 上传页面
func UploadView(w http.ResponseWriter, r *http.Request) {
	w.Write(LoadHtml("upload.html"))
}

// 图片上传
func ApiUpload(w http.ResponseWriter, r *http.Request) {
	//f, h, err := r.FormFile("file")
	//if err != nil {
	//	io.WriteString(w, "上传错误："+err.Error())
	//	return
	//}
	//t := h.Header.Get("Content-Type")
	//if !strings.Contains(t, "image") {
	//	io.WriteString(w, "文件类型错误")
	//}
	//os.Mkdir("/static", os.ModePerm)
	//timeName := time.Now()

}
func LoadHtml(name string) []byte {
	buf, err := ioutil.ReadFile("./views/" + name)
	if err != nil {
		return []byte("")
	}
	return buf
}
