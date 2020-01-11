#!/usr/bin/python
# encoding=utf-8

import sys
import os
import re

active = 'dev'
latest = 'latest'
region = None
app = None
repo = None


try:
    region=sys.argv[1]
except:
    # github只能编译 hk 的镜像
    region = "hongkong"
try:
    app = sys.argv[2]
except:
    pass


env = os.environ

branch = os.getenv('GITHUB_REF', 'release')
commit = os.getenv('GITHUB_SHA', 'abcd123')
tag = 'latest'

if re.search(r'master', branch):
    tag = 'prod'
    region = 'hongkong'
    active = 'prod'
elif re.search(r'test|staging', branch):
    tag = branch
    region = 'hongkong'
    active = 'test'
elif re.search(r'dev', branch):
    tag = 'dev'
    active = 'dev'
elif re.search(r'release|master', branch):
    tag = 'latest'
    active = 'prod'
elif re.search(r'tag', branch):
    pring("tag version ref====", branch)
    tag = re.search(r'tag.(v[\d\.]+)', branch).group(1)
    active = 'prod'
else:
    tag = branch
    active = 'dev'

print(os.environ['HOME'])
print "active = {0}, region = {1}, tag={2}, commit={3}".format(active, region, tag, commit)

project = 'cloudfeet'
app_list = {
    'api': { 'port': 8082},
}

app_file = os.path.join(os.getcwd(),"conf","app.json")
env_file = os.path.join(os.getcwd(),"conf","app.json." + active)
cp_cmd = "cp -f {profile} {app}".format(profile=env_file, app=app_file)
print "change profile config -> " + env_file

os.system(cp_cmd)


# build go project
for app in app_list:
    # build docker image and push
    cwd = os.getcwd()
    build_path = os.path.join(cwd)
    os.chdir(build_path)
    print "build_path = ", build_path

    # push docker images
    for item in (tag, commit):
        image_name = "registry.cn-{0}.aliyuncs.com/1024w/cloudfeet-{1}:{2}".format(region, app, item)
        cmd1 = "docker build -t {0} .".format(image_name)
        print cmd1
        rc = os.system(cmd1)
        if rc != 0:
            raise Exception("build image failed")

        cmd2 = "docker push {0}".format(image_name)
        print cmd2
        os.system(cmd2)
        if rc != 0:
            raise Exception("push image failed")

    os.chdir(cwd)
