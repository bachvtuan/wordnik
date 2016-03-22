package worknik

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

type DailyWord struct{
  Id int `json:"id" bson:"id"` 
  Word string `json:"word" bson:"word"` 
  PublishDate string `json:"publishDate" bson:"publishDate"` 
  PublishDate string `json:"publishDate" bson:"publishDate"` 
  ContentProvider ContentProvider `json:"contentProvider" bson:"contentProvider"` 
  Note string `json:"note" bson:"note"` 
  Examples []Example `json:"examples" bson:"examples"` 
  Definitions []Definition `json:"definitions" bson:"definitions"` 
}
/*
{
  "id": 520800,
  "word": "phronesis",
  "publishDate": "2016-03-22T03:00:00.000+0000",
  "contentProvider": {
    "name": "wordnik",
    "id": 711
  },
  "note": "The word 'phronesis' comes from an Ancient Greek word meaning 'practical wisdom', from a word meaning 'mind'.",
  "examples": [
    {
      "url": "http://api.wordnik.com/v4/mid/6eadc41f2472d119498b47a73e36393114e1a5e956f5cc93a0b66a68efdf60de",
      "text": "Aristotle called it \"phronesis,\" or practical wisdom.",
      "id": 570909690,
      "title": "Q&A: What Would Plato Have Done?"
    }
  ],
  "definitions": [
    {
      "text": "Practical judgment; the faculty of conducting one's self wisely.",
      "source": "century",
      "partOfSpeech": "noun"
    }
  ]
}*/