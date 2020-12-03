if [ -d $1 ]; then
    echo "Directory $1 already exists"
else 
    mkdir $1
fi

if [ -f $1/main.go ]; then
    echo "$1/main.go already exists"
else
    cp ./data/main_template.go $1/main.go
fi

if [ -f $1/main_test.go ]; then
    echo "$1/main_test.go already exists"
else
    cp ./data/main_test_template.go $1/main_test.go
fi

echo "Open adventofcode.com? [y]es/[n]o/[i]nput (Default: y): "
read open
if [ $open == "n" ]; then 
    exit
fi

dayInt=`expr $1 + 0`
link="https://adventofcode.com/2020/day/$dayInt"
openCmd="explorer $link" 

if [ $open == "i" ]; then
    openCmd+="/input"
fi

if grep -qEi "(Microsoft|WSL)" /proc/version &> /dev/null ; then
   cmd.exe /C $openCmd
else
    echo "Launching browser not supported outside of WSL. Use the link:"
    echo $link 
fi
