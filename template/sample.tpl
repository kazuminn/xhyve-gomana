[hyv]
uuid = {{.Uuid}}
ram = 1G
vcpu = 2
disk = disk.img
; iso = live.iso
; bridge = en0
[linux]
kernel = vmlinuz
kargs = earlyprintk=serial console=ttyS0 root=/dev/vda1
initrd = initrd.img
; [freebsd]
; userboot = userboot.so
; bootvol = disk.img
; kernelenv = 

