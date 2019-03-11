# go-git-http-credentials-helper
A Go library that adds a secure password/username ask function to a git push/pull/fetch command over http

## The problem:
When executing git fetch/push/pull it might ask for credentials but git craches when it detects it's ran in a tty.  
This libary fixes that problem and makes it possible to fillin the username/password fields if there they are asked for.  

## How it works:
**This is important to know** because this libary is not a 1 function solusion.  

The result of this:  
1. Your program calles `git push` in a http repo
2. This libary creates a webserver for username/password questions and adds the needed shell variables to the process
3. Git sees the `GIT_ASKPASS` and runs your program with as arugment: `username for repo "https://example.com/you/yourprogram"`
4. Now the proccess tree looks a bit like this: `yourProgram -> git -> yourProgram`
5. This libary detects if `yourProgram` is started by git and if so runs it's own startup script 
6. This libary will create a keypair and do a network request to the main program that spawned git in the first place
7. The main program will pickup the network request and check if it is not a fake request. After that it will ask the ask function for a password/username and will encrypt the data with a public key from `yourProgram` and send the encrypted message back to the program that was created by git.
8. `yourProgram` spawned by git will decrypt the message and printout the contents of the message, after that it will exit the program.  
(There are still some things i've not mentioned but those things are mostly useless to know)  

## Q and A:
> Why all this if you can just use a pty?
Mostly because of Windows.  
Windows does have a dll to create a PTY but there are no inplementations yet and i would need to included the ddl because a lot of users don't have the ddl.  
Also PTY support on the Windows 10 subsystem *(linux on windows)* as non-root user is completely broken.  
Beside all of that this is the offical way to do these kinds of things with git so..

> Is this secure?
TL;DR Yes.  
The long answer is no, it's probebly possible to break this libary. Though i could not success in breaking this and it's quite a bit of work because of the security measures and if someone successes it's still a 50/50 change to get any input from a user. It's easier for someone to add a keylogger in your terminal than to break this.

> Why so menny functions to inplement this?
Most things are the result of security measures to check if it's really your program that is asking for someones password.
