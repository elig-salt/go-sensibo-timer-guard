# Sensibo Timer Guard #


The Sensibo ecosystem lets you create an API key which you can use to query data / send changes to your devices.

This image uses an API key and an amount of minutes (default is 60) to set a timer for any devices currently running without one.

Schedule this image to run periodically (say, every 5 minutes) to make sure your devices will turn off after the specified time.

## Supported environment variables:

An API key (30 characters) - get this from the Sensibo device control web site (under "Create an API Key")

This is <b>required</b> !

>SENSIBO_API_KEY=AAAAbbbbCCCCdddd11112222333344

A timer schedule period (in minutes) - default is 60

This is optional

>SENSIBO_TIMER_SCHEDULE_MINS=60

## Docker

The Docker image is available at:

> https://hub.docker.com/repository/docker/slygon/sensibo-timer-guard

### Pulling the image

```sh
docker pull slygon/sensibo-timer-guard
```
