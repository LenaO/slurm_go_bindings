/*These are some extra functions to work with slurm in go
** They are seperated, since they don't use the slurm-API
** but wrap arround the SLURM comand line tools */
package extra


import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "errors"
    "path/filepath"
    "slurm/jobinfo"
    "slurm"
    "strconv"

)

var slurm_path string


func find_slurm_path () {
	var err error
	var path string
  path=os.Getenv("SLURM_PATH")
  if path == " "{
	   path, err = exec.LookPath("sinfo")
	   if err != nil {
		    fmt.Printf("could not find slurm executables\n Either add slum-bins to your PATH or define SLURM_PATH\n")
	   } else {
		     slurm_path=strings.TrimSuffix(path, "bin/sinfo")
	}
} else {
  test_path := filepath.Join(path, "bin/sinfo")
  _, err := os.Stat(test_path)
  if os.IsNotExist(err) {
    fmt.Printf("Slurm executable sinfo does no exist at %s\n", test_path)
  } else {
    slurm_path = path

  }
}
}

func Cancel_job( JobId uint32) error{
	find_slurm_path()
  if slurm_path == "" {
      return errors.New("Cannot find slurm executable")
  }
 job_list := job_info.Get_job(JobId)
 if job_list.Error_code != 0 {
		msg := slurm.GetErrorString(job_list.Error_code)
    fmt.Printf(msg)
    return errors.New(msg)
}
  path := filepath.Join(slurm_path,"bin","scancel")
  cmd := exec.Command(path,  strconv.FormatInt(int64(JobId), 10))
  fmt.Print(cmd.String())
  out, err := cmd.CombinedOutput()
  if err!= nil {
    msg := string(out) + err.Error()
    return errors.New(msg)
  }
  return nil
}

type Acc_Job_info struct {
  JobId uint32;
  User string;
  Account string;
  State string;
  JobName string;
}
 var sacct_format_string string

func parse_sacct_output(input string) []Acc_Job_info {
  var job_list []Acc_Job_info
  lines := strings.Split(string(input), "\n")
  fmt.Printf("len %d\n",len(lines)-1)
  for l := range lines {
    var job_info Acc_Job_info
    elements := strings.Split(lines[l], "|")
    if len(elements) < 5 {
      break //Well, this is not clean, but keep it like this for Now
    }
    id, ierr := strconv.Atoi(elements[0])

    if ierr != nil {
      break  //we have no useable entry here but something like 323.batch . Ignore these for now
    }
    job_info.JobId =uint32(id)
    job_info.User = elements[1]
    job_info.Account = elements[2]
    job_info.State = elements[3]
    job_info.JobName =elements[4]
    job_list = append(job_list, job_info)
  }
  return job_list
}

func Get_job_info_accounting(JobId uint32 ) ([]Acc_Job_info, error) {

  sacct_format_string = "JobId,user,account,state,JobName"
	find_slurm_path()
  if slurm_path == "" {
      return nil, errors.New("Cannot find slurm executable")
  }
  path := filepath.Join(slurm_path,"bin","sacct")
  cmd:= exec.Command(path, "-j", strconv.FormatInt(int64(JobId), 10),"--format", sacct_format_string,"-p","-n")
  //fmt.Printf(cmd.String())
  out, err := cmd.CombinedOutput()
  if err!= nil {
    msg := string(out) + err.Error()
    return nil, errors.New(msg)
  }
  list := parse_sacct_output(string(out))


  return list, nil
}
