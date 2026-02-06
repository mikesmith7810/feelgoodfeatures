# FeelGoodFeatures

A project to help me learn Go (Golang) :-)

This is an API that uses a local LLM (Ollama) to scrape a pre-configured list of websites for 'feel good' news stories.

A list of story summaries will then be returned to the client in JSON.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Local Ollama Usage](#local-ollama-usage)
- [Features](#features)

## Installation

To run this locally:

```bash
go run .
```

To build and run it for shipping:

```bash
go build -o feelgoodfeatures
./feelgoodfeatures
```

## Usage

Start the API server as described above. Ensure you have a local Ollama instance running (see below).

## API Endpoints

### `GET /api/v1/articles`

Fetches a list of 'feel good' news story summaries.

**Response Example:**
```json
[
  {
    "id": 0,
    "content": "Elton John described the Daily Mail's breach of his family's privacy in relation to the birth of his son Zachary as \"truly sickening\" and \"outside even the most basic standards of human decency.\" He is suing the publisher, Associated Newspapers Limited, along with six others, including Prince Harry and Elizabeth Hurley."
  },
  {
    "id": 1,
    "content": "The UK has issued another recall notice for some baby formula products due to potential contamination with a toxin called cereulide, which can cause vomiting and stomach cramps in babies. 36 children in the UK have already been suspected to be affected by food poisoning from contaminated formula, but none are seriously ill. Food manufacturers Danone and Nestle have recalled several batches of infant formula across over 60 countries since December, with the latest recall covering Aptamil and Cow & Gate first infant and follow-on milks."
  }
]
```

## Local Ollama Usage

This project relies on a local [Ollama](https://ollama.com/) instance to provide LLM capabilities.  
Make sure Ollama is installed and running on your machine before starting the API.

- By default, the API expects Ollama to be accessible at `http://localhost:11434`.
- You can download and run Ollama from [https://ollama.com/download](https://ollama.com/download).
- Ensure the required model (e.g., `llama2`) is available in your Ollama instance.

## Features

- Fetches and summarizes 'feel good' news stories from a pre-configured list of websites.
- Uses a local LLM (Ollama) for natural language processing.
- Simple JSON API for integration with other services or frontends.
