# wiki
Command line tool to get Wikipedia summaries written in Golang.

Credit: Adapted from PHP version - thedouglenz (https://gist.github.com/thedouglenz/193defdb711e0e54d68a)

#Usage 
```
Usage: wiki <query> 
Example: wiki python
```
#Examples
```
$ ./wiki Eros

 Eros

  In Greek mythology, Eros (/ˈɪərɒs/ or US /ˈɛrɒs/, /ˈɛroʊs/; Greek: Ἔρως, "Desire"), was the Greek god of love. His Roman counterpart was Cupid ("desire"). Some myths make him a primordial god, while in other myths, he is the son of Aphrodite.
```
```
$ ./wiki golang

 Go (programming language)

 Go, also commonly referred to as golang, is a programming language initially developed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a statically-typed language with syntax loosely derived from that of C, adding garbage collection, type safety, some dynamic-typing capabilities, additional built-in types such as variable-length arrays and key-value maps, and a large standard library.
The language was announced in November 2009 and is now used in some of Google's production systems. Go's "gc" compiler targets the Linux, Mac OS X, FreeBSD, NetBSD, OpenBSD, Plan 9, and Microsoft Windows operating systems and the i386, amd64, ARM and IBM POWER processor architectures. A second compiler, gccgo, is a GCC frontend.
```
#TODO 
* Increase to more than one word [DONE]
* Better formatting of output 
* Search Feature 
* More error checking 
