package wordnik

import(
  "encoding/json"
  "errors"
  "strconv"
)

type WordService struct{
  Service *Service
  ExampleService *ExampleService
  DefinitionService *DefinitionService
  RelativedWordService *RelativedWordService
  PronunciationService *PronunciationService
  HyphenationService *HyphenationService
  AudioService *AudioService
}

func NewWordService( s *Service  ) *WordService{
  //with empty date
  exampleService := ExampleService{ }
  exampleService.Service = s
  exampleService.New()

  definitionService := DefinitionService{ }
  definitionService.Service = s
  definitionService.New()

  relativedWordService := RelativedWordService{ }
  relativedWordService.Service = s
  relativedWordService.New()
  
  pronunciationService := PronunciationService{ }
  pronunciationService.Service = s
  pronunciationService.New()

  hyphenationService := HyphenationService{ }
  hyphenationService.Service = s
  hyphenationService.New()

  audioService := AudioService{ }
  audioService.Service = s
  audioService.New()

  return &WordService{ s, &exampleService, &definitionService, &relativedWordService, &pronunciationService, &hyphenationService, &audioService}
}

type ExampleService struct{
  Service *Service
  Word  string 
  IncludeDuplicates  string
  UseCanonical  string
  Skip  int
  Limit  int
}


func ( exampleService *ExampleService ) New() {
  exampleService.IncludeDuplicates = "false"
  exampleService.UseCanonical = "false"
  exampleService.Skip = 0
  exampleService.Limit = 5
}

func ( exampleService *ExampleService ) Do() ( FullExample, error ) {
  //get 
  if exampleService.Word == ""{
    return FullExample{}, errors.New( "word is required" )
  }
  url := RootURL + "/word.json/"+ exampleService.Word +"/examples"
  
  data := []Data{}

  data = append( data, Data{"includeDuplicates", exampleService.IncludeDuplicates } )
  data = append( data, Data{"useCanonical", exampleService.UseCanonical } )
  data = append( data, Data{"skip", strconv.Itoa(exampleService.Skip) } )
  data = append( data, Data{"limit", strconv.Itoa(exampleService.Limit) } )


  response_text, err := exampleService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return FullExample{}, err
  }


  b := []byte( response_text )

  var listExample FullExample
  err = json.Unmarshal(b, &listExample)
  return listExample, err
}

type Provider struct{
  Name     string   `json:"name" bson:"name"`
  Id     int  `json:"id" bson:"id"`
}

type Example struct{

  Id     int        `json:"id" bson:"id"`
  Provider Provider `json:"provider" bson:"provider"`
  Year     int        `json:"year" bson:"year"`
  Rating     float32        `json:"rating" bson:"rating"`
  Url     string   `json:"url" bson:"url"`
  Word     string   `json:"word" bson:"word"`
  Text     string   `json:"text" bson:"text"`
  DocumentId     int  `json:"documentId" bson:"documentId"`
  ExampleId     int  `json:"exampleId" bson:"exampleId"`
  Title     string   `json:"title" bson:"title"`
}

type FullExample struct{
  
  Examples []Example `json:"examples" bson:"examples"`
}



type DefinitionService struct{
  Service *Service
  Word  string 
  Limit  int 
  PartOfSpeech string
  IncludeRelated string
  SourceDictionaries  string
  UseCanonical  string
  IncludeTags  string
}

func ( definitionService *DefinitionService ) New()  {
  definitionService.Limit = 200
  definitionService.IncludeRelated = "true"
  definitionService.SourceDictionaries = "all"
  definitionService.UseCanonical = "false"
  definitionService.IncludeTags = "false"
}


func ( definitionService *DefinitionService ) Do() ( []Definition, error ) {
  //get 
  if definitionService.Word == ""{
    return []Definition{}, errors.New( "word is required" )
  }

  url := RootURL + "/word.json/"+ definitionService.Word +"/definitions"
  
  data := []Data{}

  data = append( data, Data{"limit", strconv.Itoa(definitionService.Limit) } )
  if definitionService.PartOfSpeech != ""{
    data = append( data, Data{"partOfSpeech", definitionService.PartOfSpeech } )
  }
  
  data = append( data, Data{"includeRelated", definitionService.IncludeRelated } )
  data = append( data, Data{"sourceDictionaries", definitionService.SourceDictionaries } )
  data = append( data, Data{"useCanonical", definitionService.UseCanonical } )
  data = append( data, Data{"includeTags", definitionService.IncludeTags } )


  response_text, err := definitionService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return []Definition{}, err
  }


  b := []byte( response_text )

  var definitions []Definition
  err = json.Unmarshal(b, &definitions)
  return definitions, err
}

