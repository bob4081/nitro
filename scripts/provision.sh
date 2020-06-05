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

# create the directory if it does not exist
if [ ! -d "$DIRECTORY" ]; then
  mkdir -p "$DIRECTORY"
fi

# mounting the workspace through nfs
BOOT_MOUNT="$HOST_SHARE $DIRECTORY nfs rw,hard,intr,nolock,lookupcache=none,actimeo=1 0 0"

# putting it in /etc/fstab so it will mount on boot.
echo "$BOOT_MOUNT" >>/etc/fstab
mount -a
