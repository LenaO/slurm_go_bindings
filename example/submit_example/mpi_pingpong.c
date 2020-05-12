#include <mpi.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#define MAX_ITER 1000
int main (int argc, char **argv) {
        int rc;
        int size;
        int myrank;
	size_t max_send = 1<<22;
	char *send_buf = (char*)malloc(sizeof(char)*max_send);
	char *recv_buf = (char*)malloc(sizeof(char)*max_send);
	size_t send_size;
        clock_t start, end;
       	rc = MPI_Init (&argc, &argv);
        if (rc != MPI_SUCCESS) {
                fprintf (stderr, "MPI_Init() failed");
                return EXIT_FAILURE;
        }

        rc = MPI_Comm_size (MPI_COMM_WORLD, &size);
        if (rc != MPI_SUCCESS) {
                fprintf (stderr, "MPI_Comm_size() failed");
                goto exit_with_error;
        }

	if(size!= 2) {

		fprintf(stderr, "This process requieres exact two processes\n");
	}
        rc = MPI_Comm_rank (MPI_COMM_WORLD, &myrank);
        if (rc != MPI_SUCCESS) {
                fprintf (stderr, "MPI_Comm_rank() failed");
                goto exit_with_error;
        }
	if(myrank==0)
       		fprintf (stdout, "Size\t Time(ms)\n");

	for(send_size=1 ; send_size<= max_send; send_size*=2){
		for (int i = 0; i<MAX_ITER+2; i++) {
		if(i == 2)
			start = clock();
		if(myrank == 0){
			MPI_Send(send_buf,  send_size, MPI_CHAR, 1, 0x4, MPI_COMM_WORLD);
			MPI_Recv(recv_buf, send_size, MPI_CHAR, 1, 0x5, MPI_COMM_WORLD, MPI_STATUS_IGNORE); 
		}
		else {
			MPI_Recv(recv_buf, send_size, MPI_CHAR, 0, 0x4, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
			MPI_Send(send_buf,  send_size, MPI_CHAR, 0, 0x5, MPI_COMM_WORLD);
		}
		}
		end= clock();
		double time_taken = (double)(end-start)/CLOCKS_PER_SEC;
		if(myrank == 0 )
			fprintf(stdout, "%ld\t %f\n", send_size, time_taken);

	}
	MPI_Finalize();

        return EXIT_SUCCESS;

 exit_with_error:
        MPI_Finalize();
        return EXIT_FAILURE;
}
