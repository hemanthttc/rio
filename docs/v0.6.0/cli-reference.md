# CLI Reference


## build

Build a docker image using buildkitd

##### Usage

```
rio build command [command options] [arguments...]
```

##### Options

| flag         | aliases | description                                                             | default |
|--------------|---------|-------------------------------------------------------------------------|---------|
| --file value | -f      | Name of the file to look for build, support both Riofile and Dockerfile |         |
| --tag value  | -t      | Name and optionally a tag in the 'name:tag' format                      |         |
| --build-arg  |         | Set build-time variables                                                |         |
| --no-cache   |         | Do not use cache when building the image                                |         |
| --help       | -h      | show help                                                               |         |


##### Examples

```shell script
# see previous builds from stacks or workloads
rio build history

# Navigate to directory with Dockerfile and build it into local registry
rio build -t test:v1

# See image that was build
rio image

# Build from riofile insted
rio build -t test:v1 --no-cache -f Riofile.yaml

# Now run the image
rio run -n test -p 8080 localhost:5442/default/test:v1

# Build the image again with new tag
rio build -t test:v2

# Now stage the 2nd image
rio stage --image localhost:5442/default/test:v2 test v2
```


## ps

List services

##### Usage
```
rio ps [OPTIONS]
```

##### Options

| flag        | aliases | description                                                            | default |
|-------------|---------|------------------------------------------------------------------------|---------|
| --quiet     | -q      | Only display Names                                                     |         |
| --format    |         | 'json' or 'yaml' or Custom format: '{{.Name}} {{.Obj.Name}}' [$FORMAT] |         |
| --all       | -a      | print all resources, including router and externalservice              |         |
| --workloads | -w      | include apps/v1 Deployments and DaemonSets in output                   |         |


##### Examples

```shell script
# show services and workloads
rio ps -w

# output json
rio ps --format json

# display name and weight in custom format
rio ps --format "{{.Obj.Name}} -> {{.Data.Weight}}" 
```

## image

List images built from local registry

##### Usage
```
rio image
```


## run

Create and run a new service

##### Usage
```
rio run [OPTIONS] IMAGE [COMMAND] [ARG...]
```

##### Options

