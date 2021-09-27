# Online Learning Platform Admin
## Deskripsi
Untuk repo ini digunakan sebagai microservice admin, terdapat fitur crud course, delete user dan statistic

## Cara Install
### Pertama 
Jalankan perintah dibawah ini pada cmd:

	go get -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && \
	go mod vendor && swag init                                                                                                      
atau ketikkan (menggunakan make file)

    make install

### Selanjutnya
Copy main.example.json lalu ubah namanya menjadi
main.json dan terakhir setting database mysql

    "database": {
        "mysql": {
            "host": "localhost",
            "port": "3306",
            "dbname": "olp",
            "user": "root",
            "password": ""
        }
    }

setelah itu buka dan clone repo berikut
https://github.com/fauzanmh/olp-user
https://github.com/fauzanmh/olp-migration



### Cara Menjalankan
Jalankan perintah dibawah ini pada cmd:
    
    make run


### Dokumentasi API (Swagger)

    http://localhost:8099/api/swagger/index.html
