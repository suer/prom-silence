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
}' | ./amp-silence -e https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/ -a

{"silenceID": "yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy"}
```

### Delete silence

```bash
$ ./amp-silence -e https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/ -d -s yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy
```

### Query with JMESPath

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
}' | ./amp-silence -e https://aps-workspaces.ap-northeast-1.amazonaws.com/workspaces/ws-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/ -a -q 'silenceID'

yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy
```

## Build

```bash
$ go build
```
