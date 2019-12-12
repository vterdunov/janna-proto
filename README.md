# Janna API Protocol buffers

This repo contains protocol buffer service definitions and generated interfaces for Janna.

## Workflow
- make changes in proto files
- generate code `make proto`. Need Docker installed.
- commit new changes `git add . && git commit -m '<Commit description>'`
- push the changes `git push`

Please remember about backward compatibility. Heres is a [article](https://developers.google.com/protocol-buffers/docs/proto3#updating)

## Repository structure

### `/gen`
This folder contains generated code from `.proto` files. Currently generates files for:
- go
- swagger

### `/proto`
Protocol buffer service definitions. To generate code run `make proto`. Also, this directory contains config file for [uber/protool](https://github.com/uber/prototool) - `prototool.yaml`.

### `/third_party`
Third party resources/definition. Need to generate code.
