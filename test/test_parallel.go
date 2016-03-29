package main 

import (
  "github.com/bachvtuan/wordnik"
  "fmt"
  "time"
  "sync"
)

//This is the demo to run multiple requests at the same time that reduce total executed time
func main() {
  
  start := time.Now()
  var wg sync.WaitGroup
  wg.Add(1)


  apiKey := "a2a73e7b926c924fad7001ca3111acd55af2ffabf50eb4ae5"
  service, err := wordnik.New( apiKey )
  if err != nil{
    panic(err)
  }

  c := make(chan string)
  count_done := 0
 
  //workers are 10
  for i:=0; i< 10;i++{
 
    go func ( c chan string ) {
 
      for {
        select{
          case action := <-c:
            
            
            if action == "WordOfDay"{

              service.WordsService.WordOfTheDayService.Date = "2015-01-20"
              result, err := service.WordsService.WordOfTheDayService.Do()
              count_done = count_done + 1
              if err != nil{
                panic(err)
              }else{
                fmt.Printf("Word of day is %+v\n\n\n", result)
              }
            }else if action == "Search"{
              service.WordsService.SearchService.Query = "honey*"
              searchResult, err := service.WordsService.SearchService.Do()
              count_done = count_done + 1
              if err != nil{
                panic(err)
              }else{
                fmt.Printf("Search result is %+v\n\n\n", searchResult)
              }
            }else if action == "RandomWord"{
              randomWord, err := service.WordsService.RandomWordService.Do()
              count_done = count_done + 1
              if err != nil{
                panic(err)
              }else{
                fmt.Printf("Random word is %+v\n\n\n", randomWord)
              }
            }else if action == "RandomWords"{
              randomWords, err := service.WordsService.RandomWordsService.Do()
              count_done = count_done + 1

              if err != nil{
                panic(err)
              }else{
                fmt.Printf("Random words are %+v\n\n\n", randomWords)
              }
            }else if action == "ExampleService"{
              service.WordService.ExampleService.Word = "appropriate"
              exampleResult, err := service.WordService.ExampleService.Do()
              count_done = count_done + 1
              if err != nil{
                panic(err)
              }else{
                fmt.Printf("exampleResult is %+v\n\n\n", exampleResult)
              }

            }else if action == "DefinitionService"{

              service.WordService.DefinitionService.Word = "appropriate"
              definitions, err := service.WordService.DefinitionService.Do()
              count_done = count_done + 1
              if err != nil{
                panic(err)
              }else{
                fmt.Printf("definitions is %+v\n\n\n", definitions)
              }

            }else if action == "RelativedWordService"{
              service.WordService.RelativedWordService.Word = "appropriate"
              relativedWords, err := service.WordService.RelativedWordService.Do()
              count_done = count_done + 1
              if err != nil{
                panic(err)
              }else{
                fmt.Printf("relativedWords is %+v\n\n\n", relativedWords)
              }

            }else if action == "PronunciationService"{
              service.WordService.PronunciationService.Word = "appropriate"
              pronunciations, err := service.WordService.PronunciationService.Do()
              count_done = count_done + 1
              if err != nil{
                panic(err)
              }else{
                fmt.Printf("pronunciations is %+v\n\n\n", pronunciations)
              }
            }else if action == "HyphenationService"{

              service.WordService.HyphenationService.Word = "appropriate"
              hyphenations, err := service.WordService.HyphenationService.Do()
              count_done = count_done + 1
              if err != nil{
                panic(err)
              }else{
                fmt.Printf("hyphenations is %+v\n\n\n", hyphenations)
              }
            }else if action == "AudioService"{

              service.WordService.AudioService.Word = "appropriate"
              audios, err := service.WordService.AudioService.Do()
              count_done = count_done + 1

              if err != nil{
                panic(err)
              }else{
                fmt.Printf("audios is %+v\n\n\n", audios)
              }
            }

 
            if count_done == 10{
              end := time.Now()
              fmt.Printf("Executed time %v\n",  end.Sub(start))
              wg.Done()
              return
            }
        }
      }
 
    }(c)
  }
 
  
  c <- "WordOfDay"
  c <- "Search"
  c <- "RandomWord"
  c <- "RandomWords"

  c <- "ExampleService"
  c <- "DefinitionService"
  c <- "RelativedWordService"
  c <- "PronunciationService"
  c <- "HyphenationService"
  c <- "AudioService"

  wg.Wait()

  

}