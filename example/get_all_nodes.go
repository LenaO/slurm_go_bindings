package main

import "slurm/nodeinfo"
import "fmt"

func main(){
	node_list := node_info.Get_all_nodes()
	fmt.Printf("Found %d nodes \n", node_list.Record_count)


/* a little bit nicer*/ 
	fmt.Printf("name\t State\t\t\t Reason\t\t Tres\n")
	fmt.Printf("________________________________________\n")
	for i := range node_list.Node_list {
		node := node_list.Node_list[i]
		fmt.Printf("%s\t %s\t %s\t %s\n", node.Node_hostname, node_info.State_to_string(node.Node_state), node.Reason, node.Tres_fmt_str)

	}

}
