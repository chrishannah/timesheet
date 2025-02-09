# Timesheet

I have to fill in a weekly timesheet at work, where I detail my time spent on
each project down to roughly the nearest hour.

I want it to be accurate, but I don't want to spend too much time managing it.
So this is my attempt at making the process a bit more efficient.

## Usage

**List all available tasks**:

````bash
timesheet list
````

Example output:
````
1: Example project (Total time: 0h 0m)
2: Super important project (Total time: 0h 0m)
3: Random task (Total time: 0h 0m)
````

**Add a new task to the list**:

````bash
timesheet add "Important Task"
````

**Delete a task** (by it's ID found via list):

````bash
timesheet delete 1
````

**Rename a task** (by it's ID found via list):

````bash
timesheet rename 1 "New project name"
````

**Start a task timer** (by it's ID found via list):

````bash
timesheet start 1
````

**Stop the current task timer**:

````bash
timesheet stop
````

**Get current task**:

````bash
timesheet current
````

Example output:

````
Current task: Example project (ID: 1)
Running for: 2s
````

**Reset all tasks**:

````bash
timesheet reset
````

By default this will reset all tracked time of tasks to 0.

Use the `--force` flag to also delete all tasks.


