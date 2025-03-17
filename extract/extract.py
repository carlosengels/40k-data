import requests
from bs4 import BeautifulSoup
import os

"""
This script reads a list of URLs from 'links.txt', fetches their text content, and saves each page as a .txt file.
Usage:
1. Create 'links.txt' and list each URL on a new line.
2. Run this script to download and save the text content of each page in the 'space_marines_data' folder.
"""

INPUT_FILE = "extract/ultramarines_links.txt"  # File containing the list of URLs
OUTPUT_DIR = "space_marines_data"

def fetch_text(url):
    """Fetch and extract text content from a URL."""
    try:
        response = requests.get(url)
        response.raise_for_status()
        soup = BeautifulSoup(response.text, "html.parser")
        return soup.get_text()
    except requests.exceptions.RequestException as e:
        print(f"Error fetching {url}: {e}")
        return None

def save_text(filename, text):
    """Save text content to a file."""
    os.makedirs(OUTPUT_DIR, exist_ok=True)
    filepath = os.path.join(OUTPUT_DIR, filename)

    with open(filepath, "w", encoding="utf-8") as file:
        file.write(text)

    print(f"Saved: {filepath}")

def main():
    # Read URLs from the file
    try:
        with open(INPUT_FILE, "r", encoding="utf-8") as file:
            urls = [line.strip() for line in file if line.strip()]
    except FileNotFoundError:
        print(f"Error: {INPUT_FILE} not found.")
        return

    # Process each URL
    for url in urls:
        text = fetch_text(url)
        if text:
            filename = url.rstrip("/").split("/")[-1] + ".txt"  # Extract last part of URL as filename
            save_text(filename, text)

if __name__ == "__main__":
    main()