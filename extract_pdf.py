#!/usr/bin/env python3
import pypdf
import sys

def extract_text_from_pdf(pdf_path):
    reader = pypdf.PdfReader(pdf_path)
    text = ""
    for page in reader.pages:
        text += page.extract_text() + "\n"
    return text

if __name__ == "__main__":
    pdf_path = "/Users/xuhaolan/Documents/Code/Go/LearnGo/GoLang.pdf"
    text = extract_text_from_pdf(pdf_path)
    print(text)
