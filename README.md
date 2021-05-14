# template
CLI wrapper around Go's templating library.

## Usage
`template JSON_FILE SOURCE_FILE

__*JSON_FILE*__		-	JSON formatted data which should be inserted into the template.
__*SOURCE_FILE*__	-	Path to the file which should have data inserted into it.

Output is written to stdout.

## Examples
Given a 'data.json' file containing:
```
{"key":"value"}
```
And a 'source.txt' file containing:
```
template inserted '{{.key}}' into this file.
```
Calling:
```
./template data.json source.txt
```
Outputs:
```
template inserted 'value' into this file.
```
You can write the output data to a file like so:
```
$ touch output.txt -- create the output file
$ ./template data.json source.txt > output.txt
```

The templating library used is very powerful, allowing conditional templating of values, as well as iterating over lists etc etc.

Full documentation of what's possible can be found [here](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet).

## Custom functions
The template library supports providing custom functions to template more complex data. Because this is a command line tool, we 
can't make use of the custom fuctions, however it's easy enough to add some pre-defined ones.
