package main

import (
  "fmt"
  "encoding/csv"
  "os"
  "flag"
)

func main(){

  csv_filename := flag.String("csv", "problems.csv", "a csv file in the format 'question, answer' ")
  flag.Parse()

  problem_file, err := os.Open(*csv_filename)
  if (err != nil){
    panic(err)
  }

  reader := csv.NewReader(problem_file)
  reader.FieldsPerRecord = 2
  data, err := reader.ReadAll()

  if (err != nil){
    panic(err)
  }

  // iterate over csv data
  correct_answers := 0
  for _, row := range data{
    question, answer := row[0], row[1]

    fmt.Println(question)

    var userinput string
    fmt.Scanln(&userinput)

    if userinput == answer{
      correct_answers++
    }
  }

  fmt.Printf("you got %d / %d answers correct\n", correct_answers, len(data))
}
