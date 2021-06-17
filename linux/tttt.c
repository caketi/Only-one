 #include <sys/types.h>
       #include <sys/stat.h>
       #include <fcntl.h>
#include <unistd.h>
#define ASIZE (1024*16)

void main()
{
       int fd;
       fd=open("/etc/passwd",O_RDONLY);
       char a[ASIZE];
int n=read(fd,a,ASIZE);
a[n]=0;
printf("%s",a);
close(fd);
}
