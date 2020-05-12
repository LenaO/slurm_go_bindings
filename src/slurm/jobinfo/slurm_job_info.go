package job_info
/*
#cgo LDFLAGS: -lslurm
#include<stdlib.h>
#include<slurm/slurm.h>


inline uint8_t uint8_ptr(uint8_t* pointer) { 
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
inline int8_t int8_ptr(int8_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
inline uint16_t uint16_ptr(uint16_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
inline int16_t int16_ptr(int16_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
inline uint32_t uint32_ptr(uint32_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
inline int32_t int32_ptr(int32_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
inline uint64_t uint64_ptr(uint64_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
inline int64_t int64_ptr(int16_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
 struct job_info_msg *get_job_info(){
   	struct job_info_msg* job_buffer;
      if(slurm_load_jobs ((time_t) NULL,
		&job_buffer, SHOW_ALL)) {
			return NULL;
		}
	return job_buffer;
 }
  struct job_info_msg *get_single_job_info(uint32_t id){
   	struct job_info_msg* job_buffer;
	if( slurm_load_job (&job_buffer, id, SHOW_DETAIL)) {
		return NULL;
	}
	return job_buffer;
 }
 
 struct job_info* job_from_list(struct job_info_msg *list, int i){
	 return &list->job_array[i];
}
 void free_job_buffer(void* buffer){	

	slurm_free_job_info_msg ((struct job_info_msg*)buffer);
}
 
struct job_info_msg *get_user_job_info(uint32_t id){

	struct job_info_msg* job_buffer;
     	if( slurm_load_job_user(&job_buffer, id, SHOW_DETAIL))
     		return NULL;
 	return job_buffer;

}

int64_t get_job_endtime(int32_t jobid){
	time_t end_time;

	if(slurm_get_end_time (jobid, &end_time))
		return -1;
        else
		return end_time;


}

int char_len(char* c) {
	uint i = 0;
	while(c[i]!='\0') i++;
	return i;
 }


*/
import "C"

import "fmt"
import "os/user"
import "strconv"
import "time"

