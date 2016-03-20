package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "os"
)

type ComicInfo struct {
  Transcript string
}

func SearchIssues(page string)(*ComicInfo, error){

  resp, err := http.Get("https://xkcd.com/" + page + "/info.0.json")

  if err != nil {
    return nil, err
  }

  if resp.StatusCode != http.StatusOK{
    resp.Body.Close()
    return nil, fmt.Errorf("search query failed: %s", resp.Status)
  }
  var result ComicInfo

  if err := json.NewDecoder(resp.Body).Decode(&result); err != nil{
    resp.Body.Close()
    return nil, err
  }

  resp.Body.Close()

  return &result, nil
}


func main() {

  for i, post := range os.Args[1:]{

    result , err := SearchIssues(post)
    if err != nil{
      fmt.Println("errr")
    }

    fmt.Printf("Issue Number: %d\n\n", i)
    fmt.Printf("%s\n\n", result.Transcript)
  }


}
