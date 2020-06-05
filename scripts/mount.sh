#!/bin/bash
# help from this issue, thank you! https://github.com/canonical/multipass/issues/1502
export LC_CTYPE="UTF-8"

WORKSPACE_PATH="$1"
MACHINE_NAME="$2"

if [ -z "$WORKSPACE_PATH" ] || [ ! -d "$WORKSPACE_PATH" ]; then
  echo "You must provide a path valid workspace path as the first argument (e.g. /Users/jason/sites)..."
  exit
fi

if [ -z "$MACHINE_NAME" ]; then
  echo "You must provide the machine name as the second argument (e.g. nitro-dev)..."
  exit
fi

HOST_SHARE="192.168.64.1:$WORKSPACE_PATH"
USER=$(whoami)

echo "I'll use the following as the mount folder: "
echo "$HOST_SHARE"

# now we can get the ip of the VM
INFO=$(multipass info "$MACHINE_NAME")
regex='IPv4: +([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3})'
if [[ $INFO =~ $regex ]]; then
  VM_IP=${BASH_REMATCH[1]}
fi

#Lets configure NFS
SHARE="$WORKSPACE_PATH -mapall=${USER} $VM_IP"
EXPORT_FILE=$(cat /etc/exports)

#only add share if not already present in /etc/exports
if [[ ! $EXPORT_FILE =~ $SHARE ]]; then
  echo "I'll enable NFS. I'll need a admin pass for that."

  echo "$EXPORT_FILE" >/tmp/exports
  echo "$SHARE" >>/tmp/exports

  sudo cp /tmp/exports /etc/exports

  # enable and restart nfs sharing
  sudo nfsd enable
  sudo nfsd restart
fi

# now we copy a provision script into the vm which besides provisioning also sets up the nfs share
# by giving the share address as a argument to that script.
multipass copy-files ./provision.sh "$MACHINE_NAME":/tmp/provision.sh
multipass exec "$MACHINE_NAME" sudo chmod +x /tmp/provision.sh
multipass exec "$MACHINE_NAME" sudo /tmp/provision.sh "$HOST_SHARE" "$WORKSPACE_PATH"