func Reason_to_string(state uint16) string{
	switch s := C.uint16_t(state); s {
	case C.WAIT_NO_REASON:
		 return "wait no reason"
	case C.WAIT_PRIORITY:
		 return "wait priority"
	case C.WAIT_DEPENDENCY:
		 return "wait dependency"
	case C.WAIT_RESOURCES:
		 return "wait resources"
	case C.WAIT_PART_NODE_LIMIT:
		 return "wait part node limit"
	case C.WAIT_PART_TIME_LIMIT:
		 return "wait part time limit"
	case C.WAIT_PART_DOWN:
		 return "wait part down"
	case C.WAIT_PART_INACTIVE:
		 return "wait part inactive"
	case C.WAIT_HELD:
		 return "wait held"
	case C.WAIT_TIME:
		 return "wait time"
	case C.WAIT_LICENSES:
		 return "wait licenses"
	case C.WAIT_ASSOC_JOB_LIMIT:
		 return "wait assoc job limit"
	case C.WAIT_ASSOC_RESOURCE_LIMIT:
		 return "wait assoc resource limit"
	case C.WAIT_ASSOC_TIME_LIMIT:
		 return "wait assoc time limit"
	case C.WAIT_RESERVATION:
		 return "wait reservation"
	case C.WAIT_NODE_NOT_AVAIL:
		 return "wait node not avail"
	case C.WAIT_HELD_USER:
		 return "wait held user"
	case C.WAIT_FRONT_END:
		 return "wait front end"
	case C.FAIL_DOWN_PARTITION:
		 return "fail down partition"
	case C.FAIL_DOWN_NODE:
		 return "fail down node"
	case C.FAIL_BAD_CONSTRAINTS:
		 return "fail bad constraints"
	case C.FAIL_SYSTEM:
		 return "fail system"
	case C.FAIL_LAUNCH:
		 return "fail launch"
	case C.FAIL_EXIT_CODE:
		 return "fail exit code"
	case C.FAIL_TIMEOUT:
		 return "fail timeout"
	case C.FAIL_INACTIVE_LIMIT:
		 return "fail inactive limit"
	case C.FAIL_ACCOUNT:
		 return "fail account"
	case C.FAIL_QOS:
		 return "fail qos"
	case C.WAIT_QOS_THRES:
		 return "wait qos thres"
	case C.WAIT_QOS_JOB_LIMIT:
		 return "wait qos job limit"
	case C.WAIT_QOS_RESOURCE_LIMIT:
		 return "wait qos resource limit"
	case C.WAIT_QOS_TIME_LIMIT:
		 return "wait qos time limit"
	case C.WAIT_BLOCK_MAX_ERR:
		 return "wait block max err"
	case C.WAIT_BLOCK_D_ACTION:
		 return "wait block d action"
	case C.WAIT_CLEANING:
		 return "wait cleaning"
	case C.WAIT_PROLOG:
		 return "wait prolog"
	case C.WAIT_QOS:
		 return "wait qos"
	case C.WAIT_ACCOUNT:
		 return "wait account"
	case C.WAIT_DEP_INVALID:
		 return "wait dep invalid"
	case C.WAIT_QOS_GRP_CPU:
		 return "wait qos grp cpu"
	case C.WAIT_QOS_GRP_CPU_MIN:
		 return "wait qos grp cpu min"
	case C.WAIT_QOS_GRP_CPU_RUN_MIN:
		 return "wait qos grp cpu run min"
	case C.WAIT_QOS_GRP_JOB:
		 return "wait qos grp job"
	case C.WAIT_QOS_GRP_MEM:
		 return "wait qos grp mem"
	case C.WAIT_QOS_GRP_NODE:
		 return "wait qos grp node"
	case C.WAIT_QOS_GRP_SUB_JOB:
		 return "wait qos grp sub job"
	case C.WAIT_QOS_GRP_WALL:
		 return "wait qos grp wall"
	case C.WAIT_QOS_MAX_CPU_PER_JOB:
		 return "wait qos max cpu per job"
	case C.WAIT_QOS_MAX_CPU_MINS_PER_JOB:
		 return "wait qos max cpu mins per job"
	case C.WAIT_QOS_MAX_NODE_PER_JOB:
		 return "wait qos max node per job"
	case C.WAIT_QOS_MAX_WALL_PER_JOB:
		 return "wait qos max wall per job"
	case C.WAIT_QOS_MAX_CPU_PER_USER:
		 return "wait qos max cpu per user"
	case C.WAIT_QOS_MAX_JOB_PER_USER:
		 return "wait qos max job per user"
	case C.WAIT_QOS_MAX_NODE_PER_USER:
		 return "wait qos max node per user"
	case C.WAIT_QOS_MAX_SUB_JOB:
		 return "wait qos max sub job"
	case C.WAIT_QOS_MIN_CPU:
		 return "wait qos min cpu"
	case C.WAIT_ASSOC_GRP_CPU:
		 return "wait assoc grp cpu"
	case C.WAIT_ASSOC_GRP_CPU_MIN:
		 return "wait assoc grp cpu min"
	case C.WAIT_ASSOC_GRP_CPU_RUN_MIN:
		 return "wait assoc grp cpu run min"
	case C.WAIT_ASSOC_GRP_JOB:
		 return "wait assoc grp job"
	case C.WAIT_ASSOC_GRP_MEM:
		 return "wait assoc grp mem"
	case C.WAIT_ASSOC_GRP_NODE:
		 return "wait assoc grp node"
	case C.WAIT_ASSOC_GRP_SUB_JOB:
		 return "wait assoc grp sub job"
	case C.WAIT_ASSOC_GRP_WALL:
		 return "wait assoc grp wall"
	case C.WAIT_ASSOC_MAX_JOBS:
		 return "wait assoc max jobs"
	case C.WAIT_ASSOC_MAX_CPU_PER_JOB:
		 return "wait assoc max cpu per job"
	case C.WAIT_ASSOC_MAX_CPU_MINS_PER_JOB:
		 return "wait assoc max cpu mins per job"
	case C.WAIT_ASSOC_MAX_NODE_PER_JOB:
		 return "wait assoc max node per job"
	case C.WAIT_ASSOC_MAX_WALL_PER_JOB:
		 return "wait assoc max wall per job"
	case C.WAIT_ASSOC_MAX_SUB_JOB:
		 return "wait assoc max sub job"
	case C.WAIT_MAX_REQUEUE:
		 return "wait max requeue"
	case C.WAIT_ARRAY_TASK_LIMIT:
		 return "wait array task limit"
	case C.WAIT_BURST_BUFFER_RESOURCE:
		 return "wait burst buffer resource"
	case C.WAIT_BURST_BUFFER_STAGING:
		 return "wait burst buffer staging"
	case C.FAIL_BURST_BUFFER_OP:
		 return "fail burst buffer op"
	case C.WAIT_POWER_NOT_AVAIL:
		 return "wait power not avail"
	case C.WAIT_POWER_RESERVED:
		 return "wait power reserved"
	case C.WAIT_ASSOC_GRP_UNK:
		 return "wait assoc grp unk"
	case C.WAIT_ASSOC_GRP_UNK_MIN:
		 return "wait assoc grp unk min"
	case C.WAIT_ASSOC_GRP_UNK_RUN_MIN:
		 return "wait assoc grp unk run min"
	case C.WAIT_ASSOC_MAX_UNK_PER_JOB:
		 return "wait assoc max unk per job"
	case C.WAIT_ASSOC_MAX_UNK_PER_NODE:
		 return "wait assoc max unk per node"
	case C.WAIT_ASSOC_MAX_UNK_MINS_PER_JOB:
		 return "wait assoc max unk mins per job"
	case C.WAIT_ASSOC_MAX_CPU_PER_NODE:
		 return "wait assoc max cpu per node"
	case C.WAIT_ASSOC_GRP_MEM_MIN:
		 return "wait assoc grp mem min"
	case C.WAIT_ASSOC_GRP_MEM_RUN_MIN:
		 return "wait assoc grp mem run min"
	case C.WAIT_ASSOC_MAX_MEM_PER_JOB:
		 return "wait assoc max mem per job"
	case C.WAIT_ASSOC_MAX_MEM_PER_NODE:
		 return "wait assoc max mem per node"
	case C.WAIT_ASSOC_MAX_MEM_MINS_PER_JOB:
		 return "wait assoc max mem mins per job"
	case C.WAIT_ASSOC_GRP_NODE_MIN:
		 return "wait assoc grp node min"
	case C.WAIT_ASSOC_GRP_NODE_RUN_MIN:
		 return "wait assoc grp node run min"
	case C.WAIT_ASSOC_MAX_NODE_MINS_PER_JOB:
		 return "wait assoc max node mins per job"
	case C.WAIT_ASSOC_GRP_ENERGY:
		 return "wait assoc grp energy"
	case C.WAIT_ASSOC_GRP_ENERGY_MIN:
		 return "wait assoc grp energy min"
	case C.WAIT_ASSOC_GRP_ENERGY_RUN_MIN:
		 return "wait assoc grp energy run min"
	case C.WAIT_ASSOC_MAX_ENERGY_PER_JOB:
		 return "wait assoc max energy per job"
	case C.WAIT_ASSOC_MAX_ENERGY_PER_NODE:
		 return "wait assoc max energy per node"
	case C.WAIT_ASSOC_MAX_ENERGY_MINS_PER_JOB:
		 return "wait assoc max energy mins per job"
	case C.WAIT_ASSOC_GRP_GRES:
		 return "wait assoc grp gres"
	case C.WAIT_ASSOC_GRP_GRES_MIN:
		 return "wait assoc grp gres min"
	case C.WAIT_ASSOC_GRP_GRES_RUN_MIN:
		 return "wait assoc grp gres run min"
	case C.WAIT_ASSOC_MAX_GRES_PER_JOB:
		 return "wait assoc max gres per job"
	case C.WAIT_ASSOC_MAX_GRES_PER_NODE:
		 return "wait assoc max gres per node"
	case C.WAIT_ASSOC_MAX_GRES_MINS_PER_JOB:
		 return "wait assoc max gres mins per job"
	case C.WAIT_ASSOC_GRP_LIC:
		 return "wait assoc grp lic"
	case C.WAIT_ASSOC_GRP_LIC_MIN:
		 return "wait assoc grp lic min"
	case C.WAIT_ASSOC_GRP_LIC_RUN_MIN:
		 return "wait assoc grp lic run min"
	case C.WAIT_ASSOC_MAX_LIC_PER_JOB:
		 return "wait assoc max lic per job"
	case C.WAIT_ASSOC_MAX_LIC_MINS_PER_JOB:
		 return "wait assoc max lic mins per job"
	case C.WAIT_ASSOC_GRP_BB:
		 return "wait assoc grp bb"
	case C.WAIT_ASSOC_GRP_BB_MIN:
		 return "wait assoc grp bb min"
	case C.WAIT_ASSOC_GRP_BB_RUN_MIN:
		 return "wait assoc grp bb run min"
	case C.WAIT_ASSOC_MAX_BB_PER_JOB:
		 return "wait assoc max bb per job"
	case C.WAIT_ASSOC_MAX_BB_PER_NODE:
		 return "wait assoc max bb per node"
	case C.WAIT_ASSOC_MAX_BB_MINS_PER_JOB:
		 return "wait assoc max bb mins per job"
	case C.WAIT_QOS_GRP_UNK:
		 return "wait qos grp unk"
	case C.WAIT_QOS_GRP_UNK_MIN:
		 return "wait qos grp unk min"
	case C.WAIT_QOS_GRP_UNK_RUN_MIN:
		 return "wait qos grp unk run min"
	case C.WAIT_QOS_MAX_UNK_PER_JOB:
		 return "wait qos max unk per job"
	case C.WAIT_QOS_MAX_UNK_PER_NODE:
		 return "wait qos max unk per node"
	case C.WAIT_QOS_MAX_UNK_PER_USER:
		 return "wait qos max unk per user"
	case C.WAIT_QOS_MAX_UNK_MINS_PER_JOB:
		 return "wait qos max unk mins per job"
	case C.WAIT_QOS_MIN_UNK:
		 return "wait qos min unk"
	case C.WAIT_QOS_MAX_CPU_PER_NODE:
		 return "wait qos max cpu per node"
	case C.WAIT_QOS_GRP_MEM_MIN:
		 return "wait qos grp mem min"
	case C.WAIT_QOS_GRP_MEM_RUN_MIN:
		 return "wait qos grp mem run min"
	case C.WAIT_QOS_MAX_MEM_MINS_PER_JOB:
		 return "wait qos max mem mins per job"
	case C.WAIT_QOS_MAX_MEM_PER_JOB:
		 return "wait qos max mem per job"
	case C.WAIT_QOS_MAX_MEM_PER_NODE:
		 return "wait qos max mem per node"
	case C.WAIT_QOS_MAX_MEM_PER_USER:
		 return "wait qos max mem per user"
	case C.WAIT_QOS_MIN_MEM:
		 return "wait qos min mem"
	case C.WAIT_QOS_GRP_ENERGY:
		 return "wait qos grp energy"
	case C.WAIT_QOS_GRP_ENERGY_MIN:
		 return "wait qos grp energy min"
	case C.WAIT_QOS_GRP_ENERGY_RUN_MIN:
		 return "wait qos grp energy run min"
	case C.WAIT_QOS_MAX_ENERGY_PER_JOB:
		 return "wait qos max energy per job"
	case C.WAIT_QOS_MAX_ENERGY_PER_NODE:
		 return "wait qos max energy per node"
	case C.WAIT_QOS_MAX_ENERGY_PER_USER:
		 return "wait qos max energy per user"
	case C.WAIT_QOS_MAX_ENERGY_MINS_PER_JOB:
		 return "wait qos max energy mins per job"
	case C.WAIT_QOS_MIN_ENERGY:
		 return "wait qos min energy"
	case C.WAIT_QOS_GRP_NODE_MIN:
		 return "wait qos grp node min"
	case C.WAIT_QOS_GRP_NODE_RUN_MIN:
		 return "wait qos grp node run min"
	case C.WAIT_QOS_MAX_NODE_MINS_PER_JOB:
		 return "wait qos max node mins per job"
	case C.WAIT_QOS_MIN_NODE:
		 return "wait qos min node"
	case C.WAIT_QOS_GRP_GRES:
		 return "wait qos grp gres"
	case C.WAIT_QOS_GRP_GRES_MIN:
		 return "wait qos grp gres min"
	case C.WAIT_QOS_GRP_GRES_RUN_MIN:
		 return "wait qos grp gres run min"
	case C.WAIT_QOS_MAX_GRES_PER_JOB:
		 return "wait qos max gres per job"
	case C.WAIT_QOS_MAX_GRES_PER_NODE:
		 return "wait qos max gres per node"
	case C.WAIT_QOS_MAX_GRES_PER_USER:
		 return "wait qos max gres per user"
	case C.WAIT_QOS_MAX_GRES_MINS_PER_JOB:
		 return "wait qos max gres mins per job"
	case C.WAIT_QOS_MIN_GRES:
		 return "wait qos min gres"
	case C.WAIT_QOS_GRP_LIC:
		 return "wait qos grp lic"
	case C.WAIT_QOS_GRP_LIC_MIN:
		 return "wait qos grp lic min"
	case C.WAIT_QOS_GRP_LIC_RUN_MIN:
		 return "wait qos grp lic run min"
	case C.WAIT_QOS_MAX_LIC_PER_JOB:
		 return "wait qos max lic per job"
	case C.WAIT_QOS_MAX_LIC_PER_USER:
		 return "wait qos max lic per user"
	case C.WAIT_QOS_MAX_LIC_MINS_PER_JOB:
		 return "wait qos max lic mins per job"
	case C.WAIT_QOS_MIN_LIC:
		 return "wait qos min lic"
	case C.WAIT_QOS_GRP_BB:
		 return "wait qos grp bb"
	case C.WAIT_QOS_GRP_BB_MIN:
		 return "wait qos grp bb min"
	case C.WAIT_QOS_GRP_BB_RUN_MIN:
		 return "wait qos grp bb run min"
	case C.WAIT_QOS_MAX_BB_PER_JOB:
		 return "wait qos max bb per job"
	case C.WAIT_QOS_MAX_BB_PER_NODE:
		 return "wait qos max bb per node"
	case C.WAIT_QOS_MAX_BB_PER_USER:
		 return "wait qos max bb per user"
	case C.WAIT_QOS_MAX_BB_MINS_PER_JOB:
		 return "wait qos max bb mins per job"
	case C.WAIT_QOS_MIN_BB:
		 return "wait qos min bb"
	case C.FAIL_DEADLINE:
		 return "fail deadline"
	case C.WAIT_QOS_MAX_BB_PER_ACCT:
		 return "wait qos max bb per acct"
	case C.WAIT_QOS_MAX_CPU_PER_ACCT:
		 return "wait qos max cpu per acct"
	case C.WAIT_QOS_MAX_ENERGY_PER_ACCT:
		 return "wait qos max energy per acct"
	case C.WAIT_QOS_MAX_GRES_PER_ACCT:
		 return "wait qos max gres per acct"
	case C.WAIT_QOS_MAX_NODE_PER_ACCT:
		 return "wait qos max node per acct"
	case C.WAIT_QOS_MAX_LIC_PER_ACCT:
		 return "wait qos max lic per acct"
	case C.WAIT_QOS_MAX_MEM_PER_ACCT:
		 return "wait qos max mem per acct"
	case C.WAIT_QOS_MAX_UNK_PER_ACCT:
		 return "wait qos max unk per acct"
	case C.WAIT_QOS_MAX_JOB_PER_ACCT:
		 return "wait qos max job per acct"
	case C.WAIT_QOS_MAX_SUB_JOB_PER_ACCT:
		 return "wait qos max sub job per acct"
	case C.WAIT_PART_CONFIG:
		 return "wait part config"
	case C.WAIT_ACCOUNT_POLICY:
		 return "wait account policy"
	case C.WAIT_FED_JOB_LOCK:
		 return "wait fed job lock"
	case C.FAIL_OOM:
		 return "fail oom"
	case C.WAIT_PN_MEM_LIMIT:
		 return "wait pn mem limit"
	case C.WAIT_ASSOC_GRP_BILLING:
		 return "wait assoc grp billing"
	case C.WAIT_ASSOC_GRP_BILLING_MIN:
		 return "wait assoc grp billing min"
	case C.WAIT_ASSOC_GRP_BILLING_RUN_MIN:
		 return "wait assoc grp billing run min"
	case C.WAIT_ASSOC_MAX_BILLING_PER_JOB:
		 return "wait assoc max billing per job"
	case C.WAIT_ASSOC_MAX_BILLING_PER_NODE:
		 return "wait assoc max billing per node"
	case C.WAIT_ASSOC_MAX_BILLING_MINS_PER_JOB:
		 return "wait assoc max billing mins per job"
	case C.WAIT_QOS_GRP_BILLING:
		 return "wait qos grp billing"
	case C.WAIT_QOS_GRP_BILLING_MIN:
		 return "wait qos grp billing min"
	case C.WAIT_QOS_GRP_BILLING_RUN_MIN:
		 return "wait qos grp billing run min"
	case C.WAIT_QOS_MAX_BILLING_PER_JOB:
		 return "wait qos max billing per job"
	case C.WAIT_QOS_MAX_BILLING_PER_NODE:
		 return "wait qos max billing per node"
	case C.WAIT_QOS_MAX_BILLING_PER_USER:
		 return "wait qos max billing per user"
	case C.WAIT_QOS_MAX_BILLING_MINS_PER_JOB:
		 return "wait qos max billing mins per job"
	case C.WAIT_QOS_MAX_BILLING_PER_ACCT:
		 return "wait qos max billing per acct"
	case C.WAIT_QOS_MIN_BILLING:
		 return "wait qos min billing"
	case C.WAIT_RESV_DELETED:
		 return "wait resv deleted"
	}
	return "unkown reason"

}
func state_to_string(state uint32) string{
	switch s := C.uint32_t(state); s {
		case C.JOB_PENDING:
			return "Pending"		/* queued waiting for initiation */
		case C.JOB_RUNNING:
			return "Running"		/* allocated resources and executing */
		case C.JOB_SUSPENDED:
			return "Suspended"		/* allocated resources, execution suspended */
	        case C.JOB_COMPLETE:
			return "Complete"		/* completed execution successfully */
		case C.JOB_CANCELLED:
			return "Cancelled"		/* cancelled by user */
		case C.JOB_FAILED:
			return "Failed"			/* completed execution unsuccessfully */
		case C.JOB_TIMEOUT:
			return "Timeout"		/* terminated on reaching time limit */
		case C.JOB_NODE_FAIL:
			return "Node Fail"		/* terminated on node failure */
		case C.JOB_PREEMPTED:
			return "Preempted"		/* terminated due to preemption */
		case C.JOB_BOOT_FAIL:
			return "Boot Fail"	/* terminated due to node boot failure */
		case C.JOB_DEADLINE:
			return "Term. Deadline"    /* terminated on deadline */
		case C.JOB_OOM:
			return "Out of memory"	      /* experienced out of memory error */
	}

		return "Unknow status";		/* not a real state, last entry in table */
}


