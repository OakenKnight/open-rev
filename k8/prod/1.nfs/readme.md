## NFS Server

sudo apt update
sudo apt install nfs-kernel-server
sudo mkdir -p /mnt/nfs_share
sudo chown -R nobody:nogroup /mnt/nfs_share/
sudo chmod 777 /mnt/nfs_share/
echo "/mnt/nfs_share \*(rw,sync,no_subtree_check,insecure)" | sudo tee -a /etc/exports
sudo exportfs -a
sudo systemctl restart nfs-kernel-server

## NFS Client (Ubuntu)

1. sudo apt update
2. sudo apt install nfs-common
3. sudo mkdir -p /mnt/nfs_clientshare
4. sudo mount <IP>:/mnt/nfs_share /mnt/nfs_clientshare
5. ls -l /mnt/nfs_clientshare/

## NFS Client (MacOS)

1. mkdir nfs_clientshare
2. sudo mount -o nolocks -t nfs <IP>:/mnt/nfs_share ./nfs_clientshare

sudo sshfs -o allow_other,IdentityFile=~/.ssh/id_rsa -p1044 korisnik4@188.255.182.6:/mnt/nfs_share /mnt/nfs_clientshare