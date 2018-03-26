:warning: **This has been made public but it is still very dirty and hacky. I promise I'll tidy it up soon**

# fdctl
Handy Go app for managing firedrills within Core

## What is the point of this?

Some context around firedrills

## Before you start

1. Choose the channel you will use for firedrills
1. Get an incoming webhook set up in Slack
1. Set webhook URL as env variable `FIREDRILL_WEBHOOK_URL`

## Install

Get the latest [release](https://github.com/skybet/fdctl/releases) and use it

## Build

1. `git clone git@github.com:skybet/fdctl.git`
1. `cd fdctl && go install fdctl`

## Usage

### Start

`fdctl -operation start`

### Stop

`fdctl -operation stop`

### Say

`fdctl -role techops -message "Hello, World!"`

## Todo

Pagerduty integration
