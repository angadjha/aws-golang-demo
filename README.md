1: Create build file: 
GOOS=linux GOARCH=amd64 go build -o demolambda ./cmd/main.go
2: Create bootstrap file
++++++++++++++++++++++++++++++++++++++++++++++
#!/bin/sh

# This is the bootstrap file that will invoke your Go binary

# Set the permissions to ensure the main binary is executable
chmod +x /var/task/demolambda
# Execute the Go binary
exec /var/task/demolambda
+++++++++++++++++++++++++++++++++++++++++++++
give permission: chmod +x bootstrap

3: Create zip file include build file and bootstrap file in powershell
cmd: Compress-Archive -LiteralPath .\demolambda, .\bootstrap -DestinationPath .\demofunction.zip 

4: check created file 
cmd : Expand-Archive -Path .\demofunction.zip -DestinationPath .\temp

