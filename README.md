# SMS Manager

A simple daemon for SMS redirection from modems(Huawei E3372) to Telegram

## Overview

SMS Manager is a Go application that:

- **Monitors SMS messages** from modems
- **Redirects SMS to Telegram** through a bot
- **Stores messages** in SQLite database

## Key Features

- 🔄 **Automatic SMS checking** with configurable intervals
- 📱 **Multi-user support** with different modems
- 💬 **Telegram integration** for notifications
- 💾 **Local message storage** in SQLite
- 📊 **Prometheus metrics** for monitoring
- 🐧 **Ready systemd service** for Linux
- 📦 **Debian package** for easy installation
- 🏥 **Health checks** for Kubernetes/containers

## Installation

### From Source

```bash
# Clone repository
git clone https://github.com/telepuz/smsmanager.git
cd smsmanager

# Build and install
make install

# Or create Debian package
make deb-package
```

## Configuration

### Main Configuration File

Default location: `/etc/smsmanager/smsmanager.yml`

```yaml
---
messenger:
  token: YOUR_BOT_TOKEN                 # Telegram bot token

# User list
users:
- name: Scooby-Doo                      # User name
  chat_id: 123456789                    # Telegram chat ID
  modem_url: 192.168.8.1                # Modem IP address
- name: Shaggy
  chat_id: 123456789
  modem_url: 192.168.9.1
```

### Configuration Parameters

#### Main Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `check_interval` | duration | `1m` | SMS check interval |
| `users` | array | - | User list |

#### Users (`users`)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `name` | string | - | User name |
| `chat_id` | int64 | - | Telegram chat ID |
| `modem_url` | string | - | Modem IP address |
| `modem_type` | string | `huaweie3372` | Modem type |

#### Logging (`logger`)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `format` | string | `plaintext` | Log format (`plaintext`, `json`) |
| `level` | string | `info` | Log level (`debug`, `info`, `warn`, `error`) |

#### Messenger (`messenger`)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `type` | string | `telegram` | Messenger type (`telegram`, `stdout`) |
| `token` | string | - | Telegram bot token |

#### Storage (`storage`)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `type` | string | `sqlite3` | Storage type (`sqlite3`, `stdout`) |
| `file_path` | string | `/var/lib/smsmanager/db.sql` | Database file path |

#### Exporter (`exporter`)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `type` | string | `none` | Exporter type (`prom`, `none`) |
| `listen_port` | int | - | Metrics port |
| `metrics_path` | string | - | Metrics endpoint path |

#### Health Check (`health_check`)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `enable` | bool | `false` | Enable health checks |
| `listen_port` | int | `3000` | Health check port |

## Running the Application

### Development

```bash
# Run in development mode
make run
```

### Production

```bash
# Run as systemd service
sudo systemctl enable smsmanager
sudo systemctl start smsmanager
sudo systemctl status smsmanager

# View logs
sudo journalctl -u smsmanager -f
```

### Manual Run

```bash
# Run with custom config
./smsmanager -config_file=/path/to/config.yml
```

## Telegram Bot Setup

1. Create a bot via [@BotFather](https://t.me/BotFather)
2. Get bot token
3. Add token to configuration
4. Get your Chat ID (use [@userinfobot](https://t.me/userinfobot))
5. Add Chat ID to configuration

## Huawei E3372 Modem Setup

### Udev Rules

Thanks [this thread](https://forums.raspberrypi.com/viewtopic.php?t=18996)

Create `/etc/udev/rules.d/10-HuaweiFlashCard.rules`:

```
SUBSYSTEMS=="usb", ATTRS{modalias}=="usb:v12D1p1F01*", SYMLINK+="hwcdrom", RUN+="/usr/bin/sg_raw /dev/hwcdrom 11 06 20 00 00 00 00 00 01 00"
```

### Dependencies

```bash
apt-get install sg3-utils
udevadm control --reload-rules
```

### Change network config

Thanks [this repo](https://github.com/tnbhaeufi/Huawei-E3372h-320-IP-Range-change)

## Monitoring

### Prometheus Metrics

Available at: `http://localhost:2112/metrics`

| Metric | Type | Description |
|--------|------|-------------|
| `message_receive_count` | Counter | Total received SMS count |
| `message_send_count` | Counter | Total sent message count |
| `database_messages_count` | Gauge | Messages in database count |
| `message_error_receive_count` | Counter | SMS receive error count |
| `message_error_send_count` | Counter | Message send error count |
| `message_error_database_count` | Counter | Database error count |

### Health Checks

- **Liveness**: `http://localhost:3000/livez`
- **Readiness**: `http://localhost:3000/readyz`

## Database

Application uses SQLite for message storage. Table structure:

```sql
CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    message_id INT64,
    chat_id INT64,
    phone TEXT,
    content TEXT,
    dt DATETIME DEFAULT current_timestamp
);
```

## Building with Make

### Available Make Commands

```bash
make help            # Show help
make run             # Run application
make build           # Build application
make deb-package     # Create Debian package
make install         # Install application
make uninstall       # Uninstall application
make clean           # Clean temporary files
make linter-golangci # Run code linter
```

### Build Process

```bash
# Build application
make build

# Create Debian package
make deb-package
# Result: smsmanager_0.0.1_amd64.deb

# Install system-wide
make install
```

### Build Details

The `make build` command:
- Creates build directory structure
- Compiles Go application with CGO disabled
- Copies configuration files
- Copies systemd service file

The `make deb-package` command:
- Runs build first
- Creates Debian package structure
- Generates .deb file

## Logging

Application supports various log levels:

- `debug` - Detailed debug information
- `info` - General operation information
- `warn` - Warnings
- `error` - Errors only

Log formats:
- `plaintext` - Text format
- `json` - Structured JSON logging

## Usage Examples

### Simple Single User Configuration

```yaml
---
check_interval: 30s
messenger:
  type: telegram
  token: "1234567890:ABCdefGHIjklMNOpqrsTUVwxyz"
storage:
  type: sqlite3
  file_path: /var/lib/smsmanager/messages.db
logger:
  level: info
users:
- name: Admin
  chat_id: 123456789
  modem_url: 192.168.1.100
```

### Monitoring Configuration

```yaml
---
check_interval: 1m
exporter:
  type: prom
  listen_port: 2112
health_check:
  enable: true
  listen_port: 3000
messenger:
  type: telegram
  token: "YOUR_BOT_TOKEN"
storage:
  type: sqlite3
  file_path: /var/lib/smsmanager/messages.db
logger:
  level: info
  format: json
users:
- name: User1
  chat_id: 123456789
  modem_url: 192.168.8.1
- name: User2
  chat_id: 987654321
  modem_url: 192.168.8.2
```

## Troubleshooting

### Common Issues

1. **Modem Unreachable**
   - Check modem IP address
   - Ensure modem is connected to network

2. **Telegram Bot Not Sending**
   - Verify bot token
   - Ensure bot is added to chat
   - Check Chat ID

3. **Database Errors**
   - Check database file permissions
   - Ensure directory exists

### Viewing Logs

```bash
# Systemd logs
sudo journalctl -u smsmanager -f

# Application logs (if run manually)
./smsmanager -config_file=config.yml 2>&1 | tee smsmanager.log
```

## License

Project is distributed under the MIT license.

## Support

- GitHub Issues: [https://github.com/telepuz/smsmanager/issues](https://github.com/telepuz/smsmanager/issues)
- Documentation: [https://github.com/telepuz/smsmanager](https://github.com/telepuz/smsmanager)
