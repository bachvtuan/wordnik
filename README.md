# wordnik
Golang client for WORDNIK API.

## API Status:

#### Account
- [ ]  Pending all

#### Word
- [ ]  Pending all

#### Words
- [x] /words.json/search/{query}  
- [x] /words.json/wordOfTheDay  
- [ ] /words.json/reverseDictionary 
- [x] /words.json/randomWords  
- [x] /words.json/randomWord  

#### WordList
- [x] /word.json/{word}/examples
- [ ] /word.json/{word}
- [x] /word.json/{word}/definitions
- [ ] /word.json/{word}/topExample
- [x] /word.json/{word}/relatedWords
- [x] /word.json/{word}/pronunciations
- [x] /word.json/{word}/hyphenation
- [ ] /word.json/{word}/frequency
- [ ] /word.json/{word}/phrases
- [ ] /word.json/{word}/etymologies
- [x]  /word.json/{word}/audio

#### WordLists
- [ ]  Pending all

## Example 
Check out test/test.go or test/test_parallel.go .

This's the content fo test/test.go
```
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

  end := time.Now()
  fmt.Printf("Executed time %v\n",  end.Sub(start))
}
```