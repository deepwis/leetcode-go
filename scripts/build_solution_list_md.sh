#! /bin/sh

dir=`basename $(pwd)`
md_file="./solutions.md"
cat > $md_file <<EOF
LeetCode Problem Solutions
---

| Title | Algorithm | Runtime (ms) | Runtime: (faster than) | Memory: (MB) | Memory: (less than) |
|:--------------|:-----:|:-----:|:-----:|:-----:|:-----:|
EOF

SOLUTION_DIRS=`ls | grep -E '^[a-zA-Z-]*$'`

for dir in $SOLUTION_DIRS
do
  echo "Directory: $dir"
  solution_files=`ls $dir | grep -E '^[0-9]*\.go$'`
  for file in $solution_files
  do
    problem=`echo "$file" | cut -f1 -d '.'`
    title=`sed -n -r "s/^\/\/ ($problem\. .*)$/\1/p" "$dir/$file"`
    echo "  Title: $title"

    runtime=`sed -n -r "s/^\/\/ Runtime: ([0-9]*) ms, faster than .*/\1 ms/p" "$dir/$file" | sed -n '$!N;s/\n/ <br> /;p'`
    faster=`sed -n -r "s/^\/\/ Runtime: [0-9]* ms, faster than ([0-9.]*)% of Go.*/\1%/p" "$dir/$file" | sed -n '$!N;s/\n/ <br> /;p'`
    memory=`sed -n -r "s/^\/\/ Memory Usage: ([0-9.]*) MB, less than .*/\1 MB/p" "$dir/$file" | sed -n '$!N;s/\n/ <br> /;p'` 
    less=`sed -n -r "s/^\/\/ Memory Usage: [0-9.]* MB, less than ([0-9.]*)% of Go.*/\1%/p" "$dir/$file" | sed -n '$!N;s/\n/ <br> /;p'` 
    echo "  Runtime: $runtime"
    echo "  Memory: $memory"
    echo "| $title | [$dir](./$dir) | $runtime | $faster | $memory | $less |" >> $md_file
  done
done
