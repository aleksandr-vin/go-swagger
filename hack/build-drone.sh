#!/bin/bash
set -e -x -o pipefail

mkdir -p /usr/share/{testresults,coverage,dist}
go test -v -race $(go list ./... | grep -v vendor) | go-junit-report -dir /usr/share/testresults

# Run test coverage on each subdirectories and merge the coverage profile.
echo "mode: ${GOCOVMODE}" > profile.cov

repo_pref="${CI_BUILD_DIR##${GOPATH/%:*/}/src/}/"
# Standard go tooling behavior is to ignore dirs with leading underscores
for dir in $(go list ./... | grep -v -E 'vendor|generator')
do
  pth="${dir//*$repo_pref}"
  go test -covermode=${GOCOVMODE} -coverprofile=${pth}/profile.tmp $dir
  if [ -f $pth/profile.tmp ]
  then
      cat $pth/profile.tmp | tail -n +2 >> profile.cov
      rm $pth/profile.tmp
  fi
done

go tool cover -func profile.cov
gocov convert profile.cov | gocov report
gocov convert profile.cov | gocov-html > /usr/share/coverage/coverage-${CI_BUILD_NUM-"0"}.html
go build -o /usr/share/dist/swagger ./cmd/swagger
