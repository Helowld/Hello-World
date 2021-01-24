//  System is a namespace which contains the commonly used types. It is specified with a 'using' System directive.
using System;

//  The 'namespace' keyword is used to define our own namespace. Here we are creating a namespace called HelloWorld.
//  namespace acts as a container which consists of classes, methods and other namespaces.
namespace HelloWorld
{
    //  The statement creates a class named - 'Hello' in C#.
    // Since, C# is an object-oriented programming language, creating a class is mandatory for the program's execution.
    public class Hello
    {
        //  Here 'static' keyword tells us that this method is accessible without instantiating the class.
        // The 'void' keyword tells that this method will not return anything.
        //  The execution of every C# program starts from the Main() method.
        public static void Main(string[] args)
        {
            // WriteLine() is a method of the 'Console' class defined in the System namespace.
            // This is the piece of code that prints Hello World! to the output screen.
            Console.WriteLine("Hello, World!");
        }
    }
}

// For more info check this microsoft's official website https://docs.microsoft.com/en-us/dotnet/csharp/
// For more info check this https://www.programiz.com/csharp-programming/hello-world
// For refernces to that check the site https://www.geeksforgeeks.org/hello-world-in-c-sharp/

// OUTPUT:
// Hello, World!
