# ys
Yaml Search

Search for path to or child of desired string.

```
go get -d ./...
go run ys.go -d account -m pathonly -r test.yaml 
go run ys.go -d account -m childonly -r test.yaml 
```

LIMITATION: WILL NOT WORK WITH YAML FILES CONTAINING LISTS.
