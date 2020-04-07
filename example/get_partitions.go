package main

import "slurm/partitioninfo"
import "fmt"

func main(){
	partition_list := partition_info.Get_partitions()
	fmt.Printf("Found %d partions \n", partition_list.Record_count)


/* a little bit nicer */
	fmt.Printf("Name\t Nodes\t\t\t Max_time(min)\t\t Tres\n")
	fmt.Printf("________________________________________\n")
	for i := range partition_list.Partition_list {
		partition := partition_list.Partition_list[i]
		fmt.Printf("%s\t %s\t %d\t %d\n", partition.Name, partition.Nodes, partition.Max_time, partition.Node_inx )

	}

}
