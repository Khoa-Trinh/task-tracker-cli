# 📝 Task CLI – A Simple Task Manager in Go

A lightweight **command-line task management tool** written in Go that lets you add, update, delete, and mark tasks as in-progress or done. Tasks are stored in a local `tasks.json` file for persistence.

> 📌 This project is inspired by the [roadmap.sh Task Tracker](https://roadmap.sh/projects/task-tracker) challenge.

---

## 🚀 Features
- **Add tasks** with a description.
- **Update tasks** with new descriptions.
- **Delete tasks** by ID.
- **Mark tasks** as `in-progress` or `done`.
- **List tasks** with optional filtering by status.
- **Persistent storage** using a JSON file.

---

## 📂 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Khoa-Trinh/task-tracker-cli.git
   cd task-tracker-cli
    ```
   
2. **Build the application**
    ```bash
    go build -o task-cli main.go
    ```
   
3. **Run the application**
    ```bash
    ./task-cli <command> [arguments]
    ```
   
---
   
## 📋 Usage

### General Syntax
```bash
./task-cli <command> [arguments]
```

### Commands
| Command                         | Description                                                     | Example                              |
|---------------------------------|-----------------------------------------------------------------|--------------------------------------|
| `add <description>`             | Add a new task                                                  | `task-cli add "Write documentation"` |
| `update <id> <new_description>` | Update an existing task’s description                           | `task-cli update 1 "Review PRs"`     |
| `delete <id>`                   | Delete a task by ID                                             | `task-cli delete 1`                  |
| `mark-in-progress <id>`         | Mark a task as in-progress                                      | `task-cli mark-in-progress 2`        |
| `mark-done <id>`                | Mark a task as done                                             | `task-cli mark-done 2`               |
| `list [status]`                 | List all tasks (optional filter: `todo`, `in-progress`, `done`) | `task-cli list done`                 |

### Examples
```bash
# Add tasks
./task-cli add "Finish Go project"
./task-cli add "Read Go documentation"

# List all tasks
./task-cli list

# Mark a task as in-progress
./task-cli mark-in-progress 1

# Mark a task as done
./task-cli mark-done 1

# Delete a task
./task-cli delete 2
```
Sample output:
```yaml
Task added successfully (ID: 1)
Task added successfully (ID: 2)
ID: 1 | Finish Go project | todo | Created: 2025-08-08T10:00:00Z | Updated: 2025-08-08T10:00:00Z
ID: 2 | Read Go documentation | todo | Created: 2025-08-08T10:01:00Z | Updated: 2025-08-08T10:01:00Z
Task marked as in-progress (ID: 1)
Task marked as done (ID: 1)
Task deleted successfully (ID: 2)
```

---

## 🛠 How It Works

### Data Storage
Tasks are saved in a file called `tasks.json` in the current working directory.

Example structure:
```json
[
  {
    "id": 1,
    "description": "Finish Go project",
    "status": "done",
    "created_at": "2025-08-08T10:00:00Z",
    "updated_at": "2025-08-08T10:05:00Z"
  }
]
```

### Status Values
- `todo`: Task is pending. **(default)**
- `in-progress`: Task is currently being worked on.
- `done`: Task has been completed.

---

## 📄 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details
```pgsql

This `.md` file contains **everything** — description, features, installation, usage, examples, data format, and license — all in a single file.  

If you want, I can also add **ASCII art banners** at the top so the CLI feels cooler when running. That would make your README and tool pop visually.

```