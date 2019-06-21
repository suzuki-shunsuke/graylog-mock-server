cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

find . -type d -name node_modules -prune -o \
  -type d -name .git -prune -o \
  -type f -name "*.go" -print \
  | xargs gofmt -l -s -w
