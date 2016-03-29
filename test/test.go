package main 

import (
  "github.com/bachvtuan/wordnik"
  "fmt"
  "time"
)

func main() {
  
  start := time.Now()
  apiKey := "a2a73e7b926c924fad7001ca3111acd55af2ffabf50eb4ae5"
  service, err := wordnik.New( apiKey )

  if err != nil{
    panic(err)
  }

  //select word on specific date
  service.WordsService.WordOfTheDayService.Date = "2015-01-20"
  result, err := service.WordsService.WordOfTheDayService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("Word of day is %+v\n\n\n", result)
  }

  ////////////////////////////////////////////////////////////

  //Search with word begin with honey
  service.WordsService.SearchService.Query = "honey*"
  searchResult, err := service.WordsService.SearchService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("Search result is %+v\n\n\n", searchResult)
  }

  ////////////////////////////////////////////////////////////  

  //Get random words
  randomWords, err := service.WordsService.RandomWordsService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("Random words are %+v\n\n\n", randomWords)
  }

  ////////////////////////////////////////////////////////////

  //Get single random word
  randomWord, err := service.WordsService.RandomWordService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("Random word is %+v\n\n\n", randomWord)
  }

  /**
   * Word service section
   */
  service.WordService.ExampleService.Word = "appropriate"
  exampleResult, err := service.WordService.ExampleService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("exampleResult is %+v\n\n\n", exampleResult)
  }


  
  service.WordService.DefinitionService.Word = "appropriate"
  definitions, err := service.WordService.DefinitionService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("definitions is %+v\n\n\n", definitions)
  }

  service.WordService.RelativedWordService.Word = "appropriate"
  relativedWords, err := service.WordService.RelativedWordService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("relativedWords is %+v\n\n\n", relativedWords)
  }
  
  service.WordService.PronunciationService.Word = "appropriate"
  pronunciations, err := service.WordService.PronunciationService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("pronunciations is %+v\n\n\n", pronunciations)
  }


  service.WordService.HyphenationService.Word = "appropriate"
  hyphenations, err := service.WordService.HyphenationService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("hyphenations is %+v\n\n\n", hyphenations)
  }

  service.WordService.AudioService.Word = "appropriate"
  audios, err := service.WordService.AudioService.Do()

  if err != nil{
    panic(err)
  }else{
    fmt.Printf("audios is %+v\n\n\n", audios)
  }


  end := time.Now()
  fmt.Printf("Executed time %v\n",  end.Sub(start))
}