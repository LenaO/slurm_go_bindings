
package submit_job
/*
#cgo LDFLAGS: -lslurm
#include<stdlib.h>
#include<slurm/slurm.h>
#include<slurm/slurm_errno.h>
#include <stdint.h> 
#ifndef ptr_convert
#define ptr_convert 
inline uint8_t uint8_ptr(uint8_t* pointer) { 
    if (NULL == pointer) {
    return 0;}
    return *pointer;
}
inline int8_t int8_ptr(int8_t* pointer) {
    if (NULL == pointer) {
    return 0;}
    return *pointer;
}
inline uint16_t uint16_ptr(uint16_t* pointer) {
    if (NULL == pointer) {
    return 0;}
    return *pointer;
}
inline int16_t int16_ptr(int16_t* pointer) {
    if (NULL == pointer) {
    return 0;}
    return *pointer;
}
inline uint32_t uint32_ptr(uint32_t* pointer) {
    if (NULL == pointer) {
    return 0;}
    return *pointer;
}
inline int32_t int32_ptr(int32_t* pointer) {
    if (NULL == pointer) {
    return 0;}
    return *pointer;
}
inline uint64_t uint64_ptr(uint64_t* pointer) {
    if (NULL == pointer) {
    return 0;}
    return *pointer;
}
inline int64_t int64_ptr(int16_t* pointer) {
    if (NULL == pointer) {
    return 0;}
    return *pointer;
}

#endif
struct submit_response_msg *submit_job(struct job_descriptor *desc)
{

	struct submit_response_msg *resp_msg;
	if (slurm_submit_batch_job(desc,
				&resp_msg)) {
		return NULL;
	}
	return resp_msg;

}
int update_job (struct job_descriptor *msg) {

	 return slurm_update_job (msg);
}

void free_submit_response_msg(struct submit_response_msg *msg)
{
	 slurm_free_submit_response_response_msg(msg);
}
*/
import "C"

import "fmt"
import "unsafe"
import "slurm/jobinfo" 

