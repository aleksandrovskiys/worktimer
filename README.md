# Terminal Work Timer

A lightweight terminal-based work timer that plays an alarm sound after the timer ends. Designed to help you stay focused during work periods and remind you to take breaks.

## Features

- Customizable work and break durations.
- Visual progress bar with time remaining.
- Plays an alarm sound when the timer ends.
- Easy-to-use command-line interface.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/aleksandrovskiys/worktimer
   cd worktimer
   ```

2. Build the project:
   ```bash
   make build
   ```

3. Run the timer:
   ```bash
   ./bin/work -t <work_duration> -b <break_duration>
   ```

## Usage

- **Start the timer**: 
  ```bash
  ./bin/work -t 25 -b 5
  ```
  - `-t` sets the work duration in minutes (default: 25 minutes).
  - `-b` sets the break duration in minutes (default: 5 minutes).

- **Controls**:
  - Press `Ctrl-C` to quit the timer at any time.

## Example

To start a 45-minute work session followed by a 10-minute break:
```bash
./bin/work -t 45 -b 10
```

## How It Works

1. The timer starts with a work period.
2. Once the work period ends, an alarm sound plays, and the break period begins automatically.
3. The cycle repeats until you manually quit the program using `Ctrl-C`.

## Dependencies

- [Charm Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI framework.
- [Charm Lip Gloss](https://github.com/charmbracelet/lipgloss) for styled text.
- [Charm Bubbles](https://github.com/charmbracelet/bubbles) for the progress bar.

## File Structure
- **`main.go`**: Contains the main program logic.
- **`Makefile`**: A simple Makefile for building the project.

