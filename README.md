# Calorie Tracker CLI

A command-line tool to track daily calories, food, weight, and goals. All data is stored locally as JSON — no account, no network.

> A personal project built to learn Go and practice clean architecture (layered `domain` / `service` / `storage` / `cli`, dependency injection, and unit testing).

## Features

- Create a user profile with body metrics and a weight goal
- Log food items with automatic per-serving calorie calculation
- See remaining calories for any given day
- Track weight over time and plot progress as a PNG
- Edit or delete food entries, weight entries, and goals

## Requirements

- [Go 1.25+](https://go.dev/dl/)

## Installation

```bash
git clone https://github.com/acosio14/calories-tracker-cli
cd calories-tracker-cli
go build
```

This produces a `calories-tracker-cli` binary in the project directory.

## Usage

| Command | Description |
|---|---|
| `create <name>` | Create a user (prompts for metrics interactively) |
| `add food-item <name> --meal --calories --serving-size --quantity` | Log a food item |
| `add weight <value> [--date]` | Log a weight measurement |
| `edit goal` | Update your goal (prompts interactively) |
| `edit daily-calories <value>` | Update your daily calorie limit |
| `edit food-item --id [--date --meal --name --serving-size --calories --quantity]` | Edit fields of a food item |
| `delete food-item --id <id>` | Delete a food item |
| `delete weight-entry --id <id>` | Delete a weight entry |
| `view calories [--date]` | Show logged food and remaining calories for a day |
| `view plot` | Save a weight-progress chart to `weight.png` |

Dates use the format `MM/DD/YYYY`. Date flags also accept `today` (the default).

### Create a user

```bash
$ ./calories-tracker-cli create john
Enter Birthday (MM/DD/YYYY):06/25/1995
Gender (M/F): M
Height (inches): 72
Goal (lose/maintain/gain): lose
Current weight (lbs): 200
Goal weight (lbs): 180
Duration (weeks): 12
Daily Calories intake: 2000
```

This writes a user file (`john.json`):

```json
{
	"name": "john",
	"age": 31,
	"gender": "M",
	"height": 72,
	"goal": {
		"type": "lose",
		"weight": 180,
		"rate": -1.6666666666666667,
		"daily_calories": 2000
	},
	"weight": [
		{
			"id": 0,
			"date": "2026-06-25T05:54:23.842674-05:00",
			"weight_value": 200
		}
	],
	"food_items": []
}
```

### Add food

The `--meal`, `--calories`, `--serving-size`, and `--quantity` flags are required. Total calories are computed as `calories * (quantity / serving-size)`.

```bash
$ ./calories-tracker-cli add food-item oatmeal --meal breakfast --calories 160 --serving-size 35 --quantity 43
$ ./calories-tracker-cli add food-item eggs --meal breakfast --calories 70 --serving-size 1 --quantity 3
```

Each entry is appended to `food_items`:

```json
"food_items": [
	{
		"id": 0,
		"date": "2026-06-25T06:04:27.997046-05:00",
		"meal": "breakfast",
		"name": "oatmeal",
		"serving_size": 35,
		"calories_per_serving": 160,
		"quantity": 43,
		"total_calories": 196.57142857142858
	},
	{
		"id": 1,
		"date": "2026-06-25T06:05:14.190087-05:00",
		"meal": "breakfast",
		"name": "eggs",
		"serving_size": 1,
		"calories_per_serving": 70,
		"quantity": 3,
		"total_calories": 210
	}
]
```

### Add weight

```bash
$ ./calories-tracker-cli add weight 198 --date 06/26/2026
```

### View remaining calories

```bash
$ ./calories-tracker-cli view calories --date today
oatmeal | 43.00 | 196.57
eggs | 3.00 | 210.00
Remaining Calories for 06/25/2026: 1593.43
```

### Edit and delete

```bash
# Edit only the fields you pass; others are left unchanged
$ ./calories-tracker-cli edit food-item --id 0 --quantity 50

$ ./calories-tracker-cli delete food-item --id 1
$ ./calories-tracker-cli delete weight-entry --id 0

$ ./calories-tracker-cli edit daily-calories 1800
```

### Plot weight progress

```bash
$ ./calories-tracker-cli view plot
```

Generates `weight.png` charting your logged weights over time.

## Data storage

User data is stored as JSON in `~/calories-tracker-output`, one file per user. The directory is created automatically on first `create`.

## Development

Run the test suite:

```bash
go test ./...
```

Tests use an in-memory fake store, so they run without touching the filesystem.

## Project status

This is an active personal learning project. See [`docs/design.md`](docs/design.md) for design notes and planned features.
