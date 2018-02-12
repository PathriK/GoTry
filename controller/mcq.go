package controller

import (
	"net/http"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

// MCQ is the base struct for MCQs
type MCQ struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

// MCQKey returns the key used for all guestbook entries.
func MCQKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "MCQS", "default_mcq", 0, nil)
}

// McqHandler handles MCQ User and Admin Home page display
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
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	ctx.Debugf(err.Error())
	// }
	// fullStruct := iterateJSON(dir)
	// ctx.Debugf(fullStruct)
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
