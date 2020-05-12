This are wrappers for the SLURM API in golang.
A lot of the code is generated automatic

##Installation:
Since this are simple-go files, just make sure that
the files located at src are in the GO-Path.
For this, you can copy content of the Folder src to your GOROOT.
For example:
```
sudo cp -r  src/slurm /usr/local/go/src
```

Another possibility is to add the path to your GOPATH
For example:
```
export GOPATH=$GOPATH:/home/test/go_wrapper
```

*Note: Since this depends on SLURM, make sure that the SLUM-Library is in your LD-PATH*
For example:
```
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib:/usr/local/lib64
```

if libslum.so is located at /usr/local/lib:/usr/local/lib64
_