type Job_info struct {
	Account string;
	Accrue_time int64;
	Admin_comment string;
	Alloc_node string;
	Alloc_sid uint32;
	 //array_bitmap void;
	Array_job_id uint32;
	Array_task_id uint32;
	Array_max_tasks uint32;
	Array_task_str string;
	Assoc_id uint32;
	Batch_features string;
	Batch_flag uint16;
	Batch_host string;
	Bitflags uint32;
	Boards_per_node uint16;
	Burst_buffer string;
	Burst_buffer_state string;
	Cluster string;
	Cluster_features string;
	Command string;
	Comment string;
	Contiguous uint16;
	Core_spec uint16;
	Cores_per_socket uint16;
	Cpus_per_task uint16;
	Cpu_freq_min uint32;
	Cpu_freq_max uint32;
	Cpu_freq_gov uint32;
	Cpus_per_tres string;
	Deadline int64;
	Delay_boot uint32;
	Dependency string;
	Derived_ec uint32;
	Eligible_time int64;
	End_time int64;
	Exc_nodes string;
	Exc_node_inx int32;
	Exit_code uint32;
	Features string;
	Fed_origin_str string;
	Fed_siblings_active uint64;
	Fed_siblings_active_str string;
	Fed_siblings_viable uint64;
	Fed_siblings_viable_str string;
	Gres_detail_cnt uint32;
	 //gres_detail_str char**;
	Group_id uint32;
	Job_id uint32;
	Job_state uint32;
	Job_stateS string;
	Last_sched_eval int64;
	Licenses string;
	Max_cpus uint32;
	Max_nodes uint32;
	Mcs_label string;
	Mem_per_tres string;
	Name string;
	Network string;
	Nodes string;
	Nice uint32;
	Node_inx int32;
	Ntasks_per_core uint16;
	Ntasks_per_node uint16;
	Ntasks_per_socket uint16;
	Ntasks_per_board uint16;
	Num_cpus uint32;
	Num_nodes uint32;
	Num_tasks uint32;
	Pack_job_id uint32;
	Pack_job_id_set string;
	Pack_job_offset uint32;
	Partition string;
	Pn_min_memory uint64;
	Pn_min_cpus uint16;
	Pn_min_tmp_disk uint32;
	Power_flags uint8;
	Preempt_time int64;
	Preemptable_time int64;
	Pre_sus_time int64;
	Priority uint32;
	Profile uint32;
	Qos string;
	Reboot uint8;
	Req_nodes string;
	Req_node_inx int32;
	Req_switch uint32;
	Requeue uint16;
	Resize_time int64;
	Restart_cnt uint16;
	Resv_name string;
	Sched_nodes string;
	Shared uint16;
	Show_flags uint16;
	Site_factor uint32;
	Sockets_per_board uint16;
	Sockets_per_node uint16;
	Start_time int64;
	Start_protocol_ver uint16;
	State_desc string;
	State_reason uint16;
	Std_err string;
	Std_in string;
	Std_out string;
	Submit_time int64;
	Suspend_time int64;
	System_comment string;
	Time_limit uint32;
	Time_min uint32;
	Threads_per_core uint16;
	Tres_bind string;
	Tres_freq string;
	Tres_per_job string;
	Tres_per_node string;
	Tres_per_socket string;
	Tres_per_task string;
	Tres_req_str string;
	Tres_alloc_str string;
	User_id uint32;
	User_name string;
	 //accurate void;
	Wait4switch uint32;
	Wckey string;
	Work_dir string;
}

