# quiz-game

A command-line quiz game that tests your math skills. It reads questions from a CSV file and prompts you to answer them. At the end, you receive a score based on your correct answers.

## Config

| Flag | Default | Description |
|------|---------|-------------|
| `-questions` | `problems.csv` | Path to the CSV file containing questions and answers |
| `-timer` | `30` | Time limit in seconds to complete the quiz |

### Examples

Run with defaults:
```bash
./quiz-game
```

Use a custom questions file:
```bash
./quiz-game -questions my_questions.csv
```

Set a 30-second timer:
```bash
./quiz-game -timer 30
```

Combine flags:
```bash
./quiz-game -questions math.csv -timer 20
```

## CSV Format

The questions file should be a CSV with two columns per row:
```
question,answer
```

Example:
```csv
5+5,10
7+3,10
1+1,2
```
