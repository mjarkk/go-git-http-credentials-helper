# go-git-http-credentials-helper
A Go library that adds a secure password/username ask function to a git push/pull/fetch command over http

## The problem:
When executing git fetch/push/pull it might ask for credentials but git craches when it detects it's ran in a tty.  
This libary fixes.  
**Warning: this libary is not a 1 function solusion because of the limitations.**
