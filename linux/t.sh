#!/bin/bash
read -ep "n " s1

	groupadd class01
for i in `seq 1 $s1`

do 
	grep "user-$i" /etc/passwd &> /dev/null
	if  [ $? -ne 0 ]; then
	useradd  -g class01 user-$i &>/dev/null
		if [ $? -eq 0 ];then
		echo "user-$i" >> /root/user_name
		echo "user-$i" | passwd user-$i --stdin &>/dev/null
		echo "ok"
		fi
	else
	echo "fail"
	fi
done
cat /root/user_name
