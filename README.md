# Financial Reporter CLI
Financial Reporter is a simple command-line application built in Go for analyzing financial data from CSV files and generating summary reports. The project was created for personal use to streamline financial analysis and automate the report generation process.


## 🚀 Features
- Parse and analyze financial data from CSV files.
- Generate summary reports by category and subcategory.
- Filter data by date range
- Export reports to JSON


## 🛠️ Purpose
This project was developed for personal use to simplify financial data analysis and automate report generation.


## 📄 Input CSV Format
Example CSV file format can be found in `inputfiles` directory.


## 📖 Usage
Example use to generate a summary report from a CSV file
```
./bin/finalize-cli -input inputfiles/financial_report.csv -date "2024-01-01:2024-02-01"    
```


## 🔍 Example Output
```
{
  "Losses": {
    "Product Loss": {
      "EUR": 18876.18
    }
  },
  "Marketing Costs": {
    "Promotions": {
      "USD": 8816.26
    },
    "SEO": {
      "EUR": -4670.99
    }
  },
  "Operating Costs": {
    "Rent": {
      "PLN": 6728.19
    },
    "Salaries": {
      "USD": 10284.91
    }
  },
  "Profits": {
    "Gross Profit": {
      "PLN": -20862.89
    },
    "Net Profit": {
      "EUR": 4534.77
    }
  },
  "Taxes": {
    "Income Tax": {
      "EUR": 19545.68,
      "PLN": -14732.29
    },
    "VAT": {
      "EUR": 10833.04,
      "USD": -11623.82
    }
  }
}
```


## 👨‍💻 Author
- FJ (github.com/jasaf)