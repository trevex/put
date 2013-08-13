package main

import (
	"fmt"
    "flag"
)

var (
    subCommandFlags = []string{}

    processors = map[string]Processor{
        "dropbox": new(Dropbox),
    }
)

func main() {
    flag.Parse()
    command := flag.Arg(0)
    switch command {
    case "authorize":
        what := flag.Arg(1)
        //as := "dropbox"
        if flag.Arg(2) == "as" && flag.Arg(3) != "" {
            //as = flag.Arg(3)
            subCommandFlags = flag.Args()[4:]
        } else {
            subCommandFlags = flag.Args()[2:]
        }
        
        if processor, ok := processors[what]; ok {
            if result, err := processor.Authorize(); err == nil {
                fmt.Println(result)
            } else {
                fmt.Errorf("%s", err)
            }
        } else {
            fmt.Errorf("There is no known processor for \"%s\"! ", what) 
        }
    case "on":
        where := flag.Arg(1)
        fmt.Println(where)
    }       
}
