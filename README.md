# Library App with PostgreSQL&GORM in GOLANG
NOT: Haftaiçi yoğunluktan dolayı yetiştiremedim, 26 mart gecesine kadar tamamlamaya çalışacağım.

## Build with Go. Also used Gorm , PostgreSQL , Godotenv

This app reads a csv list with workerpool and writes them to db. It work with DB querries;

•bookRepo.List --> lists all books

•bookRepo.SearchWord --> searchs book by name

•bookRepo.GetbyID --> prints book details by ID

•bookRepo.SoftDeletebyID --> delete books by ID

•bookRepo.Buy --> buys books by ID and quantity

