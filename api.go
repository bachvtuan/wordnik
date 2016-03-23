package wordnik

import(
  "net/http"
  //"fmt"
  "io/ioutil"
  "encoding/json"
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
  return &s, nil
}

type WordsService struct{
  Service *Service
  WordOfTheDayService *WordOfTheDayService
  SearchService *SearchService
  RandomWordsService *RandomWordsService
  RandomWordService *RandomWordService
}

func NewWordsService( s *Service  ) *WordsService{
  //with empty date
  wordOfTheDayService := &WordOfTheDayService{ s, "" }
  searchService := SearchService{ Service : s }
  //init default value
  searchService.CaseSensitive = "true" 
  searchService.MinCorpusCount = 5
  searchService.MaxCorpusCount = -1
  searchService.MinDictionaryCount = 1
  searchService.MaxDictionaryCount = -1
  searchService.MinLength = 1
  searchService.MaxLength = -1
  searchService.Skip = 0
  searchService.Limit = 10

  //int randomWords Service
  randomWordsService := RandomWordsService{ Service:s }

  //init default value
  randomWordsService.HasDictionaryDef = "false" 
  randomWordsService.MinCorpusCount = 0
  randomWordsService.MaxCorpusCount = -1
  randomWordsService.MinDictionaryCount = 1
  randomWordsService.MaxDictionaryCount = -1
  randomWordsService.MinLength = 5
  randomWordsService.MaxLength = -1
  randomWordsService.Limit = 10


  randomWordService := RandomWordService{ Service:s }

  //init default value
  randomWordService.HasDictionaryDef = "false" 
  randomWordService.MinCorpusCount = 0
  randomWordService.MaxCorpusCount = -1
  randomWordService.MinDictionaryCount = 1
  randomWordService.MaxDictionaryCount = -1
  randomWordService.MinLength = 5
  randomWordService.MaxLength = -1
  

  return &WordsService{ s, wordOfTheDayService, &searchService, &randomWordsService, &randomWordService }
}

type WordOfTheDayService struct{
  Service *Service
  Date  string `json:"date" bson:"date"`
}

func ( wodService *WordOfTheDayService ) Do() ( WordOfTheDay, error ) {
  //get 
  url := RootURL + "/words.json/wordOfTheDay"
  data := []Data{}

  if wodService.Date != ""{
    data = append( data, Data{"date", wodService.Date } )
  }

  response_text, err := wodService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return WordOfTheDay{}, err
  }

  b := []byte( response_text )
  var wod WordOfTheDay
  err = json.Unmarshal(b, &wod)
  return wod, err
}


type Example struct{
  Url     string        `json:"url" bson:"url"`
  Text     string        `json:"text" bson:"text"`
  Id     int        `json:"id" bson:"id"`
  Title     string        `json:"title" bson:"title"`
}

type Definition struct{
  Text     string        `json:"text" bson:"text"` 
  Source     string        `json:"source" bson:"source"` 
  PartOfSpeech     string        `json:"partOfSpeech" bson:"partOfSpeech"` 
}

type ContentProvider struct{
  Id int `json:"id" bson:"id"`
  Name string `json:"name" bson:"name"`
}

type WordOfTheDay struct{
  Id int `json:"id" bson:"id"` 
  Word string `json:"word" bson:"word"` 
  PublishDate string `json:"publishDate" bson:"publishDate"` 
  ContentProvider ContentProvider `json:"contentProvider" bson:"contentProvider"` 
  Note string `json:"note" bson:"note"` 
  Examples []Example `json:"examples" bson:"examples"` 
  Definitions []Definition `json:"definitions" bson:"definitions"` 
}



type SearchItem struct{
  Lexicality float32 `json:"lexicality" bson:"lexicality"`
  Count int `json:"count" bson:"count"`
  Word string `json:"word" bson:"word"`
}

type SearchResult struct{
  TotalResults int `json:"totalResults" bson:"totalResults"`
  SearchResults []SearchItem `json:"searchResults" bson:"searchResults"`
}

type SearchService struct{
  Service *Service

  Query string
  CaseSensitive string 
  IncludePartOfSpeech string
  ExcludePartOfSpeech string
  MinCorpusCount int
  MaxCorpusCount int
  MinDictionaryCount int
  MaxDictionaryCount int
  MinLength int
  MaxLength int
  Skip int
  Limit int
}

