for f in ./experiment-sandbox/idl/*.thrift; do
    "$GOPATH"/bin/thriftrw "$f" --out=./experiment-sandbox/thriftgen &
done
