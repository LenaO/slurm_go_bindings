package main


import "slurm/submitjob"
import "slurm"
import "os/user"
import "strconv"
import "fmt"
func main(){
	job_desc :=  submit_job.Job_descriptor{} 
	job_desc.Script = "#! /bin/bash\n ls "

	user, _:= user.Current()
	userid , _ := strconv.Atoi(user.Uid)
	job_desc.User_id= uint32(userid)
	groupid , _ := strconv.Atoi(user.Gid)

	job_desc.Group_id= uint32(groupid)
	job_desc.Name = "test_job"
	job_desc.Partition="long"
	job_desc.Time_limit = uint32(2)
	job_desc.Min_nodes =uint32(1)
	job_desc.Std_out = ("./out.txt")
	job_desc.Std_err = ("./err.txt")
	job_desc.Work_dir = ("/home0/data/")
	answer := submit_job.Submit_job(&job_desc)
	if(answer.Error_code != 0) {
		msg := slurm.GetErrorString(answer.Error_code)
		fmt.Printf("Error: %s\n" ,msg)
		return
	}
	fmt.Printf("Submitted Job %d\n", answer.Job_id)
}
