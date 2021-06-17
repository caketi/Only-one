read -p "eee" g
if [ $g -ge 90 ]; then
echo "ss(90-100)"
elif [ $g -ge 80 ] &&  [ $g -le 89 ]; then
echo "a(80-89)"
elif [ $g -ge 70 ] && [ $g -le 79 ]; then
echo "b(70-79)"
elif [ $g -ge 60 ] && [ $g -le 69 ]; then
echo "c(60-69)"
elif [ $g -ge 0 ] && [ $g -le 59 ]; then 
echo "f(0-59)"

fi
