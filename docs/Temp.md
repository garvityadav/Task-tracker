# Tasker ><

- Helps in logging todo(s).
    - CRUD
        - Create a new task.
        - Read existing tasks.
        - Update existing tasks
        - Delete existing tasks.
- A pomodoro section
    - CRUD
        - Creates a pomodoro
        - Read existing settings
        - Update the time (focus duration , short break, long break)
        - Delete existing pomodoro.
- Timer
    - CRUD
- Prints a nice CLI table which shows existing tasks or timer.

Which language should i use ,Go , Python , JAVA or JS ? 
GOLang

---

## Go naming conventions:   
  - **Packages & directories** : lowercase, no underscores → auth, user, taskmanager.
  - Variables & functions: camelCase → taskList, addTask().

  - Exported (public) items: PascalCase → AddTask(), TaskList.

  - Constants: PascalCase for exported, camelCase for internal → MaxRetries, defaultLimit.

  - Struct fields: PascalCase if exported → UserName string.

  - Interfaces: often end with -er → Reader, Writer, Fetcher.

  - File names: all lowercase, may use underscores for clarity → user_service.go, task_handler.go.

---

Some details about CLI apps 

- Command : 
$ tasker pomo start
$ tasker todo list 
$ tasker timer add

   
    So these are the commands where “tasker is the main command and “pomo” , “todo” and “timer” are sub-commands and “start” , “list” and “add” are sub-subcommands or flags. 
----


Functions : 
To create a task : addTask()
