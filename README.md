# Purpose

Created as a mini demo for comparing traditional mocking vs monkey patching

## Packages

Packages internal/person and internal/user contain exactly the same very simplistic functionality to retrieve users/persons.

## internal/user

Implemented using monkey patching

## internal/person

Implemented using traditinal approach

## Diff

Monkey patching itself requires less "initial investment" in terms of interfacing, dependencies and general setup AND can be easily used in functional way, compared to mocking which would require interfaces and everything be hidden behind structs.
