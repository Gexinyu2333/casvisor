<h1 align="center" style="border-bottom: none;">📦⚡️ Casvisor</h1>
<h3 align="center">An open-source cloud operating system management platform developed by Go and React.</h3>
<p align="center">
  <a href="#badge">
    <img alt="semantic-release" src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
  </a>
  <a href="https://hub.docker.com/r/casbin/casvisor">
    <img alt="docker pull casbin/casvisor" src="https://img.shields.io/docker/pulls/casbin/casvisor.svg">
  </a>
  <a href="https://github.com/casvisor/casvisor/releases/latest">
    <img alt="GitHub Release" src="https://img.shields.io/github/v/release/casvisor/casvisor.svg">
  </a>
  <a href="https://hub.docker.com/r/casbin/casvisor">
    <img alt="Docker Image Version (latest semver)" src="https://img.shields.io/badge/Docker%20Hub-latest-brightgreen">
  </a>
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/casvisor/casvisor">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/casvisor/casvisor?style=flat-square">
  </a>
  <a href="https://github.com/casvisor/casvisor/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/casvisor/casvisor?style=flat-square" alt="license">
  </a>
  <a href="https://github.com/casvisor/casvisor/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/casvisor/casvisor?style=flat-square">
  </a>
  <a href="#">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/casvisor/casvisor?style=flat-square">
  </a>
  <a href="https://github.com/casvisor/casvisor/network">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/casvisor/casvisor?style=flat-square">
  </a>
</p>

![image](https://github.com/casvisor/casvisor/assets/3787410/c71e9a09-38be-4f76-99a8-595aa859ee58)

## Online demo

- Read-only site: https://door.casvisor.com (any modification operation will fail)
- Writable site: https://demo.casvisor.com (original data will be restored for every 5 minutes)

## Documentation

https://casvisor.org

## Architecture

Casvisor contains 2 parts:

| Name     | Description                      | Language               | Source code                                          |
|----------|----------------------------------|------------------------|------------------------------------------------------|
| Frontend | Web frontend UI for Casvisor     | Javascript + React     | https://github.com/casvisor/casvisor/tree/master/web |
| Backend  | RESTful API backend for Casvisor | Golang + Beego + MySQL | https://github.com/casvisor/casvisor                 |

## Installation

Casvisor uses Casdoor as the authentication system. So you need to create an organization and an application for Casvisor in a Casdoor instance.

### Necessary configuration

#### Get the code

```shell
go get github.com/casdoor/casdoor
go get github.com/casvisor/casvisor
```

or

```shell
git clone https://github.com/casdoor/casdoor
git clone https://github.com/casvisor/casvisor
```

#### Setup database

Casvisor will store its users, nodes and topics information in a MySQL database named: `casvisor`, will create it if not existed. The DB connection string can be specified at: https://github.com/casvisor/casvisor/blob/master/conf/app.conf

```ini
dataSourceName = root:123@tcp(localhost:3306)/
```

Casvisor uses XORM to connect to DB, so all DBs supported by XORM can also be used.

#### Configure Casdoor

After creating an organization and an application for Casvisor in a Casdoor, you need to update `clientID`, `clientSecret`, `casdoorOrganization` and `casdoorApplication` in app.conf.

#### Run Casvisor

- Configure and run Casvisor by yourself. If you want to learn more about casvisor.
- Open browser: http://localhost:16001/

### Optional configuration

#### Setup your Casvisor to enable some third-party login platform

  Casvisor uses Casdoor to manage members. If you want to log in with oauth, you should see [casdoor oauth configuration](https://casdoor.org/docs/provider/oauth/overview).

#### OSS, Email, and SMS

  Casvisor uses Casdoor to upload files to cloud storage, send Emails and send SMSs. See Casdoor for more details.

#### RDP

Run guacd for RDP connection.

```shell
docker run --name some-guacd -d -p 4822:4822 guacamole/guacd
```

## Contribute

For Casvisor, if you have any questions, you can give Issues, or you can also directly start Pull Requests(but we recommend giving issues first to communicate with the community).

## License

[Apache-2.0](LICENSE)
