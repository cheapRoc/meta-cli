# meta-cli

`meta-cli` is an example tool that helps manage Metadata on instances running in
Joyent's Triton cloud.

## Build

`go build`

## Usage

You'll need to load a Triton profile into your shell environ.

```sh
$ eval "$(triton env us-west-1)"
```

List metadata across multiple instances.

```sh
$ ./meta-cli -name hello-world

hello-world-app-0: root_authorized_keys = ...
hello-world-app-0: something-two = true
hello-world-app-0: author-name = John Tester

hello-world-app-1: root_authorized_keys = ...
hello-world-app-1: author-name = John Tester
```

Add/update a single metadata key/value pair across multiple instances.

```sh
$ ./meta-cli -name hello-world -key data-role -val test-box
hello-world-app-0: data-role = test-box
hello-world-app-1: data-role = test-box
```

Add/update multiple key/value pairs across multiple instances.

```sh
$ ./meta-cli -name hello-world -data data-role=initdb -data data-job=batch
hello-world-app-0: data-role = initdb
hello-world-app-0: data-job = batch
hello-world-app-1: data-role = initdb
hello-world-app-1: data-job = batch
```

Get a single metadata key's value across multiple instances (quoted value).

```sh
$ ./meta-cli -name hello-world -key data-role
hello-world-app-0: data-role = "initdb"
hello-world-app-1: data-role = "initdb"
```

Delete a single metadata key/value pair across multiple instances.

```sh
$ ./meta-cli -name hello-world -key data-role -delete
hello-world-app-0: Removed data-role
hello-world-app-1: Removed data-role
```

Delete all metadata pairs across multiple instances.

```sh
$ ./meta-cli -name hello-world -delete -all
hello-world-app-0: Cleared all metadata
hello-world-app-1: Cleared all metadata
```

