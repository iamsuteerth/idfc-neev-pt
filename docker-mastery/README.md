# Docker Commands
## docker container run --publish 80:80 nginx

- Download nginx image from docker hub
- Start a new container from that image
- Opened port 80 or whatever is specified on the LHS 
- And routes it to container IP which is port 80 on the RHS

---

`docker stop <container id>` is pretty self explanatory

`docker run <container id>` is again self explanatory

`docker container ls` or `docker ps` to list running containers

`docker container ls -a` lists all containers inclusive of closed ones

---

## docker container run --publish 80:80 --detach nginx
- The detach argument is used for running the container in the background

## docker container run --publish 80:80 --detach --name webhost nginx
- Name the container as `webhost`
- You can use `-d` for `--detach`
 
## docker rm id1 id2...
- Remove containers which are not running
  - Get an error if you try to remove a container which is running

## docker rm -f id1 id2...
- Force a shutdown and then remove all running containers as well

# What happens in `docker container run`
- Looks for an image in the image cache
- If not found, look for the image in the remote image repository
- If version is not specified, download the latest available image from the docker hub
  - Can be specified as nginx:latest or nginx:version
- Creates a new container based off the image and prepares to start
- Assigned a virtual private IP inside the docker engine
- Without the `--publish` , no ports will be opened
- A container is just a process unlike a VM
- In order to pass multiple environment variables
  - `-e VAR1 -e VAR2 and so on...`

# Excercise - Managing multiple containers
```sh
# Creating and running containers for different applications
# Creating an nginx web server
docker container run -d --publish 80:80 --name nginxtest nginx:latest
# Creating an apache web server
docker container run -d --publish 8080:80 --name httpdtest httpd:latest
# Creating a mysql server with ranndom root password environment variable set to true
docker container run -d --publish 3306:3306 --name mysqltest -e MYSQL_RANDOM_ROOT_PASSWORD=yes mysql:latest
# Checking all containers before checking for logs and removing them
docker container ls -a
# Checking for the password generated
docker container logs mysqltest > container.log 2>&1 && grep "PASSWORD" ~/container.log
# Stop the containers
docker stop 384 f5c 2d4
# Remove the containers
docker rm mysqltest httpdtest nginxtest
# Check after cleanup
docker container ls -a
```

# What's going on in containers
- `docker container top` gives process list in one container
- `docker container inspect` gives details of the container config
- `docker container stats` gives performance stats

# Getting a shell inside containers
- `docker container run -it` starts a new container with which you can interact
- `docker container exec -it` run additional command in existing container
SSH is not necessarily needed since docker cli is a good way of adding ssh to containers
- `-it` : The `-t` stands for `pseudo-tty` which simulates a real terminal like SSH whereas `-i` keeps the session open to receive more commands
- Execute ONE command without interactivity

```bash
# The format is docker exec [OPTIONS] CONTAINER COMMAND [ARG...]
# Run a single command
docker container exec proxy nginx -v
# Run an interactive shell 
docker container exec -it proxy bash
```

`docker container run -it --name proxy --publish 80:80 nginx bash` when we exit the terminal, the container stops BECAUSE we changed the default program to be executed with bash whereas originally the job is to run the nginx program itself
- Container only runs as long as the command it ran on startup runs

`docker container run -it --name ubuntu ubuntu` will give bash by default since it uses bash as its shell. Its a minimal version of ubuntu's image and you can install more software using the `apt` packet manager
- If the container is opened up once again, it will still have all the software initially installed with it

- `docker container start -ai proxy` : Here the `-ai` option is used to attach the container to STDOUT and make the container interactive
  - It is like using the `-it` option
- **The exec command** runs an ADDITIONAL process on the container, so for example if we run mysql and then run bash on it and exit, it will still be running 

# Docker Networking
- When a container is run, you're really connecting to a docker network called the bridge network
- Each virtual network routes through NAT firewall on host IP
- All containers can talk to each other without `-p`

- `Batteries included but removable`
  - Defaults work well, but you can customize things as well 
- You can create multiple virtual networks
- Attach container to one or more virtual networks or NONE
- Skip virtual network configuration and use `--net=host`, this results in loss of a few containerization benefits
- Docker network drivers can give us new abilities

```sh
# Which ports are routing traffic to where
docker container port webhost
# See the IP address by filtering out the specific JSON entry
docker container inspect --format '{{.NetworkSettings.IPAddress}}' webhost
```

Some of the CLI commands in practice

```sh
# Show the existing networks
docker network ls

# The docker0 or the bridge is the default virtual network NATed behind the host IP, all containers are by default connected to this network and just work

# Using the formatting from previous example to just get the container information instead of everything
docker network inspect --format {{.Containers}} bridge    

# The host network seen after using the `docker network ls` command is a special network which directly attaches to the container which skips the virtual networking of docker

# It prevents security boundaries of the containerization from protecting the interface of that container but can also improve performance and be what's needed in some niche cases

docker network create my_app
# Create a docker network

docker container run -d --name new_nginx --network my_app nginx:alpine 
# Create and run a new container with the newly created network

docker network connect my_app webhost 
# Connect the previous container to the new virtual network we just created

# Now if we inspect the container 
docker container inspect webhost
# It is attached to TWO networks which means that you can attach a container to multiple networks

```
# DNS
- Its a good practice to avoid using static IPs for talking to containers
- DNS Naming is a good solution for the same
- Docker uses the container names as host names for intra communication
- Docker daemon has a DNS server used by containers built in by default
```
For a NEW network created, a special new feature is present which is automatic DNS resolution for all the containers on other virtual networks to the created network's containers
```
So if a container is run on the new virtual network, then the containers will be able to find each other regardless of their IP address using the container names

The default network DOESN'T have DNS resolution so you can use --link where you have to manually link one container to another

# Excercise - CLI App Testing
## Problem Statement
*Imagine yourself being a support engineer for a server farm running Ubuntu and RockyLinux(CentOS relpacement) a ticket comes in and asks you to check on the different curl versions in those two distributions*
- *Use different linux distro containers to check **curl** cli version*
- *Use two different terminal windows to start bash in both centos:7 and ubuntu:14.04 using `-it`*
- *Clean up the containers*
- *Ensure `curl` is installed and on latest version for that distro*
  - *ubuntu : `apt-get update && apt-get install curl`*
  - *rockylinux : `yum install --allowerasing --nobest curl`*
- *Check `curl -- version`*

```sh
# Pull the necessary images from docker hub
docker image pull ubuntu:14.04
docker image pull rockylinux:9
# Check whether the images are downloaded
docker image ls -a
# Run the containers
# For Ubuntu
docker container run -it --name ubuntu ubuntu:14.04
# For CentOS
docker container run -it --name rocky rockylinux:9
# Run the shell commands
```
# Exercise - DNS RR Test
DNS Round Robin technique is the concept that you can have two different hosts with DNS aliases that respond to the same DNS name. So like there are multiple IPs to respond to the same DNS

There can be multiple containers on a created network that respond to the same DNS address. The containers respond the same way that DNS round robin does.
- *Create a new virtual network*
- *Create two containers from `elasticsearch:2` image*  
- *Research and use `--network-alias search` when creating them to give them an additional DNS name to respond to*
- *Run `alpine nslookup <dns_alias>` with `--net` to see the two containers list for the same DNS name*
- *Run `rockylinux curl -s <dns_alias>:9200` with `--net` multiple times until you see both "names" field below*

```sh
# Create a new network
docker network create dns_rr

# I am using an x86-64 machine so elasticsearch:2 works just fine
docker image pull elasticsearch:2

# Create 2 containers using the custom network AND giving them a common network alias
docker container run -d --name container_1 --network dns_rr --network-alias=elasticontainer elasticsearch:2

docker container run -d --name container_2 --network dns_rr --network-alias=elasticontainer elasticsearch:2

# Check the containers
docker container ls

# Run an alpine container which closes upon exit to lookup for the IPs associated with our two containers
docker container run --rm --network dns_rr alpine nslookup elasticontainer

# Same process but in an interactive environment using rockylinux
docker container run --rm -it --name rocky --network dns_rr rockylinux:9

curl -s elasticontainer:9200
```
# Docker Images
## What is there in an image?
- Application Binaries
- Dependencies
- Metadata about the image data and how to run the image
- Not a complete OS, so no kernel modules or drivers. The host provides the kernel
- It can be as small as one file like a golang static binary
- It can be as big as a distro with things like apt, apache, php and much more installed

It's essentially a series of changes on the file system and metadata about those changes with their identity as the SHA-256 hash.

