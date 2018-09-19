# Chapter 4 Exercises

## 4-1

The `tee` command reads it standard input until end-of-file, writing a copy of the input to standard output and to the file named in its command line argument. Implement `tee` using I/O system calls. By default, `tee` overwrites any existing file with the given name. Implement the -a command line option (`tee -a <file>`), which causes `tee` to append text to the end of a file if it already exists. 

## 4-2

Write a program like `cp` that, when used to copy a regular file that contains holes (sequences of null bytes), also creates corresobding holes in the target file.
