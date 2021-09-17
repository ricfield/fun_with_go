package main

import ("net/http"; "fmt"; "io/ioutil"; "time"; "encoding/json")

type api_data struct {
   Base string
   Currency string
   Amount string
}
   
type data struct {
   Data api_data
}

func main() {
   var j data
   
   fmt.Printf("Time                 Base   Currency    Amount\n")
   fmt.Printf("------------------------------------------------\n")
   
   for i := 1; i <= 100; i++ {
   
      resp, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
      contents, err := ioutil.ReadAll(resp.Body)
      json.Unmarshal(contents, &j)
	  
      if err != nil {
      }
	
	
      fmt.Printf("%s  %s    %s         %s     \r", time.Now().Format("01-2-2006 15:04:05"), j.Data.Base, j.Data.Currency, j.Data.Amount)

      time.Sleep(time.Second * 5)
   }
}