type Label struct{
  Text string `json:"text" bson:"text"`
  Type string `json:"type" bson:"type"`
}

type Cite struct{
  Source string `json:"source" bson:"source"`
  Cite string `json:"cite" bson:"cite"`

}

type ExampleUse struct{
  Text string `json:"text" bson:"text"`
}

type RelatedWord struct{
  Gram string `json:"gram" bson:"gram"`
  RelationshipType string `json:"relationshipType" bson:"relationshipType"`
  Words []string `json:"words" bson:"words"`
}

type  TextPron struct{
  Raw string `json:"raw" bson:"raw"`
  Seq int `json:"seq" bson:"seq"`
  RawType string `json:"rawType" bson:"rawType"`
}

type Definition struct{
  TextProns  []TextPron `json:"textProns" bson:"textProns"`
  SourceDictionary string `json:"sourceDictionary" bson:"sourceDictionary"`
  ExampleUses  []ExampleUse `json:"exampleUses" bson:"exampleUses"`
  RelatedWords  []RelatedWord `json:"relatedWords" bson:"relatedWords"`
  Labels  []Label `json:"labels" bson:"labels"`
  Citations  []Cite `json:"citations" bson:"citations"`
  Word string `json:"word" bson:"word"`
  Sequence string `json:"sequence" bson:"sequence"`
  ExtendedText string `json:"extendedText" bson:"extendedText"`
  PartOfSpeech string `json:"partOfSpeech" bson:"partOfSpeech"`
  
  AttributionText string `json:"attributionText" bson:"attributionText"`
  Text string `json:"text" bson:"text"`
  Score float32 `json:"score" bson:"score"`
}

/**
 * END DEFINITION
 */

type RelativedWordService struct{
  Service *Service

  Word  string 
  UseCanonical  string
  RelationshipTypes string
  LimitPerRelationshipType  int
}


func ( relativedWordService *RelativedWordService ) New() {
  relativedWordService.Word = ""
  relativedWordService.UseCanonical = "false"
  relativedWordService.RelationshipTypes = ""
  relativedWordService.LimitPerRelationshipType = 10
}

func ( relativedWordService *RelativedWordService ) Do() ( []RelativedWord, error ) {
  //get 
  if relativedWordService.Word == ""{
    return []RelativedWord{}, errors.New( "word is required" )
  }
  url := RootURL + "/word.json/"+ relativedWordService.Word +"/relatedWords"
  
  data := []Data{}

  
  data = append( data, Data{"useCanonical", relativedWordService.UseCanonical } )
  if relativedWordService.RelationshipTypes != ""{
    data = append( data, Data{"relationshipTypes", relativedWordService.RelationshipTypes } )
  }
  data = append( data, Data{"limitPerRelationshipType", strconv.Itoa(relativedWordService.LimitPerRelationshipType) } )
  


  response_text, err := relativedWordService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return []RelativedWord{}, err
  }


  b := []byte( response_text )

  var relatedWords []RelativedWord
  err = json.Unmarshal(b, &relatedWords)
  return relatedWords, err
}

type RelativedWord struct{

  RelationshipType  string  `json:"relationshipType" bson:"relationshipType"`
  Words  []string  `json:"words" bson:"words"`
}

/**
 * END RelativeWord
 */

 type PronunciationService struct{
  Service *Service

  Word  string 
  UseCanonical  string
  SourceDictionary string
  TypeFormat string
  Limit  int
}


func ( pronunciationService *PronunciationService ) New() {
  pronunciationService.Word = ""
  pronunciationService.UseCanonical = "false"
  pronunciationService.SourceDictionary = ""
  pronunciationService.TypeFormat = ""
  pronunciationService.Limit = 10
}

