# Container from Scratch

NOTE: you'll need linux to run this, doesnt work on macOS/Windows

## Simple Run

```bash
go run main.go run_one echo "hello from container"
go run main.go run_one /bin/bash
```

## Simple Container

Download Ubuntu Base from <http://cdimage.ubuntu.com/ubuntu-base/releases/20.04/release/>

```bash
tar -xvf ubuntu-base-20.04.1-base-amd64.tar.gz -C rootfs
```

Modify `steps/two.go`

```go
	must(syscall.Chroot("/home/satyajit/temp/container-from-scratch/rootfs"))
```

```bash
sudo go run main.go run_two /bin/bash
```

or

```
go build
sudo ./container-from-scratch run_two /bin/bash
```
