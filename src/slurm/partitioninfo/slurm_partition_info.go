package partition_info
/*
#cgo LDFLAGS: -lslurm
#include<stdlib.h>
#include<slurm/slurm.h>


uint8_t uint8_ptr(uint8_t* pointer) { 
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
int8_t int8_ptr(int8_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
uint16_t uint16_ptr(uint16_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
int16_t int16_ptr(int16_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
uint32_t uint32_ptr(uint32_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
int32_t int32_ptr(int32_t* pointer, int off) {
    if (NULL == pointer) {
    return -1;}
    pointer+=off;
    return *pointer;
}
uint64_t uint64_ptr(uint64_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}
int64_t int64_ptr(int16_t* pointer) {
    if (NULL == pointer) {
    return -1;}
    return *pointer;
}

 struct partition_info_msg *get_partition_info(){
   	struct partition_info_msg* partition_buffer;
        if( slurm_load_partitions ((time_t) NULL,
		&partition_buffer, SHOW_ALL))
		return NULL;
	return partition_buffer;
 }
 struct partition_info* partition_from_list(struct partition_info_msg *list, int i){
	 return &list->partition_array[i];
}
 void free_partition_buffer(void* buffer){	

	slurm_free_partition_info_msg ((struct partition_info_msg*)buffer);
  }
  int find_node_inx(int32_t* node){
	int ret = 0;
	while(*node != -1) { node++; ret++;};
	return ret;
}


*/
import "C"

import "fmt"

