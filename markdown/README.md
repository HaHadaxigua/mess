# toc: generate markdown index automatically

```shell
# install toc
alias toc=`
grep -E "^#{${1:-1},${2:-2}} " | \
sed -E 's/(#+) (.+)/\1:\2:\2/g' | \
awk -F ":" '{ gsub(/#/,"  ",$1); gsub(/[ ]/,"-",$3); print $1 "- [" $2 "](#" tolower($3) ")" }'
`
```

# use

```shell
cat readme.md | ./toc
```