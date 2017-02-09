## Chefstrap

Simple go program to get a shellscript for chef opsworks on AWS and bootstrap instance

set userdata:
``` shell
#!/bin/sh
chefstrap -f "output file" -u "url of script"
```