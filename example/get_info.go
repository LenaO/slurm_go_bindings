package main

import "slurm"
import "fmt"
func main(){

	version := int(0)
	var config slurm.Ctl_conf
	version = slurm.Version()
	fmt.Printf("Version is %s\n", slurm.VersionString(version));
	config = slurm.GetConfig() 
        slurm.Print_Ctl_conf(config)


}
