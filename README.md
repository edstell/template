# template
CLI wrapper around Go's templating library.

## Usage
`template JSON_FILE [--source=FILE] [--output=FILE]`

__*JSON_FILE*__	-	JSON formatted data which should be inserted into the template.
__*source*__	-	Path to the file which should have data inserted into it.
__*out*__	-	Path to the file which templated data should be written to.

If no source file is provided, source data is read from `stdin`.

If no output file is provided, output data is written to `stdout`.

## Examples
This is a simple command line wrapper around Go's templating library, there are lots of great resources for using the template syntax - [here's](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet) a good cheatsheet to get you started.

## Custom functions
Because this is a pre-built CLI you won't be able to provide your own custom functions, some useful ones have been provided and are documented below.
