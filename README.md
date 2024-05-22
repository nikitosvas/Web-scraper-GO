# Go Web Scraper: Quick Start Guide

## Step-by-Step Instructions


## Step 1: Clone the Repository

First, **clone** the repository to your local machine: 


```
git clone https://github.com/yourusername/web-scraper-go.git
cd web-scraper-go

```

## Step 2: Install Dependencies

Ensure you have Go installed on your system. Then, install the required dependencies:  


```
go mod tidy
```

## Step 3: Configure the config.yaml File
Edit the config.yaml file to set up the sites you want to scrape and configure the storage options. 


*Example configuration:*

```
sites:
  - name: "Example Site"
    url: "http://example.com"
    
    selectors:  #CSS selector for scraping data from a web page
      item: ".item-class"
      title: ".title-class"

storage:
  type: "file"
  file_path: "data/output.json"

```

## Step 4: Add Allowed Domain


Ensure you specify the allowed domain for the scraper in the scraper.go file. 
This prevents the scraper from visiting unauthorized domains: 


```
for _, site := range sites {
    c := colly.NewCollector(
        colly.AllowedDomains("example.com"),  // Replace "example.com" with the actual domain
        colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
    )
    // Remaining code...
}
```

## Step 5: Step 6: Run the Scraper

Finally, run the scraper:  


```
go run main.go
```
The scraper will read the configuration, scrape the specified sites, and save the data according to the storage **data/contains.json**
