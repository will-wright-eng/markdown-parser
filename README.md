# markdown-parser

## setup

- build
- test on `example.md`

```bash
go build -o mdparse
./mdparse parse example.md
```

- move to PATH directory
- test again

```bash
mv mdparse /usr/local/bin/
mdparse parse example.md output
```

- compare `tmp/` and `output/` directories to validate

## not fault tollerant

- file path and name are pulled from string priror to code block (doesn't matter if it's bold)
- `github.com/yuin/goldmark/ast` will split a string by formatting

```md
**docker-compose.yml**  # this works
**docker-compose.yml**: # this will fail
docker-compose.yml      # this should also work
```
