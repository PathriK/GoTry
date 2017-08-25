package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"appengine"
	"appengine/user"
)

func McqHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ctx.Debugf("Inside MCQ Handler")
	u := user.Current(ctx)
	// resp := ""
	if u == nil {
		ctx.Debugf("MCQ Handler: user nil")
		url, _ := user.LoginURL(ctx, "/")
		// resp = resp + fmt.Sprintf(`<a href="%s">Sign in or register</a>`, url)
		http.Redirect(w, r, url, http.StatusFound)
	}
	// url, _ := user.LogoutURL(ctx, "/")
	// resp = resp + fmt.Sprintf(`Welcome, %s! <a href="%s">sign out</a><br>`, u, url)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		ctx.Debugf(err.Error())
	}
	fullStruct := iterateJSON(dir)
	ctx.Debugf(fullStruct)
	//	fmt.Fprintf(w, fullStruct)
	if u.Admin {
		ctx.Debugf("MCQ Handler: Admin")
		// resp = resp + getAdminContents(ctx)
		http.ServeFile(w, r, "./frontend/dist/mcq/admin/index.html")
	} else {
		// resp = resp + getUserContents(ctx)
		ctx.Debugf("MCQ Handler: User")
		http.ServeFile(w, r, "./frontend/dist/mcq/user/index.html")
	}

	// if err := mcqpage.Execute(w, template.HTML(resp)); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

type File struct {
	ModifiedTime time.Time `json:"ModifiedTime"`
	IsLink       bool      `json:"IsLink"`
	IsDir        bool      `json:"IsDir"`
	LinksTo      string    `json:"LinksTo"`
	Size         int64     `json:"Size"`
	Name         string    `json:"Name"`
	Path         string    `json:"Path"`
	Children     []*File   `json:"Children"`
}

func iterateJSON(path string) string {
	rootOSFile, _ := os.Stat(path)
	rootFile := toFile(rootOSFile, path) //start with root file
	stack := []*File{rootFile}

	for len(stack) > 0 { //until stack is empty,
		file := stack[len(stack)-1] //pop entry from stack
		stack = stack[:len(stack)-1]
		children, _ := ioutil.ReadDir(file.Path) //get the children of entry
		for _, chld := range children {          //for each child
			child := toFile(chld, filepath.Join(file.Path, chld.Name())) //turn it into a File object
			file.Children = append(file.Children, child)                 //append it to the children of the current file popped
			stack = append(stack, child)                                 //append the child to the stack, so the same process can be run again
		}
	}

	output, _ := json.MarshalIndent(rootFile, "", "     ")
	return string(output)
}

func toFile(file os.FileInfo, path string) *File {
	JSONFile := File{ModifiedTime: file.ModTime(),
		IsDir:    file.IsDir(),
		Size:     file.Size(),
		Name:     file.Name(),
		Path:     path,
		Children: []*File{},
	}
	if file.Mode()&os.ModeSymlink == os.ModeSymlink {
		JSONFile.IsLink = true
		JSONFile.LinksTo, _ = filepath.EvalSymlinks(filepath.Join(path, file.Name()))
	} // Else case is the zero values of the fields
	return &JSONFile
}
