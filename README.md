# NOX

> A modular personal AI assistant built around LLMs, persistent memory, web search, document intelligence, and extensible tools.

NOX is an AI assistant designed to go beyond a traditional chatbot. It combines conversational AI with memory, web search, PDF intelligence, tool calling, and local AI capabilities.

The project is actively evolving toward a more capable personal AI system where the LLM can **reason, remember, retrieve information, use tools, and perform controlled actions**.

---

## ✨ Features

- 🧠 **Conversational AI** — multi-turn conversations powered by a local LLM
- 💾 **Persistent Memory** — retain useful information across conversations
- 🌐 **Web Search** — search the web through SearXNG
- 📄 **PDF Intelligence** — upload, process, summarize, and interact with documents
- 🛠️ **Tool Calling** — allow the AI to use external capabilities
- 🤖 **Local Assistant System** — voice and computer-interaction capabilities
- 🔐 **Authentication** — API authentication and protected resources
- 🗄️ **PostgreSQL** — persistent application data
- 🐳 **Docker** — containerized SearXNG deployment
- ⚡ **React + TypeScript** — web interface
- 🔵 **Go** — backend/API and AI infrastructure
- 🖥️ **LM Studio** — local LLM inference

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
                 Memory        Tools        Voice


---

📁 Project Structure

NOX/
│
├── AI/
│   ├── auth/
│   ├── brain/
│   ├── clients/
│   ├── config/
│   ├── tools/
│   └── voice/
│
├── API/
│   ├── cmd/
│   └── internal/
│
├── Web/
│   └── src/
│
├── searxng/
│   └── config/
│       └── settings.yml
│
├── docker-compose.yml
├── .gitignore
└── README.md

AI

Contains the AI/agent layer responsible for model interaction, reasoning, memory, tools, authentication-related client functionality, and voice infrastructure.

API

The Go backend responsible for authentication, conversations, persistence, PDF functionality, and communication between the frontend and AI system.

Web

The React + TypeScript frontend for interacting with NOX.

SearXNG

Provides web-search functionality to NOX.


---

🧰 Tech Stack

Component	Technology

Backend	Go
Frontend	React
Frontend Language	TypeScript
Database	PostgreSQL
Search	SearXNG
AI	Local LLM / LM Studio
Voice	whisper.cpp
Containerization	Docker / Docker Compose
Authentication	JWT
Build Tool	Vite



---

🚀 Getting Started

Prerequisites

Install the following:

Git

Go

Node.js

npm

PostgreSQL

Docker Desktop

LM Studio

whisper.cpp


Check the main installations:

git --version
go version
node --version
npm --version
docker --version
docker compose version


---

1. Clone the Repository

git clone https://github.com/Alexakk09/NOX.git

Enter the project:

cd NOX


---

2. Configure the Environment

NOX requires environment configuration for services such as PostgreSQL and the local LLM.

Create the API environment file:

API/
├── .env
└── .env.example

Copy the example:

Copy-Item API\.env.example API\.env

Then configure your local values in:

API/.env

Example:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
DB_SSLMODE=disable

LMSTUDIO_BASE_URL=http://localhost:1234/v1
LMSTUDIO_MODEL=your_model_identifier

> Never commit real API keys, passwords, database credentials, JWT secrets, or other sensitive values to GitHub.




---

3. Configure LM Studio

NOX uses LM Studio as its local LLM provider.

Install LM Studio

Download and install LM Studio:

https://lmstudio.ai/

Download a Model

Open LM Studio and download a compatible chat model.

For example:

google/gemma-4-e2b

The exact model identifier depends on the model installed in your LM Studio environment.

Start the Local Server

In LM Studio:

1. Load the model.


2. Open the Developer / Local Server section.


3. Start the server.



NOX expects the OpenAI-compatible LM Studio API at:

http://localhost:1234/v1

Configure your API/.env:

LMSTUDIO_BASE_URL=http://localhost:1234/v1
LMSTUDIO_MODEL=your_model_identifier

Example:

LMSTUDIO_BASE_URL=http://localhost:1234/v1
LMSTUDIO_MODEL=google/gemma-4-e2b:2

Once LM Studio's local server is running, NOX can send LLM requests to the local model.


---

4. Start SearXNG

NOX uses SearXNG for web search.

From the project root:

docker compose up -d

Check the container:

docker compose ps

SearXNG is exposed at:

http://localhost:8888

The Docker configuration maps:

localhost:8888 → SearXNG:8080

Open it in your browser:

http://localhost:8888

To stop SearXNG:

docker compose down


