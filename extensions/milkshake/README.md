# Milkshake Kernel Module

The milkshake module is used to pull metrics out of the kernel for `tambourine`.

This should be a simple module that will report to `/proc/milkshake`.

### Building

To build and run the kernel module.

```bash 
cd tambourine/extensions/milkshake
mkdir build
cd build
cmake ../
make
sudo insmod milkshake.ko
```