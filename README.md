# sb-go-quiz-3-cakrayudha
Quiz 3 dari bootcamp Sanbercode (Golang) Batch 66 (dikerjakan oleh Cakrayudha Kusuma Adi)

## List path users
(POST)      https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/users/signup
(Login)

(POST)      https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/users/login (SignUp)

## List path categories
(GET)       https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/categories (GetAllCategories)

(GET)       https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/categories/:id (GetCategories)

(GET)       https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/categories/:id/books (GetBooksByCategories)

(POST)      https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/categories (CreateCategories)

(PUT)       https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/categories/:id (UpdateCategories)

(DELETE)    https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/categories/:id (DeleteCategories)

## List path books
(GET)       https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/books (GetAllBooks)

(GET)       https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/books/:id (GetBooks)

(POST)      https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/books (CreateBooks)

(PUT)       https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/books/:id (UpdateBooks)

(DELETE)    https://sb-go-quiz-3-cakrayudha-production.up.railway.app/api/books/:id (DeleteBooks)

## CARA PENGGUNAAN
1.  Buat collection baru pada Postman dengan cara import file "Sanbercode Golang - Cakrayudha.postman_collection" pada repository ini
2. Buat akun baru dengan hit API "SignUp" (masukkan username, password, dan re_type_password)
3. Login pada API "Login" dengan menggunakan username dan password yang telah didaftarkan, hit API
4. Copy token ketika login, Paste pada kolom Authorization di collection, gunakan type Bearer Token
5. Create Category terlebih dahulu dengan hit API "CreateCategories"
6. Copy id category, kemudian create Book dengan hit API "CreateBooks" dengan category_id menggunakan id yang telah dicopy
7. Selesai, Anda dapat menggunakan endpoint yang lain untuk mendapatkan data, update, maupun delete data