func Job_info_convert_c_to_go(c_struct *C.struct_job_info) Job_info{
	var go_struct Job_info

	go_struct.Account = C.GoString(c_struct.account)
	go_struct.Accrue_time = int64(c_struct.accrue_time)
	go_struct.Admin_comment = C.GoString(c_struct.admin_comment)
	go_struct.Alloc_node = C.GoString(c_struct.alloc_node)
	go_struct.Alloc_sid = uint32(c_struct.alloc_sid)
	go_struct.Array_job_id = uint32(c_struct.array_job_id)
	go_struct.Array_task_id = uint32(c_struct.array_task_id)
	go_struct.Array_max_tasks = uint32(c_struct.array_max_tasks)
	go_struct.Array_task_str = C.GoString(c_struct.array_task_str)
	go_struct.Assoc_id = uint32(c_struct.assoc_id)
	go_struct.Batch_features = C.GoString(c_struct.batch_features)
	go_struct.Batch_flag = uint16(c_struct.batch_flag)
	go_struct.Batch_host = C.GoString(c_struct.batch_host)
	go_struct.Bitflags = uint32(c_struct.bitflags)
	go_struct.Boards_per_node = uint16(c_struct.boards_per_node)
	go_struct.Burst_buffer = C.GoString(c_struct.burst_buffer)
	go_struct.Burst_buffer_state = C.GoString(c_struct.burst_buffer_state)
	go_struct.Cluster = C.GoString(c_struct.cluster)
	go_struct.Cluster_features = C.GoString(c_struct.cluster_features)
	go_struct.Command = C.GoString(c_struct.command)
	go_struct.Comment = C.GoString(c_struct.comment)
	go_struct.Contiguous = uint16(c_struct.contiguous)
	go_struct.Core_spec = uint16(c_struct.core_spec)
	go_struct.Cores_per_socket = uint16(c_struct.cores_per_socket)
	go_struct.Cpus_per_task = uint16(c_struct.cpus_per_task)
	go_struct.Cpu_freq_min = uint32(c_struct.cpu_freq_min)
	go_struct.Cpu_freq_max = uint32(c_struct.cpu_freq_max)
	go_struct.Cpu_freq_gov = uint32(c_struct.cpu_freq_gov)
	go_struct.Cpus_per_tres = C.GoString(c_struct.cpus_per_tres)
	go_struct.Deadline = int64(c_struct.deadline)
	go_struct.Delay_boot = uint32(c_struct.delay_boot)
	go_struct.Dependency = C.GoString(c_struct.dependency)
	go_struct.Derived_ec = uint32(c_struct.derived_ec)
	go_struct.Eligible_time = int64(c_struct.eligible_time)
	go_struct.End_time = int64(c_struct.end_time)
	go_struct.Exc_nodes = C.GoString(c_struct.exc_nodes)
	go_struct.Exc_node_inx = int32(C.int32_ptr(c_struct.exc_node_inx))
	go_struct.Exit_code = uint32(c_struct.exit_code)
	go_struct.Features = C.GoString(c_struct.features)
	go_struct.Fed_origin_str = C.GoString(c_struct.fed_origin_str)
	go_struct.Fed_siblings_active = uint64(c_struct.fed_siblings_active)
	go_struct.Fed_siblings_active_str = C.GoString(c_struct.fed_siblings_active_str)
	go_struct.Fed_siblings_viable = uint64(c_struct.fed_siblings_viable)
	go_struct.Fed_siblings_viable_str = C.GoString(c_struct.fed_siblings_viable_str)
	go_struct.Gres_detail_cnt = uint32(c_struct.gres_detail_cnt)
	go_struct.Group_id = uint32(c_struct.group_id)
	go_struct.Job_id = uint32(c_struct.job_id)
	go_struct.Job_state = uint32(c_struct.job_state)
	go_struct.Job_stateS = state_to_string(uint32(c_struct.job_state))
	go_struct.Last_sched_eval = int64(c_struct.last_sched_eval)
	go_struct.Licenses = C.GoString(c_struct.licenses)
	go_struct.Max_cpus = uint32(c_struct.max_cpus)
	go_struct.Max_nodes = uint32(c_struct.max_nodes)
	go_struct.Mcs_label = C.GoString(c_struct.mcs_label)
	go_struct.Mem_per_tres = C.GoString(c_struct.mem_per_tres)
	go_struct.Name = C.GoString(c_struct.name)
	go_struct.Network = C.GoString(c_struct.network)
	go_struct.Nodes = C.GoString(c_struct.nodes)
	go_struct.Nice = uint32(c_struct.nice)
	go_struct.Node_inx = int32(C.int32_ptr(c_struct.node_inx))
	go_struct.Ntasks_per_core = uint16(c_struct.ntasks_per_core)
	go_struct.Ntasks_per_node = uint16(c_struct.ntasks_per_node)
	go_struct.Ntasks_per_socket = uint16(c_struct.ntasks_per_socket)
	go_struct.Ntasks_per_board = uint16(c_struct.ntasks_per_board)
	go_struct.Num_cpus = uint32(c_struct.num_cpus)
	go_struct.Num_nodes = uint32(c_struct.num_nodes)
	go_struct.Num_tasks = uint32(c_struct.num_tasks)
	go_struct.Pack_job_id = uint32(c_struct.pack_job_id)
	go_struct.Pack_job_id_set = C.GoString(c_struct.pack_job_id_set)
	go_struct.Pack_job_offset = uint32(c_struct.pack_job_offset)
	go_struct.Partition = C.GoString(c_struct.partition)
	go_struct.Pn_min_memory = uint64(c_struct.pn_min_memory)
	go_struct.Pn_min_cpus = uint16(c_struct.pn_min_cpus)
	go_struct.Pn_min_tmp_disk = uint32(c_struct.pn_min_tmp_disk)
	go_struct.Power_flags = uint8(c_struct.power_flags)
	go_struct.Preempt_time = int64(c_struct.preempt_time)
	go_struct.Preemptable_time = int64(c_struct.preemptable_time)
	go_struct.Pre_sus_time = int64(c_struct.pre_sus_time)
	go_struct.Priority = uint32(c_struct.priority)
	go_struct.Profile = uint32(c_struct.profile)
	go_struct.Qos = C.GoString(c_struct.qos)
	go_struct.Reboot = uint8(c_struct.reboot)
	go_struct.Req_nodes = C.GoString(c_struct.req_nodes)
	go_struct.Req_node_inx = int32(C.int32_ptr(c_struct.req_node_inx))
	go_struct.Req_switch = uint32(c_struct.req_switch)
	go_struct.Requeue = uint16(c_struct.requeue)
	go_struct.Resize_time = int64(c_struct.resize_time)
	go_struct.Restart_cnt = uint16(c_struct.restart_cnt)
	go_struct.Resv_name = C.GoString(c_struct.resv_name)
	go_struct.Sched_nodes = C.GoString(c_struct.sched_nodes)
	go_struct.Shared = uint16(c_struct.shared)
	go_struct.Show_flags = uint16(c_struct.show_flags)
	go_struct.Site_factor = uint32(c_struct.site_factor)
	go_struct.Sockets_per_board = uint16(c_struct.sockets_per_board)
	go_struct.Sockets_per_node = uint16(c_struct.sockets_per_node)
	go_struct.Start_time = int64(c_struct.start_time)
	go_struct.Start_protocol_ver = uint16(c_struct.start_protocol_ver)
	go_struct.State_desc = C.GoString(c_struct.state_desc)
	go_struct.State_reason = uint16(c_struct.state_reason)
	go_struct.Std_err = C.GoString(c_struct.std_err)
	go_struct.Std_in = C.GoString(c_struct.std_in)
	go_struct.Std_out = C.GoString(c_struct.std_out)
	go_struct.Submit_time = int64(c_struct.submit_time)
	go_struct.Suspend_time = int64(c_struct.suspend_time)
	go_struct.System_comment = C.GoString(c_struct.system_comment)
	go_struct.Time_limit = uint32(c_struct.time_limit)
	go_struct.Time_min = uint32(c_struct.time_min)
	go_struct.Threads_per_core = uint16(c_struct.threads_per_core)
	go_struct.Tres_bind = C.GoString(c_struct.tres_bind)
	go_struct.Tres_freq = C.GoString(c_struct.tres_freq)
	go_struct.Tres_per_job = C.GoString(c_struct.tres_per_job)
	go_struct.Tres_per_node = C.GoString(c_struct.tres_per_node)
	go_struct.Tres_per_socket = C.GoString(c_struct.tres_per_socket)
	go_struct.Tres_per_task = C.GoString(c_struct.tres_per_task)
	go_struct.Tres_req_str = C.GoString(c_struct.tres_req_str)
	go_struct.Tres_alloc_str = C.GoString(c_struct.tres_alloc_str)
	go_struct.User_id = uint32(c_struct.user_id)
	go_struct.User_name = C.GoString(c_struct.user_name)
	if len(go_struct.User_name) == 0 {
		tmp_user,_ :=  user.LookupId(strconv.Itoa(int(go_struct.User_id)))
		go_struct.User_name= tmp_user.Username
	}
	go_struct.Wait4switch = uint32(c_struct.wait4switch)
	go_struct.Wckey = C.GoString(c_struct.wckey)
	go_struct.Work_dir = C.GoString(c_struct.work_dir)
	return go_struct
 }

 func Print_Job_info(go_struct Job_info){
	 fmt.Printf("%s:\t %s\n","account", go_struct.Account)
	 fmt.Printf("%s:\t %d\n","accrue time", go_struct.Accrue_time)
	 fmt.Printf("%s:\t %s\n","admin comment", go_struct.Admin_comment)
	 fmt.Printf("%s:\t %s\n","alloc node", go_struct.Alloc_node)
	 fmt.Printf("%s:\t %d\n","alloc sid", go_struct.Alloc_sid)
	 fmt.Printf("%s:\t %d\n","array job id", go_struct.Array_job_id)
	 fmt.Printf("%s:\t %d\n","array task id", go_struct.Array_task_id)
	 fmt.Printf("%s:\t %d\n","array max tasks", go_struct.Array_max_tasks)
	 fmt.Printf("%s:\t %s\n","array task str", go_struct.Array_task_str)
	 fmt.Printf("%s:\t %d\n","assoc id", go_struct.Assoc_id)
	 fmt.Printf("%s:\t %s\n","batch features", go_struct.Batch_features)
	 fmt.Printf("%s:\t %d\n","batch flag", go_struct.Batch_flag)
	 fmt.Printf("%s:\t %s\n","batch host", go_struct.Batch_host)
	 fmt.Printf("%s:\t %d\n","bitflags", go_struct.Bitflags)
	 fmt.Printf("%s:\t %d\n","boards per node", go_struct.Boards_per_node)
	 fmt.Printf("%s:\t %s\n","burst buffer", go_struct.Burst_buffer)
	 fmt.Printf("%s:\t %s\n","burst buffer state", go_struct.Burst_buffer_state)
	 fmt.Printf("%s:\t %s\n","cluster", go_struct.Cluster)
	 fmt.Printf("%s:\t %s\n","cluster features", go_struct.Cluster_features)
	 fmt.Printf("%s:\t %s\n","command", go_struct.Command)
	 fmt.Printf("%s:\t %s\n","comment", go_struct.Comment)
	 fmt.Printf("%s:\t %d\n","contiguous", go_struct.Contiguous)
	 fmt.Printf("%s:\t %d\n","core spec", go_struct.Core_spec)
	 fmt.Printf("%s:\t %d\n","cores per socket", go_struct.Cores_per_socket)
	 fmt.Printf("%s:\t %d\n","cpus per task", go_struct.Cpus_per_task)
	 fmt.Printf("%s:\t %d\n","cpu freq min", go_struct.Cpu_freq_min)
	 fmt.Printf("%s:\t %d\n","cpu freq max", go_struct.Cpu_freq_max)
	 fmt.Printf("%s:\t %d\n","cpu freq gov", go_struct.Cpu_freq_gov)
	 fmt.Printf("%s:\t %s\n","cpus per tres", go_struct.Cpus_per_tres)
	 fmt.Printf("%s:\t %d\n","deadline", go_struct.Deadline)
	 fmt.Printf("%s:\t %d\n","delay boot", go_struct.Delay_boot)
	 fmt.Printf("%s:\t %s\n","dependency", go_struct.Dependency)
	 fmt.Printf("%s:\t %d\n","derived ec", go_struct.Derived_ec)
	 fmt.Printf("%s:\t %d\n","eligible time", go_struct.Eligible_time)
	 fmt.Printf("%s:\t %d\n","end time", go_struct.End_time)
	 fmt.Printf("%s:\t %s\n","exc nodes", go_struct.Exc_nodes)
	 fmt.Printf("%s:\t %d\n","exc node inx", go_struct.Exc_node_inx)
	 fmt.Printf("%s:\t %d\n","exit code", go_struct.Exit_code)
	 fmt.Printf("%s:\t %s\n","features", go_struct.Features)
	 fmt.Printf("%s:\t %s\n","fed origin str", go_struct.Fed_origin_str)
	 fmt.Printf("%s:\t %d\n","fed siblings active", go_struct.Fed_siblings_active)
	 fmt.Printf("%s:\t %s\n","fed siblings active str", go_struct.Fed_siblings_active_str)
	 fmt.Printf("%s:\t %d\n","fed siblings viable", go_struct.Fed_siblings_viable)
	 fmt.Printf("%s:\t %s\n","fed siblings viable str", go_struct.Fed_siblings_viable_str)
	 fmt.Printf("%s:\t %d\n","gres detail cnt", go_struct.Gres_detail_cnt)
	 fmt.Printf("%s:\t %d\n","group id", go_struct.Group_id)
	 fmt.Printf("%s:\t %d\n","job id", go_struct.Job_id)
	 fmt.Printf("%s:\t %d\n","job state", go_struct.Job_state)
	 fmt.Printf("%s:\t %d\n","last sched eval", go_struct.Last_sched_eval)
	 fmt.Printf("%s:\t %s\n","licenses", go_struct.Licenses)
	 fmt.Printf("%s:\t %d\n","max cpus", go_struct.Max_cpus)
	 fmt.Printf("%s:\t %d\n","max nodes", go_struct.Max_nodes)
	 fmt.Printf("%s:\t %s\n","mcs label", go_struct.Mcs_label)
	 fmt.Printf("%s:\t %s\n","mem per tres", go_struct.Mem_per_tres)
	 fmt.Printf("%s:\t %s\n","name", go_struct.Name)
	 fmt.Printf("%s:\t %s\n","network", go_struct.Network)
	 fmt.Printf("%s:\t %s\n","nodes", go_struct.Nodes)
	 fmt.Printf("%s:\t %d\n","nice", go_struct.Nice)
	 fmt.Printf("%s:\t %d\n","node inx", go_struct.Node_inx)
	 fmt.Printf("%s:\t %d\n","ntasks per core", go_struct.Ntasks_per_core)
	 fmt.Printf("%s:\t %d\n","ntasks per node", go_struct.Ntasks_per_node)
	 fmt.Printf("%s:\t %d\n","ntasks per socket", go_struct.Ntasks_per_socket)
	 fmt.Printf("%s:\t %d\n","ntasks per board", go_struct.Ntasks_per_board)
	 fmt.Printf("%s:\t %d\n","num cpus", go_struct.Num_cpus)
	 fmt.Printf("%s:\t %d\n","num nodes", go_struct.Num_nodes)
	 fmt.Printf("%s:\t %d\n","num tasks", go_struct.Num_tasks)
	 fmt.Printf("%s:\t %d\n","pack job id", go_struct.Pack_job_id)
	 fmt.Printf("%s:\t %s\n","pack job id set", go_struct.Pack_job_id_set)
	 fmt.Printf("%s:\t %d\n","pack job offset", go_struct.Pack_job_offset)
	 fmt.Printf("%s:\t %s\n","partition", go_struct.Partition)
	 fmt.Printf("%s:\t %d\n","pn min memory", go_struct.Pn_min_memory)
	 fmt.Printf("%s:\t %d\n","pn min cpus", go_struct.Pn_min_cpus)
	 fmt.Printf("%s:\t %d\n","pn min tmp disk", go_struct.Pn_min_tmp_disk)
	 fmt.Printf("%s:\t %d\n","power flags", go_struct.Power_flags)
	 fmt.Printf("%s:\t %d\n","preempt time", go_struct.Preempt_time)
	 fmt.Printf("%s:\t %d\n","preemptable time", go_struct.Preemptable_time)
	 fmt.Printf("%s:\t %d\n","pre sus time", go_struct.Pre_sus_time)
	 fmt.Printf("%s:\t %d\n","priority", go_struct.Priority)
	 fmt.Printf("%s:\t %d\n","profile", go_struct.Profile)
	 fmt.Printf("%s:\t %s\n","qos", go_struct.Qos)
	 fmt.Printf("%s:\t %d\n","reboot", go_struct.Reboot)
	 fmt.Printf("%s:\t %s\n","req nodes", go_struct.Req_nodes)
	 fmt.Printf("%s:\t %d\n","req node inx", go_struct.Req_node_inx)
	 fmt.Printf("%s:\t %d\n","req switch", go_struct.Req_switch)
	 fmt.Printf("%s:\t %d\n","requeue", go_struct.Requeue)
	 fmt.Printf("%s:\t %d\n","resize time", go_struct.Resize_time)
	 fmt.Printf("%s:\t %d\n","restart cnt", go_struct.Restart_cnt)
	 fmt.Printf("%s:\t %s\n","resv name", go_struct.Resv_name)
	 fmt.Printf("%s:\t %s\n","sched nodes", go_struct.Sched_nodes)
	 fmt.Printf("%s:\t %d\n","shared", go_struct.Shared)
	 fmt.Printf("%s:\t %d\n","show flags", go_struct.Show_flags)
	 fmt.Printf("%s:\t %d\n","site factor", go_struct.Site_factor)
	 fmt.Printf("%s:\t %d\n","sockets per board", go_struct.Sockets_per_board)
	 fmt.Printf("%s:\t %d\n","sockets per node", go_struct.Sockets_per_node)
	 fmt.Printf("%s:\t %d\n","start time", go_struct.Start_time)
	 fmt.Printf("%s:\t %d\n","start protocol ver", go_struct.Start_protocol_ver)
	 fmt.Printf("%s:\t %s\n","state desc", go_struct.State_desc)
	 fmt.Printf("%s:\t %d\n","state reason", go_struct.State_reason)
	 fmt.Printf("%s:\t %s\n","std err", go_struct.Std_err)
	 fmt.Printf("%s:\t %s\n","std in", go_struct.Std_in)
	 fmt.Printf("%s:\t %s\n","std out", go_struct.Std_out)
	 fmt.Printf("%s:\t %d\n","submit time", go_struct.Submit_time)
	 fmt.Printf("%s:\t %d\n","suspend time", go_struct.Suspend_time)
	 fmt.Printf("%s:\t %s\n","system comment", go_struct.System_comment)
	 fmt.Printf("%s:\t %d\n","time limit", go_struct.Time_limit)
	 fmt.Printf("%s:\t %d\n","time min", go_struct.Time_min)
	 fmt.Printf("%s:\t %d\n","threads per core", go_struct.Threads_per_core)
	 fmt.Printf("%s:\t %s\n","tres bind", go_struct.Tres_bind)
	 fmt.Printf("%s:\t %s\n","tres freq", go_struct.Tres_freq)
	 fmt.Printf("%s:\t %s\n","tres per job", go_struct.Tres_per_job)
	 fmt.Printf("%s:\t %s\n","tres per node", go_struct.Tres_per_node)
	 fmt.Printf("%s:\t %s\n","tres per socket", go_struct.Tres_per_socket)
	 fmt.Printf("%s:\t %s\n","tres per task", go_struct.Tres_per_task)
	 fmt.Printf("%s:\t %s\n","tres req str", go_struct.Tres_req_str)
	 fmt.Printf("%s:\t %s\n","tres alloc str", go_struct.Tres_alloc_str)
	 fmt.Printf("%s:\t %d\n","user id", go_struct.User_id)
	 fmt.Printf("%s:\t %s\n","user name", go_struct.User_name)
	 fmt.Printf("%s:\t %d\n","wait4switch", go_struct.Wait4switch)
	 fmt.Printf("%s:\t %s\n","wckey", go_struct.Wckey)
	 fmt.Printf("%s:\t %s\n","work dir", go_struct.Work_dir)
}

