# lexigone

A tool for purging unused translation strings from your projects.

`lexigone --lang=<lang-dir> --search=<src-dir> [--clean]`

Example output:

```
go run . --lang=src/lang --search=src 
Searching for 6 keys
key_3 in src/lang/lang.json
key_4 in src/lang/lang2.json
key_1 in src/lang/lang.json
```

---

## Examples

`lexigone --lang=src/lang/ --search=src/`

Find all the keys in `lang/` and see if they appear in `src/`.

`lexigone --lang=src/lang/ --search=src/ --clean`

Find all the keys in `lang/` and remove any that are not found in `src/`.
