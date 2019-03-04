#/bin/sh
set -e
# 中断した時とかにゴミが残らないようにcleanupする

cleanup() {
  retval=$?
  if [ $tmpprof != "" ] && [ -f $tmpprof ]; then
    rm -f $tmpprof
  fi
  exit $retval
}
trap cleanup INT QUIT TERM EXIT

# メインの処理
prof=${1:-"coverage.out"}
echo "mode: count" > $prof
gopath1=$(echo $GOPATH | cut -d: -f1)
for pkg in $(go list ./...); do
  tmpprof=$gopath1/src/$pkg/profile.tmp
  go test -covermode=count -coverprofile=$tmpprof $pkg
  if [ -f $tmpprof ]; then
    cat $tmpprof | tail -n +2 >> $prof
    rm $tmpprof
  fi
done
