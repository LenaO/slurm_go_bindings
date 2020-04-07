package main

import "slurm/jobinfo"
import "slurm"
import "fmt"
import "os"
import "strconv"

func main(){

	if len(os.Args)<2 {
		fmt.Printf("Please specify Job ID\n")
		return
	}
	id,_ := strconv.Atoi(os.Args[1])
	job_list := job_info.Get_job(uint32(id))
	if job_list.Error_code != 0 {
		msg := slurm.GetErrorString(job_list.Error_code)
		fmt.Printf("Error: %s\n" ,msg)
		return



	}
	for i := range job_list.Job_list {
		job_info.Print_Job_info(job_list.Job_list[i])
	}

	fmt.Printf("Id\tame\tPartion\tUser\tRuntime\tStatus\tReason\t\tNodes\n")
	fmt.Printf("________________________________________\n")
	for i := range job_list.Job_list {
		job := job_list.Job_list[i]
		fmt.Printf("%d\t%s\t%s\t%s %s\t%s\t%s\t%s\t\t%d\n" , 
		job.Job_id, job.Name, job.Partition, job.User_name,job_info.Get_job_runtime(job).String(), job_info.State_to_string(job.Job_state) ,
		job_info.Reason_to_string(job.State_reason), job.Nodes, job.Pre_sus_time)
	}

}
