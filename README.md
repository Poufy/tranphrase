# Tranphrase
Tranphrase is a tool that allows you to upload your excel files containing rows of text in a specific language and returns an enhanced translation of that text by using GPT in order to rephrase the original text. This gives the translation more originality and a fresher feel.

## Project Structure
```go
my_project/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   ├── config.go
│   └── config.yaml
├── controllers/
│   └── translate_controller.go
├── middlewares/
│   ├── auth.go
│   └── logger.go
│   └── excelFileCheck.go
├── models/
├── routes/
│   ├── translate_routes.go
│   └── api.go
├── services/
│   ├── translate_service.go
│   ├── check_excel_extension_service.go
├── static/
│   ├── index.html
│   ├── script.js
├── temp/

├── README.md
├── go.mod
└── go.sum
```
## Getting Started
### Prerequisites
* Go (version 1.17 or higher)

### Setup
1. Clone this repository:
```sh
git clone git@github.com:Poufy/tranphrase.git
```

2. Navigate to the project directory:
```sh
cd tranphrase
```
3. Update the config.yaml file in the config/ directory with your desired configuration settings.

## Running the server
1. Run the server:
```sh
go run cmd/server/main.go
```

The server will listen on port 8080 by default. You can change the listening port in the config.yaml file.

## Endpoints

* `POST /api/translate`: Takes an excel file and returns an excel file with the translations next to the original text.


## License
This project is licensed under the MIT License. See the LICENSE file for details.