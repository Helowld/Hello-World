//#include is a preprocessor command that to include <iostream.h> file in this program.
//iostram allows to input/output data. it acts like <stdio.h> in C.
#include <iostream>

//namespace defines C++ standard library function.
//using namespace std allows to use functions in std, 
//such as std::cin(input), std::cout(output)...etc.
using namespace std;

//Cpp program executes here, from main() function
int main(void)
{
    //cout displays the string inside quotation, after '<<' character
    //endl means break to new line
    cout << "hello, world!" << endl;

    // the return 0 gives the status of the ended program.
    // the program ends with this statement.
    return 0;
}

//for more info, check this : https://www.programiz.com/cpp-programming/input-output

// Output :
// hello, world!