func ( searchService *SearchService ) Do() ( SearchResult, error ) {
  //get 
  if searchService.Query == ""{
    return SearchResult{}, errors.New( "query is required" )
  }

  url := RootURL + "/words.json/search/" + searchService.Query
  data := []Data{}

  data = append( data, Data{"caseSensitive", searchService.CaseSensitive } )

  if searchService.IncludePartOfSpeech != ""{
    data = append( data, Data{"includePartOfSpeech", searchService.IncludePartOfSpeech } )
  }

  if searchService.ExcludePartOfSpeech != ""{
    data = append( data, Data{"excludePartOfSpeech", searchService.ExcludePartOfSpeech } )
  }

  data = append( data, Data{"minCorpusCount", strconv.Itoa(searchService.MinCorpusCount) } )
  data = append( data, Data{"maxCorpusCount", strconv.Itoa(searchService.MaxCorpusCount) } )
  data = append( data, Data{"minDictionaryCount", strconv.Itoa(searchService.MinDictionaryCount) } )
  data = append( data, Data{"maxDictionaryCount", strconv.Itoa(searchService.MaxDictionaryCount) } )
  data = append( data, Data{"minLength", strconv.Itoa(searchService.MinLength) } )
  data = append( data, Data{"maxLength", strconv.Itoa(searchService.MaxLength) } )
  data = append( data, Data{"skip", strconv.Itoa(searchService.Skip) } )
  data = append( data, Data{"limit", strconv.Itoa(searchService.Limit) } )

  //fmt.Printf("Limit is %d %s\n", searchService.Limit, strconv.Itoa(searchService.Limit))
  

  response_text, err := searchService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return SearchResult{}, err
  }

  b := []byte( response_text )
  var result SearchResult
  err = json.Unmarshal(b, &result)
  return result, err
}


type RandomWord struct{
  Id int `json:"id" bson:"id"`
  Word string `json:"word" bson:"word"`
}

type RandomWordService struct{
  Service *Service

  HasDictionaryDef string
  IncludePartOfSpeech string
  ExcludePartOfSpeech string
  MinCorpusCount int
  MaxCorpusCount int
  MinDictionaryCount int
  MaxDictionaryCount int
  MinLength int
  MaxLength int
  
}

func ( randomService *RandomWordService )Do() ( RandomWord, error ) {
  url := RootURL + "/words.json/randomWord"
  data := []Data{}

  data = append( data, Data{"hasDictionaryDef", randomService.HasDictionaryDef } )

  if randomService.IncludePartOfSpeech != ""{
    data = append( data, Data{"includePartOfSpeech", randomService.IncludePartOfSpeech } )
  }

  if randomService.ExcludePartOfSpeech != ""{
    data = append( data, Data{"excludePartOfSpeech", randomService.ExcludePartOfSpeech } )
  }

  data = append( data, Data{"minCorpusCount", strconv.Itoa(randomService.MinCorpusCount) } )
  data = append( data, Data{"maxCorpusCount", strconv.Itoa(randomService.MaxCorpusCount) } )
  data = append( data, Data{"minDictionaryCount", strconv.Itoa(randomService.MinDictionaryCount) } )
  data = append( data, Data{"maxDictionaryCount", strconv.Itoa(randomService.MaxDictionaryCount) } )
  data = append( data, Data{"minLength", strconv.Itoa(randomService.MinLength) } )
  data = append( data, Data{"maxLength", strconv.Itoa(randomService.MaxLength) } )
  

  response_text, err := randomService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return RandomWord{}, err
  }

  b := []byte( response_text )
  var result RandomWord
  err = json.Unmarshal(b, &result)
  return result, err
}

type RandomWordsService struct{
  Service *Service

  HasDictionaryDef string
  IncludePartOfSpeech string
  ExcludePartOfSpeech string
  MinCorpusCount int
  MaxCorpusCount int
  MinDictionaryCount int
  MaxDictionaryCount int
  MinLength int
  MaxLength int

  SortBy string
  SortOrder string
  Limit int
  
}

func ( randomWordsService *RandomWordsService )Do() ( []RandomWord, error ) {
  url := RootURL + "/words.json/randomWords"
  data := []Data{}

  data = append( data, Data{"hasDictionaryDef", randomWordsService.HasDictionaryDef } )

  if randomWordsService.IncludePartOfSpeech != ""{
    data = append( data, Data{"includePartOfSpeech", randomWordsService.IncludePartOfSpeech } )
  }

  if randomWordsService.ExcludePartOfSpeech != ""{
    data = append( data, Data{"excludePartOfSpeech", randomWordsService.ExcludePartOfSpeech } )
  }

  data = append( data, Data{"minCorpusCount", strconv.Itoa(randomWordsService.MinCorpusCount) } )
  data = append( data, Data{"maxCorpusCount", strconv.Itoa(randomWordsService.MaxCorpusCount) } )
  data = append( data, Data{"minDictionaryCount", strconv.Itoa(randomWordsService.MinDictionaryCount) } )
  data = append( data, Data{"maxDictionaryCount", strconv.Itoa(randomWordsService.MaxDictionaryCount) } )
  data = append( data, Data{"minLength", strconv.Itoa(randomWordsService.MinLength) } )
  data = append( data, Data{"maxLength", strconv.Itoa(randomWordsService.MaxLength) } )

  if randomWordsService.SortBy != ""{
    data = append( data, Data{"sortBy", randomWordsService.SortBy } )
  }

  if randomWordsService.SortOrder != ""{
    data = append( data, Data{"sortOrder", randomWordsService.SortOrder } )
  }
  
  data = append( data, Data{"limit", strconv.Itoa(randomWordsService.Limit) } )

  response_text, err := randomWordsService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return []RandomWord{}, err
  }

  b := []byte( response_text )
  var result []RandomWord
  err = json.Unmarshal(b, &result)
  return result, err
}