func ( pronunciationService *PronunciationService ) Do() ( []Pronunciation, error ) {
  //get 
  if pronunciationService.Word == ""{
    return []Pronunciation{}, errors.New( "word is required" )
  }
  url := RootURL + "/word.json/"+ pronunciationService.Word +"/pronunciations"
  
  data := []Data{}

  
  data = append( data, Data{"useCanonical", pronunciationService.UseCanonical } )
  if pronunciationService.SourceDictionary != ""{
    data = append( data, Data{"sourceDictionary", pronunciationService.SourceDictionary } )
  }
  if pronunciationService.TypeFormat != ""{
    data = append( data, Data{"typeFormat", pronunciationService.TypeFormat } )
  }
  data = append( data, Data{"limit", strconv.Itoa(pronunciationService.Limit) } )
  


  response_text, err := pronunciationService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return []Pronunciation{}, err
  }


  b := []byte( response_text )

  var pronunciations []Pronunciation
  err = json.Unmarshal(b, &pronunciations)
  return pronunciations, err
}

type Pronunciation struct{
  RawType  string  `json:"rawType" bson:"rawType"`
  Seq  int  `json:"seq" bson:"seq"`
  Raw  string  `json:"raw" bson:"raw"`
}

/**
 * END Pronunciation
 */
type HyphenationService struct{
  Service *Service

  Word  string 
  UseCanonical  string
  SourceDictionary string
  Limit  int
}


func ( hyphenationService *HyphenationService ) New() {
  hyphenationService.Word = ""
  hyphenationService.UseCanonical = "false"
  hyphenationService.SourceDictionary = ""
  hyphenationService.Limit = 50
}

func ( hyphenationService *HyphenationService ) Do() ( []Hyphenation, error ) {
  //get 
  if hyphenationService.Word == ""{
    return []Hyphenation{}, errors.New( "word is required" )
  }
  url := RootURL + "/word.json/"+ hyphenationService.Word +"/hyphenation"
  
  data := []Data{}

  
  data = append( data, Data{"useCanonical", hyphenationService.UseCanonical } )
  if hyphenationService.SourceDictionary != ""{
    data = append( data, Data{"sourceDictionary", hyphenationService.SourceDictionary } )
  }
  
  data = append( data, Data{"limit", strconv.Itoa(hyphenationService.Limit) } )
  


  response_text, err := hyphenationService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return []Hyphenation{}, err
  }


  b := []byte( response_text )

  var hyphenations []Hyphenation
  err = json.Unmarshal(b, &hyphenations)
  return hyphenations, err
}

type Hyphenation struct{
  Type  string  `json:"type" bson:"type"`
  Seq  int  `json:"seq" bson:"seq"`
  Text  string  `json:"text" bson:"text"`
}

/**
 * END Hypenation
 */

type AudioService struct{
  Service *Service

  Word  string 
  UseCanonical  string
  Limit  int
}


func ( audioService *AudioService ) New() {
  audioService.Word = ""
  audioService.UseCanonical = "false"
  audioService.Limit = 50
}

func ( audioService *AudioService ) Do() ( []Audio, error ) {
  //get 
  if audioService.Word == ""{
    return []Audio{}, errors.New( "word is required" )
  }
  url := RootURL + "/word.json/"+ audioService.Word +"/audio"
  
  data := []Data{}

  
  data = append( data, Data{"useCanonical", audioService.UseCanonical } )
  data = append( data, Data{"limit", strconv.Itoa(audioService.Limit) } )
  

  response_text, err := audioService.Service.GetUrlContent( GET, url, data )
  if err != nil{
    return []Audio{}, err
  }


  b := []byte( response_text )

  var audios []Audio
  err = json.Unmarshal(b, &audios)
  return audios, err
}

type Audio struct{
  CommentCount  int  `json:"commentCount" bson:"commentCount"`
  CreatedBy  string  `json:"createdBy" bson:"createdBy"`
  CreatedAt  string  `json:"createdAt" bson:"createdAt"`
  Id  int  `json:"id" bson:"id"`
  Word  string  `json:"word" bson:"word"`
  Duration  float32  `json:"duration" bson:"duration"`
  FileUrl  string  `json:"fileUrl" bson:"fileUrl"`
  AudioType  string  `json:"audioType" bson:"audioType"`
  AttributionText  string  `json:"attributionText" bson:"attributionText"`
}

/**
 * END AUDIO
 */ 