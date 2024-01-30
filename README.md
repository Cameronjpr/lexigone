# lexigone

A tool for purging unused translation strings from your projects.

`lexigone --lang=<lang-dir> --search=<src-dir> [--clean]`

e.g.

`lexigone --lang=src/lang/ --search=src/`

Find all the keys in `lang/` and see if they appear in `src/`.
