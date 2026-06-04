# cwc MCP Usage Guide

Use `list_cwc_commands` or `get_cwc_command_help` to discover commands.
Prefer dynamic tools named `cwc_<command_path>` for direct command execution.

## Natural Language To CLI Flags

- `instance`, `instance id`, `instance name`, `machine`, `machine id`, `machine name` -> `-i <instance-id>`
- `project`, `project id`, `project name` -> `-p <project-id>`
- `bucket`, `bucket id`, `bucket name` -> `-b <bucket-id>`
- `registry`, `registry id`, `registry name` -> `-r <registry-id>`
- `adapter`, `adapter id`, `adapter name` -> `-a <adapter-id>`
- `device`, `device id`, `device name` -> `-d <device-id>`
- `objectType`, `object type id`, `object type name` -> `-o <object-type-id>`
- `storage`, `storage key`, `storage id`, `storage name`, `kv key`, `kv id` -> `-k <storage-id>`

## Examples

- `redémarre la machine 553` -> `cwc instance reboot -i 553`
- `redémarre l'instance 553` -> `cwc instance reboot -i 553`
