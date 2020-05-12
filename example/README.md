This are a few examples to demonstrate the usage of the GO-SLURM API.
These are examples to explain how the SLURM-GO API Works.
There are also some additional examples in the submit folder to show, how to submit jobs, jobs with containers, and modify jobs. The examples can be build with 

**Build the examples**
```
go build example.go
```
**Run it:**

```
./example [parameter]
``` 

Or directly run it with:

```
go run example.go
```


##Description of the examples 

*get_info.go*
Returns the SLUM-Configuration and Version. Print out the configuration. Note: All fields of the config are accessible

*get_all_nodes.go*
Gets the info-struct for all nodes and prints out some information about the nodes. Note: Almost all information about the node is accessible. These are only some examples. See slurm-Documentation for more fields. 

*get_node_info.go*
Returns information about a single node. 
Synthax is:
```
 ./get_node_info nodename
```

*get_all_jobs.go*
Gets all information about all jobs that are currently running, pending or recently finished.  Outputs part of the information. Almost all fields of the jobs are accessible - for more information please refer to the Slurm documentation again.
Functions for the runtime calculation have been implemented additionally.  

*get_user_jobs.go*
Gets all information about all jobs of a specific user that are currently running, pending, or recently finished.  Synthax:
```
./get_user_jobs username
```

*get_job_by_id.go*
Returns information about a specific job, if this job is running, pending, or recently finished. Synthax:
```
./get_job_by_id jobID
```

*get_partions.go*
Print out information about the partitions in the cluster 


