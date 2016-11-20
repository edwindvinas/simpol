#!/bin/bash

DIR=$(cd $(dirname $0);pwd)

ls $DIR/*.ank |\
while read f; do
  $DIR/../simpol $DIR/lib/tester.ank $f
done
