Stayin' alive :bear:
---

Goals :

- Execute atomic actions and reporting.
- Stop share your credentials by Slack or Mail.

Promise : 

- Fault tolerance
- Secure

Start (web mode) : 
```
sta web [-p $port]
connection ...
```

Start (local mode) : 
```
sta local $url
connection ...
```

Cli 
```
sta $command
```

Provide your conf repository.

```
sta git add $gitname $repo_url_git -k $path_public_ssh_key
```

Set en env variables. Env variables are share in a bulk (AES encryption) on all servers.
And can be used on any files in repository.
``` 
sta env "$name=$value" --git $gitname
```

The repository is the action configuration. For any push on `master` the repo is reloaded on any servers.

```
sta.yaml
part1/
    part1.yaml
    ...
    action.sh
    action.php
    action.rb
    
    report.sh
    report.php
    report.rb
    ...
part2/
    part2.yaml
    ...
```

Configuration file `sta.yaml` (root)

```yaml
servers:
    sta: https://sta.org/?key=$sta_key (url of sta web server)
    web_perso_1:  https://perso.../?key=$sta_key
    web_perso_2:  https://perso.../?key=$sta_key
    local_1: null (null set a local server)
    local_2: null
parts:
    part1: ./part1/ (path to part 1)
```

Configuration file `part1.yaml` (part1)

```yaml
actions:
    action1:
        file: action.sh
        schedule: 00h:00m
        run_on : [local_1,perso_1]
report:
    report1:
        file: report.sh
        run_after_all :
          - local_1.action1:0 ($servers.$action.$exit_code)
```

### Architecture :

Consense example : Sessions: https://www.consul.io/docs/internals/sessions.html


Sync data between 3 server (web mode)
```
server : A,B(run),C

event($data)->A
A->getLock(B): wait i run ($action)
A->getLock(C): ok ($time)
A(wait B run $action for event): retry (5sec)
A->getLock(B): ok ($time)
A->getLock(C): ok ($time)
A->createSession($time(A),$time(B),$time(C)): ($sessionId)
it A->sendEvent($sessionId,$data,C): ok
Check matrix: event
   A B C
rA x o o
rB o x 
rC o   x
B->check($sessionId,$data,C): ok
Check matrix: event
   A B C
rA x o o
rB o x o
rC o o x
C->done($sessionId,B): ok
C->done($sessionId,A): ok
```

Sync data between 4 server (web mode)
```

```
