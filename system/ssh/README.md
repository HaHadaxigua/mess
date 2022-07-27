# install open ssh

```shell
sudo apt-get update
sudo apt-get install openssh-server
sudo systemctl status sshd
sudo systemctl list-unit-files | grep enabled | grep ssh
sudo systemctl enable ssh
```

## github & gitlab

### write config

```
#github
      Host github.com
      HostName github.com
      IdentityFile ~/.ssh/id_rsa_github
      User roxy

#gitlab
      Host gitlab.com
      HostName gitlab.com
      IdentityFile ~/.ssh/id_rsa_gitlab
      User roxy
```

### test connection

```shell
chmod 400 ~/.ssh/id_rsa_github

chmod 400 ~/.ssh/id_rsa_gitlab

ssh-add ~/.ssh/id_rsa_github

ssh-add ~/.ssh/id_rsa_gitlab

ssh-add -l
```