type Job_descriptor struct {
	Account string;
	Acctg_freq string;
	Admin_comment string;
	Alloc_node string;
	Alloc_resp_port uint16;
	Alloc_sid uint32;
	Argc uint32;
	Argv []string;
	Array_inx string;
	 //array_bitmap void;
	Batch_features string;
	Begin_time int64;
	Bitflags uint32;
	Burst_buffer string;
	Ckpt_interval uint16;
	Ckpt_dir string;
	Clusters string;
	Cluster_features string;
	Comment string;
	Contiguous uint16;
	Core_spec uint16;
	Cpu_bind string;
	Cpu_bind_type uint16;
	Cpu_freq_min uint32;
	Cpu_freq_max uint32 ;
	Cpu_freq_gov uint32;
	Cpus_per_tres string;
	Deadline int64;
	Delay_boot uint32;
	Dependency string;
	End_time int64;
	Environment []string;
	Env_size uint32;
	Extra string;
	Exc_nodes string;
	Features string;
	Fed_siblings_active uint64;
	Fed_siblings_viable uint64;
	Group_id uint32;
	Immediate uint16;
	Job_id uint32;
	Job_id_str string;
	Kill_on_node_fail uint16;
	Licenses string;
	Mail_type uint16;
	Mail_user string;
	Mcs_label string;
	Mem_bind string;
	Mem_bind_type uint16;
	Mem_per_tres string;
	Name string;
	Network string;
	Nice uint32;
	Num_tasks uint32;
	Open_mode uint8;
	Origin_cluster string;
	Other_port uint16;
	Overcommit uint8;
	Pack_job_offset uint32;
	Partition string;
	Plane_size uint16;
	Power_flags uint8;
	Priority uint32;
	Profile uint32;
	Qos string;
	Reboot uint16;
	Resp_host string;
	Restart_cnt uint16;
	Req_nodes string;
	Requeue uint16;
	Reservation string;
	Script string;
	 //script_buf void;
	Shared uint16;
	Site_factor uint32;
	 //spank_job_env char**;
	Spank_job_env_size uint32;
	Task_dist uint32;
	Time_limit uint32;
	Time_min uint32;
	Tres_bind string;
	Tres_freq string;
	Tres_per_job string;
	Tres_per_node string;
	Tres_per_socket string;
	Tres_per_task string;
	User_id uint32;
	Wait_all_nodes uint16;
	Warn_flags uint16;
	Warn_signal uint16;
	Warn_time uint16;
	Work_dir string;
	Cpus_per_task uint16;
	Min_cpus uint32;
	Max_cpus uint32;
	Min_nodes uint32;
	Max_nodes uint32;
	Boards_per_node uint16;
	Sockets_per_board uint16;
	Sockets_per_node uint16;
	Cores_per_socket uint16;
	Threads_per_core uint16;
	Ntasks_per_node uint16;
	Ntasks_per_socket uint16;
	Ntasks_per_core uint16;
	Ntasks_per_board uint16;
	Pn_min_cpus uint16;
	Pn_min_memory uint64;
	Pn_min_tmp_disk uint32;
	Req_switch uint32;
	Std_err string;
	Std_in string;
	Std_out string;
	Tres_req_cnt uint64;
	Wait4switch uint32;
	Wckey string;
	X11 uint16;
	X11_magic_cookie string;
	X11_target string;
	X11_target_port uint16;
}
func Job_descriptor_convert_c_to_go(c_struct *C.struct_job_descriptor) Job_descriptor{
	var go_struct Job_descriptor

	go_struct.Account = C.GoString(c_struct.account)
	go_struct.Acctg_freq = C.GoString(c_struct.acctg_freq)
	go_struct.Admin_comment = C.GoString(c_struct.admin_comment)
	go_struct.Alloc_node = C.GoString(c_struct.alloc_node)
	go_struct.Alloc_resp_port = uint16(c_struct.alloc_resp_port)
	go_struct.Alloc_sid = uint32(c_struct.alloc_sid)
	go_struct.Argc = uint32(c_struct.argc)
	go_struct.Array_inx = C.GoString(c_struct.array_inx)
	go_struct.Batch_features = C.GoString(c_struct.batch_features)
	go_struct.Begin_time = int64(c_struct.begin_time)
	go_struct.Bitflags = uint32(c_struct.bitflags)
	go_struct.Burst_buffer = C.GoString(c_struct.burst_buffer)
	go_struct.Ckpt_interval = uint16(c_struct.ckpt_interval)
	go_struct.Ckpt_dir = C.GoString(c_struct.ckpt_dir)
	go_struct.Clusters = C.GoString(c_struct.clusters)
	go_struct.Cluster_features = C.GoString(c_struct.cluster_features)
	go_struct.Comment = C.GoString(c_struct.comment)
	go_struct.Contiguous = uint16(c_struct.contiguous)
	go_struct.Core_spec = uint16(c_struct.core_spec)
	go_struct.Cpu_bind = C.GoString(c_struct.cpu_bind)
	go_struct.Cpu_bind_type = uint16(c_struct.cpu_bind_type)
	go_struct.Cpu_freq_min = uint32(c_struct.cpu_freq_min)
	go_struct.Cpu_freq_max = uint32(c_struct.cpu_freq_max)
	go_struct.Cpu_freq_gov = uint32(c_struct.cpu_freq_gov)
	go_struct.Cpus_per_tres = C.GoString(c_struct.cpus_per_tres)
	go_struct.Deadline = int64(c_struct.deadline)
	go_struct.Delay_boot = uint32(c_struct.delay_boot)
	go_struct.Dependency = C.GoString(c_struct.dependency)
	go_struct.End_time = int64(c_struct.end_time)
	go_struct.Env_size = uint32(c_struct.env_size)
	go_struct.Extra = C.GoString(c_struct.extra)
	go_struct.Exc_nodes = C.GoString(c_struct.exc_nodes)
	go_struct.Features = C.GoString(c_struct.features)
	go_struct.Fed_siblings_active = uint64(c_struct.fed_siblings_active)
	go_struct.Fed_siblings_viable = uint64(c_struct.fed_siblings_viable)
	go_struct.Group_id = uint32(c_struct.group_id)
	go_struct.Immediate = uint16(c_struct.immediate)
	go_struct.Job_id = uint32(c_struct.job_id)
	go_struct.Job_id_str = C.GoString(c_struct.job_id_str)
	go_struct.Kill_on_node_fail = uint16(c_struct.kill_on_node_fail)
	go_struct.Licenses = C.GoString(c_struct.licenses)
	go_struct.Mail_type = uint16(c_struct.mail_type)
	go_struct.Mail_user = C.GoString(c_struct.mail_user)
	go_struct.Mcs_label = C.GoString(c_struct.mcs_label)
	go_struct.Mem_bind = C.GoString(c_struct.mem_bind)
	go_struct.Mem_bind_type = uint16(c_struct.mem_bind_type)
	go_struct.Mem_per_tres = C.GoString(c_struct.mem_per_tres)
	go_struct.Name = C.GoString(c_struct.name)
	go_struct.Network = C.GoString(c_struct.network)
	go_struct.Nice = uint32(c_struct.nice)
	go_struct.Num_tasks = uint32(c_struct.num_tasks)
	go_struct.Open_mode = uint8(c_struct.open_mode)
	go_struct.Origin_cluster = C.GoString(c_struct.origin_cluster)
	go_struct.Other_port = uint16(c_struct.other_port)
	go_struct.Overcommit = uint8(c_struct.overcommit)
	go_struct.Pack_job_offset = uint32(c_struct.pack_job_offset)
	go_struct.Partition = C.GoString(c_struct.partition)
	go_struct.Plane_size = uint16(c_struct.plane_size)
	go_struct.Power_flags = uint8(c_struct.power_flags)
	go_struct.Priority = uint32(c_struct.priority)
	go_struct.Profile = uint32(c_struct.profile)
	go_struct.Qos = C.GoString(c_struct.qos)
	go_struct.Reboot = uint16(c_struct.reboot)
	go_struct.Resp_host = C.GoString(c_struct.resp_host)
	go_struct.Restart_cnt = uint16(c_struct.restart_cnt)
	go_struct.Req_nodes = C.GoString(c_struct.req_nodes)
	go_struct.Requeue = uint16(c_struct.requeue)
	go_struct.Reservation = C.GoString(c_struct.reservation)
	go_struct.Script = C.GoString(c_struct.script)
	go_struct.Shared = uint16(c_struct.shared)
	go_struct.Site_factor = uint32(c_struct.site_factor)
	go_struct.Spank_job_env_size = uint32(c_struct.spank_job_env_size)
	go_struct.Task_dist = uint32(c_struct.task_dist)
	go_struct.Time_limit = uint32(c_struct.time_limit)
	go_struct.Time_min = uint32(c_struct.time_min)
	go_struct.Tres_bind = C.GoString(c_struct.tres_bind)
	go_struct.Tres_freq = C.GoString(c_struct.tres_freq)
	go_struct.Tres_per_job = C.GoString(c_struct.tres_per_job)
	go_struct.Tres_per_node = C.GoString(c_struct.tres_per_node)
	go_struct.Tres_per_socket = C.GoString(c_struct.tres_per_socket)
	go_struct.Tres_per_task = C.GoString(c_struct.tres_per_task)
	go_struct.User_id = uint32(c_struct.user_id)
	go_struct.Wait_all_nodes = uint16(c_struct.wait_all_nodes)
	go_struct.Warn_flags = uint16(c_struct.warn_flags)
	go_struct.Warn_signal = uint16(c_struct.warn_signal)
	go_struct.Warn_time = uint16(c_struct.warn_time)
	go_struct.Work_dir = C.GoString(c_struct.work_dir)
	go_struct.Cpus_per_task = uint16(c_struct.cpus_per_task)
	go_struct.Min_cpus = uint32(c_struct.min_cpus)
	go_struct.Max_cpus = uint32(c_struct.max_cpus)
	go_struct.Min_nodes = uint32(c_struct.min_nodes)
	go_struct.Max_nodes = uint32(c_struct.max_nodes)
	go_struct.Boards_per_node = uint16(c_struct.boards_per_node)
	go_struct.Sockets_per_board = uint16(c_struct.sockets_per_board)
	go_struct.Sockets_per_node = uint16(c_struct.sockets_per_node)
	go_struct.Cores_per_socket = uint16(c_struct.cores_per_socket)
	go_struct.Threads_per_core = uint16(c_struct.threads_per_core)
	go_struct.Ntasks_per_node = uint16(c_struct.ntasks_per_node)
	go_struct.Ntasks_per_socket = uint16(c_struct.ntasks_per_socket)
	go_struct.Ntasks_per_core = uint16(c_struct.ntasks_per_core)
	go_struct.Ntasks_per_board = uint16(c_struct.ntasks_per_board)
	go_struct.Pn_min_cpus = uint16(c_struct.pn_min_cpus)
	go_struct.Pn_min_memory = uint64(c_struct.pn_min_memory)
	go_struct.Pn_min_tmp_disk = uint32(c_struct.pn_min_tmp_disk)
	go_struct.Req_switch = uint32(c_struct.req_switch)
	go_struct.Std_err = C.GoString(c_struct.std_err)
	go_struct.Std_in = C.GoString(c_struct.std_in)
	go_struct.Std_out = C.GoString(c_struct.std_out)
	go_struct.Wait4switch = uint32(c_struct.wait4switch)
	go_struct.Wckey = C.GoString(c_struct.wckey)
	go_struct.X11 = uint16(c_struct.x11)
	go_struct.X11_magic_cookie = C.GoString(c_struct.x11_magic_cookie)
	go_struct.X11_target = C.GoString(c_struct.x11_target)
	go_struct.X11_target_port = uint16(c_struct.x11_target_port)
	return go_struct
 }
 func Print_Job_descriptor(go_struct Job_descriptor){
	 fmt.Printf("%s:\t %s\n","account", go_struct.Account)
	 fmt.Printf("%s:\t %s\n","acctg freq", go_struct.Acctg_freq)
	 fmt.Printf("%s:\t %s\n","admin comment", go_struct.Admin_comment)
	 fmt.Printf("%s:\t %s\n","alloc node", go_struct.Alloc_node)
	 fmt.Printf("%s:\t %d\n","alloc resp port", go_struct.Alloc_resp_port)
	 fmt.Printf("%s:\t %d\n","alloc sid", go_struct.Alloc_sid)
	 fmt.Printf("%s:\t %d\n","argc", go_struct.Argc)
	 fmt.Printf("%s:\t %s\n","array inx", go_struct.Array_inx)
	 fmt.Printf("%s:\t %s\n","batch features", go_struct.Batch_features)
	 fmt.Printf("%s:\t %d\n","begin time", go_struct.Begin_time)
	 fmt.Printf("%s:\t %d\n","bitflags", go_struct.Bitflags)
	 fmt.Printf("%s:\t %s\n","burst buffer", go_struct.Burst_buffer)
	 fmt.Printf("%s:\t %d\n","ckpt interval", go_struct.Ckpt_interval)
	 fmt.Printf("%s:\t %s\n","ckpt dir", go_struct.Ckpt_dir)
	 fmt.Printf("%s:\t %s\n","clusters", go_struct.Clusters)
	 fmt.Printf("%s:\t %s\n","cluster features", go_struct.Cluster_features)
	 fmt.Printf("%s:\t %s\n","comment", go_struct.Comment)
	 fmt.Printf("%s:\t %d\n","contiguous", go_struct.Contiguous)
	 fmt.Printf("%s:\t %d\n","core spec", go_struct.Core_spec)
	 fmt.Printf("%s:\t %s\n","cpu bind", go_struct.Cpu_bind)
	 fmt.Printf("%s:\t %d\n","cpu bind type", go_struct.Cpu_bind_type)
	 fmt.Printf("%s:\t %d\n","cpu freq min", go_struct.Cpu_freq_min)
	 fmt.Printf("%s:\t %d\n","cpu freq max", go_struct.Cpu_freq_max)
	 fmt.Printf("%s:\t %d\n","cpu freq gov", go_struct.Cpu_freq_gov)
	 fmt.Printf("%s:\t %s\n","cpus per tres", go_struct.Cpus_per_tres)
	 fmt.Printf("%s:\t %d\n","deadline", go_struct.Deadline)
	 fmt.Printf("%s:\t %d\n","delay boot", go_struct.Delay_boot)
	 fmt.Printf("%s:\t %s\n","dependency", go_struct.Dependency)
	 fmt.Printf("%s:\t %d\n","end time", go_struct.End_time)
	 fmt.Printf("%s:\t %d\n","env size", go_struct.Env_size)
	 fmt.Printf("%s:\t %s\n","extra", go_struct.Extra)
	 fmt.Printf("%s:\t %s\n","exc nodes", go_struct.Exc_nodes)
	 fmt.Printf("%s:\t %s\n","features", go_struct.Features)
	 fmt.Printf("%s:\t %d\n","fed siblings active", go_struct.Fed_siblings_active)
	 fmt.Printf("%s:\t %d\n","fed siblings viable", go_struct.Fed_siblings_viable)
	 fmt.Printf("%s:\t %d\n","group id", go_struct.Group_id)
	 fmt.Printf("%s:\t %d\n","immediate", go_struct.Immediate)
	 fmt.Printf("%s:\t %d\n","job id", go_struct.Job_id)
	 fmt.Printf("%s:\t %s\n","job id str", go_struct.Job_id_str)
	 fmt.Printf("%s:\t %d\n","kill on node fail", go_struct.Kill_on_node_fail)
	 fmt.Printf("%s:\t %s\n","licenses", go_struct.Licenses)
	 fmt.Printf("%s:\t %d\n","mail type", go_struct.Mail_type)
	 fmt.Printf("%s:\t %s\n","mail user", go_struct.Mail_user)
	 fmt.Printf("%s:\t %s\n","mcs label", go_struct.Mcs_label)
	 fmt.Printf("%s:\t %s\n","mem bind", go_struct.Mem_bind)
	 fmt.Printf("%s:\t %d\n","mem bind type", go_struct.Mem_bind_type)
	 fmt.Printf("%s:\t %s\n","mem per tres", go_struct.Mem_per_tres)
	 fmt.Printf("%s:\t %s\n","name", go_struct.Name)
	 fmt.Printf("%s:\t %s\n","network", go_struct.Network)
	 fmt.Printf("%s:\t %d\n","nice", go_struct.Nice)
	 fmt.Printf("%s:\t %d\n","num tasks", go_struct.Num_tasks)
	 fmt.Printf("%s:\t %d\n","open mode", go_struct.Open_mode)
	 fmt.Printf("%s:\t %s\n","origin cluster", go_struct.Origin_cluster)
	 fmt.Printf("%s:\t %d\n","other port", go_struct.Other_port)
	 fmt.Printf("%s:\t %d\n","overcommit", go_struct.Overcommit)
	 fmt.Printf("%s:\t %d\n","pack job offset", go_struct.Pack_job_offset)
	 fmt.Printf("%s:\t %s\n","partition", go_struct.Partition)
	 fmt.Printf("%s:\t %d\n","plane size", go_struct.Plane_size)
	 fmt.Printf("%s:\t %d\n","power flags", go_struct.Power_flags)
	 fmt.Printf("%s:\t %d\n","priority", go_struct.Priority)
	 fmt.Printf("%s:\t %d\n","profile", go_struct.Profile)
	 fmt.Printf("%s:\t %s\n","qos", go_struct.Qos)
	 fmt.Printf("%s:\t %d\n","reboot", go_struct.Reboot)
	 fmt.Printf("%s:\t %s\n","resp host", go_struct.Resp_host)
	 fmt.Printf("%s:\t %d\n","restart cnt", go_struct.Restart_cnt)
	 fmt.Printf("%s:\t %s\n","req nodes", go_struct.Req_nodes)
	 fmt.Printf("%s:\t %d\n","requeue", go_struct.Requeue)
	 fmt.Printf("%s:\t %s\n","reservation", go_struct.Reservation)
	 fmt.Printf("%s:\t %s\n","script", go_struct.Script)
	 fmt.Printf("%s:\t %d\n","shared", go_struct.Shared)
	 fmt.Printf("%s:\t %d\n","site factor", go_struct.Site_factor)
	 fmt.Printf("%s:\t %d\n","spank job env size", go_struct.Spank_job_env_size)
	 fmt.Printf("%s:\t %d\n","task dist", go_struct.Task_dist)
	 fmt.Printf("%s:\t %d\n","time limit", go_struct.Time_limit)
	 fmt.Printf("%s:\t %d\n","time min", go_struct.Time_min)
	 fmt.Printf("%s:\t %s\n","tres bind", go_struct.Tres_bind)
	 fmt.Printf("%s:\t %s\n","tres freq", go_struct.Tres_freq)
	 fmt.Printf("%s:\t %s\n","tres per job", go_struct.Tres_per_job)
	 fmt.Printf("%s:\t %s\n","tres per node", go_struct.Tres_per_node)
	 fmt.Printf("%s:\t %s\n","tres per socket", go_struct.Tres_per_socket)
	 fmt.Printf("%s:\t %s\n","tres per task", go_struct.Tres_per_task)
	 fmt.Printf("%s:\t %d\n","user id", go_struct.User_id)
	 fmt.Printf("%s:\t %d\n","wait all nodes", go_struct.Wait_all_nodes)
	 fmt.Printf("%s:\t %d\n","warn flags", go_struct.Warn_flags)
	 fmt.Printf("%s:\t %d\n","warn signal", go_struct.Warn_signal)
	 fmt.Printf("%s:\t %d\n","warn time", go_struct.Warn_time)
	 fmt.Printf("%s:\t %s\n","work dir", go_struct.Work_dir)
	 fmt.Printf("%s:\t %d\n","cpus per task", go_struct.Cpus_per_task)
	 fmt.Printf("%s:\t %d\n","min cpus", go_struct.Min_cpus)
	 fmt.Printf("%s:\t %d\n","max cpus", go_struct.Max_cpus)
	 fmt.Printf("%s:\t %d\n","min nodes", go_struct.Min_nodes)
	 fmt.Printf("%s:\t %d\n","max nodes", go_struct.Max_nodes)
	 fmt.Printf("%s:\t %d\n","boards per node", go_struct.Boards_per_node)
	 fmt.Printf("%s:\t %d\n","sockets per board", go_struct.Sockets_per_board)
	 fmt.Printf("%s:\t %d\n","sockets per node", go_struct.Sockets_per_node)
	 fmt.Printf("%s:\t %d\n","cores per socket", go_struct.Cores_per_socket)
	 fmt.Printf("%s:\t %d\n","threads per core", go_struct.Threads_per_core)
	 fmt.Printf("%s:\t %d\n","ntasks per node", go_struct.Ntasks_per_node)
	 fmt.Printf("%s:\t %d\n","ntasks per socket", go_struct.Ntasks_per_socket)
	 fmt.Printf("%s:\t %d\n","ntasks per core", go_struct.Ntasks_per_core)
	 fmt.Printf("%s:\t %d\n","ntasks per board", go_struct.Ntasks_per_board)
	 fmt.Printf("%s:\t %d\n","pn min cpus", go_struct.Pn_min_cpus)
	 fmt.Printf("%s:\t %d\n","pn min memory", go_struct.Pn_min_memory)
	 fmt.Printf("%s:\t %d\n","pn min tmp disk", go_struct.Pn_min_tmp_disk)
	 fmt.Printf("%s:\t %d\n","req switch", go_struct.Req_switch)
	 fmt.Printf("%s:\t %s\n","std err", go_struct.Std_err)
	 fmt.Printf("%s:\t %s\n","std in", go_struct.Std_in)
	 fmt.Printf("%s:\t %s\n","std out", go_struct.Std_out)
	 fmt.Printf("%s:\t %d\n","tres req cnt", go_struct.Tres_req_cnt)
	 fmt.Printf("%s:\t %d\n","wait4switch", go_struct.Wait4switch)
	 fmt.Printf("%s:\t %s\n","wckey", go_struct.Wckey)
	 fmt.Printf("%s:\t %d\n","x11", go_struct.X11)
	 fmt.Printf("%s:\t %s\n","x11 magic cookie", go_struct.X11_magic_cookie)
	 fmt.Printf("%s:\t %s\n","x11 target", go_struct.X11_target)
	 fmt.Printf("%s:\t %d\n","x11 target port", go_struct.X11_target_port)
}

