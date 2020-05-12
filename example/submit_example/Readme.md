# Submission of jobs 
This folder shows in a few more examples of how jobs can be submitted in Slurm. Some examples use containers.
Attention: The parameters for job names and partitions probably have to be adjusted! 

# Simple Jobs

## submit_job.go
In this example, a simple bash-Jobs is submitted. The used partition is *long* (adapt probably). 
```
job_desc.Partition="long"
```

The job sets two environment variables and executes a
```
hostname
env | grep SLUM
``` 
On a single node of the cluster (single task job).
The application does not wait until the hob is completed, but dirctly returns.
The (std) output is written to 
out-jobid.txt, the std- error to err-jobid.txt

```
job_desc.Std_out = ("./out-%j.txt")
job_desc.Std_err = ("./err-%j.txt")
````


## update_job.go
This example allows to update the qos and the partition a job is running on. This can help to move the job to another queue with another partition. 
Note to users: In theory, the API allows the update of the number of nodes and the tasks per node. However, since this is only allowed by root or a slurm admin, we do not include an example here. 
Synthax
```
./update_job JobId qos partition 
``` 
(Note: This requires that the Job with the Id JobID is already submitted and in a pending state) 


# Container jobs

The following examples all submit a job that starts singulrity containers.
These containers, if they do not exist, are created. However, problems can arise if the user does not have sudo permissions..

## The containers

The first container is an MPI container. This is used by and `submit_mpi_containier.go` and `submit_mpi_and_update.go`. The definition is stored in `mpi_container.def`
 It can also be created with the command

```
sudo  singularity  build mpi_container.img mpi_container.def
```

The program mpi_pingppong (source code enclosed: `mpi_pingpong.c` ) is built into the container. It performs a ping-pong test between two processes. 

This container uses the hybrid model, which assumes that MPI is installed on the cluter (to start the job) and installs it in the container itself. Works with OpenMPI.

The second container is an openmp container, including a sample OpenMP programm  openmp_example (source code: ` openmp_example.c`).
It can also be created with the command:

```
sudo  singularity  build openmp_container.img openmp_container.def
```

This container is used bei `submit_openmp_container.go`.


## submit_mpi_containier.go
Submits a mpi-container job to the cluster. It runs to Processes on two nodes
```
job_desc.Min_nodes =uint32(2)
job_desc.Num_tasks = uint32(2)
```
The application  blocks, until the job is completed. The (std) output is written to 
jobid-out.txt, the std- error to jobId-err.txt
```
job_desc.Std_out = ("./%j-out.txt")
job_desc.Std_err = ("./%j-err.txt")
```
## submit_omp_container.go
Submits two openMP jobs  to the cluster and wait, until they are completed. 
Both jobs allocate *one process* for the job, but *two CPUs per task/process* (for multi-threading).
```
job_desc.Num_tasks = uint32(1)
job_desc.Cpus_per_task = uint16(2)

```
The first job reads the environment variable ` SLURM_JOB_CPUS_PER_NODE` and sets the number of openMP threads to exactly the number of cpus that are available per task/process.
```
job_desc.Script+= "export OMP_NUM_THREADS=$SLURM_JOB_CPUS_PER_NODE\n"
```
The second job sets the number of threads to 4 (which is oversuscribing because more threads are started than processes) and executes the same job.
```
job_desc.Script+= "export OMP_NUM_THREADS=4\n"
```

The program waits until both jobs are completed. The results are written to the two outputs files, similiar to `submit_mpi_container.go`

###  submit_mpi_and_update.go
This application is dooing the same as `submit_mpi_container.go`
```
ops.Qos = "shortjobs"
ops.Partition = "short"
```

This situation, can, for example, be created my submitting longer, other jobs bevore in the background (depending on the partion size) and than start this application:

```
./submit_mpi_containier & ./submit_mpi_containier & ./submit_mpi_and_update
```
