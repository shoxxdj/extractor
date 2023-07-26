package main

import (
 "bufio"
 "fmt"
 "os"
 "log"
 "flag"
 "strings"
)




func main(){

 startPtr:=flag.String("start","defaultValue_Th4tn3v3r_Ex1st","the first string")
 endPtr:=flag.String("end","defaultValue_Th4tn3v3r_Ex1st","the end string")

 continuePtr:=flag.Bool("continue",false,"if not set, stop after first block")

 startLikePtr:=flag.String("start-like","defaultValue_Th4tn3v3r_Ex1st","part the first string")
 endLikePtr:=flag.String("end-like","defaultValue_Th4tn3v3r_Ex1st","part of end string")

 flag.Usage = func() {
  fmt.Fprintf(os.Stderr, "Extractor : a binary to extract strings between two boundary. v:0.1\n")
    
  flag.VisitAll(func(f *flag.Flag) {
    fmt.Fprintf(os.Stderr, "\t-%v: %v\n", f.Name,f.Usage) // f.Name, f.Value
  })
 }


 flag.Parse()
  


 var hasStarted =false

 scanner := bufio.NewScanner(os.Stdin)
 for scanner.Scan() {
  text := scanner.Text()
  if(text == *startPtr || strings.Contains(text,*startLikePtr)){
   hasStarted=true
  }
  if(hasStarted){
   fmt.Println(text)
  }
  if((text == *endPtr || strings.Contains(text,*endLikePtr)) && hasStarted ){
   if(!*continuePtr){
    return 
   }
  hasStarted=false
  }
 }

 if err := scanner.Err(); err != nil {
	log.Println(err)
 }
}