## Docker Hub
- Only the official images, maintained and kept consistent with the original software developers and docker team have the repository names. They are often well documented.
```
An example of nginx image having multiple tags for the same variant

1.27.2, mainline, 1, 1.27, latest, 1.27.2-bookworm, mainline-bookworm, 1-bookworm, 1.27-bookworm, bookworm
```
- For these tags, any can be used to refer to either
- An important observation is when using `docker container ls -a` is that there maybe MULTIPLE containers with different tags but same SHA. In a loose way, the unique SHA are what take up space. So the storage aspect can be a bit deceiving.

---

Suppose you run `docker image ls` and then `docker image history nginx:latest`
- What you see is the history of image 
- Every image starts with a BLANK layer called `scratch`
- Ever set of change made subsequently on the file system in the image is another layer
- Some changes may have no changes in terms of file size where simply metadata is changed. Say which command is about to be run.
- Some can add a lot of changes adding size
  - Every layer gets it own unique SHA that helps the system differentiate two layers
  - As you make more changes, you create more layers
  - And say you want this image to be the base for more layers. You only store ONE copy of each layer
  - This avoids redundancy and you only store the image data ONCE on the file system
  - You don't necessarily need to download and upload the layers you already have

Say you had a custom image and you added an apache server on top of it as another layer. Opened a port and then told it to copy the source code for two different websites. **You have TWO images**

**BUT** *only ONE copy of each layer is stored*

```
|   A   ||  B   | (2)
|      PORT     | (1)
|     APACHE    | (1)
|     CUSTOM    | (1)
```
When you run a container, docker simply reates a read/write layer for that container on top of the chosen image.

Say you change a file in the running container, this is known as `copy on write`

The file system takes it out of the image and copies it into the differencing and store a copy of that file in the container layer.

The container is just the running process and the files different than the base image

The `<missing>` is there for the layers so they don't get their own image id 

---

```sh
# Used to return metadata about the image in form of JSON
docker image inspect nginx:alpine

# You an again check the specific parts you want by using the format option
docker image inspect --format '{{json .Config.ExposedPorts}}' nginx:alpine
```

Building on the point of `non redundancy`

```sh
docker image ls
# nginx                        latest    28402db69fec   3 weeks ago     279MB
docker image pull nginx:mainline
# nginx                        mainline   28402db69fec   3 weeks ago     279MB
```

- There are TWO different nginx with different tags BUT they have the same container ID or SHA so they're only really stored ONCE

- You can Re-TAG existing images AS WELL using this format
  - `docker image tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]`

- By default, if you don't specify a tag, it assumes `latest` But you COULD by doing something like `docker image tag nginx test/nginx:main`

- You can push the latest changes to an image registry (docker hub is used by default)
  - `docker image push <repo>`

- You can simply change the visibility of the repository by making the changes on the docker hub

----

### Important thing about credentials
- When you login using `docker login`, the credentials are stored in ~/.docker/config.json
- So do note that whenever you use a pc, the credentials get stored for the user account logged in

---

# Dockerfiles
An observation right off the bat can be that it looks like a shell script which is not true and the default name is `Dockerfile`
- `docker build -f some-dockerfile` is how you can specify a different file than default
- `FROM` command is REQUIRED to be there and is the FIRST command which is normally a minimal distribution
  - It can be alpine these days which is usually done to save a lot of time 
  - These minimal distributions have the basic package manager
  - Then you can download the software you need
  - Then there is the `ENV` stanza which is used to specify and set environment variables
    - They are the main way to set keys and values for running and building containers
    - They work everywhere, on every OS and configuration
- Every stanza is an actual layer in the docker image and the order is important since its run TOP DOWN
- You can chain commands together since each stanza is a separate layer
  - Often done by 
    ```bash
    apt-get update \
    && apt-get install curl
    ```
- A common stanza is about pointing out the log files to the STDOUT and STDERR
  - Docker handles logginf for us and we just need to point correctly
- By default, NO TCP or UDP ports are open inside a container
  - You can do so by using the `EXPORT` command
  - Which is often a stanza on its own
  - Its important to note that simply exposing the ports won't work and you have to `-p` on the containers to expose the ports to the host
- Lastly, `CMD` is a required parameter which is the final command that will be run every time a new container is launched from the image or upon a container restart

# Creating Images
```sh
docker image build -t customnginx .
```
Here, the `-t` is for tag and the `.` is for saying that build the image in the directory where the command is being run
- What can be observed is that there will be hashes after every step or command
- These hashes are kept in the build cache so that next time when this is built and the line hasn't changed, it won't be re-run
- This makes software development VERY fast due to caching after the build
- An important thing to note is that a change causes an avalanche effect so it's important to keep the order
- Keep the things which change the most in the end of the docker file

```docker
FROM nginx:latest
WORKDIR /usr/share/nginx/html
# Change working directory to root of nginx webhost
COPY index.html index.html
# No need to specify EXPOSE or CMD (FROM)
```

- `WORKDIR /usr/share/nginx/html` is better and preferred to using `RUN cd /some/path`
- The `COPY` command is used for copying your source code from your local machine/build servers into your container images
- Here the default home page is replaced by the custom one for the web server
- Everything is inherited `FROM` so there is no need for a `CMD` here due to `FROM nginx:latest`

---

# Exercise - Build your image and run it!
*A common exercise for a docker admin is to make custom images and help others make Dockerfiles for their own source code. Often you'll need to go looking at the official images and may not know that much about the app but will still have to dockerize it. Basically make a Dockerfile for an existing app and get it to run on a container. You might have to search for matching images or something that's close*

