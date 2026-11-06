## Options
- Jekyll
- Wordpress?
- Werc, a suckless setup
- Raw html setup?
- Pelycan?
- Haslkyll

## Werc setup instructions

## Buy a domain
Using https://njal.la/
Requires either email or XMPP (which can be provided anonymously, e.g., using disroot.org)
Pay for the domain, either using Paypal (as I did, for convenience), or using one of the many cryptocurrencies that njala supports.
Add a password to pass

### Buy a server
Because I'm familiar 
Create a Digital Ocean project, called nunosempere.com, and a droplet within it.
Choose debian, because Ubuntu comes with snap, which takes too much space
And also, because I hadn't tried debian alone.
Choose the Frankfurt region.
Connect with ssh keys: Add (or create) an ssh key as needed.
Or create a password with pass
Enable backups, because why not.
Choose a hostname: uriel
edit /etc/hostname to add uriel to another line similar to localhost.

### DNS records
Add a domain to your Digital Ocean project
Change DNS records to point to that domain
Point that domain to your digital ocean droplet

### Add user
ssh root@your_ip
adduser uriel
usermod -aG sudo uriel
sudo reboot

## Install necessary stuff
sudo apt install rc git tmux make gcc binutils
wget https://github.com/gokcehan/lf/releases/download/r26/lf-freebsd-amd64.tar.gz
tar -zxvf lf-linux-amd64.tar.gz
mv lf /usr/bin/lf0

### binutils contains ar, which we need

## Install a compiler
sudo apt install gcc reveals that the compiler requires 190M of disk space. Wtf?
## https://www.phoronix.com/scan.php?page=news_item&px=MTg3OTQ
Suckless compilers: 
sudo apt install tcc ## https://packages.debian.org/sid/tcc

ttc seemed to work at first, but then gave a weird error: 
> tcc: error: Unknown relocation type for got: 23
which I didn't feel like debugging, so I gave in to gcc

## Ssh keys and shit
In uriel:
mkdir /home/uriel/.ssh

## In your normal computer:
username="uriel"
host="46.101.221.237"

scp /home/loki/.ssh/nunosempere_protonmail $username@$host:/home/$username/.ssh/nunosempere_protonmail
scp /home/loki/.ssh/nunosempere_protonmail.pub $username@$host:/home/$username/.ssh/nunosempere_protonmail.pub

scp /home/loki/.ssh/nunosempere_gmail $username@$host:/home/$username/.ssh/nunosempere_gmail
scp /home/loki/.ssh/nunosempere_gmail.pub $username@$host:/home/$username/.ssh/nunosempere_gmail.pub

## In uriel
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/nunosempere_protonmail
ssh-add ~/.ssh/nunosempere_gmail
## https://www.atlassian.com/git/tutorials/dotfiles
alias config='/usr/bin/git --git-dir=$HOME/.cfg/ --work-tree=$HOME'
echo "alias config='/usr/bin/git --git-dir=$HOME/.cfg/ --work-tree=$HOME'" >> .bashrc
. .bashrc
config --help

echo ".cfg" >> .gitignore

git clone --bare git@github.com:NunoSempere/.dotfiles.git $HOME/.cfg
config checkout -f

## Remove xset from bash, in config.

## Install 9base
### Failed attempt
wget https://dl.suckless.org/tools/9base-6.tar.gz
tar -xzvf 9base-6.tar.gz
We are going to intall it into /usr/local/plan9, and then later modify the werc setup. Seems more prudent that way.
cd 9-base-6
sudo mkdir /usr/local/plan9
sudo make install # tcc should work.
Neither tcc nor gcc worked, sadly

## /usr/lib/plan9
sudo apt install plan9
(installs to usr/lib/plan9/bin)

## Install werc.
> Note: Werc by default expects the Plan 9 tools to be installed under
> /bin/, if you have installed them elsewhere you will need to edit the
> #! line in bin/werc.rc and customize the $plan9port variable in your
> etc/initrc.local.

## Ok, I'm giving up on compiling it

### Setting up nginx 
https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-20-04
sudo apt install nginx

sudo chmod 644 nunosempere.com
sudo vim /etc/nginx/sites-available/nunosempere.com
sudo ln -s /etc/nginx/sites-available/nunosempere.com /etc/nginx/sites-enabled/
Uncomment line abut hash (server_names_hash_bucket_size 64)
sudo nginx -t
sudo systemctl restart nginx

## Setting up nginx in an even more complicated way
sudo apt install cgiwrap spawn-fcgi
sudo vim /etc/nginx/nginx.conf
https://werc.cat-v.org/docs/web-server-setup/nginx

## https://askubuntu.com/questions/1250538/start-nginx-after-reboot-automatically
sudo systemctl enable nginx

## Spawn fcgi
/usr/bin/spawn-fcgi -a 127.0.0.1 -p 9000 -f /usr/sbin/fcgiwrap  # Use the path to your fcgiwrap binary here 
 PID: 764

Create a spawn-fcgi.sh somewhere with that command, and  chmod 755 spawn-fcgi.sh. Then create a systemctl service, per here:
https://askubuntu.com/questions/919054/how-do-i-run-a-single-command-at-startup-using-systemd
https://alphahydrae.com/2013/01/deploy-wordpress-with-nginx-spawn-fcgi-and-systemd/

 sudo update-rc.d spawn-fcgi defaults

## back to systemd
Didn't particularly work.

## Certbot
sudo apt install certbot python3-certbot-nginx


## rc
https://doc.cat-v.org/plan_9/4th_edition/papers/rc

## change http to https in ../bin/werc.rc
lines 37 and 99

## Change line 65 in 
u=$post_arg_comment_user':'$post_arg_comment_passwd

## replace mkdown2html with discount markdown
sudo apt install discount ## In Debian based systems
https://www.pell.portland.or.us/~orc/Code/discount/
