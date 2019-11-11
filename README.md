# url-word-counter

Simple program that counts entries of a given token (regular expressions are supported) on pages at given URLs. 
URLs are read one by one from stdin. Probably the most convenient way is to use file with URLs as shown below but you also may use echo or just type URLs one by one by hand. 

## How to use

To run the program: 
```bash
cat example_long.txt | go run main.go -jobs=4 -token="\s[Gg][Oo]\s"
```

Or with `echo` command:
```bash
echo -e 'https://golang.org\nhttps://habr.com' | go run main.go
```

Args (all args are optional):
```
-jobs int
      Max concurrent jobs (default 5)
-token string
      Token to count (regexp supported) (default "Go")
```

## Sample output

```
$ cat example_short.txt | go run main.go
2019/11/11 16:44:01 Starting with max 5 jobs
2019/11/11 16:44:01 Started processing url 'https://habr.com'
2019/11/11 16:44:01 Started processing url 'https://golang.org'
2019/11/11 16:44:01 Got result for 'https://golang.org': 20 entries of token 'Go'
2019/11/11 16:44:02 Got result for 'https://habr.com': 4 entries of token 'Go'
```


## Testing

To run offline tests:
```
go test
```

To run all tests including online tests:
```
go test -tags=online
```
