package node_info
/*
#cgo LDFLAGS: -lslurm
#include<stdlib.h>
#include<slurm/slurm.h>
#include<slurm/slurm_errno.h>
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
uint16_t uint16_ptr(uint16_t* pointer) {
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
 struct node_info_msg *get_node_info(){
   	struct node_info_msg* node_buffer;
       if(slurm_load_node ((time_t) NULL,
		&node_buffer, SHOW_ALL))
		return NULL;
	return node_buffer;
 }
  struct node_info_msg *get_single_node_info(char* name){
   	struct node_info_msg* node_buffer;
	if( slurm_load_node_single (&node_buffer, name, SHOW_DETAIL))
		return NULL;
	return node_buffer;
 }
 
 struct node_info* node_from_list(struct node_info_msg *list, int i){
	 return &list->node_array[i];
}
 void free_node_buffer(void* buffer){	

	slurm_free_node_info_msg ((struct node_info_msg*)buffer);
  }


*/
import "C"

import "fmt"
import "unsafe"

type  Node_info struct {
	Arch string;
	Boards uint16;
	Boot_time int64;
	Cluster_name string;
	Cores uint16;
	Core_spec_cnt uint16;
	Cpu_bind uint32;
	Cpu_load uint32;
	Free_mem uint64;
	Cpus uint16;
	Cpu_spec_list string;
	Features string;
	Features_act string;
	Gres string;
	Gres_drain string;
	Gres_used string;
	Mcs_label string;
	Mem_spec_limit uint64;
	Name string;
	Next_state uint32;
	Node_addr string;
	Node_hostname string;
	Node_state uint32;
	Os string;
	Owner uint32;
	Partitions string;
	Port uint16;
	Real_memory uint64;
	Reason string;
	Reason_time int64;
	Reason_uid uint32;
	Slurmd_start_time int64;
	Sockets uint16;
	Threads uint16;
	Tmp_disk uint32;
	Weight uint32;
	Tres_fmt_str string;
	Version string;
}
func Node_info_convert_c_to_go(c_struct *C.struct_node_info) Node_info{
	var go_struct Node_info

	go_struct.Arch = C.GoString(c_struct.arch)
	go_struct.Boards = uint16(c_struct.boards)
	go_struct.Boot_time = int64(c_struct.boot_time)
	go_struct.Cluster_name = C.GoString(c_struct.cluster_name)
	go_struct.Cores = uint16(c_struct.cores)
	go_struct.Core_spec_cnt = uint16(c_struct.core_spec_cnt)
	go_struct.Cpu_bind = uint32(c_struct.cpu_bind)
	go_struct.Cpu_load = uint32(c_struct.cpu_load)
	go_struct.Free_mem = uint64(c_struct.free_mem)
	go_struct.Cpus = uint16(c_struct.cpus)
	go_struct.Cpu_spec_list = C.GoString(c_struct.cpu_spec_list)
	go_struct.Features = C.GoString(c_struct.features)
	go_struct.Features_act = C.GoString(c_struct.features_act)
	go_struct.Gres = C.GoString(c_struct.gres)
	go_struct.Gres_drain = C.GoString(c_struct.gres_drain)
	go_struct.Gres_used = C.GoString(c_struct.gres_used)
	go_struct.Mcs_label = C.GoString(c_struct.mcs_label)
	go_struct.Mem_spec_limit = uint64(c_struct.mem_spec_limit)
	go_struct.Name = C.GoString(c_struct.name)
	go_struct.Next_state = uint32(c_struct.next_state)
	go_struct.Node_addr = C.GoString(c_struct.node_addr)
	go_struct.Node_hostname = C.GoString(c_struct.node_hostname)
	go_struct.Node_state = uint32(c_struct.node_state)
	go_struct.Os = C.GoString(c_struct.os)
	go_struct.Owner = uint32(c_struct.owner)
	go_struct.Partitions = C.GoString(c_struct.partitions)
	go_struct.Port = uint16(c_struct.port)
	go_struct.Real_memory = uint64(c_struct.real_memory)
	go_struct.Reason = C.GoString(c_struct.reason)
	go_struct.Reason_time = int64(c_struct.reason_time)
	go_struct.Reason_uid = uint32(c_struct.reason_uid)
	go_struct.Slurmd_start_time = int64(c_struct.slurmd_start_time)
	go_struct.Sockets = uint16(c_struct.sockets)
	go_struct.Threads = uint16(c_struct.threads)
	go_struct.Tmp_disk = uint32(c_struct.tmp_disk)
	go_struct.Weight = uint32(c_struct.weight)
	go_struct.Tres_fmt_str = C.GoString(c_struct.tres_fmt_str)
	go_struct.Version = C.GoString(c_struct.version)
	return go_struct
 }

 func State_to_string(state uint32) string{

	switch s := C.uint16_t(state); s {
 	case C.NODE_STATE_UNKNOWN:
		 return "node state unknown"
	case C.NODE_STATE_DOWN:
		 return "node state down"
	case C.NODE_STATE_IDLE:
		 return "node state idle"
	case C.NODE_STATE_ALLOCATED:
		 return "node state allocated"
	case C.NODE_STATE_ERROR:
		 return "node state error"
	case C.NODE_STATE_MIXED:
		 return "node state mixed"
	case C.NODE_STATE_FUTURE:
		 return "node state future"
	case C.NODE_STATE_END:
		 return "node state end"
	 }
	return "Unkown state"
 }

 func Print_node_info(go_struct Node_info){
	 fmt.Printf("%s:\t %s\n","arch", go_struct.Arch)
	 fmt.Printf("%s:\t %d\n","boards", go_struct.Boards)
	 fmt.Printf("%s:\t %d\n","boot time", go_struct.Boot_time)
	 fmt.Printf("%s:\t %s\n","cluster name", go_struct.Cluster_name)
	 fmt.Printf("%s:\t %d\n","cores", go_struct.Cores)
	 fmt.Printf("%s:\t %d\n","core spec cnt", go_struct.Core_spec_cnt)
	 fmt.Printf("%s:\t %d\n","cpu bind", go_struct.Cpu_bind)
	 fmt.Printf("%s:\t %d\n","cpu load", go_struct.Cpu_load)
	 fmt.Printf("%s:\t %d\n","free mem", go_struct.Free_mem)
	 fmt.Printf("%s:\t %d\n","cpus", go_struct.Cpus)
	 fmt.Printf("%s:\t %s\n","cpu spec list", go_struct.Cpu_spec_list)
	 fmt.Printf("%s:\t %s\n","features", go_struct.Features)
	 fmt.Printf("%s:\t %s\n","features act", go_struct.Features_act)
	 fmt.Printf("%s:\t %s\n","gres", go_struct.Gres)
	 fmt.Printf("%s:\t %s\n","gres drain", go_struct.Gres_drain)
	 fmt.Printf("%s:\t %s\n","gres used", go_struct.Gres_used)
	 fmt.Printf("%s:\t %s\n","mcs label", go_struct.Mcs_label)
	 fmt.Printf("%s:\t %d\n","mem spec limit", go_struct.Mem_spec_limit)
	 fmt.Printf("%s:\t %s\n","name", go_struct.Name)
	 fmt.Printf("%s:\t %d\n","next state", go_struct.Next_state)
	 fmt.Printf("%s:\t %s\n","node addr", go_struct.Node_addr)
	 fmt.Printf("%s:\t %s\n","node hostname", go_struct.Node_hostname)
	 fmt.Printf("%s:\t %d\n","node state", go_struct.Node_state)
	 fmt.Printf("%s:\t %s\n","os", go_struct.Os)
	 fmt.Printf("%s:\t %d\n","owner", go_struct.Owner)
	 fmt.Printf("%s:\t %s\n","partitions", go_struct.Partitions)
	 fmt.Printf("%s:\t %d\n","port", go_struct.Port)
	 fmt.Printf("%s:\t %d\n","real memory", go_struct.Real_memory)
	 fmt.Printf("%s:\t %s\n","reason", go_struct.Reason)
	 fmt.Printf("%s:\t %d\n","reason time", go_struct.Reason_time)
	 fmt.Printf("%s:\t %d\n","reason uid", go_struct.Reason_uid)
	 fmt.Printf("%s:\t %d\n","slurmd start time", go_struct.Slurmd_start_time)
	 fmt.Printf("%s:\t %d\n","sockets", go_struct.Sockets)
	 fmt.Printf("%s:\t %d\n","threads", go_struct.Threads)
	 fmt.Printf("%s:\t %d\n","tmp disk", go_struct.Tmp_disk)
	 fmt.Printf("%s:\t %d\n","weight", go_struct.Weight)
	 fmt.Printf("%s:\t %s\n","tres fmt str", go_struct.Tres_fmt_str)
	 fmt.Printf("%s:\t %s\n","version", go_struct.Version)
}
type Node_info_msg struct {
	Last_update int64;
	Record_count uint32;
	Error_code uint32;
	Node_list  []Node_info;
}

