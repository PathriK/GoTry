package hello

import (
    "fmt"
	"html/template"
		"net/http"
		"encoding/json"		
		"log"
	"bytes"	
	"appengine"
	"appengine/datastore"
	"appengine/user"	
)

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

type MCQ struct {
	ID string
	Question  string
	Options []string
	Answer string
}

type tab struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func init() {
	http.HandleFunc("/api/unauth/home", tabHandler)
    http.HandleFunc("/mcq", mainhandler)
	http.HandleFunc("/mcq/submit", submithandler)
	http.HandleFunc("/mcq/add", addhandler)
}

func tabHandler(w http.ResponseWriter, r *http.Request){
	tabs := []tab {
		tab{Name: "HOME", URL: "/"},
		tab{Name: "Contact Us", URL: "/contact-us"},
	}
	home := struct {
		Title string `json:"title"`
		Tabs []tab `json:"tabs"`
	}	{
		Title : "MyPage",
		Tabs : tabs,
	}
	js, err := json.Marshal(home)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
	log.Printf("JSON:%s",home)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

// mcqKey returns the key used for all guestbook entries.
func mcqKey(c appengine.Context) *datastore.Key {
        return datastore.NewKey(c, "MCQS", "default_mcq", 0, nil)
}

func mainhandler(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	resp := ""
	if u == nil {
			url, _ := user.LoginURL(ctx, "/")
			resp = resp + fmt.Sprintf(`<a href="%s">Sign in or register</a>`, url)
			return
	}
	url, _ := user.LogoutURL(ctx, "/")
	resp = resp + fmt.Sprintf(`Welcome, %s! <a href="%s">sign out</a><br>`, u, url)
	if u.Admin {
		resp = resp + getAdminContents(ctx)
    } else {
		resp = resp + getUserContents(ctx)
	}
	
	if err := mcqpage.Execute(w, template.HTML(resp)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
			
	q := datastore.NewQuery("MCQ").Ancestor(mcqKey(c)).Filter("ID =", id).Limit(1)
	mcqs := make([]MCQ, 0, 1)
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
	mcq := MCQ{
		ID: r.FormValue("id`"),
		Question: r.FormValue("question"),
		Options: []string{r.FormValue("mcq1"),r.FormValue("mcq2"),r.FormValue("mcq3"),r.FormValue("mcq4")},
		Answer: r.FormValue("answer"),
	}
	key := datastore.NewIncompleteKey(c, "MCQ", mcqKey(c))
	_, err = datastore.Put(c, key, &mcq)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	http.Redirect(w, r, "/mcq", http.StatusFound)	
}

func getAdminContents(c appengine.Context) string{
	var resp bytes.Buffer
	q := datastore.NewQuery("MCQ").Ancestor(mcqKey(c))
	cnt,err := q.Count(c);
	if err != nil {
		return err.Error()
	}
	
	mcqs := make([]MCQ, 0, cnt)
	if _, err := q.GetAll(c, &mcqs); err != nil {
		return err.Error()
	}
	if err := mcqadmin.Execute(&resp, mcqs); err != nil {
		return err.Error()
	}
	
	return resp.String()
}

func getUserContents(c appengine.Context) string{
	var resp bytes.Buffer
	q := datastore.NewQuery("MCQ").Ancestor(mcqKey(c)).Limit(10)
	mcqs := make([]MCQ, 0, 10)
	if _, err := q.GetAll(c, &mcqs); err != nil {
		return err.Error()
	}
	if err := mcquser.Execute(&resp, mcqs[:1]); err != nil {
		return err.Error()
	}
	
	return resp.String()
}
