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
    "content": "A commemorative coin released by the Royal Australian Mint to honor Queen Elizabeth II has been criticized for its unflattering portrayal of the late monarch, with some comparing her face to that of Mrs Doubtfire.\n\nGood news: false",
    "feel_good_article": true
  },
  {
    "id": 1,
    "content": "Sir Elton John testified at the High Court, describing how the Daily Mail breached his family's privacy by publishing articles about his health and the birth of his son Zachary without consent. He felt \"passionately\" about the invasion of their privacy and called it \"outside even the most basic standards of human decency.\" The musician and his husband are suing the publisher of the newspaper for breaches of privacy.\n\nGood news: false",
    "feel_good_article": false
  },
  {
    "id": 2,
    "content": "Graham Norton, Lewis Capaldi, and Taylor Swift star in a surreal music video called Opalite, which premiered on Friday. The video features Domhnall Gleeson as a lonely man who summons Swift into his life with a magic potion on his cactus. The video is set in the 1980s and also stars Cillian Murphy, Greta Lee, and Jodie Turner-Smith.\n\nGood news: true",
    "feel_good_article": true
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
