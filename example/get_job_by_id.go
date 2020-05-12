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
	fmt.Printf("Id\tName\t\tPartion\tUser\tRuntime\tStatus\t\t(Reason)\tNodes\tPriority\n")
	fmt.Printf("________________________________________________________________________________________________\n")
	for i := range job_list.Job_list {
		job := job_list.Job_list[i]
		fmt.Printf("%d\t%s\t%s\t%s %s\t%s\t%s\t%s\t%d\n" , 
		job.Job_id, job.Name, job.Partition, job.User_name,job_info.Get_job_runtime(job).String(), job.Job_stateS, 
		job_info.Reason_to_string(job.State_reason), job.Nodes,job.Priority)
	}


	end_time :=job_info.Get_job_endtime(uint32(id))
	fmt.Printf("End-Time: %s\n", end_time)

}
