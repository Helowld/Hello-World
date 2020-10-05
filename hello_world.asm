; ----------------------------------------------------------------------------------------
; Assembly Hello World for 64-bit Linux
; nasm -felf64 hello_world.asm && ld hello_world.o && ./a.out
; ----------------------------------------------------------------------------------------

          global    _start

          section   .text
_start:   mov       rax, 1                  ; write system call
          mov       rdi, 1                  ; file handle 1 is stdout
          mov       rsi, message            ; address of string
          mov       rdx, 13                 ; number of bytes
          syscall                           ; operating system write
          mov       rax, 60                 ; exit system call
          xor       rdi, rdi                ; exit code 0
          syscall                           ; operating system exit

          section   .data
message:  db        "Hello, World", 10      ; message and newline
