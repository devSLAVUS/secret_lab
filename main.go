package main

import ("fmt"
        "net/http"
			  "html/template")


type User struct {
  Name string
	Age uint16
	Money int16
	Avg_grade, Happy float64
	Hobbies []string
}
func (u User) getAll() string {
return fmt.Sprintf("User: %s %d %d", u.Name, u.Age, u.Money)
}
func (u *User) setName(newName string){
	u.Name = newName
}
func slavus_page(w http.ResponseWriter, r *http.Request) {
bob := User{Name: "bob", Age: 15, Money: 100, Avg_grade: 5, Happy: 0.1, Hobbies: []string{"dota", "cs", "star"} }
//  bob.setName("jon")
tmpl, _ := template.ParseFiles("templates/slavus_page.html")
tmpl.Execute(w, bob)
}
func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about page")
}
func mainserver(){
	http.HandleFunc("/", slavus_page)
	http.HandleFunc("/about/", about_page)
	http.ListenAndServe(":80", nil)
}
func main() {
		mainserver()
}
