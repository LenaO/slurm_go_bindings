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
