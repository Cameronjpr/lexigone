# lexigone

A tool for purging unused translation strings from your projects.

`lexigone --lang=<lang-dir> --search=<src-dir> [--clean]`

---

## Examples

`lexigone --lang=src/lang/ --search=src/`

Find all the keys in `lang/` and see if they appear in `src/`.

`lexigone --lang=src/lang/ --search=src/ --clean`

Find all the keys in `lang/` and remove any that are not found in `src/`.
