# bitsweetfont

![Font preview](https://www.quasilyte.dev/images/bitsweetfont/font_preview.png)

## Overview

A bitmap font for Go, made with [bitfontier](https://github.com/quasilyte/bitfontier). It's designed to be used in pixel art games.

Currently, it only supports two "languages" and some extra symbols:

* English
* Russian
* Extra misc symbols

It has two base sizes: `1` and `1.3`. If you take `Scale()` function into account, we get many sizes to choose from:

* `1`
* `1.3`
* `2`
* `2.6`
* `3`
* `3.9`
* `4`

If you want to contribute new images for this pack, please keep in mind that I want to keep the license as open as possible to make this font easy plug-and-play without any multi-level attribution paperwork.

This is a simple font and by no means universal (look at the tiny number of languages it supports!) At the same time, it should be a good example on how to make bitmap fonts for your games using [bitfontier](https://github.com/quasilyte/bitfontier). Good luck!

> See [fontinfo.md](fontinfo.md) to see a full list of runes supported.

## Installation

```bash
$ go get github.com/quasilyte/bitsweetfont
```

## Usage

```go
// import "github.com/quasilyte/bitsweetfont"

// ff is a font.Face of a base size 1.
ff := bitsweetfont.New1()

// It's possible to create programmatically scaled versions
// of the base fonts.
// Only whole scaling factors are available (2, 3, 4, ...)
ff2 := bitsweetfont.Scale(ff, 2)
ff3 := bitsweetfont.Scale(ff, 3)
```

## Building Font

This repository already contains a built font for you to use.

But if you want to re-generate it with some other `bitfontier` options, it's also possible.

Take a look at Makefile's `font` target.
