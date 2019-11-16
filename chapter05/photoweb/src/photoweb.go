package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	ListDir      = 0x0001
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

func init() {
	sdir, _ := os.Getwd()
	fmt.Println(sdir)
	// 加载模板文件列表
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		// 文件扩展名
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		// Must确保模板不能解析成功时触发错误处理流程
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	// 将locals渲染到tmpl模板上
	err := templates[tmpl].Execute(w, locals)
	// 检查渲染模板是否出错，错误抛出panic
	check(err)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 这段被渲染成纯文本了
		/*
			_, _ = io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
				" enctype=\"multipart/form-data\">"+
				"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
				"<input type=\"submit\" value=\"Upload\" />"+
				"</form>")
		*/
		//w.Header().Set("Content-Type", "text/html")
		renderHtml(w, "upload", nil)
	}
	if r.Method == "POST" {
		// multipart.File *multipart.FileHeader err
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		// 最后关闭临时文件
		defer f.Close()
		t, err := ioutil.TempFile(UPLOAD_DIR, filename)
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	// 将该路径下的文件从磁盘中读取并作为服务端返回信息输出给客户端
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	// 读取文件夹下的文件列表
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	check(err)
	locals := make(map[string]interface{})
	var images []string
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["image"] = images
	renderHtml(w, "list", locals)
}

// 处理动态请求，利用匿名函数闭包终止错误处理，抛出500状态码，并打印日志
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// 或者输出自定义的50x错误页面
				//w.WriteHeader(http.StatusInternalServerError)
				//renderHtml(w, "error", e)
				// logging
				log.Println("WARN: panic in %v. - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

// 静态资源
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flag int) {
	// mux是http请求复用器
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flag & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		// 将该路径下的文件从磁盘中读取并作为服务端返回信息输出给客户端
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	// 分发请求
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	fmt.Println("Please visit the address: http://localhost:8080/")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListernAndServe: ", err.Error())
	}
}
