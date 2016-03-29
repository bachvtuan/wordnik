package wordnik

import(
  "encoding/json"
  "errors"
  "strconv"
)

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
  wordOfTheDayService.New()

  searchService := SearchService{ Service : s }
  //init default value
  searchService.New()
  
  //int randomWords Service
  randomWordsService := RandomWordsService{ Service:s }
  randomWordsService.New()


  randomWordService := RandomWordService{ Service:s }
  randomWordService.New()
  

  return &WordsService{ s, wordOfTheDayService, &searchService, &randomWordsService, &randomWordService }
}

type WordOfTheDayService struct{
  Service *Service
  Date  string `json:"date" bson:"date"`
}

func ( wordOfTheDayService *WordOfTheDayService ) New() {
  wordOfTheDayService.Date = ""
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


type ShortExample struct{
  Url     string        `json:"url" bson:"url"`
  Text     string        `json:"text" bson:"text"`
  Id     int        `json:"id" bson:"id"`
  Title     string        `json:"title" bson:"title"`
}

type ShortDefinition struct{
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
  Examples []ShortExample `json:"examples" bson:"examples"` 
  Definitions []ShortDefinition `json:"definitions" bson:"definitions"` 
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

func ( searchService *SearchService ) New() {
  searchService.Query = ""
  searchService.CaseSensitive = "true" 
  searchService.IncludePartOfSpeech = ""
  searchService.ExcludePartOfSpeech = ""
  searchService.MinCorpusCount = 5
  searchService.MaxCorpusCount = -1
  searchService.MinDictionaryCount = 1
  searchService.MaxDictionaryCount = -1
  searchService.MinLength = 1
  searchService.MaxLength = -1
  searchService.Skip = 0
  searchService.Limit = 10
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

func ( randomService *RandomWordService ) New() {
  randomService.HasDictionaryDef = "false" 
  randomService.IncludePartOfSpeech = ""
  randomService.ExcludePartOfSpeech = ""
  randomService.MinCorpusCount = 0
  randomService.MaxCorpusCount = -1
  randomService.MinDictionaryCount = 1
  randomService.MaxDictionaryCount = -1
  randomService.MinLength = 5
  randomService.MaxLength = -1
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


func ( randomWordsService *RandomWordsService ) New() {
  randomWordsService.HasDictionaryDef = "false" 
  randomWordsService.IncludePartOfSpeech = ""
  randomWordsService.ExcludePartOfSpeech = ""
  randomWordsService.MinCorpusCount = 0
  randomWordsService.MaxCorpusCount = -1
  randomWordsService.MinDictionaryCount = 1
  randomWordsService.MaxDictionaryCount = -1
  randomWordsService.MinLength = 5
  randomWordsService.MaxLength = -1
  randomWordsService.Limit = 10
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
