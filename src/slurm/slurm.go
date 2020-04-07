package slurm
/*
#cgo LDFLAGS: -lslurm
#include<stdlib.h>
#include<slurm/slurm.h>
#include<slurm/slurm_errno.h>
#ifndef ptr_convert
#define ptr_convert 

struct slurm_ctl_conf* get_config(uint64_t time){
	slurm_ctl_conf_t * config = NULL;
	slurm_load_ctl_conf ((time_t) time, &config);
	return config;

}

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
#endif
*/
import "C"


import "fmt"
type Ctl_conf struct {
	Last_update int64;
	Accounting_storage_tres string;
	Accounting_storage_enforce uint16;
	Accounting_storage_backup_host string;
	Accounting_storage_host string;
	Accounting_storage_loc string;
	Accounting_storage_pass string;
	Accounting_storage_port uint32;
	Accounting_storage_type string;
	Accounting_storage_user string;
	Acctng_store_job_comment uint16;
	 //acct_gather_conf void;
	Acct_gather_energy_type string;
	Acct_gather_profile_type string;
	Acct_gather_interconnect_type string;
	Acct_gather_filesystem_type string;
	Acct_gather_node_freq uint16;
	Authalttypes string;
	Authinfo string;
	Authtype string;
	Batch_start_timeout uint16;
	Bb_type string;
	Boot_time int64;
	 //cgroup_conf void;
	Checkpoint_type string;
	Cli_filter_plugins string;
	Core_spec_plugin string;
	Cluster_name string;
	Comm_params string;
	Complete_wait uint16;
	 //control_addr char**;
	Control_cnt uint32;
	 //control_machine char**;
	Cpu_freq_def uint32;
	Cpu_freq_govs uint32;
	Cred_type string;
	Debug_flags uint64;
	Def_mem_per_cpu uint64;
	Disable_root_jobs uint16;
	Eio_timeout uint16;
	Enforce_part_limits uint16;
	Epilog string;
	Epilog_msg_time uint32;
	Epilog_slurmctld string;
	Ext_sensors_type string;
	Ext_sensors_freq uint16;
	 //ext_sensors_conf void;
	Fast_schedule uint16;
	Fed_params string;
	First_job_id uint32;
	Fs_dampening_factor uint16;
	Get_env_timeout uint16;
	Gres_plugins string;
	Group_time uint16;
	Group_force uint16;
	Gpu_freq_def string;
	Hash_val uint32;
	Health_check_interval uint16;
	Health_check_node_state uint16;
	Health_check_program string;
	Inactive_limit uint16;
	Job_acct_gather_freq string;
	Job_acct_gather_type string;
	Job_acct_gather_params string;
	Job_acct_oom_kill uint16;
	Job_ckpt_dir string;
	Job_comp_host string;
	Job_comp_loc string;
	Job_comp_pass string;
	Job_comp_port uint32;
	Job_comp_type string;
	Job_comp_user string;
	Job_container_plugin string;
	Job_credential_private_key string;
	Job_credential_public_certificate string;
	Job_file_append uint16;
	Job_requeue uint16;
	Job_submit_plugins string;
	Keep_alive_time uint16;
	Kill_on_bad_exit uint16;
	Kill_wait uint16;
	Launch_params string;
	Launch_type string;
	Layouts string;
	Licenses string;
	Licenses_used string;
	Log_fmt uint16;
	Mail_domain string;
	Mail_prog string;
	Max_array_sz uint32;
	Max_job_cnt uint32;
	Max_job_id uint32;
	Max_mem_per_cpu uint64;
	Max_step_cnt uint32;
	Max_tasks_per_node uint16;
	Mcs_plugin string;
	Mcs_plugin_params string;
	Min_job_age uint32;
	Mpi_default string;
	Mpi_params string;
	Msg_aggr_params string;
	Msg_timeout uint16;
	Tcp_timeout uint16;
	Next_job_id uint32;
	 //node_features_conf void;
	Node_features_plugins string;
	Node_prefix string;
	Over_time_limit uint16;
	Plugindir string;
	Plugstack string;
	Power_parameters string;
	Power_plugin string;
	Preempt_exempt_time uint32;
	Preempt_mode uint16;
	Preempt_type string;
	Priority_decay_hl uint32;
	Priority_calc_period uint32;
	Priority_favor_small uint16;
	Priority_flags uint16;
	Priority_max_age uint32;
	Priority_params string;
	Priority_reset_period uint16;
	Priority_type string;
	Priority_weight_age uint32;
	Priority_weight_assoc uint32;
	Priority_weight_fs uint32;
	Priority_weight_js uint32;
	Priority_weight_part uint32;
	Priority_weight_qos uint32;
	Priority_weight_tres string;
	Private_data uint16;
	Proctrack_type string;
	Prolog string;
	Prolog_epilog_timeout uint16;
	Prolog_slurmctld string;
	Propagate_prio_process uint16;
	Prolog_flags uint16;
	Propagate_rlimits string;
	Propagate_rlimits_except string;
	Reboot_program string;
	Reconfig_flags uint16;
	Requeue_exit string;
	Requeue_exit_hold string;
	Resume_fail_program string;
	Resume_program string;
	Resume_rate uint16;
	Resume_timeout uint16;
	Resv_epilog string;
	Resv_over_run uint16;
	Resv_prolog string;
	Ret2service uint16;
	Route_plugin string;
	Salloc_default_command string;
	Sbcast_parameters string;
	Sched_logfile string;
	Sched_log_level uint16;
	Sched_params string;
	Sched_time_slice uint16;
	Schedtype string;
	Select_type string;
	 //select_conf_key_pairs void;
	Select_type_param uint16;
	Site_factor_plugin string;
	Site_factor_params string;
	Slurm_conf string;
	Slurm_user_id uint32;
	Slurm_user_name string;
	Slurmd_user_id uint32;
	Slurmd_user_name string;
	Slurmctld_addr string;
	Slurmctld_debug uint16;
	Slurmctld_logfile string;
	Slurmctld_pidfile string;
	Slurmctld_plugstack string;
	 //slurmctld_plugstack_conf void;
	Slurmctld_port uint32;
	Slurmctld_port_count uint16;
	Slurmctld_primary_off_prog string;
	Slurmctld_primary_on_prog string;
	Slurmctld_syslog_debug uint16;
	Slurmctld_timeout uint16;
	Slurmctld_params string;
	Slurmd_debug uint16;
	Slurmd_logfile string;
	Slurmd_params string;
	Slurmd_pidfile string;
	Slurmd_port uint32;
	Slurmd_spooldir string;
	Slurmd_syslog_debug uint16;
	Slurmd_timeout uint16;
	Srun_epilog string;
	Srun_port_range uint16;
	Srun_prolog string;
	State_save_location string;
	Suspend_exc_nodes string;
	Suspend_exc_parts string;
	Suspend_program string;
	Suspend_rate uint16;
	Suspend_time uint32;
	Suspend_timeout uint16;
	Switch_type string;
	Task_epilog string;
	Task_plugin string;
	Task_plugin_param uint32;
	Task_prolog string;
	Tmp_fs string;
	Topology_param string;
	Topology_plugin string;
	Track_wckey uint16;
	Tree_width uint16;
	Unkillable_program string;
	Unkillable_timeout uint16;
	Use_pam uint16;
	Use_spec_resources uint16;
	Version string;
	Vsize_factor uint16;
	Wait_time uint16;
	X11_params string;
}
func Ctl_conf_convert_c_to_go(c_struct *C.struct_slurm_ctl_conf) Ctl_conf{
	var go_struct Ctl_conf

	go_struct.Last_update = int64(c_struct.last_update)
	go_struct.Accounting_storage_tres = C.GoString(c_struct.accounting_storage_tres)
	go_struct.Accounting_storage_enforce = uint16(c_struct.accounting_storage_enforce)
	go_struct.Accounting_storage_backup_host = C.GoString(c_struct.accounting_storage_backup_host)
	go_struct.Accounting_storage_host = C.GoString(c_struct.accounting_storage_host)
	go_struct.Accounting_storage_loc = C.GoString(c_struct.accounting_storage_loc)
	go_struct.Accounting_storage_pass = C.GoString(c_struct.accounting_storage_pass)
	go_struct.Accounting_storage_port = uint32(c_struct.accounting_storage_port)
	go_struct.Accounting_storage_type = C.GoString(c_struct.accounting_storage_type)
	go_struct.Accounting_storage_user = C.GoString(c_struct.accounting_storage_user)
	go_struct.Acctng_store_job_comment = uint16(c_struct.acctng_store_job_comment)
	go_struct.Acct_gather_energy_type = C.GoString(c_struct.acct_gather_energy_type)
	go_struct.Acct_gather_profile_type = C.GoString(c_struct.acct_gather_profile_type)
	go_struct.Acct_gather_interconnect_type = C.GoString(c_struct.acct_gather_interconnect_type)
	go_struct.Acct_gather_filesystem_type = C.GoString(c_struct.acct_gather_filesystem_type)
	go_struct.Acct_gather_node_freq = uint16(c_struct.acct_gather_node_freq)
	go_struct.Authalttypes = C.GoString(c_struct.authalttypes)
	go_struct.Authinfo = C.GoString(c_struct.authinfo)
	go_struct.Authtype = C.GoString(c_struct.authtype)
	go_struct.Batch_start_timeout = uint16(c_struct.batch_start_timeout)
	go_struct.Bb_type = C.GoString(c_struct.bb_type)
	go_struct.Boot_time = int64(c_struct.boot_time)
	go_struct.Checkpoint_type = C.GoString(c_struct.checkpoint_type)
	go_struct.Cli_filter_plugins = C.GoString(c_struct.cli_filter_plugins)
	go_struct.Core_spec_plugin = C.GoString(c_struct.core_spec_plugin)
	go_struct.Cluster_name = C.GoString(c_struct.cluster_name)
	go_struct.Comm_params = C.GoString(c_struct.comm_params)
	go_struct.Complete_wait = uint16(c_struct.complete_wait)
	go_struct.Control_cnt = uint32(c_struct.control_cnt)
	go_struct.Cpu_freq_def = uint32(c_struct.cpu_freq_def)
	go_struct.Cpu_freq_govs = uint32(c_struct.cpu_freq_govs)
	go_struct.Cred_type = C.GoString(c_struct.cred_type)
	go_struct.Debug_flags = uint64(c_struct.debug_flags)
	go_struct.Def_mem_per_cpu = uint64(c_struct.def_mem_per_cpu)
	go_struct.Disable_root_jobs = uint16(c_struct.disable_root_jobs)
	go_struct.Eio_timeout = uint16(c_struct.eio_timeout)
	go_struct.Enforce_part_limits = uint16(c_struct.enforce_part_limits)
	go_struct.Epilog = C.GoString(c_struct.epilog)
	go_struct.Epilog_msg_time = uint32(c_struct.epilog_msg_time)
	go_struct.Epilog_slurmctld = C.GoString(c_struct.epilog_slurmctld)
	go_struct.Ext_sensors_type = C.GoString(c_struct.ext_sensors_type)
	go_struct.Ext_sensors_freq = uint16(c_struct.ext_sensors_freq)
	go_struct.Fast_schedule = uint16(c_struct.fast_schedule)
	go_struct.Fed_params = C.GoString(c_struct.fed_params)
	go_struct.First_job_id = uint32(c_struct.first_job_id)
	go_struct.Fs_dampening_factor = uint16(c_struct.fs_dampening_factor)
	go_struct.Get_env_timeout = uint16(c_struct.get_env_timeout)
	go_struct.Gres_plugins = C.GoString(c_struct.gres_plugins)
	go_struct.Group_time = uint16(c_struct.group_time)
	go_struct.Group_force = uint16(c_struct.group_force)
	go_struct.Gpu_freq_def = C.GoString(c_struct.gpu_freq_def)
	go_struct.Hash_val = uint32(c_struct.hash_val)
	go_struct.Health_check_interval = uint16(c_struct.health_check_interval)
	go_struct.Health_check_node_state = uint16(c_struct.health_check_node_state)
	go_struct.Health_check_program = C.GoString(c_struct.health_check_program)
	go_struct.Inactive_limit = uint16(c_struct.inactive_limit)
	go_struct.Job_acct_gather_freq = C.GoString(c_struct.job_acct_gather_freq)
	go_struct.Job_acct_gather_type = C.GoString(c_struct.job_acct_gather_type)
	go_struct.Job_acct_gather_params = C.GoString(c_struct.job_acct_gather_params)
	go_struct.Job_acct_oom_kill = uint16(c_struct.job_acct_oom_kill)
	go_struct.Job_ckpt_dir = C.GoString(c_struct.job_ckpt_dir)
	go_struct.Job_comp_host = C.GoString(c_struct.job_comp_host)
	go_struct.Job_comp_loc = C.GoString(c_struct.job_comp_loc)
	go_struct.Job_comp_pass = C.GoString(c_struct.job_comp_pass)
	go_struct.Job_comp_port = uint32(c_struct.job_comp_port)
	go_struct.Job_comp_type = C.GoString(c_struct.job_comp_type)
	go_struct.Job_comp_user = C.GoString(c_struct.job_comp_user)
	go_struct.Job_container_plugin = C.GoString(c_struct.job_container_plugin)
	go_struct.Job_credential_private_key = C.GoString(c_struct.job_credential_private_key)
	go_struct.Job_credential_public_certificate = C.GoString(c_struct.job_credential_public_certificate)
	go_struct.Job_file_append = uint16(c_struct.job_file_append)
	go_struct.Job_requeue = uint16(c_struct.job_requeue)
	go_struct.Job_submit_plugins = C.GoString(c_struct.job_submit_plugins)
	go_struct.Keep_alive_time = uint16(c_struct.keep_alive_time)
	go_struct.Kill_on_bad_exit = uint16(c_struct.kill_on_bad_exit)
	go_struct.Kill_wait = uint16(c_struct.kill_wait)
	go_struct.Launch_params = C.GoString(c_struct.launch_params)
	go_struct.Launch_type = C.GoString(c_struct.launch_type)
	go_struct.Layouts = C.GoString(c_struct.layouts)
	go_struct.Licenses = C.GoString(c_struct.licenses)
	go_struct.Licenses_used = C.GoString(c_struct.licenses_used)
	go_struct.Log_fmt = uint16(c_struct.log_fmt)
	go_struct.Mail_domain = C.GoString(c_struct.mail_domain)
	go_struct.Mail_prog = C.GoString(c_struct.mail_prog)
	go_struct.Max_array_sz = uint32(c_struct.max_array_sz)
	go_struct.Max_job_cnt = uint32(c_struct.max_job_cnt)
	go_struct.Max_job_id = uint32(c_struct.max_job_id)
	go_struct.Max_mem_per_cpu = uint64(c_struct.max_mem_per_cpu)
	go_struct.Max_step_cnt = uint32(c_struct.max_step_cnt)
	go_struct.Max_tasks_per_node = uint16(c_struct.max_tasks_per_node)
	go_struct.Mcs_plugin = C.GoString(c_struct.mcs_plugin)
	go_struct.Mcs_plugin_params = C.GoString(c_struct.mcs_plugin_params)
	go_struct.Min_job_age = uint32(c_struct.min_job_age)
	go_struct.Mpi_default = C.GoString(c_struct.mpi_default)
	go_struct.Mpi_params = C.GoString(c_struct.mpi_params)
	go_struct.Msg_aggr_params = C.GoString(c_struct.msg_aggr_params)
	go_struct.Msg_timeout = uint16(c_struct.msg_timeout)
	go_struct.Tcp_timeout = uint16(c_struct.tcp_timeout)
	go_struct.Next_job_id = uint32(c_struct.next_job_id)
	go_struct.Node_features_plugins = C.GoString(c_struct.node_features_plugins)
	go_struct.Node_prefix = C.GoString(c_struct.node_prefix)
	go_struct.Over_time_limit = uint16(c_struct.over_time_limit)
	go_struct.Plugindir = C.GoString(c_struct.plugindir)
	go_struct.Plugstack = C.GoString(c_struct.plugstack)
	go_struct.Power_parameters = C.GoString(c_struct.power_parameters)
	go_struct.Power_plugin = C.GoString(c_struct.power_plugin)
	go_struct.Preempt_exempt_time = uint32(c_struct.preempt_exempt_time)
	go_struct.Preempt_mode = uint16(c_struct.preempt_mode)
	go_struct.Preempt_type = C.GoString(c_struct.preempt_type)
	go_struct.Priority_decay_hl = uint32(c_struct.priority_decay_hl)
	go_struct.Priority_calc_period = uint32(c_struct.priority_calc_period)
	go_struct.Priority_favor_small = uint16(c_struct.priority_favor_small)
	go_struct.Priority_flags = uint16(c_struct.priority_flags)
	go_struct.Priority_max_age = uint32(c_struct.priority_max_age)
	go_struct.Priority_params = C.GoString(c_struct.priority_params)
	go_struct.Priority_reset_period = uint16(c_struct.priority_reset_period)
	go_struct.Priority_type = C.GoString(c_struct.priority_type)
	go_struct.Priority_weight_age = uint32(c_struct.priority_weight_age)
	go_struct.Priority_weight_assoc = uint32(c_struct.priority_weight_assoc)
	go_struct.Priority_weight_fs = uint32(c_struct.priority_weight_fs)
	go_struct.Priority_weight_js = uint32(c_struct.priority_weight_js)
	go_struct.Priority_weight_part = uint32(c_struct.priority_weight_part)
	go_struct.Priority_weight_qos = uint32(c_struct.priority_weight_qos)
	go_struct.Priority_weight_tres = C.GoString(c_struct.priority_weight_tres)
	go_struct.Private_data = uint16(c_struct.private_data)
	go_struct.Proctrack_type = C.GoString(c_struct.proctrack_type)
	go_struct.Prolog = C.GoString(c_struct.prolog)
	go_struct.Prolog_epilog_timeout = uint16(c_struct.prolog_epilog_timeout)
	go_struct.Prolog_slurmctld = C.GoString(c_struct.prolog_slurmctld)
	go_struct.Propagate_prio_process = uint16(c_struct.propagate_prio_process)
	go_struct.Prolog_flags = uint16(c_struct.prolog_flags)
	go_struct.Propagate_rlimits = C.GoString(c_struct.propagate_rlimits)
	go_struct.Propagate_rlimits_except = C.GoString(c_struct.propagate_rlimits_except)
	go_struct.Reboot_program = C.GoString(c_struct.reboot_program)
	go_struct.Reconfig_flags = uint16(c_struct.reconfig_flags)
	go_struct.Requeue_exit = C.GoString(c_struct.requeue_exit)
	go_struct.Requeue_exit_hold = C.GoString(c_struct.requeue_exit_hold)
	go_struct.Resume_fail_program = C.GoString(c_struct.resume_fail_program)
	go_struct.Resume_program = C.GoString(c_struct.resume_program)
	go_struct.Resume_rate = uint16(c_struct.resume_rate)
	go_struct.Resume_timeout = uint16(c_struct.resume_timeout)
	go_struct.Resv_epilog = C.GoString(c_struct.resv_epilog)
	go_struct.Resv_over_run = uint16(c_struct.resv_over_run)
	go_struct.Resv_prolog = C.GoString(c_struct.resv_prolog)
	go_struct.Ret2service = uint16(c_struct.ret2service)
	go_struct.Route_plugin = C.GoString(c_struct.route_plugin)
	go_struct.Salloc_default_command = C.GoString(c_struct.salloc_default_command)
	go_struct.Sbcast_parameters = C.GoString(c_struct.sbcast_parameters)
	go_struct.Sched_logfile = C.GoString(c_struct.sched_logfile)
	go_struct.Sched_log_level = uint16(c_struct.sched_log_level)
	go_struct.Sched_params = C.GoString(c_struct.sched_params)
	go_struct.Sched_time_slice = uint16(c_struct.sched_time_slice)
	go_struct.Schedtype = C.GoString(c_struct.schedtype)
	go_struct.Select_type = C.GoString(c_struct.select_type)
	go_struct.Select_type_param = uint16(c_struct.select_type_param)
	go_struct.Site_factor_plugin = C.GoString(c_struct.site_factor_plugin)
	go_struct.Site_factor_params = C.GoString(c_struct.site_factor_params)
	go_struct.Slurm_conf = C.GoString(c_struct.slurm_conf)
	go_struct.Slurm_user_id = uint32(c_struct.slurm_user_id)
	go_struct.Slurm_user_name = C.GoString(c_struct.slurm_user_name)
	go_struct.Slurmd_user_id = uint32(c_struct.slurmd_user_id)
	go_struct.Slurmd_user_name = C.GoString(c_struct.slurmd_user_name)
	go_struct.Slurmctld_addr = C.GoString(c_struct.slurmctld_addr)
	go_struct.Slurmctld_debug = uint16(c_struct.slurmctld_debug)
	go_struct.Slurmctld_logfile = C.GoString(c_struct.slurmctld_logfile)
	go_struct.Slurmctld_pidfile = C.GoString(c_struct.slurmctld_pidfile)
	go_struct.Slurmctld_plugstack = C.GoString(c_struct.slurmctld_plugstack)
	go_struct.Slurmctld_port = uint32(c_struct.slurmctld_port)
	go_struct.Slurmctld_port_count = uint16(c_struct.slurmctld_port_count)
	go_struct.Slurmctld_primary_off_prog = C.GoString(c_struct.slurmctld_primary_off_prog)
	go_struct.Slurmctld_primary_on_prog = C.GoString(c_struct.slurmctld_primary_on_prog)
	go_struct.Slurmctld_syslog_debug = uint16(c_struct.slurmctld_syslog_debug)
	go_struct.Slurmctld_timeout = uint16(c_struct.slurmctld_timeout)
	go_struct.Slurmctld_params = C.GoString(c_struct.slurmctld_params)
	go_struct.Slurmd_debug = uint16(c_struct.slurmd_debug)
	go_struct.Slurmd_logfile = C.GoString(c_struct.slurmd_logfile)
	go_struct.Slurmd_params = C.GoString(c_struct.slurmd_params)
	go_struct.Slurmd_pidfile = C.GoString(c_struct.slurmd_pidfile)
	go_struct.Slurmd_port = uint32(c_struct.slurmd_port)
	go_struct.Slurmd_spooldir = C.GoString(c_struct.slurmd_spooldir)
	go_struct.Slurmd_syslog_debug = uint16(c_struct.slurmd_syslog_debug)
	go_struct.Slurmd_timeout = uint16(c_struct.slurmd_timeout)
	go_struct.Srun_epilog = C.GoString(c_struct.srun_epilog)
	go_struct.Srun_port_range = uint16(C.uint16_ptr(c_struct.srun_port_range))
	go_struct.Srun_prolog = C.GoString(c_struct.srun_prolog)
	go_struct.State_save_location = C.GoString(c_struct.state_save_location)
	go_struct.Suspend_exc_nodes = C.GoString(c_struct.suspend_exc_nodes)
	go_struct.Suspend_exc_parts = C.GoString(c_struct.suspend_exc_parts)
	go_struct.Suspend_program = C.GoString(c_struct.suspend_program)
	go_struct.Suspend_rate = uint16(c_struct.suspend_rate)
	go_struct.Suspend_time = uint32(c_struct.suspend_time)
	go_struct.Suspend_timeout = uint16(c_struct.suspend_timeout)
	go_struct.Switch_type = C.GoString(c_struct.switch_type)
	go_struct.Task_epilog = C.GoString(c_struct.task_epilog)
	go_struct.Task_plugin = C.GoString(c_struct.task_plugin)
	go_struct.Task_plugin_param = uint32(c_struct.task_plugin_param)
	go_struct.Task_prolog = C.GoString(c_struct.task_prolog)
	go_struct.Tmp_fs = C.GoString(c_struct.tmp_fs)
	go_struct.Topology_param = C.GoString(c_struct.topology_param)
	go_struct.Topology_plugin = C.GoString(c_struct.topology_plugin)
	go_struct.Track_wckey = uint16(c_struct.track_wckey)
	go_struct.Tree_width = uint16(c_struct.tree_width)
	go_struct.Unkillable_program = C.GoString(c_struct.unkillable_program)
	go_struct.Unkillable_timeout = uint16(c_struct.unkillable_timeout)
	go_struct.Use_pam = uint16(c_struct.use_pam)
	go_struct.Use_spec_resources = uint16(c_struct.use_spec_resources)
	go_struct.Version = C.GoString(c_struct.version)
	go_struct.Vsize_factor = uint16(c_struct.vsize_factor)
	go_struct.Wait_time = uint16(c_struct.wait_time)
	go_struct.X11_params = C.GoString(c_struct.x11_params)
	return go_struct
 }

 func Print_Ctl_conf(go_struct Ctl_conf){
	 fmt.Printf("%s:\t %d\n","last update", go_struct.Last_update)
	 fmt.Printf("%s:\t %s\n","accounting storage tres", go_struct.Accounting_storage_tres)
	 fmt.Printf("%s:\t %d\n","accounting storage enforce", go_struct.Accounting_storage_enforce)
	 fmt.Printf("%s:\t %s\n","accounting storage backup host", go_struct.Accounting_storage_backup_host)
	 fmt.Printf("%s:\t %s\n","accounting storage host", go_struct.Accounting_storage_host)
	 fmt.Printf("%s:\t %s\n","accounting storage loc", go_struct.Accounting_storage_loc)
	 fmt.Printf("%s:\t %s\n","accounting storage pass", go_struct.Accounting_storage_pass)
	 fmt.Printf("%s:\t %d\n","accounting storage port", go_struct.Accounting_storage_port)
	 fmt.Printf("%s:\t %s\n","accounting storage type", go_struct.Accounting_storage_type)
	 fmt.Printf("%s:\t %s\n","accounting storage user", go_struct.Accounting_storage_user)
	 fmt.Printf("%s:\t %d\n","acctng store job comment", go_struct.Acctng_store_job_comment)
	 fmt.Printf("%s:\t %s\n","acct gather energy type", go_struct.Acct_gather_energy_type)
	 fmt.Printf("%s:\t %s\n","acct gather profile type", go_struct.Acct_gather_profile_type)
	 fmt.Printf("%s:\t %s\n","acct gather interconnect type", go_struct.Acct_gather_interconnect_type)
	 fmt.Printf("%s:\t %s\n","acct gather filesystem type", go_struct.Acct_gather_filesystem_type)
	 fmt.Printf("%s:\t %d\n","acct gather node freq", go_struct.Acct_gather_node_freq)
	 fmt.Printf("%s:\t %s\n","authalttypes", go_struct.Authalttypes)
	 fmt.Printf("%s:\t %s\n","authinfo", go_struct.Authinfo)
	 fmt.Printf("%s:\t %s\n","authtype", go_struct.Authtype)
	 fmt.Printf("%s:\t %d\n","batch start timeout", go_struct.Batch_start_timeout)
	 fmt.Printf("%s:\t %s\n","bb type", go_struct.Bb_type)
	 fmt.Printf("%s:\t %d\n","boot time", go_struct.Boot_time)
	 fmt.Printf("%s:\t %s\n","checkpoint type", go_struct.Checkpoint_type)
	 fmt.Printf("%s:\t %s\n","cli filter plugins", go_struct.Cli_filter_plugins)
	 fmt.Printf("%s:\t %s\n","core spec plugin", go_struct.Core_spec_plugin)
	 fmt.Printf("%s:\t %s\n","cluster name", go_struct.Cluster_name)
	 fmt.Printf("%s:\t %s\n","comm params", go_struct.Comm_params)
	 fmt.Printf("%s:\t %d\n","complete wait", go_struct.Complete_wait)
	 fmt.Printf("%s:\t %d\n","control cnt", go_struct.Control_cnt)
	 fmt.Printf("%s:\t %d\n","cpu freq def", go_struct.Cpu_freq_def)
	 fmt.Printf("%s:\t %d\n","cpu freq govs", go_struct.Cpu_freq_govs)
	 fmt.Printf("%s:\t %s\n","cred type", go_struct.Cred_type)
	 fmt.Printf("%s:\t %d\n","debug flags", go_struct.Debug_flags)
	 fmt.Printf("%s:\t %d\n","def mem per cpu", go_struct.Def_mem_per_cpu)
	 fmt.Printf("%s:\t %d\n","disable root jobs", go_struct.Disable_root_jobs)
	 fmt.Printf("%s:\t %d\n","eio timeout", go_struct.Eio_timeout)
	 fmt.Printf("%s:\t %d\n","enforce part limits", go_struct.Enforce_part_limits)
	 fmt.Printf("%s:\t %s\n","epilog", go_struct.Epilog)
	 fmt.Printf("%s:\t %d\n","epilog msg time", go_struct.Epilog_msg_time)
	 fmt.Printf("%s:\t %s\n","epilog slurmctld", go_struct.Epilog_slurmctld)
	 fmt.Printf("%s:\t %s\n","ext sensors type", go_struct.Ext_sensors_type)
	 fmt.Printf("%s:\t %d\n","ext sensors freq", go_struct.Ext_sensors_freq)
	 fmt.Printf("%s:\t %d\n","fast schedule", go_struct.Fast_schedule)
	 fmt.Printf("%s:\t %s\n","fed params", go_struct.Fed_params)
	 fmt.Printf("%s:\t %d\n","first job id", go_struct.First_job_id)
	 fmt.Printf("%s:\t %d\n","fs dampening factor", go_struct.Fs_dampening_factor)
	 fmt.Printf("%s:\t %d\n","get env timeout", go_struct.Get_env_timeout)
	 fmt.Printf("%s:\t %s\n","gres plugins", go_struct.Gres_plugins)
	 fmt.Printf("%s:\t %d\n","group time", go_struct.Group_time)
	 fmt.Printf("%s:\t %d\n","group force", go_struct.Group_force)
	 fmt.Printf("%s:\t %s\n","gpu freq def", go_struct.Gpu_freq_def)
	 fmt.Printf("%s:\t %d\n","hash val", go_struct.Hash_val)
	 fmt.Printf("%s:\t %d\n","health check interval", go_struct.Health_check_interval)
	 fmt.Printf("%s:\t %d\n","health check node state", go_struct.Health_check_node_state)
	 fmt.Printf("%s:\t %s\n","health check program", go_struct.Health_check_program)
	 fmt.Printf("%s:\t %d\n","inactive limit", go_struct.Inactive_limit)
	 fmt.Printf("%s:\t %s\n","job acct gather freq", go_struct.Job_acct_gather_freq)
	 fmt.Printf("%s:\t %s\n","job acct gather type", go_struct.Job_acct_gather_type)
	 fmt.Printf("%s:\t %s\n","job acct gather params", go_struct.Job_acct_gather_params)
	 fmt.Printf("%s:\t %d\n","job acct oom kill", go_struct.Job_acct_oom_kill)
	 fmt.Printf("%s:\t %s\n","job ckpt dir", go_struct.Job_ckpt_dir)
	 fmt.Printf("%s:\t %s\n","job comp host", go_struct.Job_comp_host)
	 fmt.Printf("%s:\t %s\n","job comp loc", go_struct.Job_comp_loc)
	 fmt.Printf("%s:\t %s\n","job comp pass", go_struct.Job_comp_pass)
	 fmt.Printf("%s:\t %d\n","job comp port", go_struct.Job_comp_port)
	 fmt.Printf("%s:\t %s\n","job comp type", go_struct.Job_comp_type)
	 fmt.Printf("%s:\t %s\n","job comp user", go_struct.Job_comp_user)
	 fmt.Printf("%s:\t %s\n","job container plugin", go_struct.Job_container_plugin)
	 fmt.Printf("%s:\t %s\n","job credential private key", go_struct.Job_credential_private_key)
	 fmt.Printf("%s:\t %s\n","job credential public certificate", go_struct.Job_credential_public_certificate)
	 fmt.Printf("%s:\t %d\n","job file append", go_struct.Job_file_append)
	 fmt.Printf("%s:\t %d\n","job requeue", go_struct.Job_requeue)
	 fmt.Printf("%s:\t %s\n","job submit plugins", go_struct.Job_submit_plugins)
	 fmt.Printf("%s:\t %d\n","keep alive time", go_struct.Keep_alive_time)
	 fmt.Printf("%s:\t %d\n","kill on bad exit", go_struct.Kill_on_bad_exit)
	 fmt.Printf("%s:\t %d\n","kill wait", go_struct.Kill_wait)
	 fmt.Printf("%s:\t %s\n","launch params", go_struct.Launch_params)
	 fmt.Printf("%s:\t %s\n","launch type", go_struct.Launch_type)
	 fmt.Printf("%s:\t %s\n","layouts", go_struct.Layouts)
	 fmt.Printf("%s:\t %s\n","licenses", go_struct.Licenses)
	 fmt.Printf("%s:\t %s\n","licenses used", go_struct.Licenses_used)
	 fmt.Printf("%s:\t %d\n","log fmt", go_struct.Log_fmt)
	 fmt.Printf("%s:\t %s\n","mail domain", go_struct.Mail_domain)
	 fmt.Printf("%s:\t %s\n","mail prog", go_struct.Mail_prog)
	 fmt.Printf("%s:\t %d\n","max array sz", go_struct.Max_array_sz)
	 fmt.Printf("%s:\t %d\n","max job cnt", go_struct.Max_job_cnt)
	 fmt.Printf("%s:\t %d\n","max job id", go_struct.Max_job_id)
	 fmt.Printf("%s:\t %d\n","max mem per cpu", go_struct.Max_mem_per_cpu)
	 fmt.Printf("%s:\t %d\n","max step cnt", go_struct.Max_step_cnt)
	 fmt.Printf("%s:\t %d\n","max tasks per node", go_struct.Max_tasks_per_node)
	 fmt.Printf("%s:\t %s\n","mcs plugin", go_struct.Mcs_plugin)
	 fmt.Printf("%s:\t %s\n","mcs plugin params", go_struct.Mcs_plugin_params)
	 fmt.Printf("%s:\t %d\n","min job age", go_struct.Min_job_age)
	 fmt.Printf("%s:\t %s\n","mpi default", go_struct.Mpi_default)
	 fmt.Printf("%s:\t %s\n","mpi params", go_struct.Mpi_params)
	 fmt.Printf("%s:\t %s\n","msg aggr params", go_struct.Msg_aggr_params)
	 fmt.Printf("%s:\t %d\n","msg timeout", go_struct.Msg_timeout)
	 fmt.Printf("%s:\t %d\n","tcp timeout", go_struct.Tcp_timeout)
	 fmt.Printf("%s:\t %d\n","next job id", go_struct.Next_job_id)
	 fmt.Printf("%s:\t %s\n","node features plugins", go_struct.Node_features_plugins)
	 fmt.Printf("%s:\t %s\n","node prefix", go_struct.Node_prefix)
	 fmt.Printf("%s:\t %d\n","over time limit", go_struct.Over_time_limit)
	 fmt.Printf("%s:\t %s\n","plugindir", go_struct.Plugindir)
	 fmt.Printf("%s:\t %s\n","plugstack", go_struct.Plugstack)
	 fmt.Printf("%s:\t %s\n","power parameters", go_struct.Power_parameters)
	 fmt.Printf("%s:\t %s\n","power plugin", go_struct.Power_plugin)
	 fmt.Printf("%s:\t %d\n","preempt exempt time", go_struct.Preempt_exempt_time)
	 fmt.Printf("%s:\t %d\n","preempt mode", go_struct.Preempt_mode)
	 fmt.Printf("%s:\t %s\n","preempt type", go_struct.Preempt_type)
	 fmt.Printf("%s:\t %d\n","priority decay hl", go_struct.Priority_decay_hl)
	 fmt.Printf("%s:\t %d\n","priority calc period", go_struct.Priority_calc_period)
	 fmt.Printf("%s:\t %d\n","priority favor small", go_struct.Priority_favor_small)
	 fmt.Printf("%s:\t %d\n","priority flags", go_struct.Priority_flags)
	 fmt.Printf("%s:\t %d\n","priority max age", go_struct.Priority_max_age)
	 fmt.Printf("%s:\t %s\n","priority params", go_struct.Priority_params)
	 fmt.Printf("%s:\t %d\n","priority reset period", go_struct.Priority_reset_period)
	 fmt.Printf("%s:\t %s\n","priority type", go_struct.Priority_type)
	 fmt.Printf("%s:\t %d\n","priority weight age", go_struct.Priority_weight_age)
	 fmt.Printf("%s:\t %d\n","priority weight assoc", go_struct.Priority_weight_assoc)
	 fmt.Printf("%s:\t %d\n","priority weight fs", go_struct.Priority_weight_fs)
	 fmt.Printf("%s:\t %d\n","priority weight js", go_struct.Priority_weight_js)
	 fmt.Printf("%s:\t %d\n","priority weight part", go_struct.Priority_weight_part)
	 fmt.Printf("%s:\t %d\n","priority weight qos", go_struct.Priority_weight_qos)
	 fmt.Printf("%s:\t %s\n","priority weight tres", go_struct.Priority_weight_tres)
	 fmt.Printf("%s:\t %d\n","private data", go_struct.Private_data)
	 fmt.Printf("%s:\t %s\n","proctrack type", go_struct.Proctrack_type)
	 fmt.Printf("%s:\t %s\n","prolog", go_struct.Prolog)
	 fmt.Printf("%s:\t %d\n","prolog epilog timeout", go_struct.Prolog_epilog_timeout)
	 fmt.Printf("%s:\t %s\n","prolog slurmctld", go_struct.Prolog_slurmctld)
	 fmt.Printf("%s:\t %d\n","propagate prio process", go_struct.Propagate_prio_process)
	 fmt.Printf("%s:\t %d\n","prolog flags", go_struct.Prolog_flags)
	 fmt.Printf("%s:\t %s\n","propagate rlimits", go_struct.Propagate_rlimits)
	 fmt.Printf("%s:\t %s\n","propagate rlimits except", go_struct.Propagate_rlimits_except)
	 fmt.Printf("%s:\t %s\n","reboot program", go_struct.Reboot_program)
	 fmt.Printf("%s:\t %d\n","reconfig flags", go_struct.Reconfig_flags)
	 fmt.Printf("%s:\t %s\n","requeue exit", go_struct.Requeue_exit)
	 fmt.Printf("%s:\t %s\n","requeue exit hold", go_struct.Requeue_exit_hold)
	 fmt.Printf("%s:\t %s\n","resume fail program", go_struct.Resume_fail_program)
	 fmt.Printf("%s:\t %s\n","resume program", go_struct.Resume_program)
	 fmt.Printf("%s:\t %d\n","resume rate", go_struct.Resume_rate)
	 fmt.Printf("%s:\t %d\n","resume timeout", go_struct.Resume_timeout)
	 fmt.Printf("%s:\t %s\n","resv epilog", go_struct.Resv_epilog)
	 fmt.Printf("%s:\t %d\n","resv over run", go_struct.Resv_over_run)
	 fmt.Printf("%s:\t %s\n","resv prolog", go_struct.Resv_prolog)
	 fmt.Printf("%s:\t %d\n","ret2service", go_struct.Ret2service)
	 fmt.Printf("%s:\t %s\n","route plugin", go_struct.Route_plugin)
	 fmt.Printf("%s:\t %s\n","salloc default command", go_struct.Salloc_default_command)
	 fmt.Printf("%s:\t %s\n","sbcast parameters", go_struct.Sbcast_parameters)
	 fmt.Printf("%s:\t %s\n","sched logfile", go_struct.Sched_logfile)
	 fmt.Printf("%s:\t %d\n","sched log level", go_struct.Sched_log_level)
	 fmt.Printf("%s:\t %s\n","sched params", go_struct.Sched_params)
	 fmt.Printf("%s:\t %d\n","sched time slice", go_struct.Sched_time_slice)
	 fmt.Printf("%s:\t %s\n","schedtype", go_struct.Schedtype)
	 fmt.Printf("%s:\t %s\n","select type", go_struct.Select_type)
	 fmt.Printf("%s:\t %d\n","select type param", go_struct.Select_type_param)
	 fmt.Printf("%s:\t %s\n","site factor plugin", go_struct.Site_factor_plugin)
	 fmt.Printf("%s:\t %s\n","site factor params", go_struct.Site_factor_params)
	 fmt.Printf("%s:\t %s\n","slurm conf", go_struct.Slurm_conf)
	 fmt.Printf("%s:\t %d\n","slurm user id", go_struct.Slurm_user_id)
	 fmt.Printf("%s:\t %s\n","slurm user name", go_struct.Slurm_user_name)
	 fmt.Printf("%s:\t %d\n","slurmd user id", go_struct.Slurmd_user_id)
	 fmt.Printf("%s:\t %s\n","slurmd user name", go_struct.Slurmd_user_name)
	 fmt.Printf("%s:\t %s\n","slurmctld addr", go_struct.Slurmctld_addr)
	 fmt.Printf("%s:\t %d\n","slurmctld debug", go_struct.Slurmctld_debug)
	 fmt.Printf("%s:\t %s\n","slurmctld logfile", go_struct.Slurmctld_logfile)
	 fmt.Printf("%s:\t %s\n","slurmctld pidfile", go_struct.Slurmctld_pidfile)
	 fmt.Printf("%s:\t %s\n","slurmctld plugstack", go_struct.Slurmctld_plugstack)
	 fmt.Printf("%s:\t %d\n","slurmctld port", go_struct.Slurmctld_port)
	 fmt.Printf("%s:\t %d\n","slurmctld port count", go_struct.Slurmctld_port_count)
	 fmt.Printf("%s:\t %s\n","slurmctld primary off prog", go_struct.Slurmctld_primary_off_prog)
	 fmt.Printf("%s:\t %s\n","slurmctld primary on prog", go_struct.Slurmctld_primary_on_prog)
	 fmt.Printf("%s:\t %d\n","slurmctld syslog debug", go_struct.Slurmctld_syslog_debug)
	 fmt.Printf("%s:\t %d\n","slurmctld timeout", go_struct.Slurmctld_timeout)
	 fmt.Printf("%s:\t %s\n","slurmctld params", go_struct.Slurmctld_params)
	 fmt.Printf("%s:\t %d\n","slurmd debug", go_struct.Slurmd_debug)
	 fmt.Printf("%s:\t %s\n","slurmd logfile", go_struct.Slurmd_logfile)
	 fmt.Printf("%s:\t %s\n","slurmd params", go_struct.Slurmd_params)
	 fmt.Printf("%s:\t %s\n","slurmd pidfile", go_struct.Slurmd_pidfile)
	 fmt.Printf("%s:\t %d\n","slurmd port", go_struct.Slurmd_port)
	 fmt.Printf("%s:\t %s\n","slurmd spooldir", go_struct.Slurmd_spooldir)
	 fmt.Printf("%s:\t %d\n","slurmd syslog debug", go_struct.Slurmd_syslog_debug)
	 fmt.Printf("%s:\t %d\n","slurmd timeout", go_struct.Slurmd_timeout)
	 fmt.Printf("%s:\t %s\n","srun epilog", go_struct.Srun_epilog)
	 fmt.Printf("%s:\t %d\n","srun port range", go_struct.Srun_port_range)
	 fmt.Printf("%s:\t %s\n","srun prolog", go_struct.Srun_prolog)
	 fmt.Printf("%s:\t %s\n","state save location", go_struct.State_save_location)
	 fmt.Printf("%s:\t %s\n","suspend exc nodes", go_struct.Suspend_exc_nodes)
	 fmt.Printf("%s:\t %s\n","suspend exc parts", go_struct.Suspend_exc_parts)
	 fmt.Printf("%s:\t %s\n","suspend program", go_struct.Suspend_program)
	 fmt.Printf("%s:\t %d\n","suspend rate", go_struct.Suspend_rate)
	 fmt.Printf("%s:\t %d\n","suspend time", go_struct.Suspend_time)
	 fmt.Printf("%s:\t %d\n","suspend timeout", go_struct.Suspend_timeout)
	 fmt.Printf("%s:\t %s\n","switch type", go_struct.Switch_type)
	 fmt.Printf("%s:\t %s\n","task epilog", go_struct.Task_epilog)
	 fmt.Printf("%s:\t %s\n","task plugin", go_struct.Task_plugin)
	 fmt.Printf("%s:\t %d\n","task plugin param", go_struct.Task_plugin_param)
	 fmt.Printf("%s:\t %s\n","task prolog", go_struct.Task_prolog)
	 fmt.Printf("%s:\t %s\n","tmp fs", go_struct.Tmp_fs)
	 fmt.Printf("%s:\t %s\n","topology param", go_struct.Topology_param)
	 fmt.Printf("%s:\t %s\n","topology plugin", go_struct.Topology_plugin)
	 fmt.Printf("%s:\t %d\n","track wckey", go_struct.Track_wckey)
	 fmt.Printf("%s:\t %d\n","tree width", go_struct.Tree_width)
	 fmt.Printf("%s:\t %s\n","unkillable program", go_struct.Unkillable_program)
	 fmt.Printf("%s:\t %d\n","unkillable timeout", go_struct.Unkillable_timeout)
	 fmt.Printf("%s:\t %d\n","use pam", go_struct.Use_pam)
	 fmt.Printf("%s:\t %d\n","use spec resources", go_struct.Use_spec_resources)
	 fmt.Printf("%s:\t %s\n","version", go_struct.Version)
	 fmt.Printf("%s:\t %d\n","vsize factor", go_struct.Vsize_factor)
	 fmt.Printf("%s:\t %d\n","wait time", go_struct.Wait_time)
	 fmt.Printf("%s:\t %s\n","x11 params", go_struct.X11_params)
}

func Version() int {

	ver := C.long(0)
	ver = C.slurm_api_version()


	return int(ver)
}
func VersionString(v int) string {

	var major, minor, micro int 
	var version string
	major = ((v >> 16) & 0xff)
	minor =  ((v >> 8) & 0xff)
	micro = ( v & 0xff)
	version = fmt.Sprintf("%d.%d.%d",major, minor, micro);
	return version
}


func GetConfig() Ctl_conf  {

	c_config := C.get_config (C.uint64_t(0))
	ret_config :=  Ctl_conf_convert_c_to_go(c_config) 
	return ret_config
}

func GetErrorString( errno uint32 ) string {
	msg := C.GoString(C.slurm_strerror(C.int(errno)))
	return msg
}
