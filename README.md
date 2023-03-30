# backend-test

## Instructions
Clone
```
git clone https://github.com/lutzid/backend-test.git
```

Go into project dir
```
cd backend-test
```

Go into each project (2 terminals)
```
cd golang-auth

cd node-fetch
```

Run Auth App

```
go run main.go
```

Run Fetch App

```
node app.js
```
## C4 Diagram
![](https://raw.githubusercontent.com/lutzid/backend-test/main/images/Deployment%20Diagram.png)
![](https://raw.githubusercontent.com/lutzid/backend-test/main/images/Context%20Diagram.png)

## Goals
- [x] 1. Servers bisa dinyalakan di port berbeda
- [x] 2. Semua endpoint berfungsi dengan semestinya (3 endpoint auth, 3 endpoint fetch)
- [x] 3. Wajib dokumentasi endpoint dengan format OpenAPI/swagger (API.yaml / API.md), atau postman/insomnia collection
- [x] 4. Dokumentasi system diagram-nya dalam format C4 Model (Context dan Deployment)
- [x] 5. Pergunakan satu repo git untuk semua apps (mono repo)
- [] 6. Dockerfile untuk masing-masing app dan wajib menggunakan docker-compose
- [x] 7. Petunjuk penggunaan dan instalasi di README.md yang memudahkan
- [x] 8. kirim video demo / cara penggunaan