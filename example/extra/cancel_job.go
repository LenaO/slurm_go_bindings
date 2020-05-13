package main

import "slurm/extra"
import "fmt"
import "os"
import "strconv"
func main(){

	if len(os.Args)<2 {
		fmt.Printf("Please specify Job ID\n")
		return
	}
	id,_ := strconv.Atoi(os.Args[1])
	fmt.Printf("try to cancel %d\n", id)
	err:= extra.Cancel_job(uint32( id))
	if(err!= nil){
		fmt.Printf(err.Error())

	}
}
