# math-skills Project

This project aims to provide functionality to calculate several statistical measures for a given dataset. By providing functions to calculate these statistical metrics, users can gain insights into the central tendency, spread, and distribution of their data. This project facilitates statistical analysis, aiding in decision-making processes across various domains such as finance, science, engineering, and more.

## Features

- **Average**: The arithmetic mean of all the values in the dataset, calculated by summing all values and dividing by the total count.
- **Median**: The middle value in the dataset when it is sorted in ascending order. If the dataset has an odd number of elements, the median is the middle value; if it has an even number of elements, it is the average of the two middle values.
- **Variance**: A measure of the dispersion or spread of the values in the dataset. It quantifies how much the values in the dataset differ from the mean.
- **Standard Deviation**: The square root of the variance. It provides a measure of the amount of variation or dispersion of a set of values.
- **Error Handling**: the program also handles errors like number of arguments given is not as required

## Getting Started

### Prerequisites

- Go compiler installed ([Download here](https://go.dev/dl/))

### Installation

1. Clone the repository

```bash
 git clone https://learn.zone01kisumu.ke/git/coketch/math-skills.git
 ```
2. Navigate to the project directory: 
```bash
cd math-skills
```

### Usage

- Run the program as below
```bash
go run main.go data.txt
```
  ```

Average: 588
Median: 553
Variance: 54575
Standard Deviation: 234


```

```

```

### Files

- `data.txt`: contains data representing a statistical population: each line contains one value.

### Tests

- Unit tests were done for each function in the Functions package and can be run in the Functions directory using: `go test`

## Roadmap

- Input Improvements: Enhance user input handling for various data formats and edge cases.
- Output Enhancements: Improve output formatting for readability and include additional statistics such as sum and count.
- Performance Optimization: Optimize algorithms for efficient calculations and enhance error handling for better user experience.
- Export Options: Provide options to export calculated statistics to files in various formats.
- User Interface: Introduce an interactive command-line interface with dynamic input options.

## Contributing

Contributions are welcome! Please feel free to submit pull requests for bug fixes or new features.

## License

This project is licensed under the MIT License.

## Acknowledgments

- The Go programming language community.
- Contributors:
    - coketch
