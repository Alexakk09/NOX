```md
# NOX

> A modular personal AI assistant built around LLMs, persistent memory, web search, document intelligence, and extensible tools.

NOX is an AI assistant designed to go beyond a traditional chatbot. It combines conversational AI with memory, web search, PDF intelligence, tool calling, and a Jarvis-style system for interacting with the local environment.

The project is actively evolving toward a more capable personal AI system where the LLM can **reason, remember, retrieve information, use tools, and perform controlled actions**.

---

## ✨ Features

- 🧠 **Conversational AI** — multi-turn conversations powered by an LLM
- 💾 **Persistent Memory** — retain useful information across conversations
- 🌐 **Web Search** — search the web through SearXNG
- 📄 **PDF Intelligence** — upload, summarize, explain, and interact with documents
- 🛠️ **Tool Calling** — allow the AI to use external capabilities
- 🤖 **Jarvis System** — local voice and computer-interaction capabilities
- 🔐 **Authentication** — API authentication and protected resources
- 🗄️ **PostgreSQL** — persistent application data
- 🐳 **Docker** — containerized SearXNG deployment
- ⚡ **React + TypeScript** — web interface
- 🔵 **Go** — backend/API and AI infrastructure

---

# 🏗️ Architecture

```text
                         ┌─────────────────┐
                         │      User       │
                         └────────┬────────┘
                                  │
                                  ▼
                         ┌─────────────────┐
                         │       Web       │
                         │ React + TS/Vite │
                         └────────┬────────┘
                                  │
                                  ▼
                         ┌─────────────────┐
                         │       API       │
                         │       Go        │
                         └───────┬─────────┘
                                 │
                ┌────────────────┼────────────────┐
                │                │                │
                ▼                ▼                ▼
          ┌───────────┐    ┌───────────┐   ┌───────────┐
          │ PostgreSQL│    │ AI / LLM  │   │  SearXNG  │
          │           │    │           │   │   Search  │
          └───────────┘    └─────┬─────┘   └───────────┘
                                 │
                    ┌────────────┼────────────┐
                    │            │            │
                    ▼            ▼            ▼
                 Memory        Tools        Jarvis
```

---

# 📁 Project Structure

```text
NOX/
│
├── AI/
│   ├── brain/
│   ├── config/
│   ├── memory/
│   └── tools/
│
├── API/
│   ├── cmd/
│   └── internal/
│
├── Jarvis/
│   ├── brain/
│   └── tools/
│
├── Web/
│   └── src/
│
├── searxng/
│   └── config/
│       └── settings.yml
│
└── docker-compose.yml
```

### AI

Contains the AI/agent layer responsible for model interaction, memory, reasoning, and tool execution.

### API

The Go backend responsible for authentication, conversations, persistence, PDF functionality, and communication between the frontend and AI system.

### Jarvis

The local assistant layer responsible for voice and computer-related capabilities.

### Web

The React + TypeScript frontend.

### SearXNG

Provides web-search functionality to NOX.

---

# 🧰 Tech Stack

| Component | Technology |
|---|---|
| Backend | Go |
| Frontend | React |
| Frontend Language | TypeScript |
| Database | PostgreSQL |
| Search | SearXNG |
| AI | LLM / LM Studio |
| Voice | whisper.cpp |
| Containerization | Docker / Docker Compose |
| Authentication | JWT |
| Build Tool | Vite |

---

# 🚀 Getting Started

## Prerequisites

Install the following:

- Git
- Go
- Node.js
- npm
- PostgreSQL
- Docker Desktop
- LM Studio
- whisper.cpp

Check the main installations:

```powershell
git --version
go version
node --version
npm --version
docker --version
docker compose version
```

---

# 1. Clone the Repository

```powershell
git clone https://github.com/Alexakk09/NOX.git
```

Enter the project:

```powershell
cd NOX
```

---

# 2. Configure the Environment

NOX requires environment configuration for services such as PostgreSQL and the LLM.

Create the required `.env` files based on the configuration expected by the individual components.

Example:

```text
API/
└── .env
```

Configure your local database and AI-related settings there.

> **Never commit real API keys, passwords, database credentials, JWT secrets, or other sensitive values to GitHub.**

---

# 3. Start SearXNG

NOX uses SearXNG for web search.

From the project root:

```powershell
docker compose up -d
```

Check the container:

```powershell
docker compose ps
```

SearXNG is exposed at:

```text
http://localhost:8888
```

The Docker configuration maps:

```text
localhost:8888 → SearXNG:8080
```

Open it in your browser:

```text
http://localhost:8888
```

To stop SearXNG:

```powershell
docker compose down
```

---

# 4. Start the API

Open a new terminal.

Move into the API:

```powershell
cd NOX\API
```

Install/update Go dependencies:

```powershell
go mod tidy
```

Start the server:

```powershell
go run .\cmd\server
```

The API runs locally on:

```text
http://localhost:8080
```

Keep this terminal running.

---

# 5. Start the AI Service

Open another terminal:

```powershell
cd NOX\AI
```

Install/update dependencies:

```powershell
go mod tidy
```

Start the AI service:

```powershell
go run .
```

The AI service requires the configured LLM environment to be available.

---

# 6. Start the Web Interface

Open another terminal:

```powershell
cd NOX\Web
```

Install frontend dependencies:

```powershell
npm install
```

Start the development server:

```powershell
npm run dev
```

Vite will provide a local development URL, normally:

```text
http://localhost:5173
```

Open that URL in your browser.

---

# ▶️ Running Everything

During development, NOX currently runs its components separately.

### Terminal 1 — SearXNG

```powershell
cd NOX
docker compose up -d
```

### Terminal 2 — API

```powershell
cd NOX\API
go run .\cmd\server
```

### Terminal 3 — AI

```powershell
cd NOX\AI
go run .
```

### Terminal 4 — Web

```powershell
cd NOX\Web
npm run dev
```

The resulting local system looks like:

```text
Browser
   │
   ▼