Given Instructions
- *You should use the `node` official image, with the alpine 6.x branch (`node:6-alpine`)*
- *This app listens on port 3000, but the container should listen on port 80 of the Docker host, so it will respond to [http://localhost:80](http://localhost:80) on your computer*
- *Then it should use the alpine package manager to install tini: `apk add --no-cache tini`*
- *Then it should create directory /usr/src/app for app files with `mkdir -p /usr/src/app`, or with the Dockerfile command `WORKDIR /usr/src/app`*
- *Node.js uses a "package manager", so it needs to copy in package.json file*
- *Then it needs to run 'npm install' to install dependencies from that file*  
- *To keep it clean and small, run `npm cache clean --force` after the above, in the same RUN command*
- *Then it needs to copy in all files from current directory into the image*
- *Then it needs to start the container with the command `/sbin/tini -- node ./bin/www`. Be sure to use JSON array syntax for CMD. (`CMD [ "something", "something" ]`)*
- *In the end you should be using FROM, RUN, WORKDIR, COPY, EXPOSE, and CMD commands*

```docker
# Use the image as described
FROM node:6-alpine

# Since the container needs to be exposed to port 80
EXPOSE 80

# Use alpine package manager to install tini and make the app directory
RUN apk add --no-cache tini 

# Set the working directory
WORKDIR /usr/src/app

COPY package.json package.json

# Run the commands as specified
RUN npm install \
    && npm cache clean --force

# Copy all files from current directory
COPY . .

# Run the commands as described
CMD [ "/sbin/tini", "--", "node", "./bin/www" ]
```

Now the terminal work

```sh
# Build the image with name custom_nodejs
docker image build -t custom_nodejs -f Dockerfile .

# Tagging it with something and then checking it
docker image tag custom_nodejs:latest custom_nodejs:mainline
docker image ls

# Pushing it 
docker image tag custom_nodejs:latest iamsuteerth/custom_nodejs:latest
docker image push iamsuteerth/custom_nodejs

# Run a container
# As it can be seen that we are routing the client's port 80 to the container's port 3000 which is what the app is using
docker container run -d -p 80:3000 --name webhost iamsuteerth/custom_nodejs:latest

# Cleaning up our system
docker system prune -a -f

# Pulling the pushed image
docker image pull iamsuteerth/custom_nodejs:latest

# Running a container off of it
docker container run -d -p 80:3000 --name webhost iamsuteerth/custom_nodejs:latest
```

# CLEAN CLEAN CLEAN!!!
- Check system usage as per the docker daemon using `docker system df`
- You can use "prune" commands to clean up images, volumes, build cache, and containers
  - `docker image prune` to clean up just "dangling" images   
  - `docker system prune` will clean up everything you're not currently using
  - `docker volume prune` to clean up local volumes. Using the `-a` option removes ALL and using the `-f` option removes confirmation. So this is pretty dangerous `docker volume prune -a -f`

---

# Container Lifetime and Persistent Data
- Containers are `immutable` and `ephemeral` which is basically saying that they are unchangeable and temporary/disposable
- You should always re-deploy containers and never change them following an `immutable architecture`
- A challenge because of this is unique data
- Docker gives a feature to ensure `separation of concerns` which is basically saying that the updated version of the app will still have access to the unique data when the container was recycled
  - **Data volumes** : Config options for a container that create a special location outside of container UFS (Unique File System) and makes it attacheable and detachable whatever container you want and the container treats it like a local volume
    - **Bind mounts** : Links container path to host path
- **Volumes need manual deletion, you can't just clean them up by removing the container**
  - You can cleanup unused volume by using `docker volume prune`
  - `docker image inspect --format '{{json .Config.Volumes}}' mysql` can be used to fetch the volumes created by this image
- The running container gets it's own unique location on the host to store the "unique data" and that location is mounted to the containers and can be inspected by running `docker container inspect --format '{{json .Mounts}}' mysql `
  ```json
  [
    {
      "Type":"volume",
      "name":"358a63b553310df0976b510b1d03ae4e265945ba26a9d6e35547a5151cb55e43",
      "Source":"/var/lib/docker/volumes/358a63b553310df0976b510b1d03ae4e265945ba26a9d6e35547a5151cb55e43/_data",
      "Destination":"/var/lib/mysql",
      "Driver":"local",
      "Mode":"",
      "RW":true,
      "Propagation":""
    }
  ]
  ```
  - The source is the place on the disk where the data is actually present and the destination is the "place" on the container
  - It is important to know that you cannot reach this location on a windows or mac machine since docker creates a VM where this data is stored
      - This limitation can be worked around with mounted volumes however
  - The volumes are difficult to distinguish from one another so this can be tricky with multiple containers running telling which is which    
  ```docker
  # The -v can be used to specify a new volume we want to create for this container thats about to run or create a named volume 
  docker container run -d --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=true -v named-volume-mysql:/var/lib/mysql mysql 
  ```
- You can create volumes beforehand before running `docker container run` using `docker volume create`
  - It's the only way you can specify a different driver and set driver options
  - But normally specifying it using `VOLUME` in the Dockerfile or during the run command is sufficient

### You can use a full path on any OS but EACH shell may do $(pwd) differently so take note of that

## Bind Mounts
- It maps host files/directories to container files/directories
- Its just a pointer to two locations on the disk
- Skips UFS and host files overwrite any in container
  - Any files in the container that you map the host files to, the host files win
  - It's not really overwriting, when you re-run the container without bind mount, you will see the container files underneath
- Since these are host specific, they must be run during `container run`
- The key different between a bind mount and a local volume is that of a **FORWARD slash** run as `.. run -v /some/path:/path/container`

```sh
# The pwd part is a shortcut in shell to get the current path
# We are using the DEFAULT nginx image (now the WORKDIR is set to /usr/share/nginx/html for the container)

docker container run -d --name webhost -p 80:80 -v $(pwd):/usr/share/nginx/html nginx:latest

docker container run -d --name webhost2 -p 8080:80 nginx:latest

# It can be observed that the webhost2 has the default nginx home page and webhost has the custom one

# What we did is we mapped the host directory from $(pwd) to the WORKDIR so that it can see all the files present in the host directory

# Container Commands
docker container exec -it webhost bash

# This is what we get

root@a6e9d14ca01f:/# cd /usr/share/nginx/html/
root@a6e9d14ca01f:/usr/share/nginx/html# ls
Dockerfile  index.html
root@a6e9d14ca01f:/usr/share/nginx/html# cat index.html 
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>Your 2nd Dockerfile worked!</title>

</head>

<body>
  <h1>You just successfully ran a container with a custom file copied into the image at build time!</h1>
</body>
</html>

root@a6e9d14ca01f:/usr/share/nginx/html#  ls
Dockerfile  index.html	test.txt

# Host Linux Commands
touch test.txt && echo "hi this is on host" > test.txt
```

- Now the `localhost:80/test.txt` will give the output of the file we just created and this proves that the container is using the host volume for source code
- You can just run the container logs to see if there are any errors while coding

### NOTE : You can look through on Docker Hub before making a new volume using -v for a container before running it to see where it expects the data path to be. Look out for the VOLUME command! 
Take mysql image for example
- The image is programmed in a way that it tells Docker that when we start a new container from it, to create a new volume and assign it to the directory mentioned which means that any file put there will be persistent across container lifecyle!
- Dockerfile isn't part of the image metadata so you won't see it upon inspection but the volume specified will be there.

# Assignment - Database Upgrades with named volumes
- For running postgres in 2024, you have to set environment variables `POSTGRES_PASSWORD=mypasswd` or ignore passwords as it used with `POSTGRES_HOST_AUTH_METHOD=trust`. This change was for the docker hub image, not postgres itself!
- Consider old as `postgres:15.1` and new as `postgres:15.2`

*Imagine a situation where you're running a particular database, say postgres and you need to update the patch version. Normally you could just update the software but for containers, it is not as simple*
- *Create a postgres container with named volume psql-data using v9.6.1*
- *Use docker hub to learn the VOLUME path and versions needed to run it*
- *Check logs in the container*
- *Create a new container with a different version*
- *Check logs to validate*

***This only works for patch versions, most SQL DBs require manual commands to upgrade DBs to major/minor versions which is a DB limitation***

```sh
# Pull the necessary images
docker image pull postgres:15.1
docker image pull postgres:15.2

# Run first container
docker container run -d --name pgres -e POSTGRES_PASSWORD=mypasswd -v psql-data:/var/lib/postgresql/data postgres:15.1

# Check Logs
docker container logs pgres

# Check Volume
docker volume ls

# Stop and remove first container
docker container rm -f pgres

# Run second container
docker container run -d --name pgres2 -e POSTGRES_PASSWORD=mypasswd -v psql-data:/var/lib/postgresql/data postgres:15.2  

# Check Logs
docker container logs pgres2

# Check Volume
docker volume ls

# It is observed that the logs for 2nd container were MUCH smaller compared to that of first container
```

---

# File Permissions

**Note** - You may have file permission problems with container apps not having the permissions they need. Say you want multiple containers to access the same volume(s) or bind mount existing files to a container. In a production environment with just `dockerd` where the permissions are not auto-translated
- File ownership between containers and host are just numbers and stay consistent no matter how you run them
  - Sometimes you see friendly user names in commands like ls but those are just name-to-number aliases that you'll see in `/etc/passwd` and `/etc/group`
  - Your host has those files, and usually, your containers will have their own. They are usually different. These files are really just for humans to see friendly names. The Linux Kernel only cares about IDs, which are attached to each file and directory in the file system itself, and those IDs are the same no matter which process accesses them.

When a container is just accessing its own files, this isn't usually an issue. ***But for multiple containers accessing the same volume or bind-mount, problems can arise in two ways:***
- **The `/etc/passwd` is different across containers**. Creating a named user in one container and running as that user may use ID 700, but that same name in another container with a different `/etc/passwd` may use a different ID for that same username. 
  - You'll see this confusion if you're running a container on a Linux VM and it had a volume or bind-mount. If you do an ls on those files from the host, it may show them owned by ubuntu or node or systemd, etc. Then if you run ls inside the container, it may show a different friendly username. The IDs are the same in both cases, but the host will have a different passwd file than the container, and show you different friendly names. ***Different names are fine, because it's only ID that counts. Two processes trying to access the same file must have a matching user ID or group ID.***

- **Your two containers are running as different users.** Maybe the user/group IDs and/or the USER statement in your Dockerfiles are different, and the two containers are technically running under different IDs. Different apps will end up running as different IDs. For example, the node base image creates a user called node with ID of 1000, but the NGINX image creates an nginx user as ID 101. Also, some apps spin-off sub-processes as different users. NGINX starts its main process (PID 1) as root (ID 0) but spawns sub-processes as the nginx user (ID 101), which keeps it more secure. 

## Troubleshooting
```docker
# Use command ps aux in each container to see a list of processes and usernames. The process needs a matching user ID or group ID to access the files in question.

# Find the UID/GID in each containers `/etc/passwd` and `/etc/group` to translate names to numbers. You'll likely find there a miss-match, where one containers process originally wrote the files with its UID/GID and the other containers process is running as a different UID/GID.

# Figure out a way to ensure both containers are running with either a matching user ID or group ID. This is often easier to manage in your own custom app (when using a language base image like python or node) rather than trying to change a 3rd party app's container (like nginx or postgres)... but it all depends. This may mean creating a new user in one Dockerfile and setting the startup user with USER. The node default image has a good example of the commands for creating a user and group with hard-coded IDs:

RUN groupadd --gid 1000 node \\
    && useradd --uid 1000 --gid node --shell /bin/bash --create-home node
USER 1000:1000

# When setting a Dockerfile's USER, use numbers, which work better in Kubernetes than using names.

# If ps doesn't work in your container, you may need to install it. In debian-based images with apt, you can add it with apt-get update && apt-get install procps
```

---

# Exercise - Edit code running in containers with bind mounts
- *Use a jekyll 'static site generator' to start a local web server*
- *Not about wev-dev but basically bridging the gap between local file access and apps running in containers*
- *Edit files with editor on our host using native tools*
- *Start container with `docker run -p 80:4000 -v $(pwd):/site bretfisher/jekyll-serve`*
- *Refresh browser to see changes*
- *Change the file in `_posts\` and refresh browser to see changes*

```sh
# Run the container
docker run -p 80:4000 -v $(pwd):/site bretfisher/jekyll-serve

# Check logs
docker container logs 

# Make changes under _posts

# Check logs again
    Regenerating: 1 file(s) changed at 2024-10-31 13:32:29
                    _posts/2020-07-21-welcome-to-jekyll.markdown
       Jekyll Feed: Generating feed for posts
                    ...done in 0.0664694 seconds.

# It reloaded as you can see
```

# Docker Compose
- It is used to configure relationships between containers
- Save docker container settings in easy to read files
- Create ONE liner dev environment startups
- Comprised of:
  - YAML Files for configuration of containers, networks and volumes i.e. basically describing our solution architecture.
  - A CLI tool `docker-compose` used for local dev/test automation with these YAML files
- YAML is Yet Anoteher Markup Langugage

## Some Information
- *In 2022, Docker announced the General Availability of Docker Compose V2. It supports all the same commands taught in this course and is meant to be fully backward compatible. It's auto-installed by Docker Desktop.*
- *Behind the scenes, Docker has rebuilt the old docker-compose Python binary with go, the same language as the Docker CLI, and added Compose V2 as a CLI plugin rather than a separate command. It's now faster and more stable, and should "just work" as a drop-in replacement for the V1 docker-compose CLI.*
- *`docker-compose` up becomes `docker compose up` etc.*
- ***v2.x is actually better for local docker-compose use, and v3.x is better for use in server clusters (Swarm and Kubernetes)***
- ***It creates networks by default even if not specified so that containers can communicate with each other courtesy of DNS Auto Resolution***
- ***Usernames are typically an environment-specific setting, that can change often, so they are often set in environment variables.***
***DNS names (set as aliases) for containers in a compose file come from the service name declared in the .yml***

## docker-compose.yml
- The compose YAML has its own versions: 1,2,2.1,3,3.1
- Can be used with `docker-compose` command for local docker automation 
- Can be used directly in production with swarm
- The default file name `docker-compose.yml` is there but any file can be used with the `-f` command

---

### The format

```yaml
# version isn't needed as of 2020 for docker compose CLI. 
# All 2.x and 3.x features supported
# Docker Swarm still needs a 3.x version
# version: '3.9'

services:  # containers. same as docker run
  servicename: # a friendly name. this is also DNS name inside network
    image: # Optional if you use build:
    command: # Optional, replace the default CMD specified by the image
    environment: # Optional, same as -e in docker run
    volumes: # Optional, same as -v in docker run
  servicename2:

volumes: # Optional, same as docker volume create

networks: # Optional, same as docker network create

```

### A sample file

```yaml
version: "3.9"

services:
  ghost:
    image: ghost
    ports:
      - "80:2368"
    environment:
      - URL=http://localhost
      - NODE_ENV=production
      - MYSQL_HOST=mysql-primary
      - MYSQL_PASSWORD=mypass
      - MYSQL_DATABASE=ghost
    volumes:
      - ./config.js:/var/lib/ghost/config.js
    depends_on:
      - mysql-primary
      - mysql-secondary

  proxysql:
    image: percona/proxysql
    environment:
      - CLUSTER_NAME=mycluster
      - CLUSTER_JOIN=mysql-primary,mysql-secondary
      - MYSQL_ROOT_PASSWORD=mypass
      - MYSQL_PROXY_USER=proxyuser
      - MYSQL_PROXY_PASSWORD=s3cret

  mysql-primary:
    image: percona/percona-xtradb-cluster:5.7
    environment:
      - CLUSTER_NAME=mycluster
      - MYSQL_ROOT_PASSWORD=mypass
      - MYSQL_DATABASE=ghost
      - MYSQL_PROXY_USER=proxyuser
      - MYSQL_PROXY_PASSWORD=s3cret

  mysql-secondary:
    image: percona/percona-xtradb-cluster:5.7
    environment:
      - CLUSTER_NAME=mycluster
      - MYSQL_ROOT_PASSWORD=mypass
      - CLUSTER_JOIN=mysql-primary
      - MYSQL_PROXY_USER=proxyuser
      - MYSQL_PROXY_PASSWORD=s3cret
    depends_on:
      - mysql-primary
```

---

## Docker Compose Commands
- Docker compose is not a production grade tool, its super ideal for local dev/test
- Two commands are primarily used:
  - `docker compose up` to setup volumes/networks and start all containers
  - `docker compose down` to stop all containers and remove stuff
- It is great for a dev team where there is a new developer on-boarding as it is as easy as `git clone something` and `docker compose up` so it makes things VERY easy

# Exercise -  Build a Compose File For a Multi-Container Project
*Create a compose config for a local Drupal CMS website*
- *You have to create a docker-compose.yml*
- *Use `drupal:9` image along with the `postgres:14` image*
- *Use `ports` to expose Drupal on 8080*
- *Setup the POSTGRES_PASSWORD environment variable on the postgres service*
- *Also note that the postgres official image defaults to user:postgres and database:postgres*
- *Walk through the Drupal config steps in browser at `http://localhost:8080`*
- *Drupal setup will assume the database runs on localhost, which is incorrect. You'll need to change it under Advanced settings to the name of the Docker service you gave to postgres*
- *Use Docker Hub documentation to figure out the right environment and volume settings*
- *Use volumes to store Drupal unique data*

```yaml
services:
  cms:
    image: drupal:9
    ports:
      - "8080:80"
    volumes:
      - drupal-modules:/var/www/html/modules
      - drupal-profiles:/var/www/html/profiles
      - drupal-themes:/var/www/html/themes
      - drupal-sites:/var/www/html/sites
  db:
    image: postgres:14
    environment:
      - POSTGRES_PASSWORD=pass
    
volumes:
  drupal-modules:
  drupal-profiles:
  drupal-themes:
  drupal-sites:
```

```sh
# You can pull an image and then inspect it to check out the ports it is using
docker compose up

# Change the name to the container's name from docker container ls -a under the advanced options tab

docker compose down -v
```

---

# Image Building in Compose
- Compose can build your custom images
- It will build them upon the usage of `docker compose up` if the image is not found cached
- It will rebuild the image with `docker compose build` or `docker compose up --build`
- A great use case for this is when you have lots of vars or build args

```Dockerfile
# Using build args

ARG NODE_VERSION="20"
ARG ALPINE_VERSION="3.20"

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} AS base
WORKDIR /src

FROM base AS build
COPY package*.json ./
RUN npm ci
RUN npm run build

FROM base AS production
COPY package*.json ./
RUN npm ci --omit=dev && npm cache clean --force
COPY --from=build /src/dist/ .
CMD ["node", "app.js"]

# Overriding can be done like this
# docker build --build-arg NODE_VERSION=current
```

Consider this docker-compose file
```yaml
services:
  proxy:
    build:
      context: .
      dockerfile: nginx.Dockerfile
    ports:
      - '80:80'
  web:
    image: httpd
    volumes:
      - ./html:/usr/local/apache2/htdocs/
```
- Here we will be using the Dockerfile in the present directory, hence the `.` 
- The name of the image upon building will be `nginx-custom`
- There is a second service of `apache` where we are using *bind mounts* to mount the local html files into the container
- Our use case here is that when we go into production with this, it will be sitting behind an nginx proxy and that is what we're emulating here

***This is a very common example of a developer setup where you need to build some custom images locally and mount some files into your application. In case it was a DB backed application, it can be simply added here***

---

# Exercise - Compose for Image Building
*This time imagine you're just wanting to learn Drupal's admin and GUI, or maybe you're a software tester and you need to test a new theme for Drupal. When configured properly, this will let you build a custom image and start everything with `docker-compose up` including storing important db and config data in volumes so the site will remember your changes across Compose restarts.*
### Starting off with this docker-compose.yml
```yaml
services:
  cms:
    image: drupal:9
    ports:
      - 8080:80
    volumes:
      - drupal-modules:/var/www/html/modules
      - drupal-profiles:/var/www/html/profiles
      - drupal-themes:/var/www/html/themes
      - drupal-sites:/var/www/html/sites
  db:
    image: postgres:14
    environment:
      - POSTGRES_PASSWORD=pass
    
volumes:
  drupal-modules:
  drupal-profiles:
  drupal-themes:
  drupal-sites:
```

### Dockerfile
- *First you need to build a custom Dockerfile in this directory,`FROM drupal:9`*
- *Then RUN apt package manager command to install git: `apt-get update && apt-get install -y git`*
- *Remember to cleanup after your apt install with `rm -rf /var/lib/apt/lists/*` and use `\` and `&&` properly. You can find examples of them in drupal official image. More on this below under Compose file.*
- *Then change `WORKDIR /var/www/html/themes`*
- *Then use git to clone the theme with: `RUN git clone --branch 8.x-4.x --single-branch --depth 1 https://git.drupalcode.org/project/bootstrap.git`*
- *Combine that line with this line, as we need to change permissions on files and don't want to use another image layer to do that (it creates size bloat). This drupal container runs as www-data user but the build actually runs as root, so often we have to do things like `chown` to change file owners to the proper user: `chown -R www-data:www-data bootstrap`. Remember the first line needs a `\` at end to signify the next line is included in the command, and at start of next line you should have `&&` to signify "if first command succeeds then also run this command"*
- *Then, just to be safe, change the working directory back to its default (from drupal image) at `/var/www/html`*

```Dockerfile
FROM drupal:9
RUN apt-get update \
    && apt-get install -y git \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /var/www/html/themes
RUN git clone --branch 8.x-4.x --single-branch --depth 1 https://git.drupalcode.org/project/bootstrap.git \
    && chown -R www-data:www-data bootstrap
WORKDIR /var/www/html 
```

### Compose File
- *We're going to build a custom image in this compose file for drupal service. Use Compose file from previous assignment for Drupal to start with, and we'll add to it, as well as change image name.*
- *Rename image to `custom-drupal` as we want to make a new image based on the official `drupal:9`.*
- *We want to build the default Dockerfile in this directory by adding `build: .` to the `drupal` service. When we add a build + image va  lue to a compose service, it knows to use the image name to write to in our image cache, rather then pull from Docker Hub.*
- *For the `postgres:14` service, you need the same password as in previous assignment, but also add a volume for `drupal-data:/var/lib/postgresql/data` so the database will persist across Compose restarts.*

```yaml
version: "3.9"

services:
  cms:
    image: drupal:9
    ports:
      - "8080:80"
    volumes:
      - drupal-modules:/var/www/html/modules
      - drupal-profiles:/var/www/html/profiles
      - drupal-themes:/var/www/html/themes
      - drupal-sites:/var/www/html/sites

  db:
    image: postgres:14
    environment:
      - POSTGRES_PASSWORD=pass
    volumes:
      - drupal-data:/var/lib/postgresql/data

volumes:
  drupal-modules:
  drupal-profiles:
  drupal-themes:
  drupal-sites:
  drupal-data:
```

### CLI
- *Start containers like before, configure Drupal web install like before.*
- *After site comes up, click on `Appearance` in top bar, and notice a new theme called `Bootstrap` is there. That's the one we added with our custom Dockerfile.*
- *Click `Install and set as default`. Then click `Back to site` (in top left) and the site interface should look different. You've successfully installed and activated a new theme in your own custom image without installing anything on your host other than Docker!*
- *If you exit (ctrl-c) and then `docker-compose down` it will delete containers, but not the volumes, so on next `docker-compose up` everything will be as it was.*
- *To totally clean up volumes, add `-v` to `down` command.*

### The logs are much shorter when we restart the containers

---

# Swarm Mode
- Clustering solution built inside docker
- Not enabled by default:
  - `docker swarm`
  - `docker node`
  - `docker service`
  - `docker stack`
  - `docker secret`
- There are manager nodes which have a database locally on the called as the `raft database`. It stores their configuration and gives them all the information they need to have to be the authority inside a swarm.
- They manage worker nodes
## Problems with containers everywhere
- How to automate container lifecycle?
- How to scale out/in/up/down?
- How to recreate failing containers?
- How to replace containers without downtime? *Blue-Green Deploy*
- How to control/track where containers get started?
- How to talk across across nodes on different virtual networks?
- How to ensure only trusted servers run containers?
- How to store secrets, keys, passwords and get them to the right container?

> Docker service command is used in a swarm to replace the docker run command and allows to add extra features to the container when it's run such as how many replicas to run

> These are known as tasks and a service can have multiple tasks and each task launches a container

> Manager nodes are used to decide where in the swarm to place those containers. By default, they are tried to be spread out.

## An architectural overview
```
        docker service create
                  |
                  V
            Manager Node
            - API: Accept commands from client and create service object
            - Orchestrator: Reconcilation loop for service objects and create tasks
            - Allocator: Assign IPs to tasks
            - Scheduler: Assign nodes to tasks
            - Dispatcher: Check on the workers
                  /|\
                   |
                   |
              Worker Node
              - Worker: Connects to dispatcher to check on assigned tasks
              - Executor: Executes the tasks assigned to worker nodes 
```

## Using Swarm
- Swarm is disabled by default which can be checked by `docker info`
- It can be enabled by `docker swarm init`
  - PKI (Public Key Infrastructure) and Security automation
    - Root signing certificate created for swaem
    - Certificate issued for first manager node
    - Join tokens are created which can be used for other nodes to join the swarm
    - Raft database created to store root CA(Certificate Authority*), configs and secrets
    - Encrypted on disk by default
    - No need for another key/value system to hold orchestration/secrets
    - Replicates logs amongst managers via mutual TLS in "control plane"
- Swarm has a redundant config (orchestration and automation system) built straight into docker daemon

```sh
# Check nodes
docker node ls

# It can be seen that one node has the status of a leader, and there can only be ONE leader at a time

# Create a service
docker service create alpine:latest ping 8.8.8.8
# Pinging google DNS servers is our task

# Check running services
docker service ls
# The LHS in 1/1 under REPLICAS is how many services are running vs RHS which is how many were specified

# Goal of the orchestrators is to match these numbers

# Check the containers created by the service using their ID
docker service ps y0szkey7o3ti
# The "NODE" component tells which server is it running on
# There are incrementors in the container names as well such as pendantic_bartik.1

# Make updates to a service
docker service update pedantic_bartik --replicas 5
# OR
docker service scale pedantic_bartik=27
# Change number of replicas to 5

docker service ps pedantic_bartik 
# There will be pedantic_bartik.1 pedantic_bartik.2 and ... pedantic_bartik.5 

# To remove
docker service rm pedantic_bartik 
```
### What happens if we shutdown a container in a swarm manually
```
->  ~ docker container rm -f pedantic_bartik.5.j1hx57fx8iqtevl0m3np34yr8 
pedantic_bartik.5.j1hx57fx8iqtevl0m3np34yr8
->  ~ docker service ls
ID             NAME              MODE         REPLICAS   IMAGE           PORTS
y0szkey7o3ti   pedantic_bartik   replicated   4/5        alpine:latest   
->  ~ docker service ls
ID             NAME              MODE         REPLICAS   IMAGE           PORTS
y0szkey7o3ti   pedantic_bartik   replicated   5/5        alpine:latest   
```

- ***We're actually telling an orchestration system, hey put this job in your queue. When you can get to it, please perform the actions on the swarm that I've asked here.***

- ***There's failure mitigation and a lot of intelligence built into that.*** 

- *Use the defaults if you're interactive at the CLI, typing commands yourself.*
- *Use `--detach true` if you're using automation or shell scripts to get things done.*

---

# 3-Node Swarm Cluster
- Multipass creates full Ubuntu server VM on your Host machine
- Get started by `sudo snap install multipass` (I've used gnome terminal with zsh and ubuntu as my os)
## Setting up Multipass
```sh
# Create instances
multipass launch docker --name cloud1 \
&& multipass launch docker --name cloud2 \
&& multipass launch docker --name cloud3

# List Instances
multipass list

# Once the instances are up, open 3 tabs for the three instances and run on seperate windows
multipass exec cloud1 bash
multipass exec cloud2 bash
multipass exec cloud3 bash

# Initialize swarm, here it works out of the box and no need was observed to add the --advertise-addr
# Its possible for a node to directly join as a manager and you have to run this for that
docker swarm join-token manager

# cloud1
docker swarm init
# cloud2
docker swarm join --token SWMTKN-1-42ks8j63wlbqdojrabtqlkzat6r4iz58ygkjom865jz8f5wgdc-7eerxenht97ecy3rjtah9chrc 10.217.73.214:2377
# cloud3
docker swarm join --token SWMTKN-1-42ks8j63wlbqdojrabtqlkzat6r4iz58ygkjom865jz8f5wgdc-7eerxenht97ecy3rjtah9chrc 10.217.73.214:2377

# cloud1
docker node update --role manager cloud2
# Create 9 tasks on our swarm cluster
docker service create --replicas 9 alpine:latest ping 8.8.8.8

# Output
ubuntu@cloud1:~$ docker service ls
ID             NAME                  MODE         REPLICAS   IMAGE           PORTS
ryaol6kjfvd2   youthful_sutherland   replicated   9/9        alpine:latest   

# Which services are running on cloud1
ubuntu@cloud1:~$ docker node ps
ID             NAME                    IMAGE           NODE      DESIRED STATE   CURRENT STATE           ERROR     PORTS
lk111jaxzya3   youthful_sutherland.4   alpine:latest   cloud1    Running         Running 4 minutes ago             
nfx1nrcxe2r9   youthful_sutherland.5   alpine:latest   cloud1    Running         Running 4 minutes ago             
y1l8493kcfni   youthful_sutherland.9   alpine:latest   cloud1    Running         Running 4 minutes ago       

# Which services are runnig on cloud2
ubuntu@cloud1:~$ docker node ps cloud2
ID             NAME                    IMAGE           NODE      DESIRED STATE   CURRENT STATE           ERROR     PORTS
sk1tbrrqfl51   youthful_sutherland.1   alpine:latest   cloud2    Running         Running 5 minutes ago             
ykywdlmnabuw   youthful_sutherland.6   alpine:latest   cloud2    Running         Running 5 minutes ago             
wwntl8fnk0wp   youthful_sutherland.7   alpine:latest   cloud2    Running         Running 5 minutes ago       

# Overall running status of all services
ubuntu@cloud1:~$ docker service ps youthful_sutherland 
ID             NAME                    IMAGE           NODE      DESIRED STATE   CURRENT STATE           ERROR     PORTS
sk1tbrrqfl51   youthful_sutherland.1   alpine:latest   cloud2    Running         Running 5 minutes ago             
ynqwoooswohf   youthful_sutherland.2   alpine:latest   cloud3    Running         Running 5 minutes ago             
jibn3jyorffc   youthful_sutherland.3   alpine:latest   cloud3    Running         Running 5 minutes ago             
lk111jaxzya3   youthful_sutherland.4   alpine:latest   cloud1    Running         Running 5 minutes ago             
nfx1nrcxe2r9   youthful_sutherland.5   alpine:latest   cloud1    Running         Running 5 minutes ago             
ykywdlmnabuw   youthful_sutherland.6   alpine:latest   cloud2    Running         Running 5 minutes ago             
wwntl8fnk0wp   youthful_sutherland.7   alpine:latest   cloud2    Running         Running 5 minutes ago             
rjm7jnrk8hog   youthful_sutherland.8   alpine:latest   cloud3    Running         Running 5 minutes ago             
y1l8493kcfni   youthful_sutherland.9   alpine:latest   cloud1    Running         Running 5 minutes ago     
```

### And that's it, a 3 Node Swarm Cluster is good to go!

# Multi Host Networking
- Swarm brings a new network driver called overlay which can be created by `docker network create --driver overlay test`
  - This is for creating a swarm wide bridge network.
- Containers across hosts on the same virtual network can access each other kind of like they're on a VLAN.
  - This is only for intra swarm communication.
- Full network encrytption is optional and can be done using IPSec (AES Encryption) but its `off by default` for performance reasons.

```sh
# Create the network
docker network create --driver overlay mydrupal

# Run a postgres service on the mydrupal network
docker service create --name psql --network mydrupal -e POSTGRES_PASSWORD=mypass postgres:14

# Check service created
docker service ls

# Check containers
docker service ps psql

# Check logs 
docker container logs psql.1.vrr0ks2w6fo9ezdgdj2hslksk

# Run a drupal service
docker service create --name drupal --network mydrupal -p 80:80 drupal:9

# Check services and the node drupal is running on
docker service ls
docker service ps drupal
```
- With overlay, it appears as if everything acts on the same subnet
- A cool thing to observer here is that regardless of WHICH IP address we use from either of the nodes, we get the correct website
- Using the service name for psql

## Routing Mesh
- Its an ingress network distributing packets for our service to the task for that service across all nodes in swarm.
- Uses IPVS from Linux Kernel.
- Load balances swarm services across their tasks.
- Consider our backend system (databases) are increased to two replicas and the frontends talk to the backend. We don't directly talk to the IP addresses... Routing mesh works in two ways:
  - Container to Container in an overlay network using VIP(Virtual IP)
    - Load is distrubuted amongst all tasks for a service. 
    - Additional load balancer is not needed.
  - External traffic coming into the swarm can choose to hit any of the nodes in the swarm and any of the worker nodes having that published port open and listening for that container's traffic. Then the traffic is re-routed to the proper container based on its load balancing.
- ***BASICALLY*** If it's not a different node, its router over the virtual network AND If it's on the same node, it's re-routed to the port of that container.

### All this is out of the box!

## An example!
Consider a newsform service with 3 replicas!
- 3 tasks
- 3 containers
- 3 Nodes
- In the overlay network, a VIP is created mapped to the DNS name of the service which by default is the `service name`
- Lets say the service name is `my-web` and any containers in the overlay that need to talk to that service in the swarm only have to worry about using the `my-web` DNS.
  - The VIP load balances the traffic amongst all tasks in that service.
  - This is not DNS RR, we can use it however!
- Benefits of VIP over RR:
  - DNS Caches in our apps prevent us from properly load balancing.
  - Rather than fighting DNS clients in DNS config, we rely on VIP similar to a hardware load balancer.

## Another Example!
Consider the drupal model created above!
- There is a service called `my-web` with 2 tasks applied to two different nodes.
- Each node has a built in load balancer on the external IP address.
- Any traffic that hits `IP:Published_Port` hits the load balancer and it decides which container should get the traffic
  - Whether its on the local node
  - Or it needs to send it over the network

```sh
# Cloud1
docker service create --name search --replicas 3 -p 9200:9200 elasticsearch:2
curl localhost:9200
```

We get

```json
ubuntu@cloud1:~$ curl localhost:9200
{
  "name" : "Charcoal",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "WdT9pe3_RI6TmhWDQiGndA",
  "version" : {
    "number" : "2.4.6",
    "build_hash" : "5376dca9f70f3abef96a77f4bb22720ace8240fd",
    "build_timestamp" : "2017-07-18T12:17:44Z",
    "build_snapshot" : false,
    "lucene_version" : "5.5.4"
  },
  "tagline" : "You Know, for Search"
}
ubuntu@cloud1:~$ curl localhost:9200
{
  "name" : "Psyche",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "8rOxwnIXTTGfkP73ZiKFQw",
  "version" : {
    "number" : "2.4.6",
    "build_hash" : "5376dca9f70f3abef96a77f4bb22720ace8240fd",
    "build_timestamp" : "2017-07-18T12:17:44Z",
    "build_snapshot" : false,
    "lucene_version" : "5.5.4"
  },
  "tagline" : "You Know, for Search"
}
ubuntu@cloud1:~$ curl localhost:9200
{
  "name" : "Honey Lemon",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "7Ht9DSAsQN-331zzRdKHnQ",
  "version" : {
    "number" : "2.4.6",
    "build_hash" : "5376dca9f70f3abef96a77f4bb22720ace8240fd",
    "build_timestamp" : "2017-07-18T12:17:44Z",
    "build_snapshot" : false,
    "lucene_version" : "5.5.4"
  },
  "tagline" : "You Know, for Search"
}
ubuntu@cloud1:~$ curl localhost:9200
{
  "name" : "Charcoal",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "WdT9pe3_RI6TmhWDQiGndA",
  "version" : {
    "number" : "2.4.6",
    "build_hash" : "5376dca9f70f3abef96a77f4bb22720ace8240fd",
    "build_timestamp" : "2017-07-18T12:17:44Z",
    "build_snapshot" : false,
    "lucene_version" : "5.5.4"
  },
  "tagline" : "You Know, for Search"
}
```
- Routing mesh is stateless load balancing.
- If you have to use session cookies on your application or it expects a consistent container talking to a consistent client, then you may have to add a few more things!
- This is a *Layer 3* Load Balancer acting at *IP* and *PORT* layer, not the *DNS*.
- You can't just run multiple websites on the same port on the same swarm just YET!
  - Nginx or HAProxy can be used to do this. They sit in front with your routing mesh and act as a stateful load balancer which can also do caching.
  - Docker Enterprise Edition comes with a built-in Level 4 web proxy allowing you to just throw DNS names in the web config of your swarm services and everything just works.

# Exercise - Create a multi-service multi-node web app
- *Using docker's distributed voting app.*
- *All images are on Docker Hub, so you should use editor to craft your commands locally.*
- *A `backend` and `frontend` overlay network are needed. Nothing different about them other than that backend will help protect database from the voting web app.*
- *The database server should use a named volume for preserving data. Use the new `--mount` format to do this: `--mount type=volume,source=db-data,target=/var/lib/postgresql/data`* 

### Services (names below should be service names)
- vote
  - bretfisher/examplevotingapp_vote
  - web frontend for users to vote dog/cat
  - ideally published on TCP 80. Container listens on 80
  - on frontend network
  - 2+ replicas of this container

- redis
  - redis:3.2
  - key-value storage for incoming votes
  - no public ports
  - on frontend network
  - 1 replica NOTE VIDEO SAYS TWO BUT ONLY ONE NEEDED

- worker
  - bretfisher/examplevotingapp_worker
  - backend processor of redis and storing results in postgres
  - no public ports
  - on frontend and backend networks
  - 1 replica

- db
  - postgres:9.4
  - one named volume needed, pointing to /var/lib/postgresql/data
  - on backend network
  - 1 replica
  - remember set env for password-less connections -e POSTGRES_HOST_AUTH_METHOD=trust

- result
  - bretfisher/examplevotingapp_result
  - web app that shows results
  - runs on high port since just for admins (lets imagine)
  - so run on a high port of your choosing (I choose 5001), container listens on 80
  - on backend network
  - 1 replica

### Commands
```sh
docker network create --driver overlay backend 

docker network create --driver overlay frontend 

docker service create --name vote -p 80:80 --replicas 2 --network frontend bretfisher/examplevotingapp_vote 

docker service create --name redis --network frontend redis:3.2 

docker service create --name dbpsql --network backend -e POSTGRES_HOST_AUTH_METHOD=trust --mount type=volume,source=db-data,target=/var/lib/postgresql/data postgres:9.4

docker service create --name worker --network backend --network frontend bretfisher/examplevotingapp_worker

docker service create --name result --network backend -p 5001:80 bretfisher/examplevotingapp_result
```

---

# Stacks - Production Grade Compose
- Stacks accept compose files as their declarative definition for services, networks and volumes.
- Usage is `docker stack deploy` instead of docker service create to import the compose file and run
  - It manages all of it including overlay network per stack and adds stack name to start of their name.
- There is a `deploy:` key in the compose file whereas `build:` is not there but allows swarm specific tweaking
  - How many replicas?
  - What to do on fail-over?
  - How to do rolling updates?
- *Building should be done on your CI system with something like Jenkins and push the built images directly into your repository. The stack will pull them and deploy them with `deploy` options.*
- Compose ignores `deploy:` and Swarm ignores `build:` so you don't need to change your file.
- `docker compose cli` not needed on your swarm. So the docker engine can read compose files through the stack command without `docker compose` on the server.
- Stack is for just one swarm!

```sh
# Usage
docker stack deploy -c example-voting-app-stack.yml --detach=true voteapp

# Check which services are there and what's running in the stack
docker stack ps voteapp

# The long names of the container are due to them having a GUID to guarantee unique names

# Similar to docker service ls but restricted to the stack specified
docker stack services voteapp
```

- *A great comparison between docker compose and swarm by BretFisher <a href="https://github.com/BretFisher/ama/discussions/146">docker-compose or single node swarm.</a>*

- *`dockersamples/visualizer` is a great tool to see which service is running on which node.* 

- Now you COULD manually update a service, say make it's replicas as 5 but that's an *anti-pattern* because the changes made to the `.yaml` file will overwrite those changes. ***Your config file should be the source of truth.*** Simply re-running the command on an existing stack is enough as it realises that the stack exists and makes the updates.

---

# Swarm Secrets
- An easy secure solution for storing secrets in swarm.
  - Easy because it's built into swarm.
  - Comes out of the box and only needs swarm to be initialized.
- Encrypted on disk, Encrypted during transit and only available where it's need to be.
- What's a secret?  
  - Usernames and Passwords
  - TLS Certificates and keys
  - SSH Keys
  - Any PII
- Supports 500kb strings or binaries.
- Apps don't need to re-written.
- Swarm Raft DB is encrypted on disk by default.
  - Stored on disk of manager nodes.
  - Only managers have the keys.
  - TLS + Mutual Auth control plane is the defaut.
- Secrets are first stored in swarm and then assigned to the service. Service commands or Swarm Stack can be used for doing this.
- Only the container in assigned service can see the secrets.
- Docker worker keeps the key secure in memory only and only gets it down to the containers that need them.
- They look like a file to your apps in the container like `/run/secrets/<secret_name>` or `/run/secrets/<secret_alias>` but they're not running on disks and are in memory using a ramfs (RAM based file system) so it's like a key value store.
- Local docker compose can use file based secrets but it's not secure and shouldn't be used on a production server.
- If you don't have swarm, you can't use secrets.
- `docker compose` has a workaround where it mounts the secrets in a clear text file into the local container and allows you to use secrets locally!

## Creating Secrets
### Using a file
```sh
docker secret create psql_user psql_user.txt

echo "dbpassword" | docker secret create psql_pass -
# Echo the key into the creation command where the - is tellng the command to read from the STDIN
# These need to be manually removed using docker secret rm
```
- In the first use case, the file is on the host which is an anti-pattern.
- A possibility is using the remote API from CLI and then pass the files that way.
- Another thing to know is that all this is going into the history of the root user so if someone got into root, password will be out.
- Only the containers and services to which the secret is assigned to can access the value

```sh
# Using it in a container
docker service create --name psql --secret psql_user --secret psql_pass -e POSTGRES_PASSWORD_FILE=/run/secrets/psql_pass -e POSTGRES_USER_FILE=/run/secrets/psql_user postgres:9.4

# Going into the container
docker container exec -it psql.1.nf74a1ekisoorkcwaw7vtrrud bash

# Accessing them from the container
root@6841bb4890ea:/run/secrets# ls
psql_pass  psql_user
root@6841bb4890ea:/run/secrets# cat psql_pass 
dbpassword
root@6841bb4890ea:/run/secrets# cat psql_user 
mypsqluser
```

- Secrets are part of immutable design of services so if anything in the container is changed, the container is redeployed which is not ideal for databases. Can be done with `docker service update --secret-rm psql_pass psql`

---

### Using Stack
- Version 3 was required to use stack but `3.1` is required to use secrets.
  ```yaml
  version: "3.9"

  services:
    psql:
      image: postgres
      secrets:
        - psql_user
        - psql_password
      environment:
        POSTGRES_PASSWORD_FILE: /run/secrets/psql_password
        POSTGRES_USER_FILE: /run/secrets/psql_user

  secrets:
    psql_user:
      file: ./psql_user.txt 
    psql_password:
      file: ./psql_password.txt
  ```

---
# Exercise - Create Stack with Secrets
### Start off with this
```yaml
services:
  cms:
    image: custom-drupal
    build: 
      context: .
      dockerfile: Dockerfile

    ports:
      - 8080:80
    volumes:
      - drupal-modules:/var/www/html/modules
      - drupal-profiles:/var/www/html/profiles
      - drupal-themes:/var/www/html/themes
      - drupal-sites:/var/www/html/sites
  db:
    image: postgres:14
    environment:
      - POSTGRES_PASSWORD=pass
    volumes:
      - drupal-data:/var/lib/postgresql/data
    
volumes:
  drupal-modules:
  drupal-profiles:
  drupal-themes:
  drupal-sites:
  drupal-data:
```

- *Rename drupal image to `drupal:9`*
- *Remove `build:`*
- *Add secret via `external:`*
- *Use environment variable `POSTGRES_PASSWORD_FILE`*
- *Add secret via cli `echo "pw" | docker secret create psql-pw -`*

### What you should have

```yaml
version: "3.9"

services:
  cms:
    image: custom-drupal
    build: 
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:80"
    volumes:
      - drupal-modules:/var/www/html/modules
      - drupal-profiles:/var/www/html/profiles
      - drupal-themes:/var/www/html/themes
      - drupal-sites:/var/www/html/sites

  db:
    image: postgres:14
    environment:
      - POSTGRES_PASSWORD=pass
    volumes:
      - drupal-data:/var/lib/postgresql/data

volumes:
  drupal-modules:
  drupal-profiles:
  drupal-themes:
  drupal-sites:
  drupal-data:
secrets:
  psql-pw:
    external: true
```

---

```sh
echo "password" | docker secret create psql-pw - && docker stack deploy -c docker-compose-stack.yml --detach=true drupal_cms
```

---

# Use secrets locally
- Can use the same compose file.
- Can use the same objects such as environment variables for postgres db.
- When we run `docker compose up -d`, it works.
  - The file is present with our secret which is NOT secure but it works.
  - It bind mounts at run time the actual file on the disk to the container.
  - It's just doing a `-v` with a particular file in the background.
- Only objective is to ease up the production by allowing similarity in both environments.
- Only works with file based secrets.

# Full App LifeCycle with Compose
- Single set of compose files for:
  - Local `docker compose up` - development environment.
  - Remote `docker compose up` - CI (Continous Integration) environment.
  - Remote `docker stack deploy` - production environment.

## Override
- When there is a default docker compose file and it sets the defaults that are same across all environments.
- Then there is the override file that by default if its named `docker-compose.override.yml` will automatically bring it up when `docker compose up` is performed. It overrides the defaults if there's any.
- The `.prod` and `.test` need to be specified using `-f`
- Since there won't be docker compose CLI on a production server, docker compose config command can be used.
  - The config command is actually going to do an output by combing the output of multiple config files.
  - **Note**: In a CI environment, there can be options to use volumes which can have a sample database as usually the containers are not supposed to have data persisted in CI.

```sh
# In the CI environment, assuming docker compose is present
docker compose -f docker-compose.yml -f docker-compose.test.yml up -d

# For the swarm or production mode
docker compose -f docker-compose.yml -f docker-compose.prod.yml config > output.yml
```

This `output.yml` is used in the production environment.

---

# Service Updates
- Provides rolling replacement of tasks/containers in a service.
- Limits downtime. Any service requiring persistent data storage.
- Will replace containers so it's better to be prepared.
- There are many CLI options to control the update.
- Create options will usually change, adding -add or -rm to them.
- There are rollbacks and healthchecks to consider as well.
  - `docker service scale web=4`
  - `docker service rollback web`
- A stack deploy to the same stack IS AN UPDATE.

## Swarm Update Examples
- Just update the image used to a newer version:
  - `docker service update --image myapp:1.2.1 myappservice`
- Add en environment variable and remove a port:
  - `docker service update --env-add NODE_ENV=production --publish-rm 8080
- Change number of replicals of 2 services:
  - `docker service scale web=8 api=6`
- For stack updates, just edit the `.yaml`
  - `docker stack deploy -c file.yml stackname`
- Updates can be forced to reissue tasks to get services running on least used nodes (whatever the reason was for this) which is a form of load balancing.

```sh
docker service create -p 8088:80 --name web nginx:1.13.7
# Scale up the service
docker service scale web=5
# Scale service
docker service update --image nginx:1.13.6 web
# Change ports
docker service update --publish-rm 8088 --publih-add 8080:80 web
# Force update
docker service update --force web
docker service -rm web
```

---

# Healthchecks
- Supported in Dockerfile, Compose YML, `docker run` and swarm services.
- Executes the command as if you ran `exec` on the container. Even simple workers without exposed ports can run a simple command to validate whether they're returning good data.
- Expects 0 (OK) or 1 (Error)
- Three container states:
  - Starting
  - Healthy
  - Unhealthy
- Better than "application is running".
- Not an external monitoring replacement.
- Its about docker understanding if the container itself has a basic level of healthy.
- Shows up in `docker container ls`
- Check last 5 healthchecks in `docker container inspect`
- Docker run does nothing with healthchecks.
- Services will replace task if they fail this healthcheck. 
- Service updates wait for them before continuing.

## Example - Docker Run
```sh
# For error codes other than 1
docker run \
  --health-cmd="curl -f localhost:9200/_cluster/health || false" \
  --health-interval=5s \
  --health-retries=3 \
  --health-timeout=2s \
  --health-start-period=15s \ 
  elasticsearch:2

docker container run --name p2 -d --health-cmd="pg_isready -U postgres || exit 1" postgres  
```

## Example - Dockerfile
- Options:
  - `--interval=30s`
  - `--timeout=30s`
  - `--start-period=12` (Default 0 and add a start period of considering broken health checks)
  - `--retries=3` (How many chances to make a comeback)
- Basic command usind default options:
  - HEALTHCHECK curl -f http://localhost || false
- Custom options with the command
  - `HEALTHCHECK --timeout=2s --interval=3s --retries=3 \ CMD curl -f http://localhost || exit 1` 
- It depends on the app as well:
  - ```Dockerfile
    FROM postgres
    # Specifying real user with -u to preven errors in logs
    HEALTHCHECK --interval=5s --timeout=3s \
    CMD pg_isready -u postgres || exit 1
    ```

## Example - Compose File
```yaml
version: "2.1"
services:
  web: 
    image: nginx
    healthcheck:
      test: ["CMD", "curl","-f","http://localhost"]
      interval: 1m30s
      timeout: 10s
      retries: 3 # Requires 2.1
      start_period: 1m # Requires 3.4

```

## Example - Services
```sh
# Doesnt take much time
docker service create --name p1 postgres
# Takes time
docker service create --name p2 --health-cmd="pg_isready -U postgres || exit 1" posgres
# Then shifts to running state
```

---

# Container Registries
- An image registry needs to be part of your container plan.

## Docker Hub
- Most popular and is the docker registry + lightweight image building.
- Can link GitHub/BitBucket to docker hub and auto build images on commit.
- Image building can be chained together to keep images up to date.
- Webhooks can be setup to other service like Jenkins or some other CI platform to have automated builds. So the allow to automate your code from something like GitHub to Docker and then to the servers.
- Don't create a repository using the button provided if your code is on github, you want to create an automated build from the CI path based on code commits. It's like a reverse webhook.
  - You can put repository links which will cause rebuilds when those change as well. 
  - It's like the dependency array `[]` in useEffect hook in react.
  - You can setup build triggers as well.

## Docker Registry
- It is a private image registry for your network.
- Is a part of the docker/distrubution GitHub repo so it's as simple as `docker pull registry`
- The de facto in private container registries. It's very barebones and has basic authentication only. It's everything or none, not RBAC. 
- Usually useful for only small teams. 
- Its just a web API and storage system written in GO at its heart.
- Storage supports local, S3, Azure, Alibaba, Google Cloud etc.
- Some things to consider while setting this up:
  - Setup TLS and secure your registry.
  - Storage cleanup with garbage collection.
  - Enable hub caching via `--registry-mirror`
- Docker wants TLS except `localhost` but there is still an `insecure registry` option

### Localhost

```sh
# Runs on port 5000 by default
docker container run -d -p 5000:5000 --name registry registry
# Root of local registry
docker tag hello-world 127.0.0.1:5000/hello-world
# Push to local registry
docker push 127.0.0.1:5000/hello-world
# After deleting and untagging images
docker pull 127.0.0.1:5000/hello-world
# You can create a volume for this as well to persist data
docker container run -d -p 5000:5000 --name registry -v $(pwd)/registry-data:/var/lib/registry registry
```

### Swarm
- Works the same way as localhost.
- Docker service command or stack can be used.
- Because of mesh routing, all nodes can see `127.0.0.1:5000`
- Care needs to taken for volumes and volume drivers.
- A swarm has to be able to pull images on all nodes from some repository in a registry that they can all reach. I didn't mention that in the swarm section, but it's an important point because that means that you can't just build an image of your own on a local.
- They all have to do push and pulls from a central repository somewhere. But because I ran the registry as a service here, they can all know how to access `127.0.0.1`, the routing mesh will actually steer them towards the correct container, wherever that registry is running on one of the nodes, and then it'll pull the required image to all the nodes as they need to be there.
- Use a hosted SaaS registry if possible, a good option is `Amazon ECS : Elastic Container Service`

```sh
docker service create --name registry --publish 5000:5000 registry

docker service ps registry

docker pull hello-world

docker tag hello-world 127.0.0.1:5000/hello-world

docker push 127.0.0.1:5000/hello-world
```