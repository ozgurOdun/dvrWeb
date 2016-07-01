# Usage of dvrWebServer Rest Service

## Deleting a server
Make a **"DELETE"** call to 8088 port with the link [/dvr/:name/delete] [linkDel]
* :name represents the name of the desired entry

## Updating server status
Make a **"POST"** call to 8088 port with the link [/dvr/:name/:newstatus/update] [linkPost]
* :name represents the name ofthe desired server
* :newstatus represents the desired status to attend to server
* "alive" and "dead" are avalilable as status

## Getting server list
Make a **"GET"** call to 8088 port with the link [/dvr/:query/query] [linkGet]
* :query can be "all" and "alive"

## Adding new server 
Make a **"PUT"** call to 8088 port with the link [/dvr/:name/:ipstring/:version/:status/add] [linkPut]
* :name is the given name to server
* :version represents the version number running on that server
* :status represents the current status of the server


[linkDel]: <http://127.0.0.1:8088/dvr/:name/delete>
[linkPost]: <http://127.0.0.1:8088/dvr/:name/:newstatus/update>
[linkGet]: <http://127.0.0.1:8088/dvr/:query/query>
[linkPut]: <http://127.0.0.1:8088/dvr/:name/:ipstring/:version/:status/add>