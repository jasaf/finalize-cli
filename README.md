# Financial Reporter CLI
Financial Reporter is a simple command-line application built in Go for analyzing financial data from CSV files and generating summary reports. The project was created for personal use to streamline financial analysis and automate the report generation process.

## 🚀 Features
- Parse and analyze financial data from CSV files.
- Generate summary reports by category and subcategory.
- Filter data by date range
- Export reports to JSON

  
## 📄 Input CSV Format
Example CSV file format can be found in `inputfiles` directory.

## 📖 Usage
Example use to generate a summary report from a CSV file
```
./bin/finalize-cli -input inputfiles/financial_report.csv -date "2024-01-01:2024-02-01"    
```

## 🛠️ Purpose
This project was developed for personal use to simplify financial data analysis and automate report generation.

## 👨‍💻 Author
- FJ (github.com/jasaf)