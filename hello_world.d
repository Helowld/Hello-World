// To learn more: https://en.wikibooks.org/wiki/D_(The_Programming_Language)

#!/usr/bin/env rdmd // use this line to use D language as traditional scripting set your main.d file to be executable

// importing standard input/output functions
import std.stdio; 

// All executable D programs contain a main function. This function does not return a value, so it is declared "void".
void main() 
// This function is executed when the program runs.
{
    // built-in function to print the quoted string
    writeln("Hello, World!"); 
}

// $ rdmd main.d

// or 

// $ chmod +x main.d
// $ ./main.d
