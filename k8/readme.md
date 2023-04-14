## NFS Server

sudo apt update
sudo apt install nfs-kernel-server
sudo mkdir -p /mnt/nfs_share
sudo chown -R nobody:nogroup /mnt/nfs_share/
sudo chmod 777 /mnt/nfs_share/
echo "/mnt/nfs_share *(rw,sync,no_subtree_check,insecure)" | sudo tee -a /etc/exports
sudo exportfs -a
sudo systemctl restart nfs-kernel-server

## NFS Client (Ubuntu)

sudo apt update
sudo apt install nfs-common
sudo mkdir -p /mnt/nfs_clientshare
sudo mount <IP>:/mnt/nfs_share /mnt/nfs_clientshare
ls -l /mnt/nfs_clientshare/



After mounting copy prerequisite folder to NFS storage.
