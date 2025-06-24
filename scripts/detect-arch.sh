#!/bin/bash

arch=$(uname -m)
case "$arch" in
  x86_64) echo amd64 ;;
  arm64) echo arm64 ;;
  aarch64) echo arm64 ;;
  armv7l) echo arm/v7 ;;
  armv6l) echo arm/v6 ;;
  i386|i686) echo 386 ;;
  *) echo "Unsupported architecture: $arch"; exit 1 ;;
esac