package main

import "slurm/submitjob"
import "slurm"
import "os/user"
import "os"
import "strconv"
import "fmt"
import "os/exec"
import "path/filepath"
import "slurm/jobinfo"
import "time"


func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
	return false
    }
    return !info.IsDir()
}
func build_container(file_name,container_name string){

	cmd := exec.Command("sudo", "/usr/local/bin/singularity", "build",container_name, file_name)
	fmt.Print("Now build new container")
	fmt.Printf("%s\n", cmd.String())
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
			fmt.Printf("error in creating container %s \n", err)
			
			fmt.Printf("%s\n", stdoutStderr)
//		return
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func main(){
	job_desc :=  submit_job.Job_descriptor{}

	dir, _ := os.Getwd()
		container := filepath.Join(dir, "openmp_container.img")
	definition := filepath.Join(dir, "openmp_container.def")
	if !fileExists(container){
		build_container(definition,container)
	}
	
	if !fileExists(container){
		return
	}
	/* use Cmd to create our script */

	job_desc.Script = "#!/bin/bash\n export PATH=$PATH:/usr/local/bin\n hostname \n"
	job_desc.Script+= "export OMP_NUM_THREADS=$SLURM_JOB_CPUS_PER_NODE\n"
	cmd := exec.Command( "/usr/local/bin/singularity", "exec",container, "/opt/openmp_example" )

	job_desc.Script+= cmd.String() 
	fmt.Printf("cmd %s\n", 	job_desc.Script)
	user, _:= user.Current()
	userid , _ := strconv.Atoi(user.Uid)
	job_desc.User_id= uint32(userid)
	groupid , _ := strconv.Atoi(user.Gid)

	job_desc.Group_id= uint32(groupid)
	job_desc.Name = "test_job"
	job_desc.Partition="long"
	job_desc.Time_limit = uint32(60)
	job_desc.Min_nodes =uint32(1)
	job_desc.Num_tasks = uint32(1)

	job_desc.Cpus_per_task = uint16(2)
	job_desc.Std_out = ("./%j-out.txt")
	job_desc.Std_err = ("./%j-err.txt")
	job_desc.Work_dir = dir

	answer := submit_job.Submit_job(&job_desc)
	if(answer.Error_code != 0) {
		msg := slurm.GetErrorString(answer.Error_code)
		fmt.Printf("Error: %s\n" ,msg)
		return
	}
	fmt.Printf("Submitted Job %d\n", answer.Job_id)

	/*Now, we submit the same jon again, ut with some oversubsciption */ 
	job_desc.Script = "#!/bin/bash\n export PATH=$PATH:/usr/local/bin\n hostname \n"
	job_desc.Script+= "export OMP_NUM_THREADS=4\n"

	job_desc.Script+= cmd.String() 
	fmt.Printf("cmd %s\n", 	job_desc.Script)
	answer2 := submit_job.Submit_job(&job_desc)
	if(answer2.Error_code != 0) {
		msg := slurm.GetErrorString(answer.Error_code)
		fmt.Printf("Error: %s\n" ,msg)
		return
	}
	fmt.Printf("Submitted Job %d\n", answer2.Job_id)



	job_list := job_info.Get_job(answer.Job_id)
	if job_list.Error_code != 0 {
		msg := slurm.GetErrorString(job_list.Error_code)
		fmt.Printf("Error: %s\n" ,msg)
		return

	}
	job := job_list.Job_list[0]

	fmt.Printf("job is %s\n",job.Job_stateS)
	state := job.Job_stateS
	for state == "Pending" || state == "Running" {
		time.Sleep(2 * time.Second)
		job_list = job_info.Get_job(answer.Job_id)
		if job_list.Error_code != 0 {
			msg := slurm.GetErrorString(job_list.Error_code)
			fmt.Printf("Error: %s\n" ,msg)
			return

		}
		job = job_list.Job_list[0]

		state = job.Job_stateS

		fmt.Printf("job is %s\n",job.Job_stateS)


	}

	fmt.Printf("Total runtime first job  %s\n",job_info.Get_job_runtime(job).String() )
	/*wait for second job */
	job_list = job_info.Get_job(answer2.Job_id)
	if job_list.Error_code != 0 {
		msg := slurm.GetErrorString(job_list.Error_code)
		fmt.Printf("Error: %s\n" ,msg)
		return

	}
	job = job_list.Job_list[0]

	fmt.Printf("job is %s\n",job.Job_stateS)
	state = job.Job_stateS
	for state == "Pending" || state == "Running" {
		time.Sleep(2 * time.Second)
		job_list = job_info.Get_job(answer2.Job_id)
		if job_list.Error_code != 0 {
			msg := slurm.GetErrorString(job_list.Error_code)
			fmt.Printf("Error: %s\n" ,msg)
			return

		}
		job = job_list.Job_list[0]

		state = job.Job_stateS

		fmt.Printf("job is %s\n",job.Job_stateS)


	}


	fmt.Printf("Total runtime second job  %s\n",job_info.Get_job_runtime(job).String() )

}
