! Hello World program written using the Fortran programming language.
!
! For information on the Fortran language visit: https://www.fortran.com/
! For more information on the Fortran "Hello, World!" program, visit: https://en.wikibooks.org/wiki/Fortran/Hello_world

program hello_world
	! The following line prints "Hello, World!". The asterisk indicates to the print statement that the items immediately following the comma after the asterisk should be printed in a format compatible with their type.
	print *, "Hello, World!"
end program

! The extension for this file is .f, though, to compile this, any of the Fortran extensions (.for, .f90, .f95, etc.) can be used.
!
! If compiling on a Unix system using the GNU Fortran compiler, execute:
! gfortran hello_world.f
!
! If compiling on Windows, execute:
! f95 hello.f
!
! The Windows compilation command may also be used on Unix, assuming that the appropriate compiler has been installed.
