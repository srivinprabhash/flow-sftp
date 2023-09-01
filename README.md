# 🌊 Flow SFTP

Flow SFTP can watch a directory for new files and upload them to a remote SFTP host in near real time.

(*Written in Go*)

## 🪶 Features
- Copy new files almost in real time.
- Configurable via a yaml config file.
- Clear backlog files.
- Backup files to a separate directory.

## ⚙️ Configuration

```yaml
flow:
  source: "/path/to/directory/" # Path to the directory you want to watch for
  backupDir: "/path/to/directory/" # Path to your backup directory 
  enableBackups: false # Enable/Disable backups
  clearBacklog: false # Enable/Disable backlog clear. It is recomended to keep this disabled if backup is not enabled.

sftp:
  host: "" # Remote SFTP Host Name/IP
  port: 22 # Remote SFTP Host Port 
  user: "srivin" # Remote SFTP User
  privateKey: "/path/to/file" # SSH Private Key file path 
  remotePath: "/path/to/directory/" # Destination directory in the remote SFTP host
```