Web :5173
   │
   ▼
API :8080
   │
   ├──────────► PostgreSQL
   │
   └──────────► AI / LLM
                    │
                    ├── Memory
                    ├── Tools
                    ├── PDF processing
                    └── SearXNG :8888
```

---

# 🌐 Web Search

NOX uses SearXNG as its web-search provider.

SearXNG configuration:

```text
searxng/
└── config/
    └── settings.yml
```

JSON responses are enabled for programmatic search:

```yaml
search:
  formats:
    - html
    - json
```

NOX communicates with SearXNG through:

```text
http://localhost:8888
```

---

# 🧠 Memory

NOX includes a memory system designed to give the assistant persistent context beyond a single conversation.

The goal is to eventually support:

- Long-term memory
- Temporary memory
- Memory scopes
- Memory confidence
- Contradiction detection
- Memory timeline
- User-controlled memory

---

# 🛠️ Tool System

NOX uses a tool-based architecture to allow the AI to interact with capabilities outside the LLM itself.

The general flow is:

```text
User Request
     │
     ▼
    LLM
     │
     ├── Normal response
     │
     └── Tool required
              │
              ▼
          Tool Registry
              │
              ▼
          Tool Execution
              │
              ▼
          Tool Result
              │
              ▼
             LLM
              │
              ▼
           Response
```

This makes it possible to add new capabilities without changing the fundamental conversation system.

---

# 📄 PDF Intelligence

NOX includes PDF-related AI functionality.

Current capabilities include:

- PDF upload
- PDF processing
- PDF summarization
- PDF explanation
- AI interaction with document content

The document system is intended to evolve toward richer retrieval and multi-document workflows.

---

# 🤖 Jarvis

The Jarvis component extends NOX toward a local personal-assistant experience.

It is designed to support capabilities such as:

- Voice interaction
- Speech recognition
- Local application interaction
- Computer tools
- Extensible actions

Voice functionality uses `whisper.cpp`.

---

# 🔐 Security

NOX uses authentication and protected API resources.

Because NOX is intended to eventually perform real-world actions, tool execution is designed with future permission and approval mechanisms in mind.

Potential consequential actions include:

- Sending messages
- Sending emails
- Modifying files
- Running system commands
- Interacting with applications
- Performing external actions

The long-term goal is to separate **reasoning** from **permission to act**.

---

# 🗺️ Roadmap

## AI

- [x] Conversational AI
- [x] Local LLM integration
- [x] Tool architecture
- [x] Memory system
- [ ] Advanced memory controls
- [ ] Memory timeline
- [ ] Memory confidence
- [ ] Contradiction detection
- [ ] Temporary memory
- [ ] Project-scoped memory

## Web Intelligence

- [x] SearXNG integration
- [x] Web search
- [ ] Source reliability
- [ ] Claim/evidence tracking
- [ ] Multi-step research
- [ ] Research memory

## Documents

- [x] PDF processing
- [x] PDF summarization
- [x] PDF explanation
- [ ] Multiple file support
- [ ] Drag-and-drop files
- [ ] Improved document retrieval
- [ ] Image understanding

## Agents

- [x] Tool registry
- [x] Tool execution
- [ ] Permission system
- [ ] Action approval
- [ ] Agent checkpoints
- [ ] Undoable actions
- [ ] Agent activity timeline
- [ ] More external integrations

## Voice / Jarvis

- [x] Voice infrastructure
- [x] Local tools
- [ ] Improved voice interaction
- [ ] Expanded computer control
- [ ] Safer action execution

---

# 🚧 Development Status

NOX is currently under active development.

The architecture and APIs may change as new capabilities are introduced.

The project is being developed incrementally with the long-term goal of building a modular personal AI system rather than a simple chatbot interface.

---

# 🎯 Vision

Most AI applications stop at:

```text
Prompt → LLM → Response
```

NOX is being built toward:

```text
                    ┌───────────┐
                    │    User   │
                    └─────┬─────┘
                          │
                          ▼
                    ┌───────────┐
                    │    NOX    │
                    └─────┬─────┘
                          │
          ┌───────────────┼────────────────┐
          │               │                │
          ▼               ▼                ▼
       Remember        Research          Reason
          │               │                │
          └───────────────┼────────────────┘
                          │
                          ▼
                        Tools
                          │
                          ▼
                     Take Action
```

The goal is to build an assistant that can **understand, remember, research, reason, use tools, and eventually act — while keeping those actions controlled by the user.**

---
