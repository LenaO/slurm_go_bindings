package main

import "slurm/jobinfo"
import "fmt"

func main(){
	job_list := job_info.Get_all_jobs()
	fmt.Printf("Found %d jobs \n", job_list.Record_count)
/* a little bit nicer */
	fmt.Printf("Id\tName\t\tPartion\tUser\tRuntime\tStatus\t\t(Reason)\tNodes\tPriority\n")
	fmt.Printf("________________________________________________________________________________________________\n")
	for i := range job_list.Job_list {
		job := job_list.Job_list[i]
		fmt.Printf("%d\t%s\t%s\t%s %s\t%s\t%s\t%s\t%d\n" , 
		job.Job_id, job.Name, job.Partition, job.User_name,job_info.Get_job_runtime(job).String(), job_info.State_to_string(job.Job_state) ,
		job_info.Reason_to_string(job.State_reason), job.Nodes,job.Priority)
	}

}
