#!/usr/bin/python

import sys
import base64
import paramiko
import select

# import the library to decrypt the secret
import SaveSecret

host = '192.168.0.1'

ssh = paramiko.SSHClient()
ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())


# call the function GetSecret with passing the script filename as argument will return the password in base64 encoding
ssh.connect(host, username='root', password=base64.b64decode(SaveSecret.GetSecret(sys.argv[0])))
# call the Clean function after calling GetSecret or your script will act a weird way :)
print base64.b64decode(hg.GetSecret(sys.argv[0]))
hg.Clean()

print "Connected to %s" % host

print "Sending command"
stdin, stdout, stderr = ssh.exec_command("cat /proc/version")
# Wait for the command to terminate
while not stdout.channel.exit_status_ready():
    # Only print data if there is data to read in the channel
    if stdout.channel.recv_ready():
        rl, wl, xl = select.select([stdout.channel], [], [], 0.0)
        if len(rl) > 0:
            # Print data from stdout
            print stdout.channel.recv(1024),

#
# Disconnect from the host
#
print "Command done, closing SSH connection"
ssh.close()
