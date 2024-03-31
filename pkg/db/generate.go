package db

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Transactor -o ./mocks/ -s "_minimock.go"
//go:generate minimock -g -i github.com/jackc/pgx/v5.Tx -o ./mocks/ -s "_minimock.go"