---

5. Start the API

Open a new terminal.

Move into the API:

cd NOX\API

Install/update Go dependencies:

go mod tidy

Start the server:

go run .\cmd\server

The API runs locally on:

http://localhost:8080

Keep this terminal running.


---

6. Start the AI Service

Open another terminal:

cd NOX\AI

Install/update dependencies:

go mod tidy

Start the AI service:

go run .

The AI service requires the configured LM Studio server to be available.


---

7. Start the Web Interface

Open another terminal:

cd NOX\Web

Install frontend dependencies:

npm install

Start the development server:

npm run dev

Vite will provide a local development URL, normally:

http://localhost:5173

Open that URL in your browser.


---

▶️ Running Everything

During development, NOX currently runs its components separately.

Terminal 1 — SearXNG

cd NOX
docker compose up -d

Terminal 2 — API

cd NOX\API
go run .\cmd\server

Terminal 3 — AI

Make sure LM Studio is running first, then:

cd NOX\AI
go run .

Terminal 4 — Web

cd NOX\Web
npm run dev

The resulting local system looks like:

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
                    ├── LM Studio
                    ├── Memory
                    ├── Tools
                    ├── PDF processing
                    └── SearXNG :8888


---

🌐 Web Search

NOX uses SearXNG as its web-search provider.

SearXNG configuration:

searxng/
└── config/
    └── settings.yml

JSON responses are enabled for programmatic search:

search:
  formats:
    - html
    - json

NOX communicates with SearXNG through:

http://localhost:8888


---

🧠 Memory

NOX includes a memory system designed to give the assistant persistent context beyond a single conversation.

The system is intended to support:

Long-term memory

Temporary memory

Memory scopes

Memory confidence

Contradiction detection

Memory timeline

User-controlled memory



---

🛠️ Tool System

NOX uses a tool-based architecture to allow the AI to interact with capabilities outside the LLM itself.

The general flow is:

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

This architecture makes it possible to add new capabilities without changing the fundamental conversation system.


---

📄 PDF Intelligence

NOX includes PDF-related AI functionality.

Current document functionality includes:

PDF upload

PDF processing

PDF summarization

PDF explanation

AI interaction with document content


The document system is intended to evolve toward richer retrieval and multi-document workflows.


---

🤖 Voice / Local Assistant

NOX includes local assistant infrastructure intended to support voice and computer-related capabilities.

The system is designed to support capabilities such as:

Voice interaction

Speech recognition

Local application interaction

Computer tools

Extensible actions


Voice functionality uses:

whisper.cpp


---

🔐 Security

NOX uses authentication and protected API resources.

Because NOX is intended to eventually perform real-world actions, tool execution is designed with future permission and approval mechanisms in mind.

Potential consequential actions include:

Sending messages

Sending emails

Modifying files

Running system commands

Interacting with applications

Performing external actions


The long-term goal is to separate:

Reasoning
    ↓
Permission
    ↓
Action

so that the AI's ability to reason does not automatically mean it has unrestricted permission to act.


---

🗺️ Roadmap

AI

[x] Conversational AI

[x] Local LLM integration

[x] Tool architecture

[x] Memory system

[ ] Advanced memory controls

[ ] Memory timeline

[ ] Memory confidence

[ ] Contradiction detection

[ ] Temporary memory

[ ] Project-scoped memory


Web Intelligence

[x] SearXNG integration

[x] Web search

[ ] Source reliability

[ ] Claim/evidence tracking

[ ] Multi-step research

[ ] Research memory


Documents

[x] PDF processing

[x] PDF summarization

[x] PDF explanation

[ ] Multiple file support

[ ] Drag-and-drop files

[ ] Improved document retrieval

[ ] Image understanding


Agents

[x] Tool registry

[x] Tool execution

[ ] Permission system

[ ] Action approval

[ ] Agent checkpoints

[ ] Undoable actions

[ ] Agent activity timeline

[ ] More external integrations


Voice / Local Assistant

[x] Voice infrastructure

[x] Local tools

[ ] Improved voice interaction

[ ] Expanded computer control

[ ] Safer action execution



---

🚧 Development Status

NOX is currently under active development.

The architecture and APIs may change as new capabilities are introduced.

The project is being developed incrementally with the long-term goal of building a modular personal AI system rather than a simple chatbot interface.


---

🎯 Vision

Most AI applications stop at:

Prompt → LLM → Response

NOX is being built toward:

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

The goal is to build an assistant that can:

understand → remember → research → reason → use tools → act

while keeping consequential actions controlled by the user.