func Get_all_nodes() Node_info_msg {
	var go_node_buffer Node_info_msg
	c_node_buffer := C.get_node_info()
	if c_node_buffer == nil {
		go_node_buffer.Last_update = int64(0)
		go_node_buffer.Record_count = uint32(0)
		go_node_buffer.Error_code = uint32(C.slurm_get_errno())
		return go_node_buffer
	}
	go_node_buffer.Last_update = int64(c_node_buffer.last_update)
	go_node_buffer.Record_count = uint32(c_node_buffer.record_count)
	go_node_buffer.Node_list =make([]Node_info,c_node_buffer.record_count, c_node_buffer.record_count)
	for i:=uint32(0); i<go_node_buffer.Record_count; i++ {
		node := C.node_from_list(c_node_buffer,  C.int(i))
		go_node :=  Node_info_convert_c_to_go(node)
		go_node_buffer.Node_list[i]=go_node
	}
	C.slurm_free_node_info_msg (c_node_buffer);

	return go_node_buffer
}

func Get_node_info (name string) Node_info_msg {

	var go_node_buffer Node_info_msg
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_node_buffer := C.get_single_node_info(c_name)
	if c_node_buffer == nil {
		go_node_buffer.Last_update = int64(0)
		go_node_buffer.Record_count = uint32(0)
		go_node_buffer.Error_code = uint32(C.slurm_get_errno())

		return go_node_buffer;
	}
	go_node_buffer.Last_update = int64(c_node_buffer.last_update)
	go_node_buffer.Record_count = uint32(c_node_buffer.record_count)
	go_node_buffer.Node_list =make([]Node_info,c_node_buffer.record_count, c_node_buffer.record_count)
	for i:=uint32(0); i<go_node_buffer.Record_count; i++ {
		node := C.node_from_list(c_node_buffer,  C.int(i))
		go_node :=  Node_info_convert_c_to_go(node)
		go_node_buffer.Node_list[i]=go_node
	}
	C.slurm_free_node_info_msg (c_node_buffer);

	return go_node_buffer


}
