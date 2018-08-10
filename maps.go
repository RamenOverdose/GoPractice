package main

import "fmt"
import "strings"

type stats struct{
  points int
  assists int
  rebounds int
}

func main(){
  var scores map[string]stats
  scores = make(map[string]stats)
  scores["james"] = stats{51, 8, 8}
  scores["durant"] = stats{35, 6, 10}
  scores["curry"] = stats{27, 6, 4}
  scores["thompson"] = stats{21, 4, 5}
  scores["green"] = stats{12, 9, 9}
  scores["korver"] = stats{15, 2, 2}
  scores["smith"] = stats{3, 1, 3}

  var name string
  var tempStats stats
  var statChoice string
  fmt.Print("Which player would you like information on? Please enter their last name: ")
  fmt.Scanln(&name)
  name = strings.ToLower(name)
  tempStats = scores[name]
  fmt.Print("Which stat would you like to know about them? (Points, Assists, Rebounds): ")
  fmt.Scanln(&statChoice)
  statChoice = strings.ToLower(statChoice)

  if strings.Compare(statChoice, "points") == 0{
      fmt.Println(tempStats.points)
  } else if strings.Compare(statChoice, "assists") == 0{
      fmt.Println(tempStats.assists)
  } else if strings.Compare(statChoice, "rebounds") == 0{
      fmt.Println(tempStats.rebounds)
  } else {
      fmt.Println("Not a valid statistic.")
  }

}
