#! /bin/sh

dir=`basename $(pwd)`
SOURCE_FILE="./README.md"
OUTPUT_FILE="./README_output.md"
head -n 5 $SOURCE_FILE > $OUTPUT_FILE
tail -n +6 $SOURCE_FILE | while read line
do
  title=`echo $line | cut -f2 -d '|' | sed -r 's/^ *//' | sed -r 's/ *$//'`
  url=`echo $title | sed -r 's/^[0-9.]* //' | sed -r 's/[()]//g' | sed -r 's/ +/-/g' | tr [A-Z] [a-z]`
  url="https:\/\/leetcode.com\/problems\/${url}\/"
  title="[$title]($url)"
  pattern="^\| [^|]* \|"
  new_line=`echo $line | sed -r "s/$pattern/| $title |/"`
  echo $new_line >> $OUTPUT_FILE

#title=`sed -r 'p'`
done
