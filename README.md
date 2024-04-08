# bitsweetfont

![Font preview](https://private-user-images.githubusercontent.com/6286655/320440050-d73913f5-edea-4861-9961-78dc2f171ac2.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTI1NzEwOTcsIm5iZiI6MTcxMjU3MDc5NywicGF0aCI6Ii82Mjg2NjU1LzMyMDQ0MDA1MC1kNzM5MTNmNS1lZGVhLTQ4NjEtOTk2MS03OGRjMmYxNzFhYzIucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI0MDQwOCUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNDA0MDhUMTAwNjM3WiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9ODdjYjEyMzQyMzY4YWE3ZWE3ZmViN2QxN2YwYzhhNDYyMjc2Y2YwNTM4N2I3MzZiZWNiN2JhMWY5NmY5ZWY1YyZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmYWN0b3JfaWQ9MCZrZXlfaWQ9MCZyZXBvX2lkPTAifQ.pReFdUgWb1rynBQ2dzFg8O0aWMmysCsDbPnkrV2Ahp8)

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
