#!/bin/bash
sudo bash build.sh
docker push cjzhang/tournaments

ssh -i ~/.ssh/JosephMacbookPro.pem ec2-user@ec2-18-219-40-208.us-east-2.compute.amazonaws.com 'bash -s' < commands.sh
