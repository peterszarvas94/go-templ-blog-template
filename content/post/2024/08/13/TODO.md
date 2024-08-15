---
title: "Todo"
category: "dev"
tags: ["may", "day", "test"]
time: "12:49"
---

## TODO

- [x] add time to post header
- [x] fix subdir static serving
- [x] tag links
- [x] main page
- [ ] generalizing:
  - [x] file based routing in content
  - [ ] cool default [theme(s)](#theme)

## Theme {#theme}

Index page is a special page, should have its own `templ` component in every
[theme](#theme).

_Maybe `static` should be just an optional folder under `content`, like any
other folder?..._

A theme contains `templ` + `css` files.

You can select which them you want to use, in the config file.

If there is no such template, it will fall back to the default one.
