package main
import (
  "fmt"
  "net/http"
  "html/template"
  "github.com/urfave/negroni"
  "encoding/xml"
  "encoding/json"
  "io/ioutil"
)

type DefWord struct{
  Name string
  Definition []string `xml:"entry>def>dt"`
}
type Page struct{
  Name string
}
func getDefinition(query string) (DefWord, error){
  var w DefWord

  fmt.Println("The query:" + query)
  body, err := dictionaryAPI("http://www.dictionaryapi.com/api/v1/references/collegiate/xml/"+query+"?key="+getKey())
  if err != nil{
    fmt.Println("error returning from dictionaryAPI")
    return DefWord{}, err
  }
  fmt.Println(body)
  err = xml.Unmarshal(body, &w)
  w.Name = query
  fmt.Println(w)
  return w, err
}
func dictionaryAPI(url string) ([]byte, error){
  fmt.Println("got to Dictionary")

  var err error
  resp, err := http.Get(url)
  if err != nil{
    fmt.Println("failed at Get")
    return []byte{},err
  }
  fmt.Println(resp.StatusCode)
  defer resp.Body.Close()
  fmt.Println("got to end of dictionaryAPI")

  return ioutil.ReadAll(resp.Body)
}
func main() {
    mux := http.NewServeMux()
    templates := template.Must(template.ParseFiles("static/templates/index.html"))
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
      p := Page{Name: "test"}
      if err := templates.ExecuteTemplate(w,"index.html",p); err != nil{
        http.Error(w,err.Error(),http.StatusInternalServerError)
      }
    })
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    mux.HandleFunc("/search",func(w http.ResponseWriter, r *http.Request){
      fmt.Println("got reques2t")

      var result DefWord
      var err error
        if result, err = getDefinition(r.FormValue("text")); err != nil{
        fmt.Println("failed get def")
        http.Error(w, err.Error(), http.StatusInternalServerError)
      }
      encoder := json.NewEncoder(w)
      if err := encoder.Encode(result); err != nil{
        fmt.Println("failed encode")
        http.Error(w,err.Error(), http.StatusInternalServerError)
      }
    })
    n := negroni.Classic()

    n.UseHandler(mux)
    n.Run(":8080")

}
