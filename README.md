# 🌊 cli-flow

Modular CLI framework for automating local and cloud-based workflows.  
Written in Go. Adding services is super easy, made it as customisable as possible 

---

## ✨ Features

- 🔧 **Modular Services**: Easily add custom runners like `s3`, `lambda`, `http.fetch`, or `shell`.
-  ☁️ **AWS Integration**: Run cloud tasks using injected AWS credentials from your local config.
- 📦 **Single Entrypoint**: Chain services in a JSON-based pipeline (coming soon).
- 🔐 **Secure Credential Handling**: Load AWS creds from `~/.cli_flow/config.json` and clean them up automatically.

---

## ⚙️ Getting Started

### 1. Build

```
go build -o cli-flow
```

### 2. Configure AWS

```
./cli-flow configure
```

Stores your AWS credentials in:

`~/.cli_flow/config.json`

Example structure:

```
{
  "access_key": "AKIA...",
  "secret_key": "abc123...",
  "region": "us-west-2"
}
```

### 3. Run a Service

```
./cli-flow run s3 --config '{"action":"upload","bucket":"my-bucket","key":"path/file.txt","path":"./file.txt"}'
```

Supports `upload` and `download` actions using the AWS CLI.

---

## 🧱 Architecture

- `internal/structures`: Shared types (like `Step`, `AWSConfig`)
- `internal/services`: All registered services (`s3`, `shell`, `http.fetch`, etc.)
- `internal/config`: AWS credential loader and env var manager
- `cmd/`: Cobra CLI commands (`configure`, `run`, etc.)

---

## 📦 Example Service Interface

```
type Service interface {
  Run(step structures.Step) error
  Name() string
  ConfigSpec() []string
}
```

Each service registers itself to the global `Registry` via `init()`.

---

## 🧼 Environment Handling

- `configure` saves AWS creds locally (not globally)
- `runflow` commands load and inject credentials before executing
- Credentials are unset after execution using `defer`

---

## 🚧 Roadmap

- [ ] Add inline scripting (`code.runner`)
- [ ] Auto-doc each service with `ConfigSpec()`
- [ ] Add logger helper to have uniform formatting

---

## 🧠 Philosophy

CLI tooling should be dead simple to run and stupid easy to extend.

---

## 🧊 License

MIT. Feel free to use and contribute
