# Documentation

## 1. Clone or extract file

extract this zip file into your computer

## 2. Installation

setup your local machine by install 

- Docker 

## 3. Setup Envvar

in your local machine, make sure that .env file is already created in folder backend. this is to load certain values for database connection.

```
  RDS_POSTGRES_HOST=localhost
  RDS_POSTGRES_PORT=5432
  RDS_POSTGRES_DATABASE=kumparan
  RDS_POSTGRES_USERNAME=root
  RDS_POSTGRES_PASSWORD=secret
  RDS_POSTGRES_SSL_CERT=disable
  TIMEZONE=Asia/Jakarta
  PORT=50050
```
## 4. Run Command

once everyhing is done please execute below command: 
`docker-compose up` this command is to create 3 container image for postgres, redis and golang

## 5. Test Program

- from your browser or postman, please run below endpoint
    ```localhost:50050/v1/article/add``` 
    to add article by input : title, author, body.
    sample curl to add article :

      ```
        curl --location --request POST 'localhost:50050/v1/article/add' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "title" : "Hujan yang jatuh tak pernah membenci angin",
            "author" : "Rafa",
             "body" : "Daun yang jatuh tak pernah membenci angina, Dia membiarkan dirinya jatuh begitu saja, Tak Melawan. Mengikhlaskan Semuanya"
          }'
      ```

    ```localhost:50050/v1/article/list``` to display all articles
      sample curl to display all articles

      ```
        curl --location --request GET 'localhost:50050/v1/article/list'
      ```
    

    ```localhost:50050/v1/article/list?author=Abdul``` to display article by paramater author
      sample curl to get article by input author parameter

      ```
        curl --location --request GET 'localhost:50050/v1/article/list?author=Abdul'
      ```
      
    ```localhost:50050/v1/article/list?query=drama korea``` to display article by parameter query
      sample curl to get article by input query parameter

      ```
        curl --location --request GET 'localhost:50050/v1/article/list?query=drama korea'
      ```

    ```localhost:50050/v1/article/list?query=drama korea&author=Maria Ulfa``` to display article by both parameters (query and author)
      sample curl to get article by author and query
      
        ```
          curl --location --request GET 'localhost:50050/v1/article/list?query=drama korea&author=Maria Ulfa'
        ```
     sample response:
     
        ```
          {
         "success": true,
         "error": null,
          "msg": "success",
         "data": [
        {
            "ID": 5,
            "Author": "Maria Ulfa",
            "Title": "botania garden",
            "Body": "nongsa garden adalah sebuah drama korea yang diperankan oleh aktor ternama, siapa lagi kalau bukan kim jong un",
            "Created": "2022-07-03T09:04:04.81721Z"
        },
        {
            "ID": 4,
            "Author": "Maria Ulfa",
            "Title": "botania garden",
            "Body": "nongsa garden adalah sebuah drama korea yang diperankan oleh aktor ternama, siapa lagi kalau bukan kim jong un",
            "Created": "2022-07-03T08:59:09.945473Z"
        },
        {
            "ID": 1,
            "Author": "Maria Ulfa",
            "Title": "botania garden",
            "Body": "nongsa garden adalah sebuah drama korea yang diperankan oleh aktor ternama, siapa lagi kalau bukan kim jong un",
            "Created": "2022-07-03T08:47:27.697622Z"
           }
         ]
        }



