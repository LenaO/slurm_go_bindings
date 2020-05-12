package main


import "slurm/submitjob"
import "slurm"
import "os"
import "strconv"
import "fmt"
import "os/exec"
import "path/filepath"
import "slurm/jobinfo"
import "time"
import "os/user"



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
			fmt.Printf("error in creating container %s \n", err);
//		return
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func main(){
	job_desc :=  submit_job.Job_descriptor{}

	dir, _ := os.Getwd()
		container := filepath.Join(dir, "mpi_container.img")
	definition := filepath.Join(dir, "mpi_container.def")
	if !fileExists(container){
		build_container(definition,container)
	}

	if !fileExists(container){
		return
	}
	/* use Cmd to create our script */

	job_desc.Script = "#!/bin/bash\n export PATH=$PATH:/usr/local/bin\n srun hostname \n"
	cmd := exec.Command( "/home0/opt/openmpi/bin/mpirun", "-mca  btl_tcp_if_include eth1", "/usr/local/bin/singularity", "exec",container, "/opt/mpi_pingpong" )
	job_desc.Script+= cmd.String()
	fmt.Printf("cmd %s\n",	job_desc.Script)
	user, _:= user.Current()
	userid , _ := strconv.Atoi(user.Uid)
	job_desc.User_id= uint32(userid)
	groupid , _ := strconv.Atoi(user.Gid)

	job_desc.Group_id= uint32(groupid)
	job_desc.Name = "test_job"
	job_desc.Partition="long"
	job_desc.Time_limit = uint32(60)
//	job_desc.Min_nodes =uint32(2)
	job_desc.Ntasks_per_node = uint16(1)
	job_desc.Num_tasks = uint32(2)
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

	time.Sleep(5 * time.Second)

	job_list := job_info.Get_job(answer.Job_id)
	if job_list.Error_code != 0 {
		msg := slurm.GetErrorString(job_list.Error_code)
		fmt.Printf("Error: %s\n" ,msg)
		return

	}
	job := job_list.Job_list[0]

	fmt.Printf("job is %s\n",job.Job_stateS)
	state := job.Job_stateS
	if state == "Pending" {
		fmt.Printf("Move job to another partition \n")
		var ops submit_job.Update_job_options

		ops.Qos = "shortjobs"
		ops.Partition = "short"
		ops.Min_nodes = 1
//		ops.Ntasks_per_node = 2
		err2 := submit_job.Update_job(ops, uint32(answer.Job_id))
		if err2!= uint32(0) {
			fmt.Printf("error %s \n", slurm.GetErrorString(err2))
		}
	}

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

	fmt.Printf("Total runtime  %s\n",job_info.Get_job_runtime(job).String() )
}



