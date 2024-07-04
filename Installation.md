To install the filecommander CLI tool to your system's PATH, you need to move the executable file to a directory that is included in the PATH environment variable. Here are the commands to do that for each operating system:

## Windows(powershell):

```shell
$code = @"
# Download the ZIP file
Invoke-WebRequest -Uri 'https://github.com/geoffrey1330/filecommander/releases/download/v0.1.6/filecommander_Windows_x86_64.zip' -OutFile 'filecommander_Windows_x86_64.zip'

# Extract the contents of the ZIP file
Expand-Archive -Path filecommander_Windows_x86_64.zip -DestinationPath filecommander -Force

# Move the filecommander.exe executable to a directory in PATH (e.g., C:\Windows\System32)
Move-Item -Path filecommander\filecommander.exe -Destination C:\Windows\System32
"@

Start-Process powershell.exe -ArgumentList "-NoProfile", "-ExecutionPolicy Bypass", "-Command $code" -Verb RunAs

```

## Linux:

```shell
# Download the Binary file
curl -LO https://github.com/geoffrey1330/filecommander/releases/download/v0.1.6/filecommander_Linux_x86_64

# Move the filecommander executable to a directory in PATH (e.g., /usr/local/bin)
sudo install filecommander_Linux_x86_64 /usr/local/bin/filecommander
```

## macOS:

#### For Apple Silicon (M1) Macs:

```shell
# Download the Binary file
curl -LO https://github.com/geoffrey1330/filecommander/releases/download/v0.1.6/filecommander_Darwin_arm64

# Move the filecommander executable to a directory in PATH (e.g., /usr/local/bin)
sudo install filecommander_Darwin_arm64 /usr/local/bin/filecommander
```

#### For Intel-based Macs:

```shell
# Download the Binary file
curl -LO https://github.com/geoffrey1330/filecommander/releases/download/v0.1.6/filecommander_Darwin_x86_64

# Move the filecommander executable to a directory in PATH (e.g., /usr/local/bin)
sudo install filecommander_Darwin_x86_64 /usr/local/bin/filecommander
```

#### For Snap Users:

```shell
# Simply run the below command. You can as well use sudo to initiate the installation process If your system privileges requires you to use sudo
snap install filecommander
```
