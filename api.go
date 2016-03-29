package wordnik

import(
  "net/http"
  //"fmt"
  "io/ioutil"
  "errors"
  "strconv"
)

var RootURL = "http://api.wordnik.com/v4"

const (
  GET    = "GET"
  POST   = "POST"
  PUT    = "PUT"
  DELETE = "DELETE"
)

type Service struct {
  Key string

  WordsService *WordsService
  WordService *WordService
}

type Data struct{
  Key string
  Value string
}

func ( s *Service )GetUrlContent( method string, url string, list []Data  ) (string, error){

  client := &http.Client{
    //CheckRedirect: redirectPolicyFunc,
  }
  req, err := http.NewRequest( method ,  url , nil)

  q := req.URL.Query()

  for _, item := range list{
    q.Add( item.Key ,  item.Value )
  }
  
  req.URL.RawQuery = q.Encode()

  //fmt.Println(req.URL.String())

  
  // ...
  req.Header.Add("api_key", s.Key )
  
  response, err := client.Do(req)
  

  //contentLength := response.ContentLength
  //contentType := response.Header.Get("Content-Type") 
  //fmt.Printf("Content length is %d\n",  contentLength)

  if err != nil {
    //fmt.Printf("error here \n")
    return  "", err
  }
 
  defer response.Body.Close()

  byteBody, _ := ioutil.ReadAll(response.Body)
  
  body := string(byteBody)

  //fmt.Printf("body %s \n", body)

  statusCode := response.StatusCode
  //fmt.Printf("statusCode %d\n", statusCode)

  if statusCode != 200{
    
    return body, errors.New( strconv.Itoa(statusCode) ) 
  }

  return body, nil
  
}

func New( apiKey string ) ( *Service , error ){
  s := Service{ Key: apiKey }
  s.WordsService = NewWordsService( &s )
  s.WordService =  NewWordService( &s )
  return &s, nil
}
