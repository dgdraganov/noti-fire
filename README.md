# noti-fire

A simple notification system that exposes a single endpoint `/notify` in order to receive messages that will be dispatched to different channels. The system consists of two services - `server` and `consumer` that communicate through `kafka` messages.

## Diagram of the solution?

Here is an ascii diagram of the system. It shows its components and how it can be horizotally scaled:


           ┌────────────────────────┐           ┌──────────────────────────┐
           │                        │           │                          │
           │ noti-fire web server 1 │    ┌─────►│    noti-fire consumer 1  │
           │                        │    │      │                          │
           └──────────┬─────────────┘    │      └──────────────────────────┘
                      │                  │
                      │                  │
                ┌─────▼─────┐            │
                │           ├────────────┘      ┌──────────────────────────┐
                │           │                   │                          │
   ...  ────────►   kafka   ├──────────────────►│    noti-fire consumer 2  │
                │           │                   │                          │
                │           ├────────────┐      └──────────────────────────┘
                └─────▲─────┘            │
                      │                  │
                      │                  └─────► ...
           ┌──────────┴─────────────┐
           │                        │
           │ noti-fire web server 2 │
           │                        │
           └────────────────────────┘


## How to run?

The project is equipped with `docker-compose.yaml` file together with all the needed configurations in `dev.env` in order to be started within a docker environmen.

The following command will run the required services - `server`, `consumer` and `kafka` with a single broker:

```
    make compose
```

The `server` will be expecting requests on `localhost:9205`

When finished one can use the below command in order to stop the running containers: 

```
    make decompose
```

## How to use? 

The `/notify` endpoint expects messages in the following format:

```
{ "message": "this is an important message" }
```

## What about tests?

Tests can be run with the following command:

```
    make tests
```

