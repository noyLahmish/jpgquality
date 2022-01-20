
package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
     pwd, err := os.Getwd()
            if err != nil {
                fmt.Println(err)
            }
     files, err := ioutil.ReadDir(pwd+"/cmd/images/")
        if err != nil {
            fmt.Println(err)
        }
    for _, f := range files {
    os.Remove(pwd+"/cmd/images/"+f.Name())
    }
}



