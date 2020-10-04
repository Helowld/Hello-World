//standard Cpp website : https://isocpp.org/get-started

//#include <iostream> includes standard input/output library to cpp file
//it allows to display data, or read your input
#include <iostream>

//namespace defines C++ standard library function.
//using namespace std allows to use functions in std, 
//such as std::cin(input), std::cout(output)...etc.
using namespace std;

//Cpp program executes here, from main() function
int main(void)
{
    //cout displays the string inside quotation, after '<<' character
    //endl means 'break to new line' (it acts like pressing enter)
    cout << "hello, world!" << endl; //Cpp program needs semicolon after each command

    // at the end of program, Cpp returns finished program status by 'return 0'
    return 0;
}

//for more info, check this : https://www.programiz.com/cpp-programming/input-output

// Output :
// hello, world!