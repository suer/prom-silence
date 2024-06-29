# amp-silence

This is a simple CLI to add/delete silences in AMP(Amazon Managed Service for Prometheus) Alertmanager.

## Usage

### Add silence

```bash
$ echo '{
    "startsAt": "2024-06-30T15:00:00.000Z",
    "endsAt": "2024-06-30T23:59:59.000Z",
    "comment": "Maintenance",
    "createdBy": "suer",
    "matchers": [        {
            "name": "host",
            "value": "www.example.com",
            "isEqual": true,
            "isRegex": false
        }
    ]
}' | ./amp-silence add -e https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/

{"silenceID": "yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy"}
```

### Delete silence

```bash
$ ./amp-silence delete -e https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/ -s yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy
```

### List silences

```bash
$ ./amp-silence list -e https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/
[{"comment":"Maintenance","createdBy":"suer","endsAt":"2024-06-30T23:59:59.000Z","id":"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx","matchers":[{"isEqual":true,"isRegex":false,"name":"host","value":"www.example.com"}],"startsAt":"2024-06-30T15:00:00.000Z","status":{"state":"pending"},"updatedAt":"2024-06-29T10:57:20.518Z"}]
```

## Advanced Usage

### Query with JMESPath

```bash
$ ./amp-silence list -e https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/ -q "[0].id"
"yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy"
```

### Raw output

```bash
$ ./amp-silence list -e https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/ -q "[0].id" -r 
yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy
```

## Build

```bash
$ go build
```
