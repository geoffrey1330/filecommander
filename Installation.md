To install the filecommander CLI tool to your system's PATH, you need to move the executable file to a directory that is included in the PATH environment variable. Here are the commands to do that for each operating system:

## Windows:

```shell
# Download the ZIP file
curl -LOk https://github.com/geoffrey1330/filecommander/releases/download/v0.1.1/filecommander_Windows_x86_64.zip

# Extract the contents of the ZIP file
Expand-Archive -Path filecommander_Windows_x86_64.zip -DestinationPath filecommander

# Move the filecommander.exe executable to a directory in PATH (e.g., C:\Windows\System32)
Move-Item -Path filecommander\filecommander.exe -Destination C:\Windows\System32
```

## Linux:

```shell
# Download the TAR.GZ file
wget https://github.com/geoffrey1330/filecommander/releases/download/v0.1.1/filecommander_Linux_x86_64.tar.gz

# Extract the contents of the TAR.GZ file
tar -xvf filecommander_Linux_x86_64.tar.gz

# Move the filecommander executable to a directory in PATH (e.g., /usr/local/bin)
sudo mv filecommander/filecommander /usr/local/bin
```

## macOS:

#### For Apple Silicon (M1) Macs:

```shell
# Download the TAR.GZ file
curl -LOk https://github.com/geoffrey1330/filecommander/releases/download/v0.1.1/filecommander_Darwin_arm64.tar.gz

# Extract the contents of the TAR.GZ file
tar -xvf filecommander_Darwin_arm64.tar.gz

# Move the filecommander executable to a directory in PATH (e.g., /usr/local/bin)
sudo mv filecommander/filecommander /usr/local/bin
```

#### For Intel-based Macs:

```shell
# Download the TAR.GZ file
curl -LOk https://github.com/geoffrey1330/filecommander/releases/download/v0.1.1/filecommander_Darwin_x86_64.tar.gz

# Extract the contents of the TAR.GZ file
tar -xvf filecommander_Darwin_x86_64.tar.gz

# Move the filecommander executable to a directory in PATH (e.g., /usr/local/bin)
sudo mv filecommander/filecommander /usr/local/bin
```
