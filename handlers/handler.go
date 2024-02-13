package handlers

import (
	utils "forum/internal"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//Inscriptions := false
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	var username, email, mdp, mdp1, mdp2 string
	username = r.FormValue("username")
	email = r.FormValue("email")
	mdp = r.FormValue("mdp")
	mdp1 = r.FormValue("mdp1")
	mdp2 = r.FormValue("mdp2")
	if email != "" && mdp1 == mdp2 {
		if !utils.Insertion(username, email, mdp1) {
			http.Redirect(w, r, "/home", 200)
			return
		}
		http.Redirect(w, r, "/", 200)
	} else if email == "" && username != "" {
		if utils.Connexion(username, email, mdp) {
			http.Redirect(w, r, "/home", 200)
			return
		}
	}

	RenderTemplate(w, "login", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home", nil)
}

func RenderTemplate(w http.ResponseWriter, page string, value interface{}) {
	filecontent, rr := template.ParseFiles("./template/" + page + ".html")
	if rr != nil {
		http.Error(w, "500 InternalServerError", http.StatusInternalServerError)
		return
	}
	filecontent.Execute(w, value)
}
