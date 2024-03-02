Band Protocol assignment (Software engineer)

Problem 1: Boss Baby's Revenge
Example :
str := bossBaby("SRSSRRR")
Problem 2: Superman's Chicken Rescue
Example :
n := 5
k := 5
chicken := []int{2, 5, 10, 12, 15}
maxSave := superman(n, k, chicken)
Problem 3: Transaction Broadcasting and Monitoring Client

1. Broadcast Transaction:
   urlPost := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
   payload := jsonCode{
   Symbol: "ETH",
   Price: 4500,
   Timestamp: 1678912345,
   }
   jsonBytes, \_ := json.Marshal(payload)
   respP := callApi(urlPost, "POST", jsonBytes)
   var dataP respPOST
   errP := json.Unmarshal([]byte(respP), &dataP)
   if errP != nil {
   fmt.Println(errP)
   }
   fmt.Println(dataP.Tx_hash)

2. Transaction Status Monitoring:
   Tx_hash := "5dfcbf86eba4ca4473eb20cc05d7448fa5da7b4263d1ccf674379caf3c022fa3"
   urlGet := "https://mock-node-wgqbnxruha-as.a.run.app/check/" + Tx_hash
   respG := callApi(urlGet, "GET", nil)
   var dataG respGET
   errG := json.Unmarshal([]byte(respG), &dataG)
   if errG != nil {
   fmt.Println(errG)
   }
   fmt.Println(RepInfo(dataG.Tx_status))
3. Documentation:

- `CONFIRMED`: Transaction has been processed and confirmed
- `FAILED`: Transaction failed to process
- `PENDING`: Transaction is awaiting processing
- `DNE`: Transaction does not exist
  respData := "DNE"
  fmt.Println(RepInfo(respData))
