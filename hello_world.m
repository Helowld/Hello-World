// Official website: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html
#import <Foundation/Foundation.h>

int main(int argc, const char * argv[]) {
    @autoreleasepool {
// @autoreleasepool creates a scoped area and makes it clearer what’s within the pool. 
// Inside of the @autoreleasepool block is where we write our code
        NSLog(@"Hello, World!");

        // The next line calls NSLog which is a function brought in by the Foundation Framework. This function is a lot like printf() function in c. It accepts a format string and can have replaceable tokens. The main noticeable difference is that NSLog automatically creates a newline after a string
        // NSLog() prefaces its output with the date, time, program name, and process ID
        // “@” is an Objective-C shorthand for creating a NSString(Another class of Foundation framework that we will discuss later) object from the given character string
    }
    return 0;
}

