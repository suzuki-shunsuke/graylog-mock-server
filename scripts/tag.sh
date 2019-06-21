# Usage
#   bash scripts/tag.sh v0.6.1

if [ $# -gt 1 ]; then
  echo "too many arguments" > /dev/stderr
  echo 'Usage tag.sh $TAG' > /dev/stderr
  exit 1
fi

if [ $# -lt 1 ]; then
  echo "TAG argument is required" > /dev/stderr
  echo 'Usage tag.sh $TAG' > /dev/stderr
  exit 1
fi

TAG=$1
echo "TAG: $TAG"
VERSION=${TAG#v}

if [ "$TAG" = "$VERSION" ]; then
  echo "TAG must start with 'v'"
  exit 1
fi

echo "cd `dirname $0`/.."
cd `dirname $0`/..

VERSION_FILE=mockserver/version.go

echo "create $VERSION_FILE"
cat << EOS > $VERSION_FILE
package mockserver

// Version is the graylog-mock-server's version.
const Version = "$VERSION"
EOS

git add $VERSION_FILE
git commit -m "build: update version to $VERSION"
npm run release -- --release-as $TAG
