# govenv
Govenv: An GOPATH activate maker just like a python virtual env.

# How to use?
Make sure you have downloaded the source code and build with this command.  
```bash
go build govenv.go
```
Done! Just move the govenv binary file to your $PATH directory.  
Now, after making a directory for your golang project, just type following command to create an activate script just like you've seen on the python virtual env.  
```
govenv
```
Then place your source code in the src/ directory. Done!