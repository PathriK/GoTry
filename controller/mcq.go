package controller

import (
	"net/http"

	"appengine"
	"appengine/user"
)

func McqHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	// resp := ""
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		// resp = resp + fmt.Sprintf(`<a href="%s">Sign in or register</a>`, url)
		http.Redirect(w, r, url, http.StatusFound)
	}
	// url, _ := user.LogoutURL(ctx, "/")
	// resp = resp + fmt.Sprintf(`Welcome, %s! <a href="%s">sign out</a><br>`, u, url)
	if u.Admin {
		// resp = resp + getAdminContents(ctx)
		http.ServeFile(w, r, "frontend/dist/mcq/admin/")
	} else {
		// resp = resp + getUserContents(ctx)
		http.ServeFile(w, r, "frontend/dist/mcq/user/")
	}

	// if err := mcqpage.Execute(w, template.HTML(resp)); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