type Job_info_msg struct {
	Last_update int64;
	Record_count uint32;
	Error_code uint32;
	Job_list  []Job_info;
}

func Get_job_runtime(job Job_info) time.Duration{

	start_time := time.Unix(job.Start_time,0)
	current_time := time.Now()
	diff := current_time.Sub(start_time).Round(time.Second)
	if int64(diff)<0 {
		return 0
	}
	end_time := time.Unix(job.End_time,0)
	diff2 := current_time.Sub(end_time)
	if int64(diff2)<0 {
		return diff
	}
	return end_time.Sub(start_time)
}

func Get_job_endtime (id uint32 ) time.Time {
	c_time := C.get_job_endtime(C.int32_t(id))
	return time.Unix(int64(c_time),0)

}

func Get_all_jobs() Job_info_msg {
	var go_job_buffer Job_info_msg
	c_job_buffer := C.get_job_info()
	if c_job_buffer == nil {
		go_job_buffer.Last_update = int64(0)
		go_job_buffer.Record_count = uint32(0)
		go_job_buffer.Error_code = uint32(C.slurm_get_errno())
		return go_job_buffer
	}

	go_job_buffer.Last_update = int64(c_job_buffer.last_update)
	go_job_buffer.Record_count = uint32(c_job_buffer.record_count)
	go_job_buffer.Job_list =make([]Job_info,c_job_buffer.record_count, c_job_buffer.record_count)
	for i:=uint32(0); i<go_job_buffer.Record_count; i++ {
		job := C.job_from_list(c_job_buffer,  C.int(i))
		go_job :=  Job_info_convert_c_to_go(job)
		go_job_buffer.Job_list[i]=go_job
	}
	C.slurm_free_job_info_msg (c_job_buffer);

	return go_job_buffer
}

func Get_job( id uint32) Job_info_msg {
	var go_job_buffer Job_info_msg
	c_job_buffer := C.get_single_job_info(C.uint32_t(id))
	if c_job_buffer == nil {
		go_job_buffer.Last_update = int64(0)
		go_job_buffer.Record_count = uint32(0)
		go_job_buffer.Error_code = uint32(C.slurm_get_errno())
		return go_job_buffer
	}
	go_job_buffer.Last_update = int64(c_job_buffer.last_update)
	go_job_buffer.Record_count = uint32(c_job_buffer.record_count)
	go_job_buffer.Job_list =make([]Job_info,c_job_buffer.record_count, c_job_buffer.record_count)
	for i:=uint32(0); i<go_job_buffer.Record_count; i++ {
		job := C.job_from_list(c_job_buffer,  C.int(i))
		go_job :=  Job_info_convert_c_to_go(job)
		go_job_buffer.Job_list[i]=go_job
	}
	C.slurm_free_job_info_msg (c_job_buffer);

	return go_job_buffer
}

func Get_user_jobs (name string) Job_info_msg {


	var go_job_buffer Job_info_msg
	user, err := user.Lookup(name);

	if  err  != nil {
		fmt.Printf("Error %s\n", err.Error())
		go_job_buffer.Last_update = int64(0)
		go_job_buffer.Record_count = uint32(0)
		go_job_buffer.Error_code =  C.ESLURMD_UID_NOT_FOUND

		return go_job_buffer


	}

	userid , _ := strconv.Atoi(user.Uid)
	c_job_buffer := C.get_user_job_info(C.uint32_t(userid))

	if c_job_buffer == nil {
		go_job_buffer.Last_update = int64(0)
		go_job_buffer.Record_count = uint32(0)
		go_job_buffer.Error_code = uint32(C.slurm_get_errno())
		return go_job_buffer
	}
	go_job_buffer.Last_update = int64(c_job_buffer.last_update)
	go_job_buffer.Record_count = uint32(c_job_buffer.record_count)
	go_job_buffer.Job_list =make([]Job_info,c_job_buffer.record_count, c_job_buffer.record_count)
	for i:=uint32(0); i<go_job_buffer.Record_count; i++ {
		job := C.job_from_list(c_job_buffer,  C.int(i))
		go_job :=  Job_info_convert_c_to_go(job)
		go_job_buffer.Job_list[i]=go_job
	}
	C.slurm_free_job_info_msg (c_job_buffer);

	return go_job_buffer


}
