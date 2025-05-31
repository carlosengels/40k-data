# 🧮 MathHammer 40K

A Go-based tool to simulate average damage output in **Warhammer 40,000**. This tool helps players evaluate the statistical effectiveness of units and weapons by modeling core mechanics like **hit rolls, wound rolls, saving throws**, and **damage calculations**.

---

## 📁 Project Structure

```

mathammer/
├── cmd/
│   └── main.go                    # Entry point for running the simulation
├── internal/
│   ├── calcMean/                  # Placeholder for advanced statistical functions
│   ├── calcHits.go                # Hit calculation logic
│   ├── calcHits\_test.go
│   ├── calcWounds.go              # Wound calculation logic
│   ├── calcWounds\_test.go
│   ├── calcSaves.go               # Save logic (armor, invuln)
│   ├── calcSaves\_test.go
│   ├── reRolls.go                 # Re-roll rules logic
│   ├── reRolls\_test.go
│   ├── simulate.go                # Full combat simulation orchestration
│   └── units/
│       └── units.go               # Unit definitions (Attacker, Defender)
├── go.mod                         # Go module file
├── README.md                      # Project documentation
├── .gitignore
├── tmp/                           # Temporary scratch/testing
└── utility/
└── extract.py                 # Python utility (e.g. Wahapedia scraper)

````

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://gitlab.com/your-username/mathammer.git
cd mathammer
````

### 2. Run the simulation

```bash
go run ./cmd/main.go
```

---

## 🧪 Running Unit Tests

Each calculation module includes its own unit tests:

```bash
go test ./internal/...
```

You can also run individual test files:

```bash
go test ./internal -run TestCalcHits
```

---

## 🛠 Features

* Modular calculation of:

  * Ballistic skill (BS) → hits
  * Strength vs. Toughness → wounds
  * Saves (armor, invuln, FNP)
  * Re-roll logic (in progress)
* Uses deterministic math (no random rolling)
* Easily extensible for additional Warhammer rules
* Organized, testable, idiomatic Go code

---

## ✅ Example Output

```
ATTACKER
Attacks: 20, BS: 3, Strength: 4, AP: 2, Damage: 1

DEFENDER
Model Count: 1, Toughness: 5, Wounds: 5, Save: 3, Invuln: 0, FNP: 0

13 out of 20 attacks hit
7 out of 13 hits wound
4 out of 7 wounds go through
Total damage is: 4
```

---

## 🔮 Future Plans

* [ ] Support for D6/D3 damage rolls
* [ ] Advanced invuln/FNP logic
* [ ] Custom CLI input for attacker/defender profiles
* [ ] Output damage distribution charts

---

## 📌 Tech Stack

* [Go](https://golang.org/) — backend simulator
* [Python](https://www.python.org/) — scraping utility (e.g. Wahapedia parsing)

---

## 📄 License

MIT — feel free to use, modify, and contribute!

---

## 🤝 Contributing

Pull requests welcome! If you'd like to help implement new rules (like re-rolls, mortal wounds, or stratagems), feel free to fork and submit a PR.

```

Let me know if you'd like a GitLab badge section, CI instructions, or API usage added!
```

README is AI generated and human verified.
