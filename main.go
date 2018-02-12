package gotry

import (
	"bytes"
	"controller"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

const homeType = "Home"
const homeParentType = "HomeParent"
const homeParentName = "home_parent"

var mcqpage = template.Must(template.New("mcqpage").Parse(`
<html>
  <head>
    <title>MCQ</title>
  </head>
  <body>
    {{.}}
  </body>
</html>
`))

var mcqadmin = template.Must(template.New("mcqadmin").Parse(`
{{range .}}<br>
  {{.Question}}
  {{range .Options}}
    {{.}}<br>
  {{end}}
  Answer::{{.Answer}}<br>
{{end}}
<form action="/mcq/add" method="post">
	Question: ​<textarea name="question" rows="10" cols="70"></textarea><br>
	Option1: <input type="text" name="mcq1" /><br>
	Option2: <input type="text" name="mcq2" /><br>
	Option3: <input type="text" name="mcq3" /><br>
	Option4: <input type="text" name="mcq4" /><br>
	<input type="hidden" name="id" value="{{len .}}" /><br>
	Answer: <input type="text" name="answer" /> (1,2,3 or 4)<br>
	<input type="submit" value="Submit" />
</form>
`))

var mcquser = template.Must(template.New("mcquser").Parse(`
<form action="/mcq/submit" method="post">
{{range .}}
  {{.Question}}<br>  
  {{range .Options}}
	<input type="radio" name="mcq" value="{{.}}">{{.}}<br>
  {{end}}
  <input type="hidden" name="id" value="{{.ID}}">
  <input type="submit" value="Submit">
{{end}}
</form>
`))

type tab struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type home struct {
	Title string   `json:"title"`
	Tabs  []string `json:"tabs"`
}

func init() {
	// ctx := appengine.NewContext(r)
	// if correctobj, ok := controller.(interface {
	// 	McqHandler(http.ResponseWriter, *http.Request)
	// }); ok {
	// 	ctx.Debugf("init: MCQ Handler exist")
	// } else {
	// 	ctx.Debugf("init: MCQ Handler NOT exist")
	// }
	http.HandleFunc("/unauth/api/home/add", homeaddhandler)
	http.HandleFunc("/unauth/api/home", homeHandler)
	// http.HandleFunc("/mcq/", mainhandler)
	http.HandleFunc("/api/admin/mcq/list", controller.ListHandler)
	http.HandleFunc("/mcq/", controller.McqHandler)
	// http.Handle("/mcq/", http.StripPrefix("/mcq/", http.FileServer(http.Dir("./testing"))))
	http.HandleFunc("/mcq/submit", submithandler)
	http.HandleFunc("/mcq/add", addhandler)
}

func homeaddhandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	parentKey := datastore.NewKey(c, homeParentType, homeParentName, 0, nil)
	key := datastore.NewIncompleteKey(c, homeType, parentKey)
	homeJSON := home{
		Title: "My Page",
		Tabs:  []string{"Home", "Gallery", "Courses", "Contact Us", "Demo Test"},
	}
	_, err := datastore.Put(c, key, &homeJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Done!")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// tabs := []tab {
	// 	tab{Name: "HOME", URL: "/"},
	// 	tab{Name: "Contact Us", URL: "/contact-us"},
	// }
	// home := struct {
	// 	Title string `json:"title"`
	// 	Tabs []tab `json:"tabs"`
	// }	{
	// 	Title : "MyPage",
	// 	Tabs : tabs,
	// }

	c := appengine.NewContext(r)
	parentKey := datastore.NewKey(c, homeParentType, homeParentName, 0, nil)
	q := datastore.NewQuery(homeType).Ancestor(parentKey).Limit(1)
	var homes []home
	if _, err := q.GetAll(c, &homes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(homes[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//ctx.Printf("JSON:%s", homes[0])
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func mainhandler(w http.ResponseWriter, r *http.Request) {
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

func submithandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := r.PostFormValue("id")
	ansExp := r.PostFormValue("mcq")

	q := datastore.NewQuery("MCQ").Ancestor(controller.MCQKey(c)).Filter("ID =", id).Limit(1)
	mcqs := make([]controller.MCQ, 0, 1)
	if _, err := q.GetAll(c, &mcqs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ans := mcqs[0].Answer
	status := "Wrong Answer!"
	if ansExp == ans {
		status = "Correct Answer!"
	}

	if err := mcqpage.Execute(w, template.HTML(status)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addhandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	mcq := controller.MCQ{
		ID:       r.FormValue("id`"),
		Question: r.FormValue("question"),
		Options:  []string{r.FormValue("mcq1"), r.FormValue("mcq2"), r.FormValue("mcq3"), r.FormValue("mcq4")},
		Answer:   r.FormValue("answer"),
	}
	key := datastore.NewIncompleteKey(c, "MCQ", controller.MCQKey(c))
	_, err = datastore.Put(c, key, &mcq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/mcq", http.StatusFound)
}

func getAdminContents(c appengine.Context) string {
	var resp bytes.Buffer
	q := datastore.NewQuery("MCQ").Ancestor(controller.MCQKey(c))
	cnt, err := q.Count(c)
	if err != nil {
		return err.Error()
	}

	mcqs := make([]controller.MCQ, 0, cnt)
	if _, err := q.GetAll(c, &mcqs); err != nil {
		return err.Error()
	}
	if err := mcqadmin.Execute(&resp, mcqs); err != nil {
		return err.Error()
	}

	return resp.String()
}

func getUserContents(c appengine.Context) string {
	var resp bytes.Buffer
	q := datastore.NewQuery("MCQ").Ancestor(controller.MCQKey(c)).Limit(10)
	mcqs := make([]controller.MCQ, 0, 10)
	if _, err := q.GetAll(c, &mcqs); err != nil {
		return err.Error()
	}
	if err := mcquser.Execute(&resp, mcqs[:1]); err != nil {
		return err.Error()
	}

	return resp.String()
}
