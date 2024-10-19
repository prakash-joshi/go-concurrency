package main

import (
	"fmt"
	"html/template"
	"net/http"
	"subscription-service/data"
)

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = app.Sessions.RenewToken(r.Context())

	// parse from post
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// get email and password from form post
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		app.Sessions.Put(r.Context(), "error", "invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// check password
	validPassword, err := user.PasswordMatches(password)

	if err != nil {
		app.Sessions.Put(r.Context(), "error", "invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword {
		msg := Message{
			To:      "testmail@test.com",
			Subject: "Failed login Attempt",
			Data:    "Invalid Password while trying to login.",
		}
		app.sendEmail(msg)

		app.Sessions.Put(r.Context(), "error", "invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// okay to log user in
	app.Sessions.Put(r.Context(), "userID", user.ID)
	app.Sessions.Put(r.Context(), "user", user)
	app.Sessions.Put(r.Context(), "flash", fmt.Sprintln("Successful logged in ", user.FirstName+" "+user.LastName))

	// redirect the user
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {

	app.Sessions.Destroy(r.Context())
	app.Sessions.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// create a user
	u := data.User{
		Email:     r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Password:  r.Form.Get("password"),
		Active:    0,
		IsAdmin:   0,
	}
	_, err = u.Insert(u)
	if err != nil {
		app.ErrorLog.Println(err)
		app.Sessions.Put(r.Context(), "error", "Unable to create user.")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// send an activation email
	url := fmt.Sprintf("http://localhost:8080/activate?email=%s", u.Email)
	signedURL := GenerateTokenFromString(url)
	app.InfoLog.Println(signedURL)

	msg := Message{
		To:       u.Email,
		Subject:  "Activate your account",
		Template: "confirmation-email",
		Data:     template.HTML(signedURL),
	}

	app.sendEmail(msg)
	app.Sessions.Put(r.Context(), "flash", "Confirmation email sent")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate the url
	url := r.RequestURI
	app.InfoLog.Println("url :", url)
	testURL := fmt.Sprintf("http://localhost:8080%s", url)
	ok := VerifyToken(testURL)

	if !ok {
		app.Sessions.Put(r.Context(), "error", "Invalid token.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// activate the account
	app.InfoLog.Println("email id : ", r.URL.Query().Get("email"))
	u, err := app.Models.User.GetByEmail(r.URL.Query().Get("email"))
	if err != nil {
		app.Sessions.Put(r.Context(), "error", "No user found.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u.Active = 1
	err = u.Update()
	if err != nil {
		app.Sessions.Put(r.Context(), "error", "Unable to update user.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.Sessions.Put(r.Context(), "flash", "Account activated. You can log in now.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
