#!/bin/bash
set -e -o pipefail
trap '[ "$?" -eq 0 ] || echo "Error Line:<$LINENO> Error Function:<${FUNCNAME}>"' EXIT

cd `dirname $0` && cd ..
CURRENT=`pwd`

function test
{
   go test -v $(go list ./... | grep -v vendor) --count 1 -race
}

function go-bindata
{
  cd $CURRENT/internal/assets
  go-bindata -pkg assets ./...
}

function bench
{
  go test -v $(go list ./... | grep -v vendor) -run none -bench . -benchtime 3s -benchmem
}

function bench_pprof
{
  go test github.com/gjbae1212/go-counter-badge/badge -run none -bench=BenchmarkBadgeWriter_RenderIconBadge -benchtime 3s -benchmem -memprofile mem.out
  go tool pprof -http=127.0.0.1:8000 mem.out
}

function escape
{
  go build -gcflags '-m -m' github.com/gjbae1212/go-counter-badge/badge
}

CMD=$1
shift
$CMD $*
