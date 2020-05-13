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

	jobs, err := extra. Get_job_info_accounting(uint32( id))
	if err!= nil {
		fmt.Printf(err.Error())
    return
	}
  fmt.Printf("JobId\tuser\taccount\tstate\t\tJobName\n")
  for i := range(jobs) {
    fmt.Printf("%d\t%s\t%s\t%s\t%s\n", jobs[i].JobId, jobs[i].User, jobs[i].Account, jobs[i].State, jobs[i].JobName)

  }
}