type Update_job_options struct {
	Partition string;
	Qos string;
	Num_tasks uint32;
	Ntasks_per_node  uint16;
	Ntasks_per_socket uint16;
	Ntasks_per_core uint16;

}

type Submit_response_msg struct {
	Job_id uint32;
	Step_id uint32;
	Error_code uint32;
	Job_submit_user_msg string;
}
func submit_response_msg_convert_c_to_go(c_struct *C.struct_submit_response_msg) Submit_response_msg{
	var go_struct Submit_response_msg

	go_struct.Job_id = uint32(c_struct.job_id)
	go_struct.Step_id = uint32(c_struct.step_id)
	go_struct.Error_code = uint32(c_struct.error_code)
	go_struct.Job_submit_user_msg = C.GoString(c_struct.job_submit_user_msg)
	return go_struct
 }
 func Print_submit_response_msg(go_struct Submit_response_msg){
	 fmt.Printf("%s:\t %d\n","job id", go_struct.Job_id)
	 fmt.Printf("%s:\t %d\n","step id", go_struct.Step_id)
	 fmt.Printf("%s:\t %d\n","error code", go_struct.Error_code)
	 fmt.Printf("%s:\t %s\n","job submit user msg", go_struct.Job_submit_user_msg)
}
/*This is an ugly function, since we start to convert everyting back*/