type Partition_info struct {
	Allow_alloc_nodes string;
	Allow_accounts string;
	Allow_groups string;
	Allow_qos string;
	Alternate string;
	Billing_weights_str string;
	Cluster_name string;
	Cr_type uint16;
	Cpu_bind uint32;
	Def_mem_per_cpu uint64;
	Default_time uint32;
	Deny_accounts string;
	Deny_qos string;
	Flags uint16;
	Grace_time uint32;
	Job_defaults_str string;
	Max_cpus_per_node uint32;
	Max_mem_per_cpu uint64;
	Max_nodes uint32;
	Max_share uint16;
	Max_time uint32;
	Min_nodes uint32;
	Name string;
	Node_inx[] int32;
	Nodes string;
	Over_time_limit uint16;
	Preempt_mode uint16;
	Priority_job_factor uint16;
	Priority_tier uint16;
	Qos_char string;
	State_up uint16;
	Total_cpus uint32;
	Total_nodes uint32;
	Tres_fmt_str string;
}
func Partition_info_convert_c_to_go(c_struct *C.struct_partition_info) Partition_info{
	var go_struct Partition_info

	go_struct.Allow_alloc_nodes = C.GoString(c_struct.allow_alloc_nodes)
	go_struct.Allow_accounts = C.GoString(c_struct.allow_accounts)
	go_struct.Allow_groups = C.GoString(c_struct.allow_groups)
	go_struct.Allow_qos = C.GoString(c_struct.allow_qos)
	go_struct.Alternate = C.GoString(c_struct.alternate)
	go_struct.Billing_weights_str = C.GoString(c_struct.billing_weights_str)
	go_struct.Cluster_name = C.GoString(c_struct.cluster_name)
	go_struct.Cr_type = uint16(c_struct.cr_type)
	go_struct.Cpu_bind = uint32(c_struct.cpu_bind)
	go_struct.Def_mem_per_cpu = uint64(c_struct.def_mem_per_cpu)
	go_struct.Default_time = uint32(c_struct.default_time)
	go_struct.Deny_accounts = C.GoString(c_struct.deny_accounts)
	go_struct.Deny_qos = C.GoString(c_struct.deny_qos)
	go_struct.Flags = uint16(c_struct.flags)
	go_struct.Grace_time = uint32(c_struct.grace_time)
	go_struct.Job_defaults_str = C.GoString(c_struct.job_defaults_str)
	go_struct.Max_cpus_per_node = uint32(c_struct.max_cpus_per_node)
	go_struct.Max_mem_per_cpu = uint64(c_struct.max_mem_per_cpu)
	go_struct.Max_nodes = uint32(c_struct.max_nodes)
	go_struct.Max_share = uint16(c_struct.max_share)
	go_struct.Max_time = uint32(c_struct.max_time)
	go_struct.Min_nodes = uint32(c_struct.min_nodes)
	go_struct.Name = C.GoString(c_struct.name)
	t := C.find_node_inx(c_struct.node_inx)

	fmt.Printf("%d", t)
	go_struct.Node_inx = make([]int32, t,t)
	for i:=int32(0); i<int32(t) ; i++{
		go_struct.Node_inx[i] = int32(C.int32_ptr(c_struct.node_inx,C.int(i)))

	}
	go_struct.Nodes = C.GoString(c_struct.nodes)
	go_struct.Over_time_limit = uint16(c_struct.over_time_limit)
	go_struct.Preempt_mode = uint16(c_struct.preempt_mode)
	go_struct.Priority_job_factor = uint16(c_struct.priority_job_factor)
	go_struct.Priority_tier = uint16(c_struct.priority_tier)
	go_struct.Qos_char = C.GoString(c_struct.qos_char)
	go_struct.State_up = uint16(c_struct.state_up)
	go_struct.Total_cpus = uint32(c_struct.total_cpus)
	go_struct.Total_nodes = uint32(c_struct.total_nodes)
	go_struct.Tres_fmt_str = C.GoString(c_struct.tres_fmt_str)
	return go_struct
 }
 func Print_Partition_info(go_struct Partition_info){
	 fmt.Printf("%s:\t %s\n","allow alloc nodes", go_struct.Allow_alloc_nodes)
	 fmt.Printf("%s:\t %s\n","allow accounts", go_struct.Allow_accounts)
	 fmt.Printf("%s:\t %s\n","allow groups", go_struct.Allow_groups)
	 fmt.Printf("%s:\t %s\n","allow qos", go_struct.Allow_qos)
	 fmt.Printf("%s:\t %s\n","alternate", go_struct.Alternate)
	 fmt.Printf("%s:\t %s\n","billing weights str", go_struct.Billing_weights_str)
	 fmt.Printf("%s:\t %s\n","cluster name", go_struct.Cluster_name)
	 fmt.Printf("%s:\t %d\n","cr type", go_struct.Cr_type)
	 fmt.Printf("%s:\t %d\n","cpu bind", go_struct.Cpu_bind)
	 fmt.Printf("%s:\t %d\n","def mem per cpu", go_struct.Def_mem_per_cpu)
	 fmt.Printf("%s:\t %d\n","default time", go_struct.Default_time)
	 fmt.Printf("%s:\t %s\n","deny accounts", go_struct.Deny_accounts)
	 fmt.Printf("%s:\t %s\n","deny qos", go_struct.Deny_qos)
	 fmt.Printf("%s:\t %d\n","flags", go_struct.Flags)
	 fmt.Printf("%s:\t %d\n","grace time", go_struct.Grace_time)
	 fmt.Printf("%s:\t %s\n","job defaults str", go_struct.Job_defaults_str)
	 fmt.Printf("%s:\t %d\n","max cpus per node", go_struct.Max_cpus_per_node)
	 fmt.Printf("%s:\t %d\n","max mem per cpu", go_struct.Max_mem_per_cpu)
	 fmt.Printf("%s:\t %d\n","max nodes", go_struct.Max_nodes)
	 fmt.Printf("%s:\t %d\n","max share", go_struct.Max_share)
	 fmt.Printf("%s:\t %d\n","max time", go_struct.Max_time)
	 fmt.Printf("%s:\t %d\n","min nodes", go_struct.Min_nodes)
	 fmt.Printf("%s:\t %s\n","name", go_struct.Name)
	 fmt.Printf("%s:\t %d\n","node inx", go_struct.Node_inx)
	 fmt.Printf("%s:\t %s\n","nodes", go_struct.Nodes)
	 fmt.Printf("%s:\t %d\n","over time limit", go_struct.Over_time_limit)
	 fmt.Printf("%s:\t %d\n","preempt mode", go_struct.Preempt_mode)
	 fmt.Printf("%s:\t %d\n","priority job factor", go_struct.Priority_job_factor)
	 fmt.Printf("%s:\t %d\n","priority tier", go_struct.Priority_tier)
	 fmt.Printf("%s:\t %s\n","qos char", go_struct.Qos_char)
	 fmt.Printf("%s:\t %d\n","state up", go_struct.State_up)
	 fmt.Printf("%s:\t %d\n","total cpus", go_struct.Total_cpus)
	 fmt.Printf("%s:\t %d\n","total nodes", go_struct.Total_nodes)
	 fmt.Printf("%s:\t %s\n","tres fmt str", go_struct.Tres_fmt_str)
}

type Partition_info_msg struct {
	Last_update int64;
	Record_count uint32;
	Partition_list  []Partition_info;
}
func Get_partitions() Partition_info_msg {
	var go_partition_buffer Partition_info_msg
	c_partition_buffer := C.get_partition_info()
	if c_partition_buffer == nil{
		go_partition_buffer.Last_update = int64(0)
		go_partition_buffer.Record_count = uint32(0)
		return 	go_partition_buffer; 
	}
	go_partition_buffer.Last_update = int64(c_partition_buffer.last_update)
	go_partition_buffer.Record_count = uint32(c_partition_buffer.record_count)
	go_partition_buffer.Partition_list =make([]Partition_info,c_partition_buffer.record_count, c_partition_buffer.record_count)
	for i:=uint32(0); i<go_partition_buffer.Record_count; i++ {
		partition := C.partition_from_list(c_partition_buffer,  C.int(i))
		go_partition :=  Partition_info_convert_c_to_go(partition)
		go_partition_buffer.Partition_list[i]=go_partition
	}
	C.slurm_free_partition_info_msg (c_partition_buffer);

	return go_partition_buffer
}
