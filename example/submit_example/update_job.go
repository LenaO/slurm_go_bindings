package main


import "slurm/submitjob"
import "slurm"
import "os"
import "strconv"
import "fmt"
func main(){
	if len(os.Args)<4 {
		fmt.Printf("Synthax specify JobID, qos and partition \n")
		return
	}
	var ops submit_job.Update_job_options
	id,err := strconv.Atoi(os.Args[1])
	if err != nil  {
		fmt.Printf("Invalid job id (no int)  %s\n", os.Args[1] )
		return
	}

	ops.Qos = os.Args[2]
	ops.Partition = os.Args[3]

	err2 := submit_job.Update_job(ops, uint32(id))
	if err2!= uint32(0) {
		fmt.Printf("error %s \n", slurm.GetErrorString(err2))
	}
}