| flag                       | aliases | description                                                                                   | default                |
|----------------------------|---------|-----------------------------------------------------------------------------------------------|------------------------|
| --add-host                 |         | Add a custom host-to-IP mapping (host:ip)                                                     |                        |
| --annotations              |         | Annotations to attach to this service                                                         |                        |
| --build-branch             |         | Build repository branch                                                                       | master                 |
| --build-dockerfile         |         | Set Dockerfile name                                                                           | defaults to Dockerfile |
| --build-context            |         | Set build context                                                                             | .                      |
| --build-webhook-secret     |         | Set GitHub webhook secret name                                                                |                        |
| --build-docker-push-secret |         | Set docker push secret name                                                                   |                        |
| --build-clone-secret       |         | Set git clone secret name                                                                     |                        |
| --build-image-name         |         | Specify custom image name to push                                                             |                        |
| --build-registry           |         | Specify to push image to                                                                      |                        |
| --build-revision           |         | Build git commit or tag                                                                       |                        |
| --build-pr                 |         | Enable pull request builds                                                                    |                        |
| --build-timeout            |         | Timeout for build, ( (ms/s/m/h))                                                              | 10m                    |
| --command                  |         | Overwrite the default ENTRYPOINT of the image                                                 |                        |
| --config                   |         | Configs to expose to the service (format: name[/key]:target)                                  |                        |
| --concurrency              |         | The maximum concurrent request a container can handle (autoscaling)                           | 10                     |
| --cpus                     |         | Number of CPUs                                                                                |                        |
| --dns                      |         | Set custom DNS servers                                                                        |                        |
| --dnsoption                |         | Set DNS options (format: key:value or key)                                                    |                        |
| --dnssearch                |         | Set custom DNS search domains                                                                 |                        |
| --env                      | -e      | Set environment variables                                                                     |                        |
| --env-file                 |         | Read in a file of environment variables                                                       |                        |
| --global-permission        |         | Permissions to grant to container's service account for all namespaces                        |                        |
| --group                    |         | The GID to run the entrypoint of the container process                                        |                        |
| --health-cmd               |         | Command to run to check health                                                                |                        |
| --health-failure-threshold |         | Consecutive failures needed to report unhealthy                                               | 0                      |
| --health-header            |         | HTTP Headers to send in GET request for healthcheck                                           |                        |
| --health-initial-delay     |         | Start period for the container to initialize before starting healthchecks ( (ms/s/m/h))       | "0s"                   |
| --health-interval          |         | Time between running the check ( (ms/s/m/h))                                                  | "0s"                   |
| --health-success-threshold |         | Consecutive successes needed to report healthy                                                | 0                      |
| --health-timeout           |         | Maximum time to allow one check to run ( (ms/s/m/h))                                          | "0s"                   |
| --health-url               |         | URL to hit to check health (example: http://:8080/ping)                                       |                        |
| --host-dns                 |         | Use the host level DNS and not the cluster level DNS                                          |                        |
| --hostname                 |         | Container host name                                                                           |                        |
| --image-pull-policy        |         | Behavior determining when to pull the image (never/always/not-present)                        | "not-present"          |
| --image-pull-secrets       |         | Specify image pull secrets                                                                    |                        |
| --interactive              | -i      | Keep STDIN open even if not attached                                                          |                        |
| --label-file               |         | Read in a line delimited file of labels                                                       |                        |
| --label                    | -l      | Set meta data on a container                                                                  |                        |
| --memory                   | -m      | Memory reservation (format: <number>[<unit>], where unit = b, k, m or g)                      |                        |
| --name                     | -n      | Assign a name to the container. Use format [namespace:]name[@version]                         |                        |
| --net                      |         | Set network mode (host)                                                                       |                        |
| --no-mesh                  |         | Disable service mesh                                                                          |                        |
| --permission               |         | Permissions to grant to container's service account in current namespace                      |                        |
| --ports                    | -p      | Publish a container's port(s) (format: svcport:containerport/protocol)                        |                        |
| --privileged               |         | Run container with privilege                                                                  |                        |
| --read-only                |         | Mount the container's root filesystem as read only                                            |                        |
| --rollout-duration         |         | How long the rollout should take                                                              | "0s"                   |
| --request-timeout-seconds  |         | Set request timeout in seconds                                                                | 0                      |
| --scale                    |         | The number of replicas to run or a range for autoscaling (example 1-10)                       |                        |
| --secret                   |         | Secrets to inject to the service (format: name[/key]:target)                                  |                        |
| --stage-only               |         | Only stage service when generating new services. Can only be used when template is true       |                        |
| --template                 |         | If true new version is created per git commit. If false update in-place                       |                        |
| --tty                      | -t      | Allocate a pseudo-TTY                                                                         |                        |
| --user                     | -u      | UID[:GID] Sets the UID used and optionally GID for entrypoint process (format: <uid>[:<gid>]) |                        |
| --volume                   | -v      | Specify volumes for for services                                                              |                        |
| --weight                   |         | Specify the weight for the services                                                           | 0                      |
| --workdir                  | -w      | Working directory inside the container                                                        |                        |

##### Examples

```shell script
```


## rm

Delete resources

##### Usage
```
rio rm [TYPE/]RESOURCE_NAME
```


## scale

Scale a service to a desired number, or set autoscaling params

##### Usage
```
rio scale [SERVICE=NUMBER_OR_MIN-MAX...]
```

##### Examples

```shell script
rio scale foo=5

# autoscaling
rio scale foo=1-5
```


## inspect

Inspect resources

##### Usage
```
rio inspect [TYPE/][NAMESPACE/]SERVICE_NAME
```

##### Options

| flag     | aliases | description                                           | default |
|----------|---------|-------------------------------------------------------|---------|
| --format |         | Edit the raw API object, not the pretty formatted one |         |


##### Examples

```shell script
rio inspect svc@v2

# inspect a build
rio inspect taskrun/affectionate-mirzakhani-mfp5q-ee709-4e40c
```


## edit

Edit resources

##### Usage
```
rio edit [TYPE/]RESOURCE_NAME
```

##### Options

| flag  | aliases | description                                           | default |
|-------|---------|-------------------------------------------------------|---------|
| --raw |         | Edit the raw API object, not the pretty formatted one |         |


##### Examples

```shell script
rio edit demo@v4

rio edit router/myrouter
```

## export

Export a namespace or service

##### Usage
```
rio export [TYPE/]NAMESPACE_OR_SERVICE
```

##### Options

| flag           | aliases | description                                        | default |
|----------------|---------|----------------------------------------------------|---------|
| --format value |         | Specify output format, yaml/json. Defaults to yaml | yaml    |
| --riofile      |         | Export riofile format                              |         |


##### Examples

```shell script
# export a service
rio export demo

# export a namespace in riofile format
rio export --riofile namespace/default
```


## cat

Print the contents of a config

##### Usage
```
rio cat [OPTIONS] [NAME...]
```

##### Options

| flag  | aliases | description             | default |
|-------|---------|-------------------------|---------|
| --key | -k      | The values which to cat |         |

##### Examples

```shell script
# cat a configmap
rio cat configmap/config-foo

# cat a key from a configmap
rio cat --key=a configmap/config-foo
```


## exec

Run a command in a running container

##### Usage
```
rio exec [OPTIONS] CONTAINER COMMAND [ARG...]
```

##### Options

| flag              | aliases  | description                                          | default |
|-------------------|----------|------------------------------------------------------|---------|
| --stdin           | -i       | Pass stdin to the container                          |         |
| --tty             | -t       | Stdin is a TTY                                       |         |
| --container value | -c value | Specify container in pod, default is first container |         |
| --pod value       |          | Specify pod, default is first pod found              |         |

##### Examples

```shell script
# ssh into running container
rio exec -it demo sh

# this is equivalent of doing
rio exec --tty --stdin demo sh

# choose pod and container
rio exec -it --pod mypod --container server demo sh
```

## attach

Attach to a process running in a container

##### Usage
```
rio attach [OPTIONS] CONTAINER
```

##### Options
| flag            | aliases | description                                                  | default |
|-----------------|---------|--------------------------------------------------------------|---------|
| --timeout value |         | Timeout waiting for the container to be created to attach to | 1m      |
| --pod value     |         | Specify pod, default is first pod found                      |         |

##### Examples

```shell script
rio attach demo

rio attach --timeout 30s --pod mydemopod demo
```

## logs

Print logs from services or containers

##### Usage
```
rio logs [OPTIONS] SERVICE/BUILD
```

##### Options

```| flag              | aliases  | description                                                                                       | default |
|-------------------|----------|---------------------------------------------------------------------------------------------------|---------|
| --since value     | -s value | Logs since a certain time, either duration (5s, 2m, 3h) or RFC3339                                | "24h"   |
| --timestamps      | -t       | Print the logs with timestamp                                                                     |         |
| --tail value      | -n value | Number of recent lines to print, -1 for all                                                       | 200     |
| --container value | -c value | Print the logs of a specific container, use -a for system containers                              |         |
| --previous        | -p       | Print the logs for the previous instance of the container in a pod if it exists, excludes running |         |
| --init-containers |          | Include or exclude init containers                                                                |         |
| --all             | -a       | Include hidden or systems logs when logging                                                       |         |
| --no-color        | --nc     | Dont show color when logging                                                                      |         |
| --output value    | -o value | Output format: [default, raw, json]                                                               | default |
```

##### Examples

```shell script
# get logs from a service
rio logs demo

# Get logs from a build
rio build history
rio logs taskrun/affectionate-mirzakhani-mfp5q-ee709-4e40c

# get 1 previous log line for the linkerd-proxy in demo service
rio logs --tail 1 --container linkerd-proxy -a demo

# ignore init-containers and filter to waiting or terminated pods, include timestamps
rio logs --container-state "terminated,waiting" --init-containers=false --timestamps demo

# target terminated pods of all kinds, format as json
rio logs -p -a  --output json demo
```


## install

Install rio management plane

##### Usage
```
rio install [OPTIONS]
```

##### Options

| flag                     | aliases | description                                                                            | default |
|--------------------------|---------|----------------------------------------------------------------------------------------|---------|
| --check                  |         | Only check status, don't deploy controller                                             |         |
| --disable-features value |         | Manually specify features to disable, supports comma separated values                  |         |
| --enable-debug           |         | Enable debug logging in controller                                                     |         |
| --ip-address value       |         | Manually specify IP addresses to generate rdns domain, supports comma separated values |         |
| --yaml                   |         | Only print out k8s yaml manifest                                                       |         |

##### Examples

```shell script
```


## uninstall

Uninstall rio

##### Usage
```
rio uninstall [OPTIONS]
```

##### Options

| flag              | aliases | description                           | default      |
|-------------------|---------|---------------------------------------|--------------|
| --namespace value |         | namespace to install system resources | "rio-system" |

##### Examples

```shell script
rio uninstall

rio uninstall --namespace alt-namespace
```


## Stage

Stage a new revision of a service


##### Usage
```
rio stage [OPTIONS] SERVICE NEW_REVISION
```

##### Options

| flag             | aliases  | description                                        | default |
|------------------|----------|----------------------------------------------------|---------|
| --image value    |          | Runtime image (Docker image/OCI image)             |         |
| --edit           |          | Edit the config to change the spec in new revision |         |
| --env value      | -e value | Set environment variables                          |         |
| --env-file value |          | Read in a file of environment variables            |         |

##### Examples

```shell script
# stage an image (tag v3) to the 2nd version of the demo service
rio stage --image ibuildthecloud/demo:v3 demo v2

# stage the same image with different env variables
rio stage --env abc=xyz demo v2
```


## weight

Set the percentage of traffic to allocate to a given service version. See also promote. 

Defaults to an immediate rollout, set duration to perform a gradual rollout

##### Usage
```
rio weight [OPTIONS] SERVICE_NAME=PERCENTAGE
```

##### Options

| flag       | aliases | description                                  | default |
|------------|---------|----------------------------------------------|---------|
| --duration | none    | How long the rollout should take             | 0s      |
| --pause    | none    | Whether to pause all rollouts on current app | false   |

##### Examples

```shell script
# immediately shift 100% of traffic to app n@v0
rio weight n=100 

# shift n@v2 to 50% of traffic gradually over 5m
rio weight --duration=5m n@v2=50 

# Pause last command at current state, will pause all rollouts on versions in app
rio weight --pause=true n@v2=50
```


## Promote

Send 100% of traffic to an app version and scale down other versions. See also weight. 

##### Usage

```
rio promote [OPTIONS] SERVICE_NAME
```

##### Options

| flag       | aliases | description                                  | default |
|------------|---------|----------------------------------------------|---------|
| --duration | none    | How long the rollout should take             | 0s      |
| --pause    | none    | Whether to pause all rollouts on current app | false   |

##### Examples

```shell script
# promote n@v2 
rio promote n@v2

# promote n@v2 over 1 hour 
rio promote --duration=1h n@v2

# pause last command
rio promote --pause=true n@v2
```


## systemlogs

Print the logs from Rio management plane

##### Usage
```
rio systemlogs
```

## up

Apply a Riofile

##### Usage
```
rio up [OPTIONS]
```

##### Options

| flag                         | aliases                                                                  | description | default |
|------------------------------|--------------------------------------------------------------------------|-------------|---------|
| --name value                 | Set stack name, defaults to current directory name                       |             |         |
| --answers value              | Set answer file                                                          |             |         |
| --file value                 | -f value        Set rio file                                             |             |         |
| --parallel                   | -p                Run builds in parallel                                 |             |         |
| --branch value               | Set branch when pointing stack to git repo (default: "master")           |             |         |
| --revision value             | Use a specific commit hash                                               |             |         |
| --build-webhook-secret value | Set GitHub webhook secret name                                           |             |         |
| --build-clone-secret value   | Set name of secret to use with git clone                                 |             |         |
| --permission value           | Permissions to grant to container's service account in current namespace |             |         |


##### Examples

```shell script
# apply a file named 'Riofile' in current directory
rio up

# apply stack.yaml as a stack named mystack as 2nd revision
rio up --name mystack -f stack.yaml -p

# apply a riofile from git repo, from a specific branch and commit, using a secret, and setup webhook.
rio up --branch branchname --build-webhook-secret=githubtoken --build-clone-secret=mysecret --revision {commit_sha}  https://github.com/exmaple/example

# Set custom permissions to give the stack, and supply answers to riofile questions
rio up  --permissions '* configmaps' --answers answerfile.yaml
```


## dashboard

Open the dashboard in a browser

##### Usage
```
rio dashboard [OPTIONS]
```

##### Options

| flag          | aliases | description          | default |
|---------------|---------|----------------------|---------|
| --reset-admin |         | Reset admin password |         |


##### Examples

```shell script
# reset admin pw
rio dashboard --reset-admin 
```


## kill

Kill pods individually or all pods belonging to a service

##### Usage
```
rio kill [SERVICE_NAME/POD_NAME]
```

##### Examples

```shell script
# kill a service
rio kill demo

# kill individual pods
rio pods # first get pod name
rio kill pod/demo-v042dxp-5fb7d8f677-f9xgn
```


## info

Show system info

##### Usage
```
rio info
```