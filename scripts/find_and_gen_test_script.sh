#! /bin/sh

dir=`basename $(pwd)`
test_file="../${dir}_test.go"
cat > $test_file <<EOF
package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
EOF

SOLUTION_FILES=`ls | grep -E '[0-9]+.go'`
for file in $SOLUTION_FILES
do
  echo "File: $file"
  problem=`echo "$file" | cut -f1 -d '.'`
  echo "Problem: $problem"
  sed -n -r "s/^(\/\/ $problem\. .*)$/    \1/p" $file >> $test_file
  FUNC=`sed -r -n 's/func ([a-zA-Z]*)[0-9\(].*/\1/p' $file | head -n 1`
  pattern="Test $FUNC"

  row1=`sed -r -n "/$pattern\(/=" ../solution_test.go`
  offset=`tail -n +"$row1" ../solution_test.go | sed -r -n '/^[[:blank:]]*\}\)$/=' | head -n 1`
  #offset=`tail -n +"$row1" ../solution_test.go | sed -r -n '/^[ \t]*\}\)$/=' | head -n 1`
  ((row2=$row1 + $offset))
  echo $row1, $row2, $offset
  sed -n -r "$row1,${row2}p" ../solution_test.go >> $test_file
done

cat >> $test_file <<EOF
}
EOF
