#!/bin/bash
HOST_SHARE="$1"
DIRECTORY="$2"

if [ -z "$HOST_SHARE" ]; then
  echo "Missing the host share as the first argument..."
  exit
fi

if [ -z "$DIRECTORY" ]; then
  echo "Missing the directory to create and mount as the second argument..."
  exit
fi

mkdir -p "$DIRECTORY"

# mounting the workspace through nfs
BOOT_MOUNT="$1 $DIRECTORY nfs rw,hard,intr,nolock,lookupcache=none 0 0"

# putting it in /etc/fstab so it will mount on boot.
echo "$BOOT_MOUNT" >>/etc/fstab
mount -a