func Submit_job (go_struct *Job_descriptor)  Submit_response_msg {

	var c_struct C.struct_job_descriptor

	C.slurm_init_job_desc_msg(&c_struct)
	if go_struct.Account!= "" {
		account_s :=C.CString(go_struct.Account)
		defer C.free(unsafe.Pointer(account_s))
		c_struct.account=account_s
	}
	if go_struct.Acctg_freq!= "" {
		acctg_freq_s :=C.CString(go_struct.Acctg_freq)
		defer C.free(unsafe.Pointer(acctg_freq_s))
		c_struct.acctg_freq=acctg_freq_s
	}
	if go_struct.Admin_comment!= "" {
		admin_comment_s :=C.CString(go_struct.Admin_comment)
		defer C.free(unsafe.Pointer(admin_comment_s))
		c_struct.admin_comment=admin_comment_s
	}
	if go_struct.Alloc_node!= "" {
		alloc_node_s :=C.CString(go_struct.Alloc_node)
		defer C.free(unsafe.Pointer(alloc_node_s))
		c_struct.alloc_node=alloc_node_s
	}
	if go_struct.Alloc_resp_port!= 0 {
		c_struct.alloc_resp_port = C.uint16_t(go_struct.Alloc_resp_port)
	}
	if go_struct.Alloc_sid!= 0 {
		c_struct.alloc_sid = C.uint32_t(go_struct.Alloc_sid)
	}
	if len(go_struct.Argv) > 0 {
		c_struct.argc = C.uint32_t(len(go_struct.Argv))
		cArray   := C.malloc(C.size_t(C.size_t(len(go_struct.Argv))*C.size_t(unsafe.Sizeof(uintptr(0)))))
		 a := (*[1<<30 - 1]*C.char)(cArray )
		 for i := 0; i < len(go_struct.Argv); i++ {
			a[i]=  C.CString(go_struct.Argv[i])
		}
		c_struct.argv=(**C.char)(cArray)
		fmt.Printf("test\n")
	}

	if go_struct.Array_inx!= "" {
		array_inx_s :=C.CString(go_struct.Array_inx)
		defer C.free(unsafe.Pointer(array_inx_s))
		c_struct.array_inx=array_inx_s
	}
	if go_struct.Batch_features!= "" {
		batch_features_s :=C.CString(go_struct.Batch_features)
		defer C.free(unsafe.Pointer(batch_features_s))
		c_struct.batch_features=batch_features_s
	}
	if go_struct.Begin_time!= 0 {
		c_struct.begin_time = C.int64_t(go_struct.Begin_time)
	}
	if go_struct.Bitflags!= 0 {
		c_struct.bitflags = C.uint32_t(go_struct.Bitflags)
	}
	if go_struct.Burst_buffer!= "" {
		burst_buffer_s :=C.CString(go_struct.Burst_buffer)
		defer C.free(unsafe.Pointer(burst_buffer_s))
		c_struct.burst_buffer=burst_buffer_s
	}
	if go_struct.Ckpt_interval!= 0 {
		c_struct.ckpt_interval = C.uint16_t(go_struct.Ckpt_interval)
	}
	if go_struct.Ckpt_dir!= "" {
		ckpt_dir_s :=C.CString(go_struct.Ckpt_dir)
		defer C.free(unsafe.Pointer(ckpt_dir_s))
		c_struct.ckpt_dir=ckpt_dir_s
	}
	if go_struct.Clusters!= "" {
		clusters_s :=C.CString(go_struct.Clusters)
		defer C.free(unsafe.Pointer(clusters_s))
		c_struct.clusters=clusters_s
	}
	if go_struct.Cluster_features!= "" {
		cluster_features_s :=C.CString(go_struct.Cluster_features)
		defer C.free(unsafe.Pointer(cluster_features_s))
		c_struct.cluster_features=cluster_features_s
	}
	if go_struct.Comment!= "" {
		comment_s :=C.CString(go_struct.Comment)
		defer C.free(unsafe.Pointer(comment_s))
		c_struct.comment=comment_s
	}
	if go_struct.Contiguous!= 0 {
		c_struct.contiguous = C.uint16_t(go_struct.Contiguous)
	}
	if go_struct.Core_spec!= 0 {
		c_struct.core_spec = C.uint16_t(go_struct.Core_spec)
	}
	if go_struct.Cpu_bind!= "" {
		cpu_bind_s :=C.CString(go_struct.Cpu_bind)
		defer C.free(unsafe.Pointer(cpu_bind_s))
		c_struct.cpu_bind=cpu_bind_s
	}
	if go_struct.Cpu_bind_type!= 0 {
		c_struct.cpu_bind_type = C.uint16_t(go_struct.Cpu_bind_type)
	}
	if go_struct.Cpu_freq_min!= 0 {
		c_struct.cpu_freq_min = C.uint32_t(go_struct.Cpu_freq_min)
	}
	if go_struct.Cpu_freq_max!= 0 {
		c_struct.cpu_freq_max = C.uint32_t(go_struct.Cpu_freq_max)
	}
	if go_struct.Cpu_freq_gov!= 0 {
		c_struct.cpu_freq_gov = C.uint32_t(go_struct.Cpu_freq_gov)
	}
	if go_struct.Cpus_per_tres!= "" {
		cpus_per_tres_s :=C.CString(go_struct.Cpus_per_tres)
		defer C.free(unsafe.Pointer(cpus_per_tres_s))
		c_struct.cpus_per_tres=cpus_per_tres_s
	}
	if go_struct.Deadline!= 0 {
		c_struct.deadline = C.int64_t(go_struct.Deadline)
	}
	if go_struct.Delay_boot!= 0 {
		c_struct.delay_boot = C.uint32_t(go_struct.Delay_boot)
	}
	if go_struct.Dependency!= "" {
		dependency_s :=C.CString(go_struct.Dependency)
		defer C.free(unsafe.Pointer(dependency_s))
		c_struct.dependency=dependency_s
	}
	if go_struct.End_time!= 0 {
		c_struct.end_time = C.int64_t(go_struct.End_time)
	}
	if len(go_struct.Environment) > 0 {
		c_struct.env_size = C.uint32_t(len(go_struct.Environment))
		cArray   := C.malloc(C.size_t(C.size_t(len(go_struct.Environment))*C.size_t(unsafe.Sizeof(uintptr(0)))))
		 a := (*[1<<30 - 1]*C.char)(cArray )
		 for i := 0; i < len(go_struct.Environment); i++ {
			a[i]=  C.CString(go_struct.Environment[i])
			defer C.free(unsafe.Pointer(a[i]))
		}
		c_struct.environment=(**C.char)(cArray)
	} else {
		c_struct.env_size = 1
		cArray   := C.malloc(C.size_t(C.size_t(1)*C.size_t(unsafe.Sizeof(uintptr(0)))))
		a := (*[1<<30 - 1]*C.char)(cArray )
		a[0]=  C.CString("SLURM_GO_JOB=TRUE")
		defer C.free(unsafe.Pointer(a[0]))
		c_struct.environment=(**C.char)(cArray)

	}
	if go_struct.Extra!= "" {
		extra_s :=C.CString(go_struct.Extra)
		defer C.free(unsafe.Pointer(extra_s))
		c_struct.extra=extra_s
	}
	if go_struct.Exc_nodes!= "" {
		exc_nodes_s :=C.CString(go_struct.Exc_nodes)
		defer C.free(unsafe.Pointer(exc_nodes_s))
		c_struct.exc_nodes=exc_nodes_s
	}
	if go_struct.Features!= "" {
		features_s :=C.CString(go_struct.Features)
		defer C.free(unsafe.Pointer(features_s))
		c_struct.features=features_s
	}
	if go_struct.Fed_siblings_active!= 0 {
		c_struct.fed_siblings_active = C.uint64_t(go_struct.Fed_siblings_active)
	}
	if go_struct.Fed_siblings_viable!= 0 {
		c_struct.fed_siblings_viable = C.uint64_t(go_struct.Fed_siblings_viable)
	}
	if go_struct.Group_id!= 0 {
		c_struct.group_id = C.uint32_t(go_struct.Group_id)
	}
	if go_struct.Immediate!= 0 {
		c_struct.immediate = C.uint16_t(go_struct.Immediate)
	}
	if go_struct.Job_id!= 0 {
		c_struct.job_id = C.uint32_t(go_struct.Job_id)
	}
	if go_struct.Job_id_str!= "" {
		job_id_str_s :=C.CString(go_struct.Job_id_str)
		defer C.free(unsafe.Pointer(job_id_str_s))
		c_struct.job_id_str=job_id_str_s
	}
	if go_struct.Kill_on_node_fail!= 0 {
		c_struct.kill_on_node_fail = C.uint16_t(go_struct.Kill_on_node_fail)
	}
	if go_struct.Licenses!= "" {
		licenses_s :=C.CString(go_struct.Licenses)
		defer C.free(unsafe.Pointer(licenses_s))
		c_struct.licenses=licenses_s
	}
	if go_struct.Mail_type!= 0 {
		c_struct.mail_type = C.uint16_t(go_struct.Mail_type)
	}
	if go_struct.Mail_user!= "" {
		mail_user_s :=C.CString(go_struct.Mail_user)
		defer C.free(unsafe.Pointer(mail_user_s))
		c_struct.mail_user=mail_user_s
	}
	if go_struct.Mcs_label!= "" {
		mcs_label_s :=C.CString(go_struct.Mcs_label)
		defer C.free(unsafe.Pointer(mcs_label_s))
		c_struct.mcs_label=mcs_label_s
	}
	if go_struct.Mem_bind!= "" {
		mem_bind_s :=C.CString(go_struct.Mem_bind)
		defer C.free(unsafe.Pointer(mem_bind_s))
		c_struct.mem_bind=mem_bind_s
	}
	if go_struct.Mem_bind_type!= 0 {
		c_struct.mem_bind_type = C.uint16_t(go_struct.Mem_bind_type)
	}
	if go_struct.Mem_per_tres!= "" {
		mem_per_tres_s :=C.CString(go_struct.Mem_per_tres)
		defer C.free(unsafe.Pointer(mem_per_tres_s))
		c_struct.mem_per_tres=mem_per_tres_s
	}
	if go_struct.Name!= "" {
		name_s :=C.CString(go_struct.Name)
		defer C.free(unsafe.Pointer(name_s))
		c_struct.name=name_s
	}
	if go_struct.Network!= "" {
		network_s :=C.CString(go_struct.Network)
		defer C.free(unsafe.Pointer(network_s))
		c_struct.network=network_s
	}
	if go_struct.Nice!= 0 {
		c_struct.nice = C.uint32_t(go_struct.Nice)
	}
	if go_struct.Num_tasks!= 0 {
		c_struct.num_tasks = C.uint32_t(go_struct.Num_tasks)
	}
	if go_struct.Open_mode!= 0 {
		c_struct.open_mode = C.uint8_t(go_struct.Open_mode)
	}
	if go_struct.Origin_cluster!= "" {
		origin_cluster_s :=C.CString(go_struct.Origin_cluster)
		defer C.free(unsafe.Pointer(origin_cluster_s))
		c_struct.origin_cluster=origin_cluster_s
	}
	if go_struct.Other_port!= 0 {
		c_struct.other_port = C.uint16_t(go_struct.Other_port)
	}
	if go_struct.Overcommit!= 0 {
		c_struct.overcommit = C.uint8_t(go_struct.Overcommit)
	}
	if go_struct.Pack_job_offset!= 0 {
		c_struct.pack_job_offset = C.uint32_t(go_struct.Pack_job_offset)
	}
	if go_struct.Partition!= "" {
		partition_s :=C.CString(go_struct.Partition)
		defer C.free(unsafe.Pointer(partition_s))
		c_struct.partition=partition_s
	}
	if go_struct.Plane_size!= 0 {
		c_struct.plane_size = C.uint16_t(go_struct.Plane_size)
	}
	if go_struct.Power_flags!= 0 {
		c_struct.power_flags = C.uint8_t(go_struct.Power_flags)
	}
	if go_struct.Priority!= 0 {
		c_struct.priority = C.uint32_t(go_struct.Priority)
	}
	if go_struct.Profile!= 0 {
		c_struct.profile = C.uint32_t(go_struct.Profile)
	}
	if go_struct.Qos!= "" {
		qos_s :=C.CString(go_struct.Qos)
		defer C.free(unsafe.Pointer(qos_s))
		c_struct.qos=qos_s
	}
	if go_struct.Reboot!= 0 {
		c_struct.reboot = C.uint16_t(go_struct.Reboot)
	}
	if go_struct.Resp_host!= "" {
		resp_host_s :=C.CString(go_struct.Resp_host)
		defer C.free(unsafe.Pointer(resp_host_s))
		c_struct.resp_host=resp_host_s
	}
	if go_struct.Restart_cnt!= 0 {
		c_struct.restart_cnt = C.uint16_t(go_struct.Restart_cnt)
	}
	if go_struct.Req_nodes!= "" {
		req_nodes_s :=C.CString(go_struct.Req_nodes)
		defer C.free(unsafe.Pointer(req_nodes_s))
		c_struct.req_nodes=req_nodes_s
	}
	if go_struct.Requeue!= 0 {
		c_struct.requeue = C.uint16_t(go_struct.Requeue)
	}
	if go_struct.Reservation!= "" {
		reservation_s :=C.CString(go_struct.Reservation)
		defer C.free(unsafe.Pointer(reservation_s))
		c_struct.reservation=reservation_s
	}
	if go_struct.Script!= "" {
		script_s :=C.CString(go_struct.Script)
		defer C.free(unsafe.Pointer(script_s))
		c_struct.script=script_s
	}
	if go_struct.Shared!= 0 {
		c_struct.shared = C.uint16_t(go_struct.Shared)
	}
	if go_struct.Site_factor!= 0 {
		c_struct.site_factor = C.uint32_t(go_struct.Site_factor)
	}
	if go_struct.Spank_job_env_size!= 0 {
		c_struct.spank_job_env_size = C.uint32_t(go_struct.Spank_job_env_size)
	}
	if go_struct.Task_dist!= 0 {
		c_struct.task_dist = C.uint32_t(go_struct.Task_dist)
	}
	if go_struct.Time_limit!= 0 {
		c_struct.time_limit = C.uint32_t(go_struct.Time_limit)
	}
	if go_struct.Time_min!= 0 {
		c_struct.time_min = C.uint32_t(go_struct.Time_min)
	}
	if go_struct.Tres_bind!= "" {
		tres_bind_s :=C.CString(go_struct.Tres_bind)
		defer C.free(unsafe.Pointer(tres_bind_s))
		c_struct.tres_bind=tres_bind_s
	}
	if go_struct.Tres_freq!= "" {
		tres_freq_s :=C.CString(go_struct.Tres_freq)
		defer C.free(unsafe.Pointer(tres_freq_s))
		c_struct.tres_freq=tres_freq_s
	}
	if go_struct.Tres_per_job!= "" {
		tres_per_job_s :=C.CString(go_struct.Tres_per_job)
		defer C.free(unsafe.Pointer(tres_per_job_s))
		c_struct.tres_per_job=tres_per_job_s
	}
	if go_struct.Tres_per_node!= "" {
		tres_per_node_s :=C.CString(go_struct.Tres_per_node)
		defer C.free(unsafe.Pointer(tres_per_node_s))
		c_struct.tres_per_node=tres_per_node_s
	}
	if go_struct.Tres_per_socket!= "" {
		tres_per_socket_s :=C.CString(go_struct.Tres_per_socket)
		defer C.free(unsafe.Pointer(tres_per_socket_s))
		c_struct.tres_per_socket=tres_per_socket_s
	}
	if go_struct.Tres_per_task!= "" {
		tres_per_task_s :=C.CString(go_struct.Tres_per_task)
		defer C.free(unsafe.Pointer(tres_per_task_s))
		c_struct.tres_per_task=tres_per_task_s
	}
	if go_struct.User_id!= 0 {
		c_struct.user_id = C.uint32_t(go_struct.User_id)
	}
	if go_struct.Wait_all_nodes!= 0 {
		c_struct.wait_all_nodes = C.uint16_t(go_struct.Wait_all_nodes)
	}
	if go_struct.Warn_flags!= 0 {
		c_struct.warn_flags = C.uint16_t(go_struct.Warn_flags)
	}
	if go_struct.Warn_signal!= 0 {
		c_struct.warn_signal = C.uint16_t(go_struct.Warn_signal)
	}
	if go_struct.Warn_time!= 0 {
		c_struct.warn_time = C.uint16_t(go_struct.Warn_time)
	}
	if go_struct.Work_dir!= "" {
		work_dir_s :=C.CString(go_struct.Work_dir)
		defer C.free(unsafe.Pointer(work_dir_s))
		c_struct.work_dir=work_dir_s
	}
	if go_struct.Cpus_per_task!= 0 {
		c_struct.cpus_per_task = C.uint16_t(go_struct.Cpus_per_task)
	}
	if go_struct.Min_cpus!= 0 {
		c_struct.min_cpus = C.uint32_t(go_struct.Min_cpus)
	}
	if go_struct.Max_cpus!= 0 {
		c_struct.max_cpus = C.uint32_t(go_struct.Max_cpus)
	}
	if go_struct.Min_nodes!= 0 {
		c_struct.min_nodes = C.uint32_t(go_struct.Min_nodes)
	}
	if go_struct.Max_nodes!= 0 {
		c_struct.max_nodes = C.uint32_t(go_struct.Max_nodes)
	}
	if go_struct.Boards_per_node!= 0 {
		c_struct.boards_per_node = C.uint16_t(go_struct.Boards_per_node)
	}
	if go_struct.Sockets_per_board!= 0 {
		c_struct.sockets_per_board = C.uint16_t(go_struct.Sockets_per_board)
	}
	if go_struct.Sockets_per_node!= 0 {
		c_struct.sockets_per_node = C.uint16_t(go_struct.Sockets_per_node)
	}
	if go_struct.Cores_per_socket!= 0 {
		c_struct.cores_per_socket = C.uint16_t(go_struct.Cores_per_socket)
	}
	if go_struct.Threads_per_core!= 0 {
		c_struct.threads_per_core = C.uint16_t(go_struct.Threads_per_core)
	}
	if go_struct.Ntasks_per_node!= 0 {
		c_struct.ntasks_per_node = C.uint16_t(go_struct.Ntasks_per_node)
	}
	if go_struct.Ntasks_per_socket!= 0 {
		c_struct.ntasks_per_socket = C.uint16_t(go_struct.Ntasks_per_socket)
	}
	if go_struct.Ntasks_per_core!= 0 {
		c_struct.ntasks_per_core = C.uint16_t(go_struct.Ntasks_per_core)
	}
	if go_struct.Ntasks_per_board!= 0 {
		c_struct.ntasks_per_board = C.uint16_t(go_struct.Ntasks_per_board)
	}
	if go_struct.Pn_min_cpus!= 0 {
		c_struct.pn_min_cpus = C.uint16_t(go_struct.Pn_min_cpus)
	}
	if go_struct.Pn_min_memory!= 0 {
		c_struct.pn_min_memory = C.uint64_t(go_struct.Pn_min_memory)
	}
	if go_struct.Pn_min_tmp_disk!= 0 {
		c_struct.pn_min_tmp_disk = C.uint32_t(go_struct.Pn_min_tmp_disk)
	}
	if go_struct.Req_switch!= 0 {
		c_struct.req_switch = C.uint32_t(go_struct.Req_switch)
	}
	if go_struct.Std_err!= "" {
		std_err_s :=C.CString(go_struct.Std_err)
		defer C.free(unsafe.Pointer(std_err_s))
		c_struct.std_err=std_err_s
	}
	if go_struct.Std_in!= "" {
		std_in_s :=C.CString(go_struct.Std_in)
		defer C.free(unsafe.Pointer(std_in_s))
		c_struct.std_in=std_in_s
	}
	if go_struct.Std_out!= "" {
		std_out_s :=C.CString(go_struct.Std_out)
		defer C.free(unsafe.Pointer(std_out_s))
		c_struct.std_out=std_out_s
	}

	if go_struct.Wait4switch!= 0 {
		c_struct.wait4switch = C.uint32_t(go_struct.Wait4switch)
	}
	if go_struct.Wckey!= "" {
		wckey_s :=C.CString(go_struct.Wckey)
		defer C.free(unsafe.Pointer(wckey_s))
		c_struct.wckey=wckey_s
	}
	if go_struct.X11!= 0 {
		c_struct.x11 = C.uint16_t(go_struct.X11)
	}
	if go_struct.X11_magic_cookie!= "" {
		x11_magic_cookie_s :=C.CString(go_struct.X11_magic_cookie)
		defer C.free(unsafe.Pointer(x11_magic_cookie_s))
		c_struct.x11_magic_cookie=x11_magic_cookie_s
	}
	if go_struct.X11_target!= "" {
		x11_target_s :=C.CString(go_struct.X11_target)
		defer C.free(unsafe.Pointer(x11_target_s))
		c_struct.x11_target=x11_target_s
	}
	if go_struct.X11_target_port!= 0 {
		c_struct.x11_target_port = C.uint16_t(go_struct.X11_target_port)
	}

	c_msg := C.submit_job(&c_struct)

	defer C.free_submit_response_msg(c_msg)
	if c_msg == nil{
		go_msg := Submit_response_msg{}
		go_msg.Job_id = 1<<31-1
		go_msg.Error_code = uint32(C.slurm_get_errno())
		return go_msg
	}
	go_msg :=  submit_response_msg_convert_c_to_go(c_msg)

	return go_msg

}


func  Update_job (update_info Update_job_options, JobId uint32 ) uint32 {

	var c_struct C.struct_job_descriptor
	C.slurm_init_job_desc_msg(&c_struct)
	if update_info.Partition !="" {
		partition_s  := C.CString(update_info.Partition)
		defer C.free(unsafe.Pointer(partition_s))
		c_struct.partition = partition_s
	}
	if update_info.Qos != "" {
		qos_s  := C.CString(update_info.Qos)
		defer C.free(unsafe.Pointer(qos_s))
		c_struct.qos = qos_s
	}
	if update_info.Num_tasks != 0 {
		c_struct.num_tasks = C.uint32_t(update_info.Num_tasks)
	}
	if update_info.Ntasks_per_core != 0 {
		c_struct.ntasks_per_core = C.uint16_t(update_info.Ntasks_per_core)
	}

	if update_info.Ntasks_per_node != 0 {
		c_struct.ntasks_per_node = C.uint16_t(update_info.Ntasks_per_node)
	}
	if update_info.Ntasks_per_socket != 0 {
		c_struct.ntasks_per_socket = C.uint16_t(update_info.Ntasks_per_socket)
	}
	job_list := job_info.Get_job(uint32(JobId))

	if job_list.Error_code != 0 {
		return uint32(job_list.Error_code)
	}

	job := job_list.Job_list[0]
	if job.Job_state != C.JOB_PENDING {
		return uint32(C.ESLURM_JOB_NOT_PENDING)
	}
	c_struct.job_id = C.uint32_t(JobId)

	err := C.update_job(&c_struct)

	return uint32(err)

}
