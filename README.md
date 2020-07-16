# Hello World (tutorial)
Hello World is a tutorial for building controllers using `dbl`. Please follow along step-by-step using the project branches.

## What is a controller?
A controller is a self-contained component which manages it's own data store and logic. The controller provides access methods, so that services utilizing it can quickly and easily access the required data (and edit it if need be). The back-end of the controller can be any DB, some examples of potential controller back-ends are:
- BoltDB
- BadgerDB
- LMDB
- SQLite
- MySQL
- Postgres
- MariaDB

In the case of our example, we will be utilizing BoltDB.

## Why would I want to utilize a controller?
Controllers provide a fantastic and scalable pattern for applications big and small. Keeping logic and data constrained to controllers enforces good separation of logic and flow. In addition, a well modularized system can easily be scaled horizontally in a micro-service architecture if